package ginorder

import (
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/modules/order/orderbiz"
	"fooddelivery-order-service/modules/order/ordermodel"
	"fooddelivery-order-service/modules/order/orderstorage"
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

		store := orderstorage.NewSqlStore(common.GetMainDb(sc))
		biz := orderbiz.NewCreateOrderBiz(store)
		if err := biz.CreateOrder(c, &dataOrder); err != nil {
			panic(err)
		}

		dataOrder.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(dataOrder.FakeId.String()))
	}
}
