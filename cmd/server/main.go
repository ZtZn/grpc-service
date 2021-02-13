package main

import (
	"log"
	"net"

	"github.com/ztzn/grpc_services/pkg/api"
	"github.com/ztzn/grpc_services/pkg/api/adder"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}
	api.RegisterAdderServer(s, srv)

	l, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
