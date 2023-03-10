package ordertrackingstorage

import (
	"context"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/modules/ordertracking/ordertrackingmodel"
)

func (s *sqlStore) Create(ctx context.Context, data *ordertrackingmodel.OrderTracking) error {
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
