package handlers

import (
	"context"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/component/asyncjob"
	"fooddelivery-order-service/modules/ordertracking/ordertrackingbiz"
	"fooddelivery-order-service/modules/ordertracking/ordertrackingmodel"
	"fooddelivery-order-service/modules/ordertracking/ordertrackingstorage"
	"fooddelivery-order-service/plugin/pubsub"
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/200Lab-Education/go-sdk/logger"
	"gorm.io/gorm"
)

func StartUpdateOrderTracking(service goservice.ServiceContext) {
	go func() {
		common.AppRecover()
		pb := service.MustGet(common.PluginNats).(pubsub.NatsPubSub)
		db := service.MustGet(common.DBMain).(*gorm.DB)
		log := service.Logger("Update order tracking").(logger.Logger)

		ctx := context.Background()

		ch, _ := pb.Subscribe(ctx, common.TopicOrderTrackingUpdate)

		for msg := range ch {
			job := asyncjob.NewJob(func(ctx context.Context) error {
				trackingType := msg.Data()["type"].(string)
				orderIdString := msg.Data()["order_id"].(string)

				uid, _ := common.FromBase58(orderIdString)

				biz := ordertrackingbiz.NewUpdateOrderBiz(ordertrackingstorage.NewSqlStore(db))

				err := biz.UpdateOrderTracking(ctx, int(uid.GetLocalID()), &ordertrackingmodel.UpdateOrderTracking{State: common.TrackingType(trackingType)})

				if err != nil {
					return err
				}

				return nil
			})

			if err := asyncjob.NewGroup(true, job).Run(ctx); err != nil {
				log.Infoln(err)
			}
		}
	}()
}
