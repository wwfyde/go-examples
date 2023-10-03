package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Service int

func (s *Service) Value(args Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func main() {
	service := new(Service)
	// register service to rpc
	rpc.Register(service)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listen error: ", err)
	}

	go http.Serve(l, nil)

}
