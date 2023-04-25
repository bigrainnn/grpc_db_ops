package dbops

import (
	"context"
	"github.com/bigrainnn/grpc_db_ops/internal/dbmanager"
	hanlder "github.com/bigrainnn/grpc_db_ops/internal/dbops/handler"
	"github.com/bigrainnn/grpc_db_ops/internal/pb"
)

type Service struct {
	*dbmanager.DbManager
	pb.UnimplementedDbOpsServiceServer
}

func NewDbOpsService() *Service {
	return &Service{}
}

func (s *Service) Insert(ctx context.Context, req *pb.InsertReq) (*pb.Res, error) {
	return hanlder.ProcessInsert(ctx, req, s.DbManager)
}

func (s *Service) Update(ctx context.Context, req *pb.UpdateReq) (*pb.Res, error) {
	return hanlder.ProcessUpdate(ctx, req, s.DbManager)
}

func (s *Service) Delete(ctx context.Context, req *pb.DeleteReq) (*pb.Res, error) {
	return hanlder.ProcessDelete(ctx, req, s.DbManager)
}

func (s *Service) GetRecord(ctx context.Context, req *pb.GetRecordReq) (*pb.GetRecordRes, error) {
	return hanlder.ProcessGetRecord(ctx, req, s.DbManager)
}
