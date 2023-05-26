package main

import (
	"flag"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	pb "tag-service/proto"
	"tag-service/server"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8003", "启动端口号")
	flag.Parse()
}

func main() {
	l, err := RunTCPServer(port)
	if err != nil {
		log.Fatalf("Run TCP Server err: %v", err)
	}

	m := cmux.New(l)
	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldPrefixSendSettings("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP1Fast())

	grpcS := RunGrpcServer()
	httpS := RunHttpServer(port)

	go grpcS.Serve(grpcL)
	go httpS.Serve(httpL)

	err = m.Serve()
	if err != nil {
		log.Fatalf("Run Serve err:%v", err)
	}
}

func RunTCPServer(port string) (net.Listener, error) {
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return nil, err
	}

	return l, err
}

func RunGrpcServer() *grpc.Server {
	s := grpc.NewServer()
	// 在服务器上注册具体服务
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)

	return s
}

func RunHttpServer(port string) *http.Server {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("pong"))
	})

	return &http.Server{
		Addr:    ":" + port,
		Handler: serveMux,
	}
}