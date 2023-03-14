package cmd

import (
	"context"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/component/asyncjob"
	"fooddelivery-order-service/modules/orderdetails/orderdetailbiz"
	"fooddelivery-order-service/modules/orderdetails/orderdetailmodel"
	"fooddelivery-order-service/modules/orderdetails/orderdetailstorage"
	"fooddelivery-order-service/modules/ordertracking/ordertrackingbiz"
	"fooddelivery-order-service/modules/ordertracking/ordertrackingmodel"
	"fooddelivery-order-service/modules/ordertracking/ordertrackingstorage"
	"fooddelivery-order-service/plugin/pubsub"
	"fooddelivery-order-service/plugin/pubsub/nats"
	"fooddelivery-order-service/plugin/sdkgorm"
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/spf13/cobra"
	"log"
)

var StartCreateOrderDetailTrackingAfterCreateOrder = &cobra.Command{
	Use:   "create-order-detail-tracking",
	Short: "Crate order detail after create order",
	Run: func(cmd *cobra.Command, args []string) {
		service := goservice.New(
			goservice.WithInitRunnable(sdkgorm.NewGormDB("main", common.DBMain)),
			goservice.WithInitRunnable(nats.NewNatsPubSub(common.PluginNats)),
		)

		if err := service.Init(); err != nil {
			log.Fatalln(err)
		}

		pb := service.MustGet(common.PluginNats).(pubsub.NatsPubSub)
		ctx := context.Background()
		ch, _ := pb.Subscribe(ctx, common.TopicUserCreateOrder)

		for msg := range ch {
			jobCreateOrderDetail := asyncjob.NewJob(func(ctx context.Context) error {
				db := common.GetMainDb(service)
				store := orderdetailstorage.NewSqlStore(db)
				biz := orderdetailbiz.NewOrderDetailBiz(store)
				return biz.CreateOrderDetail(ctx, orderdetailmodel.OrderDetail{
					OrderId:    int(msg.Data()["order_id"].(float64)),
					FoodOrigin: msg.Data()["food_origin"].(string),
					Price:      float32(msg.Data()["price"].(float64)),
					Quantity:   int(msg.Data()["quantity"].(float64)),
					Discount:   float32(msg.Data()["discount"].(float64)),
				})
			})

			jobCreateOrderTracking := asyncjob.NewJob(func(ctx context.Context) error {
				db := common.GetMainDb(service)
				store := ordertrackingstorage.NewSqlStore(db)
				biz := ordertrackingbiz.NewOrderTrackingBiz(store)
				return biz.CreateOrderTracking(ctx, ordertrackingmodel.OrderTracking{
					OrderId: int(msg.Data()["order_id"].(float64)),
					State:   common.WaitingForShipper,
				})
			})

			if err := asyncjob.NewGroup(true, jobCreateOrderDetail, jobCreateOrderTracking).Run(ctx); err != nil {
				log.Println(err)
			}
		}

	},
}
