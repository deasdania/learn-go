package main

import (
	"context"
	"fmt"
	pb "learn-go/gen/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type testApiServer struct {
	pb.UnimplementedTestApiServer
}

func (s *testApiServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{}, nil
}

func (s *testApiServer) Echo(ctx context.Context, req *pb.ResponseRequest) (*pb.ResponseRequest, error) {
	return req, nil
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8089")
	if err != nil {
		log.Println(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTestApiServer(grpcServer, &testApiServer{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("running on 8089")
}
