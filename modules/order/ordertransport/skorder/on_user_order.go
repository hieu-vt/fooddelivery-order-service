package skorder

import (
	"context"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/plugin/pubsub"
	goservice "github.com/200Lab-Education/go-sdk"
	socketio "github.com/googollee/go-socket.io"
	"log"
)

type EvenOrderMessageData struct {
	UserId    string              `json:"userId"`
	OrderId   string              `json:"orderId"`
	ShipperId string              `json:"shipperId"`
	Type      common.TrackingType `json:"type"`
}

func OnOrderTracking(sc goservice.ServiceContext, server *socketio.Server) func(s socketio.Conn, data EvenOrderMessageData) {
	return func(s socketio.Conn, data EvenOrderMessageData) {
		// Đoạn này phần shipper call socket chỉ handle test realtime
		// Thực chất khi shipper nhận socket order start
		// Shipper accept/reject package --> call request --> create pubsub để emit to room là tạo đã accept hoặc reject
		// Từ đó khi shipper cứ update trạng thái vào database thì nó sẽ emit to room cái trạng thái cho user
		// Khi nàp successfully thì update lại ở database và clear process

		// Create Pubsub update order tracking to database
		pb := sc.MustGet(common.PluginNats).(pubsub.NatsPubSub)
		roomKey := data.OrderId

		log.Println("order ", roomKey)
		//if data.Type == "" {
		//	log.Println("Tracking order", data.Type)
		//	// handle join shipper to room and update tracking type
		//	server.BroadcastToRoom("/", roomKey, common.OrderTracking, EvenOrderMessageData{
		//		UserId:  data.UserId,
		//		OrderId: data.OrderId,
		//		Type:    common.Preparing,
		//	})
		//}

		s.Join(roomKey)

		if data.Type == common.WaitingForShipper {
			log.Println("Tracking order", data.Type)
			// handle join shipper to room and update tracking type
			pb.Publish(context.Background(), common.TopicUserUpdateOrder, pubsub.NewMessage(map[string]interface{}{
				"shipper_id": data.ShipperId,
				"order_id":   data.OrderId,
			}))

			server.BroadcastToRoom("/", roomKey, common.OrderTracking, EvenOrderMessageData{
				UserId:  data.UserId,
				OrderId: data.OrderId,
				Type:    common.Preparing,
			})
		}

		if data.Type == common.Preparing {
			log.Println("Tracking order", data.Type)
			// handle find another shipper
			server.BroadcastToRoom("/", roomKey, common.OrderTracking, EvenOrderMessageData{
				OrderId: data.OrderId,
				UserId:  data.UserId,
				Type:    common.Cancel,
			})

			s.Leave(roomKey)
		}

		if data.Type == common.Cancel {
			log.Println("Tracking order", data.Type)
			// handle find another shipper
			server.BroadcastToRoom("/", roomKey, common.OrderTracking, EvenOrderMessageData{
				OrderId: data.OrderId,
				UserId:  data.UserId,
				Type:    common.Cancel,
			})

			s.Leave(roomKey)
		}

		if data.Type == common.Delivered {
			log.Println("Tracking order", data.Type)
			// handle update database
			// handle clear rooms
			server.BroadcastToRoom("/", roomKey, common.OrderTracking, EvenOrderMessageData{
				OrderId: data.OrderId,
				UserId:  data.UserId,
				Type:    common.Delivered,
			})

			s.Leave(roomKey)
		}
	}
}
