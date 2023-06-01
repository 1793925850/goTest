package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/naming"
	"google.golang.org/grpc"
	"log"
	pb "tag-service/proto"
	"time"
)

func main() {
	ctx := context.Background()
	clientConn, err := GetClientConn(ctx, "tag_service", nil)
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	defer clientConn.Close()

	tagServiceClient := pb.NewTagServiceClient(clientConn)
	resp, err := tagServiceClient.GetTagList(ctx, &pb.GetTagListRequest{Name: "Go"})
	if err != nil {
		log.Fatalf("tagServiceClient.GetTagList err: %v", err)
	}

	log.Printf("resp: %v", resp)
}

func GetClientConn(ctx context.Context, serviceName string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	config := clientv3.Config{
		Endpoints:   []string{"http://localhost:2379"},
		DialTimeout: time.Second * 60,
	}
	cli, err := clientv3.New(config)
	if err != nil {
		return nil, err
	}

	r := &naming.GRPCResolver{Client: cli}
	target := fmt.Sprintf("/etcdv3://service/grpc/%s", serviceName)
	// 注意这里得降级grpc包，否则没有 grpc.WithBalancer 这个方法
	opts = append(opts, grpc.WithInsecure(), grpc.WithBalancer(grpc.RoundRobin(r)), grpc.WithBlock())

	return grpc.DialContext(ctx, target, opts...)
}
