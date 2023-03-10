package orderbiz

import (
	"context"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/modules/order/ordermodel"
)

type orderStore interface {
	Find(ctx context.Context, userId int, paging common.Paging) ([]ordermodel.GetOrderType, error)
}

type getOrderBiz struct {
	store orderStore
}

func NewGetOrderBiz(store orderStore) *getOrderBiz {
	return &getOrderBiz{store: store}
}

func (biz *getOrderBiz) GetOrders(
	ctx context.Context,
	userId int,
	paging common.Paging,
) ([]ordermodel.GetOrderType, error) {
	orders, err := biz.store.Find(ctx, userId, paging)

	if err != nil {
		return nil, common.ErrEntityNotFound(ordermodel.TableOrderName, err)
	}

	return orders, nil
}
