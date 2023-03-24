package ginorderdetail

import (
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/modules/orderdetails/orderdetailbiz"
	"fooddelivery-order-service/modules/orderdetails/orderdetailstorage"
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetOrderDetail(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		orderId := c.Param("orderId")
		uid, _ := common.FromBase58(orderId)

		log.Println(uid.GetLocalID())

		store := orderdetailstorage.NewSqlStore(common.GetMainDb(sc))
		biz := orderdetailbiz.NewGetOrderDetailBiz(store)

		result, err := biz.GetOrderDetailByOId(c, int(uid.GetShardID()))

		if err != nil {
			panic(err)
		}

		result.Mask()

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
