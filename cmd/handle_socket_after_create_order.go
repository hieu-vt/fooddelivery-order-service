package cmd

import (
	"context"
	"errors"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/component/asyncjob"
	"fooddelivery-order-service/modules/order/ordertransport/skorder"
	"fooddelivery-order-service/plugin/appredis"
	appgrpc "fooddelivery-order-service/plugin/grpc"
	"fooddelivery-order-service/plugin/pubsub"
	"fooddelivery-order-service/plugin/pubsub/nats"
	"fooddelivery-order-service/plugin/sckio"
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/200Lab-Education/go-sdk/logger"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/spf13/cobra"
	"log"
	"time"
)

var StartHandleSocketAfterCreateOrder = &cobra.Command{
	Use:   "handle-socket-order",
	Short: "Handle socket after create order",
	Run: func(cmd *cobra.Command, args []string) {
		service := goservice.New(
			goservice.WithName("food-delivery-socket-order"),
			goservice.WithVersion("1.0.0"),
			goservice.WithInitRunnable(nats.NewNatsPubSub(common.PluginNats)),
			goservice.WithInitRunnable(appgrpc.NewAuthClient(common.PluginGrpcAuthClient)),
			goservice.WithInitRunnable(sckio.NewSocketIo(common.PluginSocket)),
			goservice.WithInitRunnable(appredis.NewAppRedis("redis", common.PluginRedis)),
			goservice.WithInitRunnable(sckio.NewSocketEngine(common.PluginSocketEngine)),
		)

		if err := service.Init(); err != nil {
			log.Fatalln(err)
		}

		service.HTTPServer().AddHandler(func(engine *gin.Engine) {
			sIo := service.MustGet(common.PluginSocket).(interface {
				StartRealtimeServer(engine *gin.Engine, sc goservice.ServiceContext, op sckio.ObserverProvider)
			})
			op := NewObserverProvider()
			engine.Use(func(c *gin.Context) {
				c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
				c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
				c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

				if c.Request.Method == "OPTIONS" {
					c.AbortWithStatus(204)
					return
				}

				c.Next()
			})

			engine.StaticFile("/user", "demo.html")
			engine.StaticFile("/shipper", "demoshipper.html")

			sIo.StartRealtimeServer(engine, service, op)
		})

		pb := service.MustGet(common.PluginNats).(pubsub.NatsPubSub)

		redisGeoService := service.MustGet(common.PluginRedis).(appredis.GeoProvider)
		ctx := context.Background()

		channel, _ := pb.Subscribe(ctx, common.TopicUserCreateOrder)

		go func(ch <-chan *pubsub.Message) {
			for msg := range ch {
				lat := msg.Data()["lat"].(float64)
				lng := msg.Data()["lng"].(float64)
				orderId := int(msg.Data()["order_id"].(float64))
				userId := int(msg.Data()["user_id"].(float64))

				log := logger.GetCurrent().GetLogger("handle-socket-order")

				se := service.MustGet(common.PluginSocketEngine).(sckio.SocketEngineProvider)

				userConn := se.GetAppSocket(userId)

				uidOrder := common.NewUID(uint32(orderId), common.DbTypeOrder, 1)
				uidUser := common.NewUID(uint32(userId), common.DbTypeUser, 1)

				orderIdString := &uidOrder
				userIdString := &uidUser

				if userConn != nil {
					userConn.Join(orderIdString.String())
				}

				job := asyncjob.NewJob(func(ctx context.Context) error {
					lo := redisGeoService.SearchDrivers(ctx, common.RedisShipperLocation, 1, lat, lng, 100)
					if lo != nil {
						uid, _ := common.FromBase58(lo.Name)
						uConn := se.GetAppSocket(int(uid.GetLocalID()))

						if uConn != nil {
							uConn.Emit(common.NewOrder+lo.Name, skorder.EvenOrderMessageData{
								UserId:    userIdString.String(),
								OrderId:   orderIdString.String(),
								ShipperId: lo.Name,
								Type:      common.WaitingForShipper,
							})

							uConn.Join(orderIdString.String())
						}

						return nil
					}
					log.Info("not found shipper continue find")
					return errors.New("not found shipper continue find")
				})

				if err := asyncjob.NewGroup(true, job).Run(ctx); err != nil {
					log.Info(err)
				}
			}
		}(channel)

		if err := service.Start(); err != nil {
			log.Fatalln(err)
		}
	},
}

type observerProvider struct {
}

func NewObserverProvider() *observerProvider {
	return &observerProvider{}
}

func (observerProvider) AddObservers(server *socketio.Server, sc goservice.ServiceContext, l logger.Logger) func(conn socketio.Conn) error {
	authClient := sc.MustGet(common.PluginGrpcAuthClient).(interface {
		ValidateToken(token string) (*common.User, error)
	})
	redis := sc.MustGet(common.PluginRedis).(appredis.GeoProvider)
	se := sc.MustGet(common.PluginSocketEngine).(sckio.SocketEngineProvider)

	server.OnEvent("/", common.Authenticated, func(s sckio.Conn, token string) *common.User {
		user, err := authClient.ValidateToken(token)

		if err != nil {
			s.Emit(common.AuthenticationFailed, err.Error())
			s.Close()
			return nil
		}

		se.SaveAppSocket(user.Id, s)

		user.Mask()

		s.Emit(common.Authenticated, user)

		return user
	})

	server.OnEvent("/", common.UserUpdateLocation, func(s socketio.Conn, location common.LocationData) {
		time.Sleep(time.Second * 2)

		switch location.Role {
		case common.RoleShipper:
			redis.AddDriverLocation(context.Background(), common.RedisShipperLocation, location.Lng, location.Lat, location.UserId)
		case common.RoleUser:
			redis.AddDriverLocation(context.Background(), common.RedisShipperLocation, location.Lng, location.Lat, location.UserId)
		}
	})

	server.OnEvent("/", common.OrderTracking, skorder.OnOrderTracking(sc, server))

	server.OnDisconnect("/", func(conn socketio.Conn, s string) {
		log.Println(s)
	})

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()

	return func(s socketio.Conn) error {
		s.SetContext("")
		l.Infoln("connected", s.ID())
		return nil
	}
}
