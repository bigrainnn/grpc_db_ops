package main

import (
	"fmt"
	config "github.com/bigrainnn/grpc_db_ops/internal/configer"
	"github.com/bigrainnn/grpc_db_ops/internal/dbmanager"
	"github.com/bigrainnn/grpc_db_ops/internal/dbops"
	"github.com/bigrainnn/grpc_db_ops/internal/pb"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 打印当前工作目录
	fmt.Println("Current directory:", currentDir)

	cfg, err := config.LoadConfig("./conf/config.toml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	dbMgr := dbmanager.NewDbManager()
	for _, v := range cfg.DatabaseConnections {
		err = dbMgr.AddConnection(v.BusinessId, v)
		if err != nil {
			log.Fatalf("Failed to add connection for %s: %v", "image_route", err)
		}
	}

	dbOpsService := &dbops.Service{
		DbManager: dbMgr,
	}

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterDbOpsServiceServer(grpcServer, dbOpsService)

	log.Println("Starting gRPC server on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
