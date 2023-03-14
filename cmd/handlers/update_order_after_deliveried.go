package handlers

import (
	"context"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/component/asyncjob"
	"fooddelivery-order-service/modules/order/orderbiz"
	"fooddelivery-order-service/modules/order/ordermodel"
	"fooddelivery-order-service/modules/order/orderstorage"
	"fooddelivery-order-service/plugin/pubsub"
	goservice "github.com/200Lab-Education/go-sdk"
	"gorm.io/gorm"
	"log"
)

func StartUpdateOrderAfterDelivered(service goservice.ServiceContext) {
	go func() {
		pb := service.MustGet(common.PluginNats).(pubsub.NatsPubSub)
		db := service.MustGet(common.DBMain).(*gorm.DB)

		ch, _ := pb.Subscribe(context.Background(), common.TopicUserUpdateOrder)

		for msg := range ch {
			job := asyncjob.NewJob(func(ctx context.Context) error {
				orderId := msg.Data()["order_id"].(string)
				shipperId := msg.Data()["shipper_id"].(string)

				store := orderstorage.NewSqlStore(db)
				biz := orderbiz.NewUpdateOrderBiz(store)

				uidShipper, _ := common.FromBase58(shipperId)
				uidOrder, _ := common.FromBase58(orderId)

				return biz.UpdateOrder(ctx, int(uidOrder.GetLocalID()), &ordermodel.UpdateOrder{
					ShipperId: int(uidShipper.GetLocalID()),
				})
			})

			if err := asyncjob.NewGroup(true, job).Run(context.Background()); err != nil {
				log.Println(err)
			}
		}

	}()
}
