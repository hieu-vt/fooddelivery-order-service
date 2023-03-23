package common

const (
	DbTypeFood       = 2
	DbTypeCategory   = 3
	DbTypeUser       = 4
	DbTypeOrder      = 5
	DbTypeCart       = 6
	DbTypeFoodRating = 7
)

const (
	CurrentUser = "user"

	DBMain                     = "mysql"
	PluginGrpcUserClient       = "Grpc_User_Client"
	PluginGrpcRestaurantClient = "Grpc_Restaurant_Client"
	PluginGrpcAuthClient       = "Grpc_auth_client"
	PluginNats                 = "nats"
	PluginRedis                = "redis"

	// Pubsub
	TopicUserCreateOrder     = "order.create"
	TopicUserUpdateOrder     = "order.update"
	TopicOrderTrackingUpdate = "ordertracking.update"
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}

type TrackingType string

const (
	WaitingForShipper TrackingType = "waiting_for_shipper"
)
