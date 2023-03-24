package orderdetailbiz

import (
	"context"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/modules/orderdetails/orderdetailmodel"
)

type OrderDetailStore interface {
	Create(ctx context.Context, orderDetail orderdetailmodel.OrderDetail) error
}

type orderDetailBiz struct {
	store OrderDetailStore
}

func NewOrderDetailBiz(store OrderDetailStore) *orderDetailBiz {
	return &orderDetailBiz{
		store: store,
	}
}

func (biz *orderDetailBiz) CreateOrderDetail(ctx context.Context, data orderdetailmodel.OrderDetail) error {
	if err := data.ValidateOrderDetailData(); err != nil {
		return common.ErrNoPermission(err)
	}

	data.Status = 1

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(orderdetailmodel.TableNameOrderDetail, err)
	}

	return nil
}
