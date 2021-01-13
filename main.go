package main

import (
  "fmt"
  "net/http"
  "net"
)

func main() {
	l, err := net.Listen("tcp", ":8585")
	if err != nil {
		panic(err)
	}
	http := http.Server{
		Handler: &WsServer{},
	}
  fmt.Println(" Listenning on 8585 for client ... ")
	err = http.Serve(l)
	if err != nil {
		panic(err)
	}
}
