package orderstorage

import (
	"context"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/modules/order/ordermodel"
	"gorm.io/gorm"
)

func (s *sqlStore) FindByCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*ordermodel.Order, error) {
	db := s.db

	var data ordermodel.Order

	for i := range moreKeys {
		db.Preload(moreKeys[i])
	}

	if err := db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, err
	}

	return &data, nil
}
