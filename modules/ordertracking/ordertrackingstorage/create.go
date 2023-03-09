package ordertrackingstorage

import (
	"context"
	"order-service/common"
	"order-service/modules/ordertracking/ordertrackingmodel"
)

func (s *sqlStore) Create(ctx context.Context, data *ordertrackingmodel.OrderTracking) error {
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
