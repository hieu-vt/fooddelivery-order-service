package main

import (
	"github.com/zishang520/engine.io/types"
	"github.com/zishang520/socket.io/socket"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	httpServer := types.CreateServer(nil)
	io := socket.NewServer(httpServer, nil)
	io.On("connection", func(clients ...any) {
		log.Println("Client connected")
		client := clients[0].(*socket.Socket)
		client.On("event", func(datas ...any) {
			log.Println("Client send event")
			client.Emit("event", "Hello client")
		})
		client.On("disconnect", func(...any) {
			log.Println("Client disconnected")
		})
	})

	log.Println("running with port 3000")

	httpServer.Listen("127.0.0.1:3000", nil)

	exit := make(chan struct{})
	SignalC := make(chan os.Signal)

	signal.Notify(SignalC, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for s := range SignalC {
			switch s {
			case os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				close(exit)
				return
			}
		}
	}()

	<-exit
	httpServer.Close(nil)
	os.Exit(0)
}
