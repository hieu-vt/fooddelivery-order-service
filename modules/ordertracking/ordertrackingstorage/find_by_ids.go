package ordertrackingstorage

import (
	"context"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/modules/ordertracking/ordertrackingmodel"
)

func (s *sqlStore) FindByIds(ctx context.Context, ids []int) ([]ordertrackingmodel.OrderTracking, error) {
	var result []ordertrackingmodel.OrderTracking
	if err := s.db.Where("order_id in (?)", ids).Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
