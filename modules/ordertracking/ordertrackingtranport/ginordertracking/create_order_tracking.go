package ginordertracking

import (
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/modules/ordertracking/ordertrackingbiz"
	"fooddelivery-order-service/modules/ordertracking/ordertrackingmodel"
	"fooddelivery-order-service/modules/ordertracking/ordertrackingstorage"
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOrderTracking(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data ordertrackingmodel.CreateOrderTrackingParams

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := ordertrackingstorage.NewSqlStore(common.GetMainDb(sc))
		biz := ordertrackingbiz.NewOrderTrackingBiz(store)

		orderId, err := common.FromBase58(data.OrderId)

		if err != nil {
			panic(err)
		}

		createOrder := ordertrackingmodel.OrderTracking{
			OrderId: int(orderId.GetLocalID()),
			State:   data.State,
		}

		if err := biz.CreateOrderTracking(c, &createOrder); err != nil {
			panic(err)
		}

		createOrder.Mask()

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(createOrder.FakeId))
	}
}
