package orderdetailbiz

import (
	"context"
	"encoding/json"
	"errors"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/modules/orderdetails/orderdetailmodel"
)

type getOrderDetailStore interface {
	FindByOId(ctx context.Context, orderId int, moreKeys ...string) (*orderdetailmodel.OrderDetail, error)
}

type getOrderDetailBiz struct {
	store getOrderDetailStore
}

func NewGetOrderDetailBiz(store getOrderDetailStore) *getOrderDetailBiz {
	return &getOrderDetailBiz{store: store}
}

func (biz getOrderDetailBiz) GetOrderDetailByOId(ctx context.Context, orderId int) (*orderdetailmodel.OrderDetail, error) {
	if orderId <= 0 {
		return nil, common.ErrNoPermission(errors.New("order_id must be provide"))
	}

	result, err := biz.store.FindByOId(ctx, orderId)

	if err != nil {
		return nil, common.ErrCannotGetEntity(orderdetailmodel.TableNameOrderDetail, err)
	}

	var food orderdetailmodel.FoodOrigin

	_ = json.Unmarshal([]byte(result.FoodOrigin), &food)

	result.Food = &food

	return result, nil
}
