package main

import (
	"context"
	"fmt"
	"net"

	proto "github.com/siddheshRajendraNimbalkar/GO-GRPC/protoc"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedNameServer
}

func (s *server) NameReply(ctx context.Context, req *proto.NameRequest) (*proto.NameResponse, error) {
	fullName := req.FirstName + " " + req.MiddleName + " " + req.LastName

	return &proto.NameResponse{FullName: fullName}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	grpcServer := grpc.NewServer()
	proto.RegisterNameServer(grpcServer, &server{})

	fmt.Println("Server running on port :4040")

	if err := grpcServer.Serve(listener); err != nil {
		fmt.Println("Error: ", err)
	}

}
