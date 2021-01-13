package main

import (
	"net/http"
  "context"
	"nhooyr.io/websocket"
)

type WsServer struct {
	manager *Manager
}

func (server *WsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	connection, err := websocket.Accept(w, r, nil)
	if err != nil {
		panic(err)
	}
	//defer c.Close(websocket.StatusInternalError, "")
	client := Client{
		conn: connection,
	}
	server.manager.AddClient(context.Background(), &client)

}
