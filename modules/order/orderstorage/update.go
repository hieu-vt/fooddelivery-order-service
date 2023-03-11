package orderstorage

import (
	"context"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/modules/order/ordermodel"
)

func (s *sqlStore) Update(ctx context.Context, orderId int, orderUpdate *ordermodel.UpdateOrder) error {
	if err := s.db.Where("id = ?", orderId).Updates(orderUpdate).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
