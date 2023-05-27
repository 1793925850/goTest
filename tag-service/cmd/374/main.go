package main

import (
	"flag"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net/http"
	"path"
	"strings"
	"tag-service/pkg/swagger"
	pb "tag-service/proto"
	"tag-service/server"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8004", "启动端口号")
	flag.Parse()
}

func main() {
	err := RunServer(port)
	if err != nil {
		log.Fatalf("Run Serve err: %v", err)
	}
}

func RunServer(port string) error {
	httpMux := runHttpServer()
	grpcS := runGrpcServer()

	// endPoint
	endPoint := "0.0.0.0:" + port
	// 给网关创建一个空路由
	gwmux := runtime.NewServeMux()
	dopts := []grpc.DialOption{grpc.WithInsecure()}
	// 在服务器上注册网关的路由gwmux
	// 当响应返回后，网关就会断开与 grpc断点的连接，此时网关指向的位置是endPoint
	// 相当于，网关代替用户发送连接请求，这个连接请求也会生成相应的子上下文
	_ = pb.RegisterTagServiceHandlerFromEndpoint(context.Background(), gwmux, endPoint, dopts)
	httpMux.Handle("/", gwmux)

	return http.ListenAndServe(":"+port, grpcHandlerFunc(grpcS, httpMux))
}

func runHttpServer() *http.ServeMux {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("又 TM ping是吧！没完了是吧！"))
	})
	prefix := "/swagger-ui/"
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})
	serveMux.Handle(prefix, http.StripPrefix(prefix, fileServer))
	serveMux.HandleFunc("/swagger/", func(writer http.ResponseWriter, request *http.Request) {
		if !strings.HasSuffix(request.URL.Path, "swagger.json") {
			http.NotFound(writer, request)
			return
		}

		p := strings.TrimPrefix(request.URL.Path, "/swagger/")
		p = path.Join("proto", p)

		// http://127.0.0.1:8004/swagger/tag.swagger.json
		http.ServeFile(writer, request, p)
	})

	return serveMux

}

func runGrpcServer() *grpc.Server {
	// 构建gRPC服务器
	s := grpc.NewServer()
	// 向gRPC服务器注册服务
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)

	return s
}

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	// 返回一个区分 gRPC 和 HTTP 的网关
	return h2c.NewHandler(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// 判断 http 请求的服务是 grpc 服务还是 http 服务
		if request.ProtoMajor == 2 && strings.Contains(request.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(writer, request)
		} else {
			otherHandler.ServeHTTP(writer, request)
		}
	}), &http2.Server{})
}
