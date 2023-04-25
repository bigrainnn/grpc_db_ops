package hanlder

import (
	"context"
	"errors"
	"github.com/bigrainnn/db_ops_common/task_generator/insert"
	"github.com/bigrainnn/db_ops_common/task_generator/query"
	"github.com/bigrainnn/db_ops_common/task_generator/update"
	"github.com/bigrainnn/db_ops_common/task_processor"
	"github.com/bigrainnn/grpc_db_ops/internal/dbmanager"
	"github.com/bigrainnn/grpc_db_ops/internal/dbops/convert"
	"github.com/bigrainnn/grpc_db_ops/internal/err_code"
	"github.com/bigrainnn/grpc_db_ops/internal/logger"
	"github.com/bigrainnn/grpc_db_ops/internal/pb"
	log "go.uber.org/zap"
	"strconv"
)

const (
	ChangeLimit = 100
	QueryLimit  = 2000
)

func NewInsertHandler(request *pb.InsertReq, tableName string) (*insert.TaskGenerator, error) {
	records := convert.InsertReqToRecords(request)
	if len(records) > ChangeLimit {
		return nil, errors.New("the number of changed items needs to be less and equal than " + strconv.Itoa(ChangeLimit))
	}

	if len(records) == 0 {
		return nil, errors.New("len of records is zero")
	}

	for _, v := range records {
		if len(v.Record) == 0 {
			return nil, errors.New("len of a record is zero")
		}
	}

	return &insert.TaskGenerator{
		Records:   records,
		TableName: tableName,
		Mode:      insert.Mode(request.GetMode()),
	}, nil
}

func NewUpdateHandler(request *pb.UpdateReq, tableName string) (*update.TaskGenerator, error) {
	updateRecord := convert.UpdateReqToUpdatedRecords(request)
	if len(updateRecord) > ChangeLimit {
		return nil, errors.New("the number of changed items needs to be less than " + string(rune(ChangeLimit)))
	}

	return &update.TaskGenerator{
		UpdateFields: updateRecord,
		TableName:    tableName,
	}, nil
}

func NewDeleteHandler(request *pb.DeleteReq, tableName string) (*update.TaskGenerator, error) {
	updateRecord := convert.DeleteReqToUpdatedRecords(request)
	if len(updateRecord) > ChangeLimit {
		return nil, errors.New("the number of changed items needs to be less than " + string(rune(ChangeLimit)))
	}

	return &update.TaskGenerator{
		UpdateFields: updateRecord,
		TableName:    tableName,
	}, nil
}

func NewQueryHandler(request *pb.GetRecordReq, table string) (*query.TaskGenerator, error) {
	if request.GetLimit() > QueryLimit {
		return nil, errors.New("the number of query items needs to be less than " + strconv.Itoa(QueryLimit))
	}
	if len(request.QueryFields) == 0 {
		return nil, errors.New("query fields can not be zero")
	}

	columnSort, whereStruct := convert.QueryQeqCovert(request)
	return &query.TaskGenerator{
		Limit:       request.GetLimit(),
		Pos:         request.GetPos(),
		QueryField:  request.GetQueryFields(),
		Sort:        columnSort,
		WhereStruct: whereStruct,
		TableName:   table,
	}, nil
}

func NewDbTaskGenerator(req interface{}, table string) (task_processor.TaskGenerator, error) {
	switch req.(type) {
	case *pb.InsertReq:
		r := req.(*pb.InsertReq)
		return NewInsertHandler(r, table)
	case *pb.UpdateReq:
		r := req.(*pb.UpdateReq)
		return NewUpdateHandler(r, table)
	case *pb.GetRecordReq:
		r := req.(*pb.GetRecordReq)
		return NewQueryHandler(r, table)
	case *pb.DeleteReq:
		r := req.(*pb.DeleteReq)
		return NewDeleteHandler(r, table)
	}
	return nil, errors.New("unsupported req type")
}

func GetBusinessId(req interface{}) string {
	switch req.(type) {
	case *pb.InsertReq:
		r := req.(*pb.InsertReq)
		return r.GetBusinessId()
	case *pb.UpdateReq:
		r := req.(*pb.UpdateReq)
		return r.GetBusinessId()
	case *pb.GetRecordReq:
		r := req.(*pb.GetRecordReq)
		return r.GetBusinessId()
	case *pb.DeleteReq:
		r := req.(*pb.DeleteReq)
		return r.GetBusinessId()
	}
	return ""
}

func NewHandler(ctx context.Context, req interface{}, dbManager *dbmanager.DbManager) (*task_processor.Processor, error) {
	businessId := GetBusinessId(req)
	taskGenerator, err := NewDbTaskGenerator(req, businessId)
	if err != nil {
		logger.Error("create task generator failed")
		return nil, err
	}

	// 获取对应的数据库实例
	db, err := dbManager.GetDB(businessId)
	if err != nil {
		logger.Error("get db failed")
		return nil, err
	}

	handler := &task_processor.Processor{}
	handler.C = ctx
	handler.Task = taskGenerator
	handler.DB = db
	handler.BusinessId = businessId
	return handler, nil
}

func ProcessUpdate(ctx context.Context, req *pb.UpdateReq, dbManager *dbmanager.DbManager) (*pb.Res, error) {
	res := pb.Res{}
	defer func() {
		logger.Info("", log.Any("Update End", res.String()))
	}()

	handler, err := NewHandler(ctx, req, dbManager)
	if err != nil {
		res.ErrCode = err_code.NewHandlerFailed
		res.ErrMsg = err.Error()
		return nil, err
	}

	err = handler.ProcessTx()
	if err != nil {
		res.ErrCode = err_code.ProcessTXFailed
		res.ErrMsg = err.Error()
		return nil, err
	}

	return &res, nil
}

func ProcessInsert(ctx context.Context, req *pb.InsertReq, dbManager *dbmanager.DbManager) (*pb.Res, error) {
	res := pb.Res{}
	defer func() {
		logger.Info("", log.Any("Delete End", res.String()))
	}()

	handler, err := NewHandler(ctx, req, dbManager)
	if err != nil {
		res.ErrCode = err_code.NewHandlerFailed
		res.ErrMsg = err.Error()
		return nil, err
	}

	err = handler.ProcessTx()
	if err != nil {
		res.ErrCode = err_code.NewHandlerFailed
		res.ErrMsg = err.Error()
		return nil, err
	}

	return &res, nil
}

func ProcessDelete(ctx context.Context, req *pb.DeleteReq, dbManager *dbmanager.DbManager) (*pb.Res, error) {
	res := pb.Res{}
	defer func() {
		logger.Info("", log.Any("Delete End", res.String()))
	}()

	handler, err := NewHandler(ctx, req, dbManager)
	if err != nil {
		res.ErrCode = err_code.NewHandlerFailed
		res.ErrMsg = err.Error()
		return nil, err
	}

	err = handler.ProcessTx()
	if err != nil {
		res.ErrCode = err_code.NewHandlerFailed
		res.ErrMsg = err.Error()
		return nil, err
	}

	return &res, nil
}

func ProcessGetRecord(ctx context.Context, req *pb.GetRecordReq, dbManager *dbmanager.DbManager) (*pb.GetRecordRes, error) {
	res := pb.GetRecordRes{}
	defer func() {
		logger.Info("", log.Any("GetRecord End", res.String()))
	}()

	handler, err := NewHandler(ctx, req, dbManager)
	if err != nil {
		res.ErrCode = err_code.NewHandlerFailed
		res.ErrMsg = err.Error()
		return nil, err
	}

	records, nums, err := handler.ProcessQuery()
	if err != nil {
		res.ErrCode = err_code.ProcessQueryFailed
		res.ErrMsg = err.Error()
		return nil, err
	}

	convert.HandlerQueryResponse(nums, records, &res)
	return &res, nil
}
