package ordertrackingbiz

import (
	"context"
	"errors"
	"order-service/common"
	"order-service/modules/ordertracking/ordertrackingmodel"
	"testing"
)

type mockUpdateOrderTrackingStore struct{}

func (s *mockUpdateOrderTrackingStore) Update(ctx context.Context, data *ordertrackingmodel.UpdateOrderTracking) error {
	if data.OrderId == 2 {
		return errors.New("something wrong with DB")
	}
	return nil
}

type testingUpdateItem struct {
	Input    ordertrackingmodel.UpdateOrderTracking
	Expected error
	Actual   error
}

func TestUpdateOrderTracking(t *testing.T) {
	store := &mockUpdateOrderTrackingStore{}
	biz := NewUpdateOrderBiz(store)

	dataTable := []testingUpdateItem{
		{
			Input: ordertrackingmodel.UpdateOrderTracking{
				OrderId: 1,
				State:   common.WaitingForShipper,
			},
			Expected: nil,
			Actual:   nil,
		},
		{
			Input: ordertrackingmodel.UpdateOrderTracking{
				OrderId: 2,
				State:   common.WaitingForShipper,
			},
			Expected: errors.New("something wrong with DB"),
			Actual:   nil,
		},
	}

	for _, item := range dataTable {
		actual := biz.UpdateOrderTracking(context.Background(), &item.Input)

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
