package nsckio

import "github.com/zishang520/socket.io/socket"

type NSocketEngineProvider interface {
	SaveAppSocket(userId int, appSck *socket.Socket)
	GetAppSocket(userId int) *socket.Socket
	RemoveAppSocket(userId int)
}
