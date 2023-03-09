package ordertrackingmodel

import (
	"errors"
	"order-service/common"
)

const (
	TableNameOrderTracking = "order_trackings"
	OrderIdIsNotEmpty      = "orderId is not empty"
	StateIsNotEmpty        = "state is not empty"
)

type OrderTracking struct {
	common.SqlModel `json:",inline"`
	OrderId         int                 `json:"orderId" gorm:"column:order_id"`
	State           common.TrackingType `json:"state" gorm:"column:state"`
}

type CreateOrderTracking struct {
	OrderId int                 `json:"orderId"`
	State   common.TrackingType `json:"state"`
}

type CreateOrderTrackingParams struct {
	OrderId string              `json:"orderId"`
	State   common.TrackingType `json:"state"`
}

type UpdateOrderTracking struct {
	OrderId int                 `json:"orderId" gorm:"column:order_id"`
	State   common.TrackingType `json:"state" gorm:"column:state"`
}

func (OrderTracking) TableName() string {
	return TableNameOrderTracking
}

func (UpdateOrderTracking) TableName() string {
	return TableNameOrderTracking
}

func (oTracking *OrderTracking) Validation() error {
	if oTracking.State == "" {
		return errors.New(StateIsNotEmpty)
	}

	if oTracking.OrderId <= 0 {
		return errors.New(OrderIdIsNotEmpty)
	}

	return nil
}

func (oTracking *OrderTracking) Mask() {
	oTracking.GenUID(common.DbTypeOrder)
}
