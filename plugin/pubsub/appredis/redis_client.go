package appredis

import "github.com/go-redis/redis/v8"

type GetRedisClient interface {
	GetClient() *redis.Client
}
