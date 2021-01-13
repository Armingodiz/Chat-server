package main

import (
	"context"
	"fmt"
	"net/http"
	"nhooyr.io/websocket"
	"time"
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
	// adding this line caues closing the connection after go in waitin mood for id .
	//defer connection.Close(websocket.StatusInternalError, "")
	client := Client{
		conn: connection,
	}
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	client.Register(ctx, server.manager)
}
