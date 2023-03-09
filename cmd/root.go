package cmd

import (
	"fmt"
	"fooddelivery-order-service/cmd/handlers"
	"fooddelivery-order-service/common"
	"fooddelivery-order-service/middleware"
	appgrpc "fooddelivery-order-service/plugin/grpc"
	"fooddelivery-order-service/plugin/pubsub/nats"
	sckio "fooddelivery-order-service/plugin/sckio"
	"fooddelivery-order-service/plugin/sdkgorm"
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
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
			engine.Use(middleware.Recover())
			//sIo := service.MustGet(common.PluginSocket).(sckio.AppSocket)
			handlers.MainRoute(engine, service)
		})

		if err := service.Start(); err != nil {
			serviceLogger.Fatalln(err)
		}
	},
}

func Execute() {
	rootCmd.AddCommand(outEnvCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)

		os.Exit(1)
	}
}
