package ginorder

import (
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/modules/order/orderbiz"
	"fooddelivery-order-service/modules/order/orderrepository"
	"fooddelivery-order-service/modules/order/orderstorage"
	"fooddelivery-order-service/modules/orderdetails/orderdetailstorage"
	"fooddelivery-order-service/modules/ordertracking/ordertrackingstorage"
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOrders(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.FullFill()

		store := orderstorage.NewSqlStore(common.GetMainDb(sc))
		storeDetail := orderdetailstorage.NewSqlStore(common.GetMainDb(sc))
		storeTracking := ordertrackingstorage.NewSqlStore(common.GetMainDb(sc))

		repo := orderrepository.NewGetOrderRepository(store, storeDetail, storeTracking, nil)
		biz := orderbiz.NewGetOrderBiz(repo)

		result, err := biz.GetOrders(c, int(requester.GetUserId()), paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask()

			if paging.Limit <= len(result) {
				paging.NextCursor = result[i].FakeId.String()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
