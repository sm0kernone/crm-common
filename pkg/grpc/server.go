package grpc

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

func Start(port int) (*grpc.Server, net.Listener) {
	log.Println("Starting GRPC server on port:", port)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Printf("Failed to start RPC server: %v", err)
	}

	s := grpc.NewServer()

	return s, listener
}
