package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	pb "grpc-demo/proto"
	"io"
	"log"
	"net"
)

type GreeterServer struct{}

func (s *GreeterServer) SayRoute(stream pb.Greeter_SayRouteServer) error {
	n := 0
	for {
		_ = stream.Send(&pb.HelloReply{Message: "say.route"})

		resp, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		n++
		log.Printf("resp: %v", resp)
	}
}

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

func (s *GreeterServer) SayRecord(stream pb.Greeter_SayRecordServer) error {
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			message := &pb.HelloReply{Message: "hello.record"}
			return stream.SendAndClose(message)
		}
		if err != nil {
			return err
		}

		log.Printf("resp: %v", resp)
	}

	return nil
}

func main() {
	server := grpc.NewServer()                         // 创建服务器实例
	pb.RegisterGreeterServer(server, &GreeterServer{}) // 在该服务器上注册 RPC服务（即 Greeter）
	lis, _ := net.Listen("tcp", ":"+port)              // 创建监听器，服务器监听端口
	err := server.Serve(lis)                           // 监听到请求便返回指定服务结果
	if err != nil {
		log.Fatalln(err)
	}
}
