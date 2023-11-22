package main

import (
	"context"
	"log"
	"net"

    //"github.com/Ericrulec/gochat/db"
	"github.com/Ericrulec/gochat/proto"

	"google.golang.org/grpc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

const (
    PORT = "3000"
)

type Server struct {
	proto.UnimplementedMessagingServiceServer
}


func (s *Server) GetUser(ctx context.Context, in *proto.QueryUser) (*proto.User, error) {
    if in.UserId == "" {
        return nil, status.Error(
            codes.InvalidArgument, "UserId must be non-empty",
            )
    }
	return &proto.User{
        UserId: "1",
        Name: "Erik Jensson",
        Status: proto.User_ONLINE,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalln("Failed to create listener:", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)

	proto.RegisterMessagingServiceServer(s, &Server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
