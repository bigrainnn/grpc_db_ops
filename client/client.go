package main

import (
	"context"
	"fmt"
	pb2 "github.com/bigrainnn/grpc_db_ops/internal/pb"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	// 连接到 gRPC 服务器
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb2.NewDbOpsServiceClient(conn)

	// 创建 InsertReq 示例
	insertReq := &pb2.InsertReq{
		BusinessId: "image_route",
		Records: []*pb2.Record{
			{
				Fields: []*pb2.Field{
					{Key: "user_id", Value: "test3"},
					{Key: "route", Value: "test3"},
				},
			},
		},
	}

	// 调用远程 Insert 方法
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := client.Insert(ctx, insertReq)
	if err != nil {
		log.Fatalf("could not call Insert: %v", err)
	}
	fmt.Printf("Insert Response: %v\n", response)
}
