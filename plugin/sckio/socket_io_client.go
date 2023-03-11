package sckio

import (
	socketio "github.com/googollee/go-socket.io"
)

type GetSocketClient interface {
	GetClient() *socketio.Server
}
