package orderdetailstorage

import (
	"context"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/modules/orderdetails/orderdetailmodel"
)

func (s *sqlStore) Create(ctx context.Context, orderDetail orderdetailmodel.OrderDetail) error {
	if err := s.db.Create(&orderDetail).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
