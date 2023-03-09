package orderdetailstorage

import (
	"context"
	"order-service/common"
	"order-service/modules/orderdetails/orderdetailmodel"
)

func (s *sqlStore) Create(ctx context.Context, orderDetail *orderdetailmodel.OrderDetail) error {
	if err := s.db.Create(orderDetail).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
