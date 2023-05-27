# 生成 tag.pb.gw.go文件

命令：

```go
protoc -I. -I%GOPATH%/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis --grpc-gateway_out=logtostderr=true:. ./proto/*.proto
```

# 心跳检测和超时检测

- 心跳检测：定时检查网络连接是否断开
- 超时检测：当数据没有就绪时，避免当前进程在某个位置无限制的阻塞

# go-bindata的安装问题

在go的版本>=1.17时，使用如下命令进行安装go-bindata：

```powershell
go install -a -v github.com/go-bindata/go-bindata/...@latest
```

# 注意

1. http/2 是没法直接请求的，只能通过反射来进行请求注入；因此，网关接收的流量全是 http/1.1 的流量，只不过它根据请求的头部信息进行判断该条请求的目的是 grpc 还是 http，从而进行路由。