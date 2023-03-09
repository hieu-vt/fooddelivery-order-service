package ordertrackingstorage

import (
	"context"
	"order-service/common"
	"order-service/modules/ordertracking/ordertrackingmodel"
)

func (s *sqlStore) Update(ctx context.Context, data *ordertrackingmodel.UpdateOrderTracking) error {
	if err := s.db.Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
