package sckio

type SocketEngineProvider interface {
	SaveAppSocket(userId int, appSck Conn)
	GetAppSocket(userId int) Conn
	RemoveAppSocket(userId int)
}
