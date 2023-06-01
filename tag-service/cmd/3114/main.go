package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/proxy/grpcproxy"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net/http"
	"path"
	"strings"
	"tag-service/internal/middleware"
	"tag-service/pkg/swagger"
	pb "tag-service/proto"
	"tag-service/server"
	"time"
)

var port string

const SERVICE_NAME = "tag_service"

func init() {
	flag.StringVar(&port, "port", "8004", "启动端口号")
	flag.Parse()
}

func main() {}

// RunServer 服务器启动
func RunServer(port string) error {
	httpMux := runHttpServer()
	grpcS := runGrpcServer()

	endPoint := "0.0.0.0:" + port
	gwmux := runtime.NewServeMux()
	dopts := []grpc.DialOption{grpc.WithInsecure()}
	_ = pb.RegisterTagServiceHandlerFromEndpoint(context.Background(), gwmux, endPoint, dopts)
	// 让所有的HTTP请求进入网关，网关在处理请求中的信息之后才会与目标服务器建立连接，所以要有个空闲的指向网址
	httpMux.Handle("/", gwmux)

	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://localhost:2379"},
		DialTimeout: time.Second * 60,
	})
	if err != nil {
		return err
	}
	defer etcdClient.Close()

	target := fmt.Sprintf("/etcdv3://service/grpc/%s", SERVICE_NAME)
	grpcproxy.Register(etcdClient, target, ":"+port, 60)

	return http.ListenAndServe(":"+port, grpcHandlerFunc(grpcS, httpMux))
}

// runHttpServer 运行HTTP服务器
func runHttpServer() *http.ServeMux {
	serveMux := http.NewServeMux()
	// 心跳检测路由
	serveMux.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("pong"))
	})

	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,            // 资源文件
		AssetDir: swagger.AssetDir,         // 资源文件目录
		Prefix:   "third_party/swagger-ui", // 路由名
	})
	prefix := "/swagger-ui/"

	serveMux.Handle(prefix, http.StripPrefix(prefix, fileServer))
	serveMux.HandleFunc("/swagger-ui/", func(writer http.ResponseWriter, request *http.Request) {
		if !strings.HasSuffix(request.URL.Path, "swagger.json") {
			http.NotFound(writer, request)
			return
		}

		p := strings.TrimPrefix(request.URL.Path, "/swagger/")
		p = path.Join("proto", p)

		http.ServeFile(writer, request, p)
	})

	return serveMux
}

// runGrpcServer 运行gRPC服务器
func runGrpcServer() *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			middleware.AccessLog,
			middleware.ErrorLog,
			middleware.Recovery,
		)),
	}
	s := grpc.NewServer(opts...)
	// 注册服务
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	// 反射gRPC服务器，用来grpcurl测试
	reflection.Register(s)

	return s
}

// grpcHandlerFunc 运行网关，网关用来连接两种不同的网络，功能是转换
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(
		// 处理 HTTP/1.1 和 HTTP/1.0 请求
		http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			if request.ProtoMajor == 2 && strings.Contains(request.Header.Get("Content-Type"), "application/grpc") {
				grpcServer.ServeHTTP(writer, request)
			} else {
				otherHandler.ServeHTTP(writer, request)
			}
		}),
		// 处理 HTTP/2 请求，没啥用，因为用反射grpcServer来进行http/2请求了
		&http2.Server{},
	)
}
