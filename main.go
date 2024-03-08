package main

import (
	"log"
	"net"

	"geearrpeecee/chat"
	"geearrpeecee/pb"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("failed to create listener:", err)
	}

	s := grpc.NewServer()
	chatserver := chat.NewServer()
	//reflection.Register(s)

	pb.RegisterChatServiceServer(s, chatserver)
	if err := s.Serve(listener); err != nil {
		log.Fatalln("failed to serve:", err)
	}
}
