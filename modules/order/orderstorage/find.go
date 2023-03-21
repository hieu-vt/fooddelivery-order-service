package orderstorage

import (
	"context"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/modules/order/ordermodel"
)

func (s *sqlStore) Find(ctx context.Context, userId int, paging common.Paging) ([]ordermodel.GetOrderType, error) {
	db := s.db

	var orders []ordermodel.GetOrderType

	if paging.FakeCursor != "" {
		if uid, err := common.FromBase58(paging.FakeCursor); err == nil {
			db = db.Where("id < ?", uid.GetLocalID())
		} else {
			return nil, common.ErrDB(err)
		}
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(offset)
	}

	if err := db.Limit(paging.Limit).Table(ordermodel.TableOrderName).
		Find(&orders).
		Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return orders, nil
}
