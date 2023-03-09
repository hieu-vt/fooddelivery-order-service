package skorder

//type DataOrder struct {
//	OrderId *common.UID `json:"orderId"`
//}
//
//type RealtimeEngine interface {
//	EmitToRoom(room string, key string, data interface{}) error
//	GetShipper(reddit reddit.RedditEngine, id int, location interface{}) int
//}
//
//type TopicEmitEvenOrderMessageData struct {
//	ordermodel.CreateOrder `json:",inline"`
//	OrderId                string              `json:"orderId"`
//	Type                   common.TrackingType `json:"type"`
//}
//
//func OnUserOrder(appCtx component.AppContext, requester common.Requester, rtEngine RealtimeEngine) func(s socketio.Conn, data DataOrder) {
//	return func(s socketio.Conn, data DataOrder) {
//		pubsub := appCtx.GetPubsub()
//		reddit := appCtx.GetReddit()
//		userId := fmt.Sprintf("%d", requester.GetUserId())
//
//		shipperId := rtEngine.GetShipper(reddit, requester.GetUserId(), reddit.Get(userId))
//
//		pubsub.Publish(context.Background(), common.TopicHandleOrderWhenUserOrderFood, pubsub2.NewMessage(ordermodel.CreateOrder{
//			SqlModel: common.SqlModel{
//				FakeId: data.OrderId,
//			},
//			UserId:    requester.GetUserId(),
//			ShipperId: shipperId,
//		}))
//	}
//}
//
//func OnOrderTracking(appCtx component.AppContext, requester common.Requester, rtEngine RealtimeEngine) func(s socketio.Conn, data TopicEmitEvenOrderMessageData) {
//	return func(s socketio.Conn, data TopicEmitEvenOrderMessageData) {
//		// Đoạn này phần shipper call socket chỉ handle test realtime
//		// Thực chất khi shipper nhận socket order start
//		// Shipper accept/reject package --> call request --> create pubsub để emit to room là tạo đã accept hoặc reject
//		// Từ đó khi shipper cứ update trạng thái vào database thì nó sẽ emit to room cái trạng thái cho user
//		// Khi nàp successfully thì update lại ở database và clear process
//
//		// Create Pubsub update order tracking to database
//
//		roomKey := common.OrderTracking + data.OrderId
//		if data.Type == common.WaitingForShipper {
//			log.Println("Tracking order", data.Type)
//			// handle join shipper to room and update tracking type
//			s.Join(roomKey)
//			rtEngine.EmitToRoom(roomKey, common.OrderTracking, TopicEmitEvenOrderMessageData{
//				CreateOrder: ordermodel.CreateOrder{
//					ShipperId: data.ShipperId,
//					UserId:    data.UserId,
//				},
//				OrderId: data.OrderId,
//				Type:    common.Preparing,
//			})
//		}
//
//		if data.Type == common.Cancel {
//			log.Println("Tracking order", data.Type)
//			// handle find another shipper
//			rtEngine.EmitToRoom(roomKey, common.OrderTracking, TopicEmitEvenOrderMessageData{
//				OrderId: data.OrderId,
//				CreateOrder: ordermodel.CreateOrder{
//					ShipperId: data.ShipperId,
//					UserId:    data.UserId,
//				},
//				Type: common.Cancel,
//			})
//
//			s.Leave(roomKey)
//		}
//
//		if data.Type == common.Delivered {
//			log.Println("Tracking order", data.Type)
//			// handle update database
//			// handle clear rooms
//			rtEngine.EmitToRoom(roomKey, common.OrderTracking, TopicEmitEvenOrderMessageData{
//				OrderId: data.OrderId,
//				CreateOrder: ordermodel.CreateOrder{
//					ShipperId: data.ShipperId,
//					UserId:    data.UserId,
//				},
//				Type: common.Delivered,
//			})
//
//			s.Leave(roomKey)
//		}
//	}
//}
