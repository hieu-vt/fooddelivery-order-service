package ginorder

import (
	"encoding/json"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/modules/order/orderbiz"
	"fooddelivery-order-service/modules/order/ordermodel"
	"fooddelivery-order-service/modules/order/orderstorage"
	"fooddelivery-order-service/plugin/pubsub"
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOrder(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataOrder ordermodel.CreateOrder

		if err := c.ShouldBind(&dataOrder); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		dataOrder.UserId = requester.GetUserId()

		jFoodOrigin, fOriginErr := json.Marshal(dataOrder.FoodOriginBody)

		if fOriginErr != nil {
			panic(fOriginErr)
		}

		dataOrder.FoodOrigin = string(jFoodOrigin)

		pb := sc.MustGet(common.PluginNats).(pubsub.NatsPubSub)

		store := orderstorage.NewSqlStore(common.GetMainDb(sc))
		biz := orderbiz.NewCreateOrderBiz(store, pb)
		if err := biz.CreateOrder(c, &dataOrder); err != nil {
			panic(err)
		}

		dataOrder.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(dataOrder.FakeId.String()))
	}
}
