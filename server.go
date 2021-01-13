package main

import (
	"net/http"
	"nhooyr.io/websocket"
)

type WsServer struct {
	manager *Manager
}

func (server *WsServer) ServeHttps(w http.ResponseWriter, r *http.Request) {
	connection, err := websocket.Accept(w, r, nil)
	if err != nil {
		panic(err)
	}
	//defer c.Close(websocket.StatusInternalError, "")
	Client := Client{
		conn: connection,
	}
	server.manager.AddClient(ctx.Background(), &client)

}
