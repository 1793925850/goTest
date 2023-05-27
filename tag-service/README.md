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

