package main

import (
	"net/http"
  "nhooyr.io/websocket"
)

type WsServer struct {
}

func (server *WsServer) ServeHttps(w http.ResponseWriter, r *http.Request) {
	_, err := websocket.Accept(w, r, nil)
	if err != nil {
		panic(err)
	}
	//defer c.Close(websocket.StatusInternalError, "")
}
