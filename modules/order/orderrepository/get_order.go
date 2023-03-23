package orderrepository

import (
	"context"
	"encoding/json"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/modules/order/ordermodel"
	"fooddelivery-order-service/modules/orderdetails/orderdetailmodel"
	"fooddelivery-order-service/modules/ordertracking/ordertrackingmodel"
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
	GetRestaurants(ctx context.Context, ids []int) ([]common.Restaurant, error)
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

	cacheResId := make(map[int]int, len(orders))
	var resIds []int

	for i, item := range orders {
		var food orderdetailmodel.FoodOrigin
		_ = json.Unmarshal([]byte(cacheTrackingDetail[item.Id].FoodOrigin), &food)

		orders[i].State = cacheTrackingDetail[item.Id].State
		orders[i].FoodOrigin = &food

		if _, ok := cacheResId[food.RestaurantId]; !ok {
			resIds = append(resIds, food.RestaurantId)
			cacheResId[food.RestaurantId] = food.RestaurantId
			orders[i].RestaurantId = food.RestaurantId
		}
	}

	restaurants, errRestaurant := repo.restaurantService.GetRestaurants(ctx, resIds)

	if errRestaurant != nil {
		log.Println("errRestaurant ", errRestaurant)
	}

	cacheRestaurant := make(map[int]common.Restaurant, len(restaurants))

	for i, item := range restaurants {
		cacheRestaurant[item.Owner.Id] = restaurants[i]
	}

	for i, item := range orders {
		if restaurant, ok := cacheRestaurant[item.RestaurantId]; ok {
			orders[i].Name = restaurant.Name
			orders[i].Logo = restaurant.Logo
		}
	}

	return orders, nil
}
