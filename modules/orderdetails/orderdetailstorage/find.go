package orderdetailstorage

import (
	"context"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/modules/orderdetails/orderdetailmodel"
)

func (s *sqlStore) FindByIds(ctx context.Context, ids []int) ([]orderdetailmodel.OrderDetail, error) {
	var result []orderdetailmodel.OrderDetail
	if err := s.db.Where("order_id in (?)", ids).Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
