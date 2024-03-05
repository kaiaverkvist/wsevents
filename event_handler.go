package wsevents

import (
	"context"
	"nhooyr.io/websocket"
)

type EventHandler interface {
	OnConnect(ctx context.Context, conn *websocket.Conn)
	OnDisconnect(ctx context.Context, conn *websocket.Conn, err error)
	OnError(ctx context.Context, conn *websocket.Conn, err error)
	OnMessage(ctx context.Context, conn *websocket.Conn, payload []byte)
}
