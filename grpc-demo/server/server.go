package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	pb "grpc-demo/proto"
	"net"
)

type GreeterServer struct{}

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()
}

func (s *GreeterServer) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello.world"}, nil
}

func (s *GreeterServer) SayList(r *pb.HelloRequest, stream pb.Greeter_SayListServer) error {
	for n := 0; n <= 6; n++ {
		_ = stream.Send(&pb.HelloReply{Message: "hello.list"})
	}

	return nil
}

func main() {
	server := grpc.NewServer()                         // 创建服务器实例
	pb.RegisterGreeterServer(server, &GreeterServer{}) // 在该服务器上注册 RPC服务（即 Greeter）
	lis, _ := net.Listen("tcp", ":"+port)              // 创建监听器，服务器监听端口
	server.Serve(lis)                                  // 监听到请求便返回指定服务结果
}
