package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	pb "grpc-demo/proto"
	"io"
	"log"
)

type GreeterServer struct{}

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()
}

func main() {
	conn, _ := grpc.Dial(":"+port, grpc.WithInsecure()) // 创建与给定目标（服务端）的连接句柄
	defer conn.Close()

	client := pb.NewGreeterClient(conn) // 创建 服务Greeter 的客户端对象
	_ = SayRoute(client)                // 发送 RPC 请求，等待同步响应，得到回调后返回响应结果
}

func SayHello(client pb.GreeterClient) error {
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{Name: "eddycjy"})
	log.Printf("client.SayHello resp: %s", resp.Message)
	return nil
}

func SayList(client pb.GreeterClient) error {
	stream, _ := client.SayList(context.Background(), &pb.HelloRequest{Name: "JinXiao"})
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("resp: %v", resp)
	}

	return nil
}

func SayRecord(client pb.GreeterClient) error {
	stream, _ := client.SayRecord(context.Background())
	for n := 0; n <= 6; n++ {
		_ = stream.Send(&pb.HelloRequest{Name: "JinXiao"})
	}
	resp, _ := stream.CloseAndRecv()

	log.Printf("resp err: %v", resp)
	return nil
}

func SayRoute(client pb.GreeterClient) error {
	stream, _ := client.SayRoute(context.Background())
	for n := 0; n <= 6; n++ {
		_ = stream.Send(&pb.HelloRequest{Name: "JinXiao"}) // 首个请求一定是由客户端发送的
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("resp err: %v", resp)
	}

	_ = stream.CloseSend()

	return nil
}
