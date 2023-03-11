package appredis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type GeoProvider interface {
	AddDriverLocation(ctx context.Context, key string, lng, lat float64, id string)
	RemoveDriverLocation(ctx context.Context, key string, id string)
	SearchDrivers(ctx context.Context, key string, limit int, lat float64, lng float64, r float64) *redis.GeoLocation
}
