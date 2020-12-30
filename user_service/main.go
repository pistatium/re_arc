package main

import (
	"google.golang.org/grpc"
	pb "github.com/pistatium/re_ark/protos"
	"log"
	"net"
)

func main() {
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	service := UserService{}
	pb.RegisterUserServer(server, service)
	server.Serve(listenPort)
}
