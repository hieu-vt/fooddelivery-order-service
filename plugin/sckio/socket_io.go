package sckio

import (
	"flag"
	"fmt"
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/200Lab-Education/go-sdk/logger"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"net"
	"net/http"
	"net/url"
)

type Conn interface {
	Close() error
	// ID returns session id
	ID() string
	URL() url.URL
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	RemoteHeader() http.Header
	Context() interface{}
	SetContext(ctx interface{})

	Namespace() string
	Emit(eventName string, v ...interface{})

	Join(room string)
	Leave(room string)
	LeaveAll()
	Rooms() []string
}

type Socket interface {
	Id() string
	Rooms() []string
	Request() *http.Request
	On(event string, f interface{}) error
	Emit(event string, args ...interface{}) error
	Join(room string) error
	Leave(room string) error
	Disconnect()
	BroadcastTo(room, event string, args ...interface{}) error
}

//type AppSocket interface {
//	ServiceContext() goservice.ServiceContext
//	Logger() logger.Logger
//	CurrentUser() sdkcm.Requester
//	SetCurrentUser(sdkcm.Requester)
//	BroadcastToRoom(room, event string, args ...interface{})
//	String() string
//	Conn
//}

type Config struct {
	Name          string
	MaxConnection int
}

type sckServer struct {
	Config
	io     *socketio.Server
	logger logger.Logger
}

func NewSocketIo(name string) *sckServer {
	return &sckServer{
		Config: Config{Name: name},
	}
}

type ObserverProvider interface {
	AddObservers(server *socketio.Server, sc goservice.ServiceContext, l logger.Logger) func(conn socketio.Conn) error
}

func (s *sckServer) StartRealtimeServer(engine *gin.Engine, sc goservice.ServiceContext, op ObserverProvider) {
	server := socketio.NewServer(nil)

	s.io = server

	s.io.OnConnect("/", op.AddObservers(server, sc, s.logger))

	engine.GET("/socket.io/", gin.WrapH(server))
	engine.POST("/socket.io/", gin.WrapH(server))
}

func (s *sckServer) GetPrefix() string {
	return s.Config.Name
}

func (s *sckServer) Get() interface{} {
	return s
}

func (s *sckServer) Name() string {
	return s.Config.Name
}

func (s *sckServer) InitFlags() {
	pre := s.GetPrefix()
	flag.IntVar(&s.MaxConnection, fmt.Sprintf("%s-max-connection", pre), 2000, "socket max connection")
}

func (s *sckServer) Configure() error {
	s.logger = logger.GetCurrent().GetLogger("io.socket")
	return nil
}

func (s *sckServer) Run() error {
	return s.Configure()
}

func (s *sckServer) GetClient() *socketio.Server {
	return s.io
}

func (s *sckServer) Stop() <-chan bool {
	c := make(chan bool)
	go func() {
		s.io.Close()
		c <- true
	}()
	return c
}
