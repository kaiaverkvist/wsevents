package wsevents

import (
	"context"
	"nhooyr.io/websocket"
	"time"
)

type WebSocketClient struct {
	handler EventHandler
}

func NewWebSocketClient(handler EventHandler) *WebSocketClient {
	return &WebSocketClient{
		handler: handler,
	}
}

func (ws *WebSocketClient) Dial(addr string, options *websocket.DialOptions) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	conn, _, err := websocket.Dial(ctx, addr, options)
	if err != nil {
		return err
	}
	defer conn.CloseNow()

	ws.handler.OnConnect(ctx, conn)

	ws.readLoop(ctx, conn)

	err = conn.Close(websocket.StatusNormalClosure, "")
	ws.handler.OnDisconnect(ctx, conn, err)

	return nil
}

func (ws *WebSocketClient) readLoop(ctx context.Context, conn *websocket.Conn) {
	for {
		_, payload, err := conn.Read(ctx)
		if err != nil {
			ws.handler.OnError(ctx, conn, err)
			return
		}

		ws.handler.OnMessage(ctx, conn, payload)
	}
}
