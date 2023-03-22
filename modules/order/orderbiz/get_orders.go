package orderbiz

import (
	"context"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/modules/order/ordermodel"
)

type orderRepository interface {
	GetOrders(
		ctx context.Context,
		userId int,
		paging common.Paging,
	) ([]ordermodel.GetOrderType, error)
}

type getOrderBiz struct {
	repo orderRepository
}

func NewGetOrderBiz(repo orderRepository) *getOrderBiz {
	return &getOrderBiz{repo: repo}
}

func (biz *getOrderBiz) GetOrders(
	ctx context.Context,
	userId int,
	paging common.Paging,
) ([]ordermodel.GetOrderType, error) {
	orders, err := biz.repo.GetOrders(ctx, userId, paging)

	if err != nil {
		return nil, common.ErrEntityNotFound(ordermodel.TableOrderName, err)
	}

	return orders, nil
}
