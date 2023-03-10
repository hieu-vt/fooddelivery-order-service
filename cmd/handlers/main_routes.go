package handlers

import (
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/modules/order/ordertransport/ginorder"
	"fooddelivery-order-service/modules/orderdetails/orderdetailtransport/ginorderdetail"
	"fooddelivery-order-service/modules/ordertracking/ordertrackingtranport/ginordertracking"
	"fooddelivery-order-service/plugin/sckio"
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/200Lab-Education/go-sdk/logger"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"log"
	"net/http"
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

	router.StaticFile("/demo", "demo.html")

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
	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		log.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})

	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		log.Println("recv " + msg)
		return "recv " + msg
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason)
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
