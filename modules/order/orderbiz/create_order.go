package orderbiz

import (
	"context"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/modules/order/ordermodel"
	"fooddelivery-order-service/plugin/pubsub"
)

type OrderStore interface {
	Create(ctx context.Context, order *ordermodel.CreateOrder) error
}

type createOrderBiz struct {
	orderStore OrderStore
	pubsub     pubsub.NatsPubSub
}

func NewCreateOrderBiz(orderStore OrderStore, pubsub pubsub.NatsPubSub) *createOrderBiz {
	return &createOrderBiz{orderStore: orderStore, pubsub: pubsub}
}

func (biz *createOrderBiz) CreateOrder(ctx context.Context, order *ordermodel.CreateOrder) error {
	if err := order.ValidateOrderData(); err != nil {
		return common.ErrNoPermission(err)
	}

	order.Status = 1

	if err := biz.orderStore.Create(ctx, order); err != nil {
		return common.ErrCannotCreateEntity(ordermodel.TableOrderName, err)
	}

	biz.pubsub.Publish(ctx, common.TopicUserCreateOrder, pubsub.NewMessage(map[string]interface{}{
		"user_id":     order.UserId,
		"order_id":    order.Id,
		"food_origin": order.FoodOrigin,
		"price":       order.Price,
		"quantity":    order.Quantity,
		"discount":    order.Discount,
		"lat":         order.Lat,
		"lng":         order.Lng,
	}))

	return nil
}
