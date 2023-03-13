package skorder

import (
	"context"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/plugin/pubsub"
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/zishang520/socket.io/socket"
	"log"
)

type EvenOrderMessageData struct {
	UserId    string              `json:"userId"`
	OrderId   string              `json:"orderId"`
	ShipperId string              `json:"shipperId"`
	Type      common.TrackingType `json:"type"`
}

func OnOrderTracking(sc goservice.ServiceContext, client *socket.Socket) func(datas ...any) {
	return func(datas ...any) {
		// Đoạn này phần shipper call socket chỉ handle test realtime
		// Thực chất khi shipper nhận socket order start
		// Shipper accept/reject package --> call request --> create pubsub để emit to room là tạo đã accept hoặc reject
		// Từ đó khi shipper cứ update trạng thái vào database thì nó sẽ emit to room cái trạng thái cho user
		// Khi nàp successfully thì update lại ở database và clear process

		// Create Pubsub update order tracking to database
		data := datas[0].(map[string]interface{})
		dataOrder := EvenOrderMessageData{
			UserId:    data["userId"].(string),
			OrderId:   data["orderId"].(string),
			ShipperId: data["shipperId"].(string),
			Type:      common.TrackingType(data["type"].(string)),
		}

		pb := sc.MustGet(common.PluginNats).(pubsub.NatsPubSub)
		roomKey := dataOrder.OrderId

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

		client.Join(socket.Room(roomKey))

		if dataOrder.Type == common.WaitingForShipper {
			log.Println("Tracking order", dataOrder.Type)
			// handle join shipper to room and update tracking type
			pb.Publish(context.Background(), common.TopicUserUpdateOrder, pubsub.NewMessage(map[string]interface{}{
				"shipper_id": dataOrder.ShipperId,
				"order_id":   dataOrder.OrderId,
			}))

			client.To(socket.Room(roomKey)).Emit(common.OrderTracking, EvenOrderMessageData{
				UserId:  dataOrder.UserId,
				OrderId: dataOrder.OrderId,
				Type:    common.Preparing,
			})

		}

		if dataOrder.Type == common.Preparing {
			log.Println("Tracking order", dataOrder.Type)
			// handle find another shipper
			client.To(socket.Room(roomKey)).Emit(common.OrderTracking, EvenOrderMessageData{
				OrderId: dataOrder.OrderId,
				UserId:  dataOrder.UserId,
				Type:    common.Cancel,
			})
		}

		if dataOrder.Type == common.Cancel {
			log.Println("Tracking order", dataOrder.Type)
			// handle find another shipper
			client.To(socket.Room(roomKey)).Emit(common.OrderTracking, EvenOrderMessageData{
				OrderId: dataOrder.OrderId,
				UserId:  dataOrder.UserId,
				Type:    common.Cancel,
			})

			client.Leave(socket.Room(roomKey))
		}

		if dataOrder.Type == common.Delivered {
			log.Println("Tracking order", dataOrder.Type)
			// handle update database
			// handle clear rooms
			client.To(socket.Room(roomKey)).Emit(common.OrderTracking, EvenOrderMessageData{
				OrderId: dataOrder.OrderId,
				UserId:  dataOrder.UserId,
				Type:    common.Delivered,
			})

			client.Leave(socket.Room(roomKey))
		}
	}
}
