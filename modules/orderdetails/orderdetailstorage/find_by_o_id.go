package orderdetailstorage

import (
	"context"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/modules/orderdetails/orderdetailmodel"
)

func (s *sqlStore) FindByOId(ctx context.Context, orderId int, moreKeys ...string) (*orderdetailmodel.OrderDetail, error) {
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	var orderDetail orderdetailmodel.OrderDetail

	if err := db.Where("status = 1").Where("order_id = ?", orderId).First(&orderDetail).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &orderDetail, nil
}
