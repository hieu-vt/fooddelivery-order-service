package orderdetailbiz

import (
	"context"
	"errors"
	"order-service/common"
	"order-service/modules/order/ordermodel"
	"order-service/modules/orderdetails/orderdetailmodel"
	"testing"
)

type mockOrderDetailStore struct {
}

type mockOrderStore struct {
}

func (s *mockOrderStore) FindByCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*ordermodel.Order, error) {
	var result *ordermodel.Order
	var error error

	result = &ordermodel.Order{
		SqlModel: common.SqlModel{
			Status: 1,
		},
		UserId:     0,
		ShipperId:  0,
		TotalPrice: 0,
	}

	for i := range condition {
		if condition[i] == 3 {
			result = nil
			error = errors.New("something wrong with db")
		}

		if condition[i] == 4 {
			result = &ordermodel.Order{
				SqlModel: common.SqlModel{
					Status: 0,
				},
				UserId:     0,
				ShipperId:  0,
				TotalPrice: 0,
			}
		}
	}

	return result, error
}

func (s *mockOrderDetailStore) Create(ctx context.Context, orderDetail *orderdetailmodel.OrderDetail) error {
	if orderDetail.OrderId == 2 {
		return errors.New("something wrong with db")
	}
	return nil
}

type testingItem struct {
	Input    orderdetailmodel.OrderDetail
	Expected error
	Actual   error
}

func TestOrderDetailBiz_CreateOrderDetail(t *testing.T) {
	store := &mockOrderDetailStore{}
	orderStore := &mockOrderStore{}
	biz := NewOrderDetailBiz(store, orderStore)

	dataTable := []testingItem{
		{
			Input: orderdetailmodel.OrderDetail{
				OrderId:    0,
				FoodOrigin: "12",
				Price:      12,
			},
			Expected: errors.New(orderdetailmodel.OrderIdINotBeEmpty),
			Actual:   nil,
		},
		{
			Input: orderdetailmodel.OrderDetail{
				OrderId:    1,
				FoodOrigin: "",
				Price:      12,
			},
			Expected: errors.New(orderdetailmodel.FoodOriginIsNotEmpty),
			Actual:   nil,
		},
		{
			Input: orderdetailmodel.OrderDetail{
				OrderId:    1,
				FoodOrigin: "222",
				Price:      0,
			},
			Expected: errors.New(orderdetailmodel.PriceMustMoreThan0),
			Actual:   nil,
		},
		{
			Input: orderdetailmodel.OrderDetail{
				OrderId:    1,
				FoodOrigin: "222",
				Price:      12,
			},
			Expected: nil,
			Actual:   nil,
		},
		{
			Input: orderdetailmodel.OrderDetail{
				OrderId:    2,
				FoodOrigin: "222",
				Price:      12,
			},
			Expected: errors.New("something wrong with db"),
			Actual:   nil,
		},
		{
			Input: orderdetailmodel.OrderDetail{
				OrderId:    3,
				FoodOrigin: "222",
				Price:      12,
			},
			Expected: errors.New("something wrong with db"),
			Actual:   nil,
		},
		{
			Input: orderdetailmodel.OrderDetail{
				OrderId:    4,
				FoodOrigin: "222",
				Price:      12,
			},
			Expected: common.ErrEntityNotFound(ordermodel.TableOrderName, nil),
			Actual:   nil,
		},
	}

	for _, item := range dataTable {
		actual := biz.CreateOrderDetail(context.Background(), &item.Input)

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
