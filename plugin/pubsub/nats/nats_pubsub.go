package nats

import (
	"context"
	"encoding/json"
	"flag"
	"fooddelivery-order-service/plugin/pubsub"
	"github.com/200Lab-Education/go-sdk/logger"
	"github.com/nats-io/nats.go"
	"time"
)

type natsPubSub struct {
	name       string
	logger     logger.Logger
	connection *nats.Conn
	url        string
}

func (natPS *natsPubSub) GetPrefix() string {
	return natPS.name
}

func (natPS *natsPubSub) Get() interface{} {
	return natPS
}

func NewNatsPubSub(name string) *natsPubSub {
	return &natsPubSub{
		name: name,
	}
}

func (natPS *natsPubSub) Name() string {
	return natPS.name
}

func (natPS *natsPubSub) InitFlags() {
	flag.StringVar(&natPS.url, natPS.name+"-url", nats.DefaultURL, "Url connect service nats pubsub")
}

func (natPS *natsPubSub) Configure() error {
	logger := logger.GetCurrent().GetLogger(natPS.name)
	natPS.logger = logger

	nc, _ := nats.Connect(natPS.url, natPS.setupConnOptions([]nats.Option{})...)

	natPS.connection = nc
	natPS.logger.Infoln("Connected to NATS service.")

	return nil
}

func (natPS *natsPubSub) Run() error {

	natPS.Configure()

	return nil
}

func (natPS *natsPubSub) Stop() <-chan bool {
	c := make(chan bool)
	go func() { c <- true }()
	return c
}

func (natPS *natsPubSub) Publish(ctx context.Context, channel string, data *pubsub.Message) error {
	dataMarshal, err := json.Marshal(data.Data())

	if err != nil {
		natPS.logger.Infoln(err)
	}

	err = natPS.connection.Publish(channel, dataMarshal)

	if err != nil {
		natPS.logger.Infoln(err)
	}

	return nil
}

func (natPS *natsPubSub) Subscribe(ctx context.Context, channel string) (ch <-chan *pubsub.Message, close func()) {
	chanMess := make(chan *pubsub.Message)

	sub, err := natPS.connection.Subscribe(channel, func(msg *nats.Msg) {
		msgData := make(map[string]interface{})
		_ = json.Unmarshal(msg.Data, &msgData)
		appMsg := pubsub.NewMessage(msgData)
		appMsg.SetChannel(channel)
		chanMess <- appMsg

	})

	if err != nil {
		natPS.logger.Infoln(err)
	}

	return chanMess, func() {
		sub.Unsubscribe()
	}
}

func (natPS *natsPubSub) setupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		natPS.logger.Infof("Disconnected due to:%s, will attempt reconnects for %.0fm", err, totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		natPS.logger.Infof("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		natPS.logger.Infof("Exiting: %v", nc.LastError())
	}))
	return opts
}
