package main

import (
	"context"
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net/http"
	"strings"
	pb "tag-service/proto"
	"tag-service/server"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8004", "启动端口号")
	flag.Parse()
}

const SERVICE_NAME = "tag-service"

func main() {

}

func RunServer(port string) error {
	httpMux := runHttpServer() // 运行 HTTP 的路由
	grpcS := runGrpcServer()   // 运行 gRPC 服务
	gatewayMux := runGrpcGatewayServer()

	httpMux.Handle("/", gatewayMux) // 网关负责监控有没有 http 请求调用 gRPC 的方法

	return http.ListenAndServe(":"+port, grpcHandlerFunc(grpcS, httpMux)) // 这个就是监控纯 gRPC 调用方法或者 HTTP 调用方法，因为此时 HTTP 调用 gRPC 的路由已经放到 HTTP 的路由里了
}

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.ProtoMajor == 2 && strings.Contains(request.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(writer, request)
		} else {
			otherHandler.ServeHTTP(writer, request)
		}
	}), &http2.Server{})
}

func runHttpServer() *http.ServeMux {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("pong"))
	})

	return serveMux
}

func runGrpcServer() *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)

	return s
}

func runGrpcGatewayServer() *runtime.ServeMux {
	endpoint := "0.0.0.0:" + port
	gwmux := runtime.NewServeMux()
	dopts := []grpc.DialOption{grpc.WithInsecure()}
	// 在“ctx”完成后自动拨叫到“endpoint”并关闭连接
	_ = pb.RegisterTagServiceHandlerFromEndpoint(context.Background(), gwmux, endpoint, dopts)

	return gwmux
}
