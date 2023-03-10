package ordertrackingstorage

import (
	"context"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/modules/ordertracking/ordertrackingmodel"
)

func (s *sqlStore) Update(ctx context.Context, data *ordertrackingmodel.UpdateOrderTracking) error {
	if err := s.db.Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
