package ordertrackingbiz

import (
	"context"
	"errors"
	"order-service/common"
	"order-service/modules/ordertracking/ordertrackingmodel"
	"testing"
)

type mockOrderTrackingStore struct{}

func (s *mockOrderTrackingStore) Create(ctx context.Context, data *ordertrackingmodel.OrderTracking) error {
	if data.OrderId == 3 {
		return errors.New("something wrong with DB")
	}
	return nil
}

type testingItem struct {
	Input    ordertrackingmodel.OrderTracking
	Expected error
	Actual   error
}

func TestCreateOrderTracking(t *testing.T) {
	store := &mockOrderTrackingStore{}
	biz := NewOrderTrackingBiz(store)

	dataTable := []testingItem{
		{
			Input: ordertrackingmodel.OrderTracking{
				OrderId: 0,
				State:   common.WaitingForShipper,
			},
			Expected: errors.New(ordertrackingmodel.OrderIdIsNotEmpty),
			Actual:   nil,
		},
		{
			Input: ordertrackingmodel.OrderTracking{
				OrderId: 1,
				State:   "",
			},
			Expected: errors.New(ordertrackingmodel.StateIsNotEmpty),
			Actual:   nil,
		},
		{
			Input: ordertrackingmodel.OrderTracking{
				OrderId: 2,
				State:   common.WaitingForShipper,
			},
			Expected: nil,
			Actual:   nil,
		},
		{
			Input: ordertrackingmodel.OrderTracking{
				OrderId: 3,
				State:   common.WaitingForShipper,
			},
			Expected: errors.New("something wrong with DB"),
			Actual:   nil,
		},
	}

	for _, item := range dataTable {
		actual := biz.CreateOrderTracking(context.Background(), &item.Input)

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
