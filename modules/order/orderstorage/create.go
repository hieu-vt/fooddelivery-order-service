package orderstorage

import (
	"context"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/modules/order/ordermodel"
)

func (s *sqlStore) Create(ctx context.Context, order *ordermodel.CreateOrder) error {
	if err := s.db.Create(order).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
