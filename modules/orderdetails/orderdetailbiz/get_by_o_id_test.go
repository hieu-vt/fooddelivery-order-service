package orderdetailbiz

import (
	"context"
	"errors"
	"fooddelivery-order-service/modules/orderdetails/orderdetailmodel"
	"testing"
)

type testingGetItem struct {
	Input    int
	Expected error
	Actual   error
}

type mockGetOrderDetailStore struct {
}

func (m mockGetOrderDetailStore) FindByOId(ctx context.Context, orderId int, moreKeys ...string) (*orderdetailmodel.OrderDetail, error) {
	if orderId == 1 {
		return nil, errors.New("error db")
	}
	return nil, nil
}

func Test_GetOrderDetailByOId(t *testing.T) {
	biz := NewGetOrderDetailBiz(mockGetOrderDetailStore{})
	dataTable := []testingGetItem{
		{
			Input:    0,
			Expected: errors.New("order_id must be provide"),
			Actual:   nil,
		},
		{
			Input:    1,
			Expected: errors.New("error db"),
			Actual:   nil,
		},
		{
			Input:    2,
			Expected: nil,
			Actual:   nil,
		},
	}

	for _, item := range dataTable {
		_, actual := biz.GetOrderDetailByOId(context.Background(), item.Input)

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
