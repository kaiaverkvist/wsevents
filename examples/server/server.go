package main

import (
	"context"
	"log"
	"nhooyr.io/websocket"
	"wsevents"
)

func main() {
	server := wsevents.NewWebSocketServer(wsHandler{})
	server.Serve(":7172")
}

type wsHandler struct{}

func (w wsHandler) OnConnect(ctx context.Context, conn *websocket.Conn) {
	log.Println("Connect")
	conn.Write(ctx, websocket.MessageText, []byte("Luke, I am your server"))
}

func (w wsHandler) OnDisconnect(ctx context.Context, conn *websocket.Conn, err error) {
	log.Println("OnDisconnect")
}

func (w wsHandler) OnError(ctx context.Context, conn *websocket.Conn, err error) {
	log.Println("Error: ", err)
}

func (w wsHandler) OnMessage(ctx context.Context, conn *websocket.Conn, payload []byte) {
	log.Println("From client: ", string(payload))
}
