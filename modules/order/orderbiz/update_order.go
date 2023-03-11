package orderbiz

import (
	"context"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/modules/order/ordermodel"
)

type OrderUpdateStore interface {
	Update(ctx context.Context, orderId int, orderUpdate *ordermodel.UpdateOrder) error
}

type updateOrderBiz struct {
	orderStore OrderUpdateStore
}

func NewUpdateOrderBiz(orderStore OrderUpdateStore) *updateOrderBiz {
	return &updateOrderBiz{orderStore: orderStore}
}

func (biz *updateOrderBiz) UpdateOrder(ctx context.Context, orderId int, order *ordermodel.UpdateOrder) error {
	if err := biz.orderStore.Update(ctx, orderId, order); err != nil {
		return common.ErrCannotCreateEntity(ordermodel.TableOrderName, err)
	}

	return nil
}
