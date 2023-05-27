package middleware

import (
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	pb "tag-service/proto"
	"tag-service/server"
)

func RunGrpcServer() *grpc.Server {
	// 拦截器是在gRPC服务器上的
	// 添加拦截器
	opts := []grpc.ServerOption{
		// 添加嵌套拦截器
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			HelloInterceptor,
			WorldInterceptor,
		)),
	}
	// 构建gRPC服务器
	s := grpc.NewServer(opts...)
	// 向gRPC服务器注册服务
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)

	return s
}

func HelloInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("你好")
	resp, err := handler(ctx, req)
	log.Println("再见")

	return resp, err
}

func WorldInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("世界")
	resp, err := handler(ctx, req)
	log.Println("世界")

	return resp, err
}
