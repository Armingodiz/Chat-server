package main

import (
	"context"
	"fmt"
	"net/http"
	"nhooyr.io/websocket"
)

type WsServer struct {
	manager *Manager
}

func (server *WsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request for connection")
	connection, err := websocket.Accept(w, r, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer connection.Close(websocket.StatusInternalError, "")
	client := Client{
		conn: connection,
	}
	server.manager.AddClient(context.Background(), &client)

}
