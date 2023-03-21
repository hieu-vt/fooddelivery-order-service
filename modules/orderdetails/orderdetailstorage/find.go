package orderdetailstorage

import (
	"context"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/modules/orderdetails/orderdetailmodel"
)

func (s *sqlStore) FindByIds(ctx context.Context, ids []int, orderDetail *orderdetailmodel.OrderDetail) error {

	if err := s.db.Where("order_id in (?)", ids).Find(orderDetail).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
