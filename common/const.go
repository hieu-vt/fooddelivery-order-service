package common

const (
	DbTypeRestaurant = 1
	DbTypeFood       = 2
	DbTypeCategory   = 3
	DbTypeUser       = 4
	DbTypeOrder      = 5
	DbTypeCart       = 6
	DbTypeFoodRating = 7
)

const (
	CurrentUser = "user"

	DBMain               = "mysql"
	PluginGrpcUserClient = "Grpc_User_Client"
	PluginGrpcAuthClient = "Grpc_auth_client"
	PluginNats           = "nats"
	PluginSocket         = "skio"
	PluginSocketEngine   = "skio_engine"
	PluginNSocket        = "n_skio"
	PluginNSocketEngine  = "n_skio_engine"
	PluginRedis          = "redis"

	// Pubsub
	TopicUserCreateOrder = "order.create"
	TopicUserUpdateOrder = "order.update"
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}

const (
	EmitUserCreateOrderSuccess = "EmitUserCreateOrderSuccess"
	EmitUserOrderFailure       = "EmitUserOrderFailure"
	EmitAuthenticated          = "EmitAuthenticated"

	// Redis key
	RedisShipperLocation = "shipper_locations"
	RedisUserLocation    = "user_locations"
)

const (
	EvenAuthenticated       = "EvenAuthenticated"
	EvenUserCreateOrder     = "EvenUserCreateOrder"
	EventUserUpdateLocation = "EventUserUpdateLocation"
)

const (
	Authenticated        = "Authenticated"
	AuthenticationFailed = "authentication_failed"

	UserUpdateLocation = "UserUpdateLocation"

	OrderTracking = "OrderTracking"
	NewOrder      = "NewOrder"
)

type TrackingType string

const (
	WaitingForShipper TrackingType = "waiting_for_shipper"
	Preparing         TrackingType = "preparing"
	Cancel            TrackingType = "cancel"
	OnTheWay          TrackingType = "on_the_way"
	Delivered         TrackingType = "delivered"
)

const (
	TRACE_SERVICE = "trace-demo"
	ENVIRONMENT   = "dev"
)

type LocationData struct {
	Lat    float64 `json:"lat"`
	Lng    float64 `json:"lng"`
	UserId string  `json:"userId"`
	Role   string  `json:"role"`
}
