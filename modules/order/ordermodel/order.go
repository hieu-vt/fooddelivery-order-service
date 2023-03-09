package ordermodel

import (
	"errors"
	"order-service/common"
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
	UserId          int     `json:"userId" gorm:"user_id"`
	ShipperId       int     `json:"shipperId" gorm:"shipper_id"`
	TotalPrice      float64 `json:"totalPrice" gorm:"column:total_price"`
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
