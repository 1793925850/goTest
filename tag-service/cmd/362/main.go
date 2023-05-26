package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	pb "tag-service/proto"
	"tag-service/server"
)

var grpcPort string
var httpPort string

func init() {
	// gRPC(HTTP/2)
	flag.StringVar(&grpcPort, "grpc_port", "8001", "gRPC启动端口号")
	// HTTP/1.1
	flag.StringVar(&httpPort, "http_port", "8002", "HTTP启动端口号")
	flag.Parse()
}

func main() {
	// 创建一个无缓冲的通道，通道元素是 error
	errs := make(chan error)

	go func() {
		err := RunHttpServer(httpPort)
		if err != nil {
			errs <- err
		}
	}()

	go func() {
		err := RunGrpcServer(grpcPort)
		if err != nil {
			errs <- err
		}
	}()

	select {
	case err := <-errs:
		log.Fatalf("Run Server err: %s", err.Error())
	}
}

func RunHttpServer(port string) error {
	// serveMux 这就是个路由
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte(`pong`))
	})

	return http.ListenAndServe(":"+port, serveMux)
}

func RunGrpcServer(port string) error {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("gRPC的服务器的网络监听发生错误: %s", err.Error())
		return err
	}

	err = s.Serve(lis)
	if err != nil {
		fmt.Println("gRPC的服务器发生错误: %s", err.Error())
		return err
	}

	return nil
}
