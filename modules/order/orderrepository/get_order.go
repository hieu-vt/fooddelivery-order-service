package orderrepository

import (
	"context"
	"encoding/json"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/modules/order/ordermodel"
	"fooddelivery-order-service/modules/orderdetails/orderdetailmodel"
	"fooddelivery-order-service/modules/ordertracking/ordertrackingmodel"
	"fooddelivery-order-service/proto/restaurant"
	"google.golang.org/grpc"
	"log"
)

type orderTrackingDetailMap struct {
	orderdetailmodel.OrderDetail
	ordertrackingmodel.OrderTracking
}

type orderStore interface {
	Find(ctx context.Context, userId int, paging common.Paging) ([]ordermodel.GetOrderType, error)
}

type orderDetailStore interface {
	FindByIds(ctx context.Context, ids []int) ([]orderdetailmodel.OrderDetail, error)
}

type orderTrackingStore interface {
	FindByIds(ctx context.Context, ids []int) ([]ordertrackingmodel.OrderTracking, error)
}

type restaurantService interface {
	GetRestaurantByIds(ctx context.Context, in *restaurant.RestaurantRequest, opts ...grpc.CallOption) (*restaurant.RestaurantResponse, error)
}

type getOrderRepository struct {
	store             orderStore
	detailStore       orderDetailStore
	trackingStore     orderTrackingStore
	restaurantService restaurantService
}

func NewGetOrderRepository(store orderStore, detailStore orderDetailStore, trackingStore orderTrackingStore, restaurantService restaurantService) *getOrderRepository {
	return &getOrderRepository{store: store, detailStore: detailStore, trackingStore: trackingStore, restaurantService: restaurantService}
}

func (repo *getOrderRepository) GetOrders(
	ctx context.Context,
	userId int,
	paging common.Paging,
) ([]ordermodel.GetOrderType, error) {
	orders, err := repo.store.Find(ctx, userId, paging)

	//resturantIds := make([]int, len(orders))

	if err != nil {
		return nil, common.ErrEntityNotFound(ordermodel.TableOrderName, err)
	}

	orderIds := make([]int, len(orders))

	for i, item := range orders {
		orderIds[i] = item.Id
	}

	orderDetail, errODetail := repo.detailStore.FindByIds(ctx, orderIds)

	if errODetail != nil {
		return nil, common.ErrEntityNotFound(orderdetailmodel.TableNameOrderDetail, errODetail)
	}

	orderTracking, errOTracking := repo.trackingStore.FindByIds(ctx, orderIds)

	if errOTracking != nil {
		return nil, common.ErrEntityNotFound(ordertrackingmodel.TableNameOrderTracking, errOTracking)
	}

	cacheTrackingDetail := make(map[int]orderTrackingDetailMap, len(orderDetail))

	for i, item := range orderDetail {
		cacheTrackingDetail[item.OrderId] = orderTrackingDetailMap{
			OrderDetail:   orderDetail[i],
			OrderTracking: orderTracking[i],
		}
	}

	for i, item := range orders {
		orders[i].State = cacheTrackingDetail[item.Id].State
		orders[i].FoodOrigin = cacheTrackingDetail[item.Id].FoodOrigin
	}

	for _, item := range orders {
		var food orderdetailmodel.FoodOrigin
		_ = json.Unmarshal([]byte(item.FoodOrigin), &food)
		log.Println(food.RestaurantId)
	}

	return orders, nil
}
