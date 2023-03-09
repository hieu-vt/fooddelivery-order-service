package ordermodel

import (
	"order-service/common"
)

type GetOrderType struct {
	common.SqlModel `json:",inline"`
	TotalPrice      int                 `json:"totalPrice" gorm:"column:total_price"`
	State           common.TrackingType `json:"state" gorm:"column:state"`
	Name            string              `json:"restaurantName" gorm:"column:name;"`
	FoodOrigin      string              `json:"foodOrigin" gorm:"column:food_origin"`
	Logo            *common.Image       `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images      `json:"cover" gorm:"column:cover;"`
}

func (gOrderType *GetOrderType) Mask() {
	gOrderType.GenUID(common.DbTypeOrder)
}
