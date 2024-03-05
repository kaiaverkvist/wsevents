package wsevents

import (
	"context"
	"log"
	"net/http"
	"nhooyr.io/websocket"
	"time"
)

type WebSocketServer struct {
	mux     *http.ServeMux
	handler EventHandler
}

func NewWebSocketServer(handler EventHandler) *WebSocketServer {
	return &WebSocketServer{
		mux:     http.NewServeMux(),
		handler: handler,
	}
}

func (ws *WebSocketServer) Serve(addr string) {
	ws.mux.HandleFunc("/", ws.acceptFunc)

	err := http.ListenAndServe(addr, ws.mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func (ws *WebSocketServer) acceptFunc(w http.ResponseWriter, req *http.Request) {
	conn, err := websocket.Accept(w, req, nil)
	if err != nil {
		ws.handler.OnError(context.Background(), nil, err)
		return
	}

	defer conn.CloseNow()

	ctx, cancel := context.WithTimeout(req.Context(), time.Second*30)
	defer cancel()

	// Connect callback
	ws.handler.OnConnect(ctx, conn)

	ws.readLoop(ctx, conn)

	err = conn.Close(websocket.StatusNormalClosure, "")
	ws.handler.OnDisconnect(ctx, conn, err)
}
func (ws *WebSocketServer) readLoop(ctx context.Context, conn *websocket.Conn) {
	for {
		_, payload, err := conn.Read(ctx)
		if err != nil {
			ws.handler.OnError(ctx, conn, err)
			return
		}

		ws.handler.OnMessage(ctx, conn, payload)
	}
}
