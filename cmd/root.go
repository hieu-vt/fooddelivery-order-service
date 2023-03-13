package cmd

import (
	"context"
	"fmt"
	"fooddelivery-order-service/cmd/handlers"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/component/asyncjob"
	"fooddelivery-order-service/middleware"
	"fooddelivery-order-service/modules/order/orderbiz"
	"fooddelivery-order-service/modules/order/ordermodel"
	"fooddelivery-order-service/modules/order/orderstorage"
	"fooddelivery-order-service/plugin/appredis"
	appgrpc "fooddelivery-order-service/plugin/grpc"
	"fooddelivery-order-service/plugin/pubsub"
	"fooddelivery-order-service/plugin/pubsub/nats"
	sckio "fooddelivery-order-service/plugin/sckio"
	"fooddelivery-order-service/plugin/sdkgorm"
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"log"
	"os"
)

func newService() goservice.Service {
	service := goservice.New(
		goservice.WithName("food-delivery-order"),
		goservice.WithVersion("1.0.0"),
		goservice.WithInitRunnable(sdkgorm.NewGormDB("main", common.DBMain)),
		goservice.WithInitRunnable(appgrpc.NewAuthClient(common.PluginGrpcAuthClient)),
		goservice.WithInitRunnable(appgrpc.NewUserClient(common.PluginGrpcUserClient)),
		goservice.WithInitRunnable(nats.NewNatsPubSub(common.PluginNats)),
		goservice.WithInitRunnable(sckio.NewSocketIo(common.PluginSocket)),
		goservice.WithInitRunnable(appredis.NewAppRedis("main-redis", common.PluginRedis)),
	)

	return service
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start an food delivery order service",
	Run: func(cmd *cobra.Command, args []string) {
		service := newService()
		serviceLogger := service.Logger("service")

		if err := service.Init(); err != nil {
			serviceLogger.Fatalln(err)
		}

		service.HTTPServer().AddHandler(func(engine *gin.Engine) {
			engine.Use(func(c *gin.Context) {
				c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
				c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
				c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

				if c.Request.Method == "OPTIONS" {
					c.AbortWithStatus(204)
					return
				}

				c.Next()
			})

			engine.Use(middleware.Recover())

			handlers.MainRoute(engine, service)
		})

		go func() {
			pb := service.MustGet(common.PluginNats).(pubsub.NatsPubSub)
			db := service.MustGet(common.DBMain).(*gorm.DB)

			ch, _ := pb.Subscribe(context.Background(), common.TopicUserUpdateOrder)

			for msg := range ch {
				job := asyncjob.NewJob(func(ctx context.Context) error {
					orderId := msg.Data()["order_id"].(string)
					shipperId := msg.Data()["shipper_id"].(string)

					store := orderstorage.NewSqlStore(db)
					biz := orderbiz.NewUpdateOrderBiz(store)

					uidShipper, _ := common.FromBase58(shipperId)
					uidOrder, _ := common.FromBase58(orderId)

					return biz.UpdateOrder(ctx, int(uidOrder.GetLocalID()), &ordermodel.UpdateOrder{
						ShipperId: int(uidShipper.GetLocalID()),
					})
				})

				if err := asyncjob.NewGroup(true, job).Run(context.Background()); err != nil {
					log.Println(err)
				}
			}

		}()

		if err := service.Start(); err != nil {
			serviceLogger.Fatalln(err)
		}
	},
}

func Execute() {
	rootCmd.AddCommand(outEnvCmd)
	rootCmd.AddCommand(StartCreateOrderDetailTrackingAfterCreateOrder)
	rootCmd.AddCommand(StartHandleSocketAfterCreateOrder)
	rootCmd.AddCommand(StartNHandleSocketAfterCreateOrder)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)

		os.Exit(1)
	}
}
