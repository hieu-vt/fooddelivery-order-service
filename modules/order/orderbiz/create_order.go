package orderbiz

import (
	"context"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/modules/order/ordermodel"
)

type OrderStore interface {
	Create(ctx context.Context, order *ordermodel.CreateOrder) error
}

type createOrderBiz struct {
	orderStore OrderStore
}

func NewCreateOrderBiz(orderStore OrderStore) *createOrderBiz {
	return &createOrderBiz{orderStore: orderStore}
}

func (biz *createOrderBiz) CreateOrder(ctx context.Context, order *ordermodel.CreateOrder) error {
	if err := order.ValidateOrderData(); err != nil {
		return common.ErrNoPermission(err)
	}

	order.Status = 1

	if err := biz.orderStore.Create(ctx, order); err != nil {
		return common.ErrCannotCreateEntity(ordermodel.TableOrderName, err)
	}

	return nil
}
