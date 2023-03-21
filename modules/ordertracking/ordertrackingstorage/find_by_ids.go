package ordertrackingstorage

import (
	"context"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/modules/ordertracking/ordertrackingmodel"
)

func (s *sqlStore) FindByIds(ctx context.Context, ids []int, orderDetail *ordertrackingmodel.OrderTracking) error {

	if err := s.db.Where("order_id in (?)", ids).Find(orderDetail).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
