package ginorderdetail

import (
	"encoding/json"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/modules/orderdetails/orderdetailbiz"
	"fooddelivery-order-service/modules/orderdetails/orderdetailmodel"
	"fooddelivery-order-service/modules/orderdetails/orderdetailstorage"
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOrderDetail(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var orderDetail orderdetailmodel.CreateOrderDetail

		if err := c.ShouldBind(&orderDetail); err != nil {
			panic(err)
		}

		store := orderdetailstorage.NewSqlStore(common.GetMainDb(sc))
		biz := orderdetailbiz.NewOrderDetailBiz(store)

		orderId, err := common.FromBase58(orderDetail.OrderId)

		if err != nil {
			panic(err)
		}

		jFoodOrigin, fOriginErr := json.Marshal(orderDetail.FoodOrigin)

		if fOriginErr != nil {
			panic(fOriginErr)
		}

		orderDetailCreated := orderdetailmodel.OrderDetail{
			OrderId:    int(orderId.GetLocalID()),
			FoodOrigin: string(jFoodOrigin),
			Price:      orderDetail.Price,
			Quantity:   orderDetail.Quantity,
			Discount:   orderDetail.Discount,
		}

		if err := biz.CreateOrderDetail(c, orderDetailCreated); err != nil {
			panic(err)
		}

		orderDetailCreated.Mask()

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(orderDetailCreated.FakeId))
	}
}
