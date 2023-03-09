package orderstorage

import (
	"context"
	"order-service/common"
	"order-service/modules/order/ordermodel"
)

func (s *sqlStore) Create(ctx context.Context, order *ordermodel.CreateOrder) error {
	if err := s.db.Create(order).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
