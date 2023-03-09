package orderdetailbiz

import (
	"context"
	"order-service/common"
	"order-service/modules/order/ordermodel"
	"order-service/modules/orderdetails/orderdetailmodel"
)

type OrderDetailStore interface {
	Create(ctx context.Context, orderDetail *orderdetailmodel.OrderDetail) error
}

type OrderStore interface {
	FindByCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*ordermodel.Order, error)
}

type orderDetailBiz struct {
	store      OrderDetailStore
	orderStore OrderStore
	//pubsub     pubsub.Pubsub
}

func NewOrderDetailBiz(store OrderDetailStore, orderStore OrderStore) *orderDetailBiz {
	return &orderDetailBiz{
		store:      store,
		orderStore: orderStore,
	}
}

func (biz *orderDetailBiz) CreateOrderDetail(ctx context.Context, data *orderdetailmodel.OrderDetail) error {
	if err := data.ValidateOrderDetailData(); err != nil {
		return common.ErrNoPermission(err)
	}

	order, err := biz.orderStore.FindByCondition(ctx, map[string]interface{}{"id": data.OrderId})

	if err != nil {
		return common.ErrEntityNotFound(ordermodel.TableOrderName, err)
	}

	if order.Status == 0 {
		return common.ErrEntityNotFound(ordermodel.TableOrderName, err)
	}

	data.Status = 1

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(orderdetailmodel.TableNameOrderDetail, err)
	}

	//biz.pubsub.Publish(ctx, common.TopicCreateOrderTrackingAfterCreateOrderDetail, pubsub.NewMessage(ordertrackingmodel.CreateOrderTracking{
	//	//SqlModel: common.SqlModel{},
	//	OrderId: data.OrderId,
	//	State:   common.WaitingForShipper,
	//}))

	return nil
}
