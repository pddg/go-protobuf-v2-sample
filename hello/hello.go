package hello

import (
	"context"

	"github.com/pddg/go-protobuf-v2-sample/hello/pb"
)

type HelloServer struct{
	pb.UnimplementedHelloServiceServer
}

func (hs HelloServer) Hello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello " + request.Name}, nil
}

func NewHelloServiceServer() pb.HelloServiceServer {
	return &HelloServer{}
}