package ordermodel

import (
	"errors"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/modules/orderdetails/orderdetailmodel"
)

type Order struct {
	common.SqlModel `json:",inline"`

	UserId         int     `json:"-" gorm:"user_id"`
	ShipperId      int     `json:"-" gorm:"shipper_id"`
	TotalPrice     float64 `json:"totalPrice" gorm:"column:total_price"`
	NameRestaurant string  `json:"name" gorm:"column:name;"`
}

const (
	TableOrderName        = "orders"
	PriceMustMoreThanZero = "Total price must more than 0"
)

func (Order) TableName() string {
	return TableOrderName
}

type CreateOrder struct {
	common.SqlModel `json:",inline"`
	UserId          int                          `json:"userId" gorm:"user_id"`
	ShipperId       int                          `json:"shipperId" gorm:"shipper_id"`
	TotalPrice      float64                      `json:"totalPrice" gorm:"column:total_price"`
	FoodOriginBody  *orderdetailmodel.FoodOrigin `json:"foodOriginBody" gorm:"-"`
	FoodOrigin      string                       `json:"foodOrigin" gorm:"-"`
	Price           float32                      `json:"price" gorm:"-"`
	Quantity        int                          `json:"quantity" gorm:"-"`
	Discount        float32                      `json:"discount" gorm:"-"`
	Lat             float64                      `json:"lat" gorm:"-"`
	Lng             float64                      `json:"lng" gorm:"-"`
}

type UpdateOrder struct {
	ShipperId int `json:"shipperId" gorm:"shipper_id"`
}

func (UpdateOrder) TableName() string {
	return Order{}.TableName()
}

func (*CreateOrder) TableName() string {
	return Order{}.TableName()
}

func (order *Order) Mask(isAdminOrOwner bool) {
	order.GenUID(common.DbTypeFood)
}

func (order *CreateOrder) Mask(isAdminOrOwner bool) {
	order.GenUID(common.DbTypeFood)
}

func (order *CreateOrder) GetTotalPrice() float64 {
	return order.TotalPrice
}

func (res *CreateOrder) ValidateOrderData() error {

	if res.TotalPrice <= 0 {
		return errors.New(PriceMustMoreThanZero)
	}

	return nil
}
