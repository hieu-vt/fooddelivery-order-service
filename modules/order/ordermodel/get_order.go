package ordermodel

import (
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/modules/orderdetails/orderdetailmodel"
)

type GetOrderType struct {
	common.SqlModel `json:",inline"`
	TotalPrice      int                          `json:"totalPrice" gorm:"column:total_price"`
	State           common.TrackingType          `json:"state" gorm:"column:state"`
	Name            string                       `json:"restaurantName" gorm:"column:name;"`
	FoodOrigin      *orderdetailmodel.FoodOrigin `json:"foodOrigin" gorm:"-"`
	Logo            *common.Image                `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images               `json:"cover" gorm:"column:cover;"`
	RestaurantId    int                          `json:"-"`
}

func (gOrderType *GetOrderType) Mask() {
	gOrderType.GenUID(common.DbTypeOrder)
}
