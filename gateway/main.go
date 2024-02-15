package main

import (
	"context"
	"github.com/bigrainnn/grpc_db_ops/internal/pb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func startGateway() {
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()

	ctx := c
	mux := runtime.NewServeMux()
	err := pb.RegisterDbOpsServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatalf("start gateway failed")
	}

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("listen failed")
	}
}

func main() {
	startGateway()
}
