package middleware

import (
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"log"
	"runtime/debug"
	"tag-service/global"
	"tag-service/pkg/errcode"
	"tag-service/pkg/metatext"
	pb "tag-service/proto"
	"tag-service/server"
	"time"
)

func RunGrpcServer() *grpc.Server {
	// 拦截器是在gRPC服务器上的
	// 添加拦截器
	opts := []grpc.ServerOption{
		// 添加嵌套拦截器
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			AccessLog,
			ErrorLog,
			Recovery,
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

// AccessLog 访问日志
func AccessLog(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	requestLog := "访问请求日志: method: %s, 开始时间: %d, 请求: %v"
	beginTime := time.Now().Local().Unix()
	log.Printf(requestLog, info.FullMethod, beginTime, req)

	resp, err := handler(ctx, req)

	responLog := "访问响应日志: method: %s, 开始时间: %d, 响应: %v"
	endTime := time.Now().Local().Unix()
	log.Printf(responLog, info.FullMethod, endTime, resp)

	return resp, err
}

// ErrorLog 错误日志
func ErrorLog(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	if err != nil {
		errLog := "错误日志: method: %s, code: %v, message: %v, details: %v"
		s := errcode.FromError(err)
		log.Printf(errLog, info.FullMethod, s.Code(), s.Err().Error(), s.Details())
	}

	return resp, err
}

// Recovery 异常捕获
func Recovery(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	defer func() {
		if e := recover(); e != nil {
			recoveryLog := "recovery log: method: %s, message: %v, stack: %s"
			log.Printf(recoveryLog, info.FullMethod, e, string(debug.Stack()[:]))
		}
	}()

	return handler(ctx, req)
}

// ServerTracing 链路追踪
func ServerTracing(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// 从提供的上下文中获得元数据，如果没有就创建
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	}

	parentSpanContext, _ := global.Tracer.Extract(opentracing.TextMap, metatext.MetadataTextMap{md})
}
