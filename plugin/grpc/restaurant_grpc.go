package appgrpc

import (
	"context"
	"encoding/json"
	"flag"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/proto/restaurant"
	"github.com/200Lab-Education/go-sdk/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type restaurantClient struct {
	prefix      string
	url         string
	gwSupported bool
	gwPort      int
	client      restaurant.RestaurantServiceClient
}

func NewRestaurantClient(prefix string) *restaurantClient {
	return &restaurantClient{
		prefix: prefix,
	}
}

func (uc *restaurantClient) GetPrefix() string {
	return uc.prefix
}

func (uc *restaurantClient) Get() interface{} {
	return uc
}

func (uc *restaurantClient) Name() string {
	return uc.prefix
}

func (uc *restaurantClient) InitFlags() {
	flag.StringVar(&uc.url, uc.GetPrefix()+"-url", "localhost:50052", "URL connect to grpc server")
}

func (uc *restaurantClient) Configure() error {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())

	cc, err := grpc.Dial(uc.url, opts)

	if err != nil {
		return err
	}

	uc.client = restaurant.NewRestaurantServiceClient(cc)

	res, err1 := uc.client.GetRestaurantByIds(context.Background(), &restaurant.RestaurantRequest{RestaurantIds: []int32{1}})
	log.Println(err1, res)
	return nil
}

func (uc *restaurantClient) Run() error {
	return uc.Configure()
}

func (uc *restaurantClient) Stop() <-chan bool {
	c := make(chan bool)

	go func() {
		c <- true
	}()
	return c
}

func (uc *restaurantClient) GetRestaurants(ctx context.Context, ids []int) ([]common.Restaurant, error) {
	logger.GetCurrent().GetLogger(uc.prefix).Infoln("GetRestaurantByIds grpc store running")
	newIds := make([]int32, len(ids))
	for i := range ids {
		newIds[i] = int32(ids[i])
	}

	res, err := uc.client.GetRestaurantByIds(ctx, &restaurant.RestaurantRequest{RestaurantIds: newIds})

	if err != nil {
		return nil, common.ErrCannotGetEntity("restaurant", err)
	}

	result := make([]common.Restaurant, len(res.Restaurants))

	for i := range res.Restaurants {
		var logo common.Image
		_ = json.Unmarshal([]byte(res.Restaurants[i].Logo), &logo)

		result[i] = common.Restaurant{
			Name:      res.Restaurants[i].Name,
			Addr:      res.Restaurants[i].Addr,
			Logo:      &logo,
			Cover:     res.Restaurants[i].Cover,
			LikeCount: int(res.Restaurants[i].LikeCount),
			Owner: common.SimpleUser{
				SqlModel: common.SqlModel{
					Id: int(res.Restaurants[i].Owner.Id),
				},
				LastName:  res.Restaurants[i].Owner.LastName,
				FirstName: res.Restaurants[i].Owner.FirstName,
				Role:      res.Restaurants[i].Owner.Role,
			},
		}
	}

	return result, nil
}
