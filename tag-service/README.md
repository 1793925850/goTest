# 生成 tag.pb.gw.go文件

命令：

```go
protoc -I. -I%GOPATH%/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis --grpc-gateway_out=logtostderr=true:. ./proto/*.proto
```

