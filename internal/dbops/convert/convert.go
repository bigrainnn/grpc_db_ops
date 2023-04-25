package convert

import (
	"database/sql"
	"github.com/bigrainnn/db_ops_common/data_struct"
	"github.com/bigrainnn/grpc_db_ops/internal/pb"
)

func ProtoFieldsToRecord(rawFields []*pb.Field) data_struct.Record {
	record := data_struct.Record{}
	for _, rawField := range rawFields {
		field := data_struct.Field{rawField.GetKey(), rawField.GetValue()}
		record = append(record, field)
	}
	return record
}

func ProtoRecordsToRecord(rawRecords []*pb.Record) []data_struct.RecordForInsert {
	var res []data_struct.RecordForInsert
	for _, rawRecord := range rawRecords {
		rawFields := rawRecord.GetFields()
		record := ProtoFieldsToRecord(rawFields)
		res = append(res, data_struct.RecordForInsert{
			Record: record,
		})
	}
	return res
}

func InsertReqToRecords(request *pb.InsertReq) []data_struct.RecordForInsert {
	rawRecords := request.GetRecords()
	return ProtoRecordsToRecord(rawRecords)
}

func UpdateReqToUpdatedRecords(request *pb.UpdateReq) []data_struct.UpdatedRecord {
	var res []data_struct.UpdatedRecord
	rawItems := request.GetItems()
	for _, value := range rawItems {
		UpdateF := ProtoFieldsToRecord(value.GetUpdateFields())
		UniqueF := ProtoFieldsToRecord(value.GetWhereFields())
		U := data_struct.UpdatedRecord{
			UpdateField: UpdateF,
			WhereField:  UniqueF,
		}
		res = append(res, U)
	}
	return res
}

func DeleteReqToUpdatedRecords(request *pb.DeleteReq) []data_struct.UpdatedRecord {
	var res []data_struct.UpdatedRecord
	UniqueF := ProtoFieldsToRecord(request.GetWhereFields())
	U := data_struct.UpdatedRecord{
		UpdateField: data_struct.Record{
			data_struct.Field{
				"enabled", "0",
			},
		},
		WhereField: UniqueF,
	}
	res = append(res, U)

	return res
}

func QueryQeqCovert(request *pb.GetRecordReq) ([]data_struct.ColumnSort, data_struct.WhereStruct) {
	var sortSlice []data_struct.ColumnSort
	rawSort := request.GetSorts()
	for _, v := range rawSort {
		sort := data_struct.ColumnSort{}
		sort.Name = v.GetColumnName()
		sort.Desc = v.GetDesc()
		sortSlice = append(sortSlice, sort)
	}

	var whereStruct data_struct.WhereStruct
	rawWhereStruct := request.GetWhereCondition()
	whereStruct.Expression = rawWhereStruct.GetExpression()
	rawImmediate := rawWhereStruct.GetImmediate()
	for _, v := range rawImmediate {
		whereStruct.Immediate = append(whereStruct.Immediate, v)
	}

	return sortSlice, whereStruct
}

func NewQueryRecord(col []string) *data_struct.QueryRecord {
	p := make([]interface{}, len(col))
	c := make([]sql.NullString, len(col))
	for i := range p {
		p[i] = &c[i]
	}
	return &data_struct.QueryRecord{
		Pointers:  p,
		Container: c,
		Columns:   col,
	}
}

func HandlerQueryResponse(nums int, record *data_struct.QueryRecord, response *pb.GetRecordRes) {
	for i := range record.Records {
		if record.Records[i][len(record.Records[i])-1].String == "0" {
			nums--
			continue
		}

		pbRecord := &pb.QueryRecord{}
		var records []string
		for j := 0; j < len(record.Records[i])-1; j++ {
			records = append(records, record.Records[i][j].String)
		}
		pbRecord.Value = records
		response.Records = append(response.Records, pbRecord)
	}

	response.QueryFields = record.Columns[:len(record.Columns)-1]
	response.Nums = uint64(nums)
}
