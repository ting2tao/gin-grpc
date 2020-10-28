package main

import (
	"log"
	"net"

	pb "gin-grpc/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/reflection"
)

// SimpleService server is used to implement GreeterServer.
type SimpleService struct{}

func (s *SimpleService) Route(ctx context.Context, request *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	res := pb.SimpleResponse{
		Code:  200,
		Value: "hello " + request.Data,
	}
	return &res, nil
}

const (
	// Address 监听地址
	Address string = ":50051"
	// Network 网络通信协议
	Network string = "tcp"
)

func main() {
	lis, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println(Address + " net.Listing...")

	// 新建gRPC服务器实例
	grpcServer := grpc.NewServer()
	// 在gRPC服务器注册我们的服务
	pb.RegisterSimpleServer(grpcServer, &SimpleService{})
	// Register reflection service on gRPC server.
	//reflection.Register(s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
