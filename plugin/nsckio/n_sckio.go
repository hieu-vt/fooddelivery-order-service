package nsckio

import (
	"flag"
	"fmt"
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/200Lab-Education/go-sdk/logger"
	"github.com/zishang520/engine.io/types"
	"github.com/zishang520/socket.io/socket"
	"log"
)

type Config struct {
	Name          string
	MaxConnection int
}

type nSckServer struct {
	Config
	io     *socket.Socket
	logger logger.Logger
}

func NewNSocketIo(name string) *nSckServer {
	return &nSckServer{
		Config: Config{Name: name},
	}
}

type NObserverProvider interface {
	AddNObservers(server *socket.Socket, sc goservice.ServiceContext, l logger.Logger)
}

func (s *nSckServer) StartNRealtimeServer(sc goservice.ServiceContext, f func(client *socket.Socket, sc goservice.ServiceContext, l logger.Logger)) {
	httpServer := types.CreateServer(nil)
	server := socket.NewServer(httpServer, nil)

	server.On("connection", func(clients ...any) {
		log.Println("Client connected")
		client := clients[0].(*socket.Socket)
		s.io = client
		f(client, sc, s.logger)
	})

	s.logger.Info("Listening with port 3006")

	httpServer.Listen("127.0.0.1:3006", nil)

	//exit := make(chan struct{})
	//SignalC := make(chan os.Signal)
	//
	//signal.Notify(SignalC, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	//go func() {
	//	for s := range SignalC {
	//		switch s {
	//		case os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
	//			close(exit)
	//			return
	//		}
	//	}
	//}()
	//
	//<-exit
	//httpServer.Close(nil)
}

func (s *nSckServer) GetPrefix() string {
	return s.Config.Name
}

func (s *nSckServer) Get() interface{} {
	return s
}

func (s *nSckServer) Name() string {
	return s.Config.Name
}

func (s *nSckServer) InitFlags() {
	pre := s.GetPrefix()
	flag.IntVar(&s.MaxConnection, fmt.Sprintf("%s-max-connection", pre), 2000, "socket max connection")
}

func (s *nSckServer) Configure() error {
	s.logger = logger.GetCurrent().GetLogger("io.socket")
	return nil
}

func (s *nSckServer) Run() error {
	return s.Configure()
}

func (s *nSckServer) GetClient() *socket.Socket {
	return s.io
}

func (s *nSckServer) Stop() <-chan bool {
	c := make(chan bool)
	go func() {
		c <- true
	}()
	return c
}
