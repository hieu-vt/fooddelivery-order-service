package ordertrackingbiz

import (
	"context"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/modules/ordertracking/ordertrackingmodel"
)

type updateOrderTrackingStore interface {
	Update(ctx context.Context, data *ordertrackingmodel.UpdateOrderTracking) error
}

type updateOrderBiz struct {
	store updateOrderTrackingStore
}

func NewUpdateOrderBiz(store updateOrderTrackingStore) *updateOrderBiz {
	return &updateOrderBiz{store: store}
}

func (biz *updateOrderBiz) UpdateOrderTracking(ctx context.Context, orderId int, data *ordertrackingmodel.UpdateOrderTracking) error {
	data.OrderId = orderId
	if err := biz.store.Update(ctx, data); err != nil {
		return common.ErrNoPermission(err)
	}

	return nil
}
