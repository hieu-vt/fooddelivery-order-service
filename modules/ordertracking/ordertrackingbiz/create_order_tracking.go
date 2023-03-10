package ordertrackingbiz

import (
	"context"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/modules/ordertracking/ordertrackingmodel"
)

type orderTrackingStore interface {
	Create(ctx context.Context, data *ordertrackingmodel.OrderTracking) error
}

type orderTrackingBiz struct {
	store orderTrackingStore
}

func NewOrderTrackingBiz(store orderTrackingStore) *orderTrackingBiz {
	return &orderTrackingBiz{
		store: store,
	}
}

func (biz *orderTrackingBiz) CreateOrderTracking(ctx context.Context, data *ordertrackingmodel.OrderTracking) error {
	if err := data.Validation(); err != nil {
		return common.ErrNoPermission(err)
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(ordertrackingmodel.TableNameOrderTracking, err)
	}

	return nil
}
