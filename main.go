package main

import ()

func main() {
	l, err := http.Listen("tcp", ":8585")
	if err != nil {
		panic(err)
	}
	http := http.Server{
		Handler: &WsServer{},
	}
	err = http.Serve(l)
	if err != nil {
		panic(err)
	}
}
