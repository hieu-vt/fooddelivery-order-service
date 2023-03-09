package handlers

import (
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"net/http"
	"order-service/common"
	"order-service/modules/order/ordertransport/ginorder"
	"order-service/modules/orderdetails/orderdetailtransport/ginorderdetail"
	"order-service/modules/ordertracking/ordertrackingtranport/ginordertracking"
)

func MainRoute(router *gin.Engine, sc goservice.ServiceContext) {
	authClient := sc.MustGet(common.PluginGrpcAuthClient).(interface {
		RequiredAuth(sc goservice.ServiceContext) func(c *gin.Context)
	})

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
