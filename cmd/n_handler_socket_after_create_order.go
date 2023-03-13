package cmd

import (
	"context"
	"errors"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/component/asyncjob"
	"fooddelivery-order-service/modules/order/ordertransport/skorder"
	"fooddelivery-order-service/plugin/appredis"
	appgrpc "fooddelivery-order-service/plugin/grpc"
	"fooddelivery-order-service/plugin/nsckio"
	"fooddelivery-order-service/plugin/pubsub"
	"fooddelivery-order-service/plugin/pubsub/nats"
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/200Lab-Education/go-sdk/logger"
	"github.com/spf13/cobra"
	"github.com/zishang520/socket.io/socket"
	"log"
	"time"
)

var StartNHandleSocketAfterCreateOrder = &cobra.Command{
	Use:   "new-handle-socket-order",
	Short: "New handle socket after create order",
	Run: func(cmd *cobra.Command, args []string) {
		service := goservice.New(
			goservice.WithInitRunnable(nats.NewNatsPubSub(common.PluginNats)),
			goservice.WithInitRunnable(appgrpc.NewAuthClient(common.PluginGrpcAuthClient)),
			goservice.WithInitRunnable(nsckio.NewNSocketIo(common.PluginNSocket)),
			goservice.WithInitRunnable(appredis.NewAppRedis("redis", common.PluginRedis)),
			goservice.WithInitRunnable(nsckio.NewNSocketEngine(common.PluginNSocketEngine)),
		)

		if err := service.Init(); err != nil {
			log.Fatalln(err)
		}

		sIo := service.MustGet(common.PluginNSocket).(interface {
			StartNRealtimeServer(sc goservice.ServiceContext, f func(client *socket.Socket, sc goservice.ServiceContext, l logger.Logger))
		})

		//engine.StaticFile("/user", "demo.html")
		//engine.StaticFile("/shipper", "demoshipper.html")

		sIo.StartNRealtimeServer(service, AddNObservers)

		// handle find shipper
		// handle update state to room
		go func() {
			pb := service.MustGet(common.PluginNats).(pubsub.NatsPubSub)

			redisGeoService := service.MustGet(common.PluginRedis).(appredis.GeoProvider)
			ctx := context.Background()

			ch, _ := pb.Subscribe(ctx, common.TopicUserCreateOrder)

			for msg := range ch {
				lat := msg.Data()["lat"].(float64)
				lng := msg.Data()["lng"].(float64)
				orderId := int(msg.Data()["order_id"].(float64))
				userId := int(msg.Data()["user_id"].(float64))

				log := logger.GetCurrent().GetLogger("handle-socket-order")

				se := service.MustGet(common.PluginNSocketEngine).(nsckio.NSocketEngineProvider)

				userConn := se.GetAppSocket(userId)

				uidOrder := common.NewUID(uint32(orderId), common.DbTypeOrder, 1)
				uidUser := common.NewUID(uint32(userId), common.DbTypeUser, 1)

				orderIdString := &uidOrder
				userIdString := &uidUser

				// join user to room
				if userConn != nil {
					userConn.Join(socket.Room(orderIdString.String()))
				}

				job := asyncjob.NewJob(func(ctx context.Context) error {
					lo := redisGeoService.SearchDrivers(ctx, common.RedisShipperLocation, 1, lat, lng, 100)
					if lo != nil {
						uid, _ := common.FromBase58(lo.Name)
						uConn := se.GetAppSocket(int(uid.GetLocalID()))

						if uConn != nil {
							uConn.Emit(common.NewOrder, skorder.EvenOrderMessageData{
								UserId:    userIdString.String(),
								OrderId:   orderIdString.String(),
								ShipperId: lo.Name,
								Type:      common.WaitingForShipper,
							})

							// join shipper to room
							uConn.Join(socket.Room(orderIdString.String()))
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
		}()

		if err := service.Start(); err != nil {
			log.Fatalln(err)
		}
	},
}

func AddNObservers(client *socket.Socket, sc goservice.ServiceContext, l logger.Logger) {
	authClient := sc.MustGet(common.PluginGrpcAuthClient).(interface {
		ValidateToken(token string) (*common.User, error)
	})
	redis := sc.MustGet(common.PluginRedis).(appredis.GeoProvider)
	se := sc.MustGet(common.PluginNSocketEngine).(nsckio.NSocketEngineProvider)

	client.On(common.Authenticated, func(datas ...any) {
		token := datas[0].(string)
		user, err := authClient.ValidateToken(token)

		if err != nil {
			client.Emit(common.AuthenticationFailed, err.Error())
		}

		se.SaveAppSocket(user.Id, client)

		user.Mask()

		client.Emit(common.Authenticated, user)
	})

	client.On(common.UserUpdateLocation, func(datas ...any) {
		data := datas[0].(map[string]interface{})
		location := common.LocationData{
			Lat:    data["lat"].(float64),
			Lng:    data["lng"].(float64),
			UserId: data["userId"].(string),
			Role:   data["role"].(string),
		}

		time.Sleep(time.Second * 2)

		log.Println("location: ", location)

		switch location.Role {
		case common.RoleShipper:
			redis.AddDriverLocation(context.Background(), common.RedisShipperLocation, location.Lng, location.Lat, location.UserId)
		case common.RoleUser:
			redis.AddDriverLocation(context.Background(), common.RedisShipperLocation, location.Lng, location.Lat, location.UserId)
		}
	})

	client.On(common.OrderTracking, skorder.OnOrderTracking(sc, client))

	client.On("disconnect", func(...any) {
		log.Println("Client disconnected")
	})

	//go func() {
	//	if err := server.Serve(); err != nil {
	//		log.Fatalf("socketio listen error: %s\n", err)
	//	}
	//}()
}
