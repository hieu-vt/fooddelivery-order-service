package orderbiz

import (
	"context"
	"errors"
	"order-service/modules/order/ordermodel"
	"testing"
)

type MockStoreCreateOrder struct {
}

func (s *MockStoreCreateOrder) Create(ctx context.Context, order *ordermodel.CreateOrder) error {
	if order.TotalPrice <= 0 {
		return errors.New(ordermodel.PriceMustMoreThanZero)
	}

	if order.UserId == -2 {
		return errors.New("that have some wrong of the database")
	}

	return nil
}

type testingItem struct {
	Input    ordermodel.CreateOrder
	Expected error
	Actual   error
}

func TestCreateOrderBiz(t *testing.T) {
	store := &MockStoreCreateOrder{}
	// imp biz
	biz := NewCreateOrderBiz(store)

	dataTable := []testingItem{
		{
			Input: ordermodel.CreateOrder{
				UserId:     1,
				TotalPrice: 120,
			},
			Expected: nil,
			Actual:   nil,
		},
		{
			Input: ordermodel.CreateOrder{
				UserId:     -1,
				TotalPrice: 0,
			},
			Expected: errors.New(ordermodel.PriceMustMoreThanZero),
			Actual:   nil,
		},
		{
			Input: ordermodel.CreateOrder{
				UserId:     -2,
				TotalPrice: 12,
			},
			Expected: errors.New("that have some wrong of the database"),
			Actual:   nil,
		},
	}

	for _, item := range dataTable {
		actual := biz.CreateOrder(context.Background(), &item.Input)

		if actual == nil {
			if item.Expected != nil {
				t.Errorf("expect error is %s but actual is %v", item.Expected.Error(), actual.Error())
			}

			continue
		}

		if actual.Error() != item.Expected.Error() {
			t.Errorf("expect error is %s but actual is %v", item.Expected.Error(), actual.Error())
		}
	}
}
