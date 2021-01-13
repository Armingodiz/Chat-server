package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
)

var ln net.Listener

func main() {
	l, err := net.Listen("tcp", ":8080")
	ln = l
	if err != nil {
		panic(err)
	}
	queue, err := NewRabitmq(context.Background())
	if err != nil {
		panic(err)
	}
	http := http.Server{
		Handler: &WsServer{
			manager: &Manager{
				clients: make(map[uint64]*Client),
				queue:   queue,
			},
		},
	}
	fmt.Println(" Listenning on 8080 for client ... ")
	err = http.Serve(l)
	if err != nil {
		panic(err)
	}
}
