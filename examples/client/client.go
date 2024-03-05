package main

import (
	"context"
	"log"
	"nhooyr.io/websocket"
	"wsevents"
)

func main() {
	client := wsevents.NewWebSocketClient(wsHandler{})
	err := client.Dial("ws://localhost:7172", nil)
	if err != nil {
		log.Println("Unable to connect: ", err)
	}
}

type wsHandler struct{}

func (w wsHandler) OnConnect(ctx context.Context, conn *websocket.Conn) {
	log.Println("Connect")

	_ = conn.Write(ctx, websocket.MessageText, []byte("Hello from client :)"))
}

func (w wsHandler) OnDisconnect(ctx context.Context, conn *websocket.Conn, err error) {
	log.Println("OnDisconnect")
}

func (w wsHandler) OnError(ctx context.Context, conn *websocket.Conn, err error) {
	log.Println("Error: ", err)
}

func (w wsHandler) OnMessage(ctx context.Context, conn *websocket.Conn, payload []byte) {
	log.Println(string(payload))
}
