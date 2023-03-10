package handlers

import (
	"context"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/modules/order/ordertransport/ginorder"
	"fooddelivery-order-service/modules/orderdetails/orderdetailtransport/ginorderdetail"
	"fooddelivery-order-service/modules/ordertracking/ordertrackingtranport/ginordertracking"
	"fooddelivery-order-service/plugin/pubsub/appredis"
	"fooddelivery-order-service/plugin/sckio"
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/200Lab-Education/go-sdk/logger"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"log"
	"net/http"
	"time"
)

func MainRoute(router *gin.Engine, sc goservice.ServiceContext) {
	authClient := sc.MustGet(common.PluginGrpcAuthClient).(interface {
		RequiredAuth(sc goservice.ServiceContext) func(c *gin.Context)
	})
	sIo := sc.MustGet(common.PluginSocket).(interface {
		StartRealtimeServer(engine *gin.Engine, sc goservice.ServiceContext, op sckio.ObserverProvider)
	})
	op := NewObserverProvider()

	sIo.StartRealtimeServer(router, sc, op)

	router.StaticFile("/user", "demo.html")
	router.StaticFile("/shipper", "demoshipper.html")

	v1 := router.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "pong go template"})
		})

		v1.GET("/auth/ping", authClient.RequiredAuth(sc), func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "Auth pong"})
		})

		// Order
		orders := v1.Group("/orders", authClient.RequiredAuth(sc))
		{
			orders.POST("", ginorder.CreateOrder(sc))
			orders.POST("/detail", ginorderdetail.CreateOrderDetail(sc))
			orders.POST("/tracking", ginordertracking.CreateOrderTracking(sc))
			orders.GET("", ginorder.GetOrders(sc))
		}
	}
}

type observerProvider struct {
}

func NewObserverProvider() *observerProvider {
	return &observerProvider{}
}

func (observerProvider) AddObservers(server *socketio.Server, sc goservice.ServiceContext, l logger.Logger) func(conn socketio.Conn) error {
	authClient := sc.MustGet(common.PluginGrpcAuthClient).(interface {
		ValidateToken(token string) *common.User
	})
	redis := sc.MustGet(common.PluginRedis).(appredis.GeoProvider)

	server.OnEvent("/", "Authenticated", func(s socketio.Conn, token string) common.User {
		user := authClient.ValidateToken(token)

		user.Mask()

		s.Emit("Authenticated", user)

		return *user
	})

	server.OnEvent("/", "UserUpdateLocation", func(s socketio.Conn, location common.LocationData) {
		time.Sleep(time.Second * 2)
		log.Println("location", location)
		redis.AddDriverLocation(context.Background(), common.RedisLocation, location.Lng, location.Lat, location.UserId)

		time.Sleep(time.Second * 4)
		if location.UserId == "e5352HrePro4" {
			log.Println("Search location", location)

			lo := redis.SearchDrivers(context.Background(), common.RedisLocation, 10, location.UserId, 100)

			log.Println("Result Search location", lo)
		}
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
