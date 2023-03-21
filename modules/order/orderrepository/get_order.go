package orderrepository

import (
	"context"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/modules/order/ordermodel"
	"fooddelivery-order-service/modules/orderdetails/orderdetailmodel"
	"fooddelivery-order-service/modules/ordertracking/ordertrackingmodel"
)

type orderStore interface {
	Find(ctx context.Context, userId int, paging common.Paging) ([]ordermodel.GetOrderType, error)
}

type orderDetailStore interface {
	FindByIds(ctx context.Context, ids []int, orderDetail *orderdetailmodel.OrderDetail) error
}

type orderTrackingStore interface {
	FindByIds(ctx context.Context, ids []int, orderDetail *ordertrackingmodel.OrderTracking) error
}

type getOrderRepository struct {
	store         orderStore
	detailStore   orderDetailStore
	trackingStore orderTrackingStore
}

func NewGetOrderRepository(store orderStore, detailStore orderDetailStore, trackingStore orderTrackingStore) *getOrderRepository {
	return &getOrderRepository{store: store, detailStore: detailStore, trackingStore: trackingStore}
}

func (repo *getOrderRepository) GetOrders(
	ctx context.Context,
	userId int,
	paging common.Paging,
) ([]ordermodel.GetOrderType, error) {
	orders, err := repo.store.Find(ctx, userId, paging)

	if err != nil {
		return nil, common.ErrEntityNotFound(ordermodel.TableOrderName, err)
	}

	orderIds := make([]int, len(orders))

	for i, item := range orders {
		orderIds[i] = item.Id
	}

	var orderDetail orderdetailmodel.OrderDetail
	var orderTracking ordertrackingmodel.OrderTracking

	errODetail := repo.detailStore.FindByIds(ctx, orderIds, &orderDetail)

	if errODetail != nil {
		return nil, common.ErrEntityNotFound(orderdetailmodel.TableNameOrderDetail, errODetail)
	}

	errOTracking := repo.trackingStore.FindByIds(ctx, orderIds, &orderTracking)

	if errOTracking != nil {
		return nil, common.ErrEntityNotFound(ordertrackingmodel.TableNameOrderTracking, errOTracking)
	}

	return orders, nil
}
