package orderdetailmodel

import (
	"errors"
	"order-service/common"
	"strings"
)

const (
	TableNameOrderDetail      = "order_details"
	FoodOriginIsNotEmpty      = "foodOrigin is not empty"
	PriceMustMoreThan0        = "price must more than 0"
	OrderIdINotBeEmpty        = "order id is not be empty"
	CannotMarshalFoodOrigin   = "cannot marshal food origin"
	CannotUnMarshalFoodOrigin = "cannot unmarshal food origin"
)

type FoodOrigin struct {
	Id           string  `json:"id" gorm:"column:id;"`
	RestaurantId int     `json:"restaurantId" gorm:"column:restaurant_id;"`
	CategoryId   string  `json:"categoryId" gorm:"column:category_id;"`
	Name         string  `json:"name" gorm:"column:name;"`
	Description  string  `json:"description" gorm:"column:description;"`
	price        float32 `json:"price" gorm:"column:price;"`
	total        int     `json:"total" gorm:"-"`
}

//
//func (fdOrigin FoodOrigin) Marshal() (error, string) {
//	 jsonFdOrigin, err := json.Marshal(fdOrigin)
//
//	if err != nil {
//		return errors.New(CannotMarshalFoodOrigin), ""
//	}
//
//	return nil, string(jsonFdOrigin)
//}

//func (fdOrigin FoodOrigin) UnMarshal() (error, FoodOrigin) {
//	jsonFdOrigin, err := json.Unmarshal(fdOrigin)
//
//	if err != nil {
//		return errors.New(CannotMarshalFoodOrigin), nil
//	}
//
//	return nil, string(jsonFdOrigin)
//}

type OrderDetail struct {
	common.SqlModel `json:",inline"`
	OrderId         int     `json:"orderId" gorm:"column:order_id"`
	FoodOrigin      string  `json:"foodOrigin" gorm:"column:food_origin"`
	Price           float32 `json:"price" gorm:"column:price"`
	Quantity        int     `json:"quantity" gorm:"column:quantity"`
	Discount        float32 `json:"discount" gorm:"column:quantity"`
}

func (OrderDetail) TableName() string {
	return TableNameOrderDetail
}

type CreateOrderDetail struct {
	OrderId    string      `json:"orderId" gorm:"column:order_id;"`
	FoodOrigin *FoodOrigin `json:"foodOrigin" gorm:"column:food_origin"`
	Price      float32     `json:"price" gorm:"column:price"`
	Quantity   int         `json:"quantity" gorm:"column:quantity"`
	Discount   float32     `json:"discount" gorm:"column:quantity"`
}

func (CreateOrderDetail) TableName() string {
	return OrderDetail{}.TableName()
}

func (orderDetail *OrderDetail) ValidateOrderDetailData() error {
	orderDetail.FoodOrigin = strings.TrimSpace(orderDetail.FoodOrigin)

	if orderDetail.OrderId <= 0 {
		return errors.New(OrderIdINotBeEmpty)
	}

	if orderDetail.FoodOrigin == "" {
		return errors.New(FoodOriginIsNotEmpty)
	}

	if orderDetail.Price <= 0 {
		return errors.New(PriceMustMoreThan0)
	}

	return nil
}

func (orderDetail *OrderDetail) Mask() {
	orderDetail.GenUID(common.DbTypeOrder)
}
