// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: db_ops.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InsertMode int32

const (
	InsertMode_COMMON                  InsertMode = 0
	InsertMode_ON_DUPLICATE_KEY_UPDATE InsertMode = 1
	InsertMode_REPLACE                 InsertMode = 2
)

// Enum value maps for InsertMode.
var (
	InsertMode_name = map[int32]string{
		0: "COMMON",
		1: "ON_DUPLICATE_KEY_UPDATE",
		2: "REPLACE",
	}
	InsertMode_value = map[string]int32{
		"COMMON":                  0,
		"ON_DUPLICATE_KEY_UPDATE": 1,
		"REPLACE":                 2,
	}
)

func (x InsertMode) Enum() *InsertMode {
	p := new(InsertMode)
	*p = x
	return p
}

func (x InsertMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InsertMode) Descriptor() protoreflect.EnumDescriptor {
	return file_db_ops_proto_enumTypes[0].Descriptor()
}

func (InsertMode) Type() protoreflect.EnumType {
	return &file_db_ops_proto_enumTypes[0]
}

func (x InsertMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InsertMode.Descriptor instead.
func (InsertMode) EnumDescriptor() ([]byte, []int) {
	return file_db_ops_proto_rawDescGZIP(), []int{0}
}

type Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Field) Reset() {
	*x = Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_ops_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field) ProtoMessage() {}

func (x *Field) ProtoReflect() protoreflect.Message {
	mi := &file_db_ops_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Field.ProtoReflect.Descriptor instead.
func (*Field) Descriptor() ([]byte, []int) {
	return file_db_ops_proto_rawDescGZIP(), []int{0}
}

func (x *Field) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Field) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type QueryRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []string `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *QueryRecord) Reset() {
	*x = QueryRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_ops_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRecord) ProtoMessage() {}

func (x *QueryRecord) ProtoReflect() protoreflect.Message {
	mi := &file_db_ops_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRecord.ProtoReflect.Descriptor instead.
func (*QueryRecord) Descriptor() ([]byte, []int) {
	return file_db_ops_proto_rawDescGZIP(), []int{1}
}

func (x *QueryRecord) GetValue() []string {
	if x != nil {
		return x.Value
	}
	return nil
}

type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []*Field `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_ops_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_db_ops_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_db_ops_proto_rawDescGZIP(), []int{2}
}

func (x *Record) GetFields() []*Field {
	if x != nil {
		return x.Fields
	}
	return nil
}

type InsertReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BusinessId string     `protobuf:"bytes,1,opt,name=business_id,json=businessId,proto3" json:"business_id,omitempty"` // 用于指定库表
	Records    []*Record  `protobuf:"bytes,2,rep,name=records,proto3" json:"records,omitempty"`
	Mode       InsertMode `protobuf:"varint,3,opt,name=mode,proto3,enum=db_common_ops.InsertMode" json:"mode,omitempty"`
}

func (x *InsertReq) Reset() {
	*x = InsertReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_ops_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertReq) ProtoMessage() {}

func (x *InsertReq) ProtoReflect() protoreflect.Message {
	mi := &file_db_ops_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertReq.ProtoReflect.Descriptor instead.
func (*InsertReq) Descriptor() ([]byte, []int) {
	return file_db_ops_proto_rawDescGZIP(), []int{3}
}

func (x *InsertReq) GetBusinessId() string {
	if x != nil {
		return x.BusinessId
	}
	return ""
}

func (x *InsertReq) GetRecords() []*Record {
	if x != nil {
		return x.Records
	}
	return nil
}

func (x *InsertReq) GetMode() InsertMode {
	if x != nil {
		return x.Mode
	}
	return InsertMode_COMMON
}

type UpdateItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdateFields []*Field `protobuf:"bytes,1,rep,name=update_fields,json=updateFields,proto3" json:"update_fields,omitempty"`
	WhereFields  []*Field `protobuf:"bytes,2,rep,name=where_fields,json=whereFields,proto3" json:"where_fields,omitempty"`
}

func (x *UpdateItem) Reset() {
	*x = UpdateItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_ops_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateItem) ProtoMessage() {}

func (x *UpdateItem) ProtoReflect() protoreflect.Message {
	mi := &file_db_ops_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateItem.ProtoReflect.Descriptor instead.
func (*UpdateItem) Descriptor() ([]byte, []int) {
	return file_db_ops_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateItem) GetUpdateFields() []*Field {
	if x != nil {
		return x.UpdateFields
	}
	return nil
}

func (x *UpdateItem) GetWhereFields() []*Field {
	if x != nil {
		return x.WhereFields
	}
	return nil
}

type UpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BusinessId string        `protobuf:"bytes,1,opt,name=business_id,json=businessId,proto3" json:"business_id,omitempty"`
	Items      []*UpdateItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *UpdateReq) Reset() {
	*x = UpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_ops_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReq) ProtoMessage() {}

func (x *UpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_db_ops_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReq.ProtoReflect.Descriptor instead.
func (*UpdateReq) Descriptor() ([]byte, []int) {
	return file_db_ops_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateReq) GetBusinessId() string {
	if x != nil {
		return x.BusinessId
	}
	return ""
}

func (x *UpdateReq) GetItems() []*UpdateItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type DeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BusinessId  string   `protobuf:"bytes,1,opt,name=business_id,json=businessId,proto3" json:"business_id,omitempty"`
	WhereFields []*Field `protobuf:"bytes,2,rep,name=where_fields,json=whereFields,proto3" json:"where_fields,omitempty"`
}

func (x *DeleteReq) Reset() {
	*x = DeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_ops_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReq) ProtoMessage() {}

func (x *DeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_db_ops_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReq.ProtoReflect.Descriptor instead.
func (*DeleteReq) Descriptor() ([]byte, []int) {
	return file_db_ops_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteReq) GetBusinessId() string {
	if x != nil {
		return x.BusinessId
	}
	return ""
}

func (x *DeleteReq) GetWhereFields() []*Field {
	if x != nil {
		return x.WhereFields
	}
	return nil
}

type DeleteRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrMsg  string `protobuf:"bytes,1,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`     // 失败信息
	ErrCode uint32 `protobuf:"varint,2,opt,name=err_code,json=errCode,proto3" json:"err_code,omitempty"` // 错误码
}

func (x *DeleteRes) Reset() {
	*x = DeleteRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_ops_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRes) ProtoMessage() {}

func (x *DeleteRes) ProtoReflect() protoreflect.Message {
	mi := &file_db_ops_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRes.ProtoReflect.Descriptor instead.
func (*DeleteRes) Descriptor() ([]byte, []int) {
	return file_db_ops_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteRes) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

func (x *DeleteRes) GetErrCode() uint32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

// ColumnSort
//
//	string column_name =1; 需要排序的字段
//	bool desc = 2;     desc降序 填 true
type ColumnSort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ColumnName string `protobuf:"bytes,1,opt,name=column_name,json=columnName,proto3" json:"column_name,omitempty"`
	Desc       bool   `protobuf:"varint,2,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *ColumnSort) Reset() {
	*x = ColumnSort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_ops_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ColumnSort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColumnSort) ProtoMessage() {}

func (x *ColumnSort) ProtoReflect() protoreflect.Message {
	mi := &file_db_ops_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColumnSort.ProtoReflect.Descriptor instead.
func (*ColumnSort) Descriptor() ([]byte, []int) {
	return file_db_ops_proto_rawDescGZIP(), []int{8}
}

func (x *ColumnSort) GetColumnName() string {
	if x != nil {
		return x.ColumnName
	}
	return ""
}

func (x *ColumnSort) GetDesc() bool {
	if x != nil {
		return x.Desc
	}
	return false
}

// WhereStruct
//
//	string expression= 1; 表达式，如"(A=? AND B=?) OR (C=? AND D=?)"
//	repeated string immediate = 2; 立即数，如[]string{"1", "2", "3", "4",}
//	                               注意，如需要使用'IN'，请将表达式设为"A IN (?,?,?)"，
//	                               这时立即数为[]string{"1", "2", "3",}
type WhereStruct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Expression string   `protobuf:"bytes,1,opt,name=expression,proto3" json:"expression,omitempty"`
	Immediate  []string `protobuf:"bytes,2,rep,name=immediate,proto3" json:"immediate,omitempty"`
}

func (x *WhereStruct) Reset() {
	*x = WhereStruct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_ops_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WhereStruct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhereStruct) ProtoMessage() {}

func (x *WhereStruct) ProtoReflect() protoreflect.Message {
	mi := &file_db_ops_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhereStruct.ProtoReflect.Descriptor instead.
func (*WhereStruct) Descriptor() ([]byte, []int) {
	return file_db_ops_proto_rawDescGZIP(), []int{9}
}

func (x *WhereStruct) GetExpression() string {
	if x != nil {
		return x.Expression
	}
	return ""
}

func (x *WhereStruct) GetImmediate() []string {
	if x != nil {
		return x.Immediate
	}
	return nil
}

type GetRecordReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BusinessId     string        `protobuf:"bytes,1,opt,name=business_id,json=businessId,proto3" json:"business_id,omitempty"`
	Pos            uint32        `protobuf:"varint,2,opt,name=pos,proto3" json:"pos,omitempty"`                                            // 当前查询位置, 初始为0
	Limit          uint32        `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`                                        // 每页记录数
	QueryFields    []string      `protobuf:"bytes,4,rep,name=query_fields,json=queryFields,proto3" json:"query_fields,omitempty"`          // 需要查询的字段（不填返回所有字段）
	WhereCondition *WhereStruct  `protobuf:"bytes,5,opt,name=where_condition,json=whereCondition,proto3" json:"where_condition,omitempty"` // where语句
	Sorts          []*ColumnSort `protobuf:"bytes,6,rep,name=sorts,proto3" json:"sorts,omitempty"`                                         // 排序字段
}

func (x *GetRecordReq) Reset() {
	*x = GetRecordReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_ops_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecordReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecordReq) ProtoMessage() {}

func (x *GetRecordReq) ProtoReflect() protoreflect.Message {
	mi := &file_db_ops_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecordReq.ProtoReflect.Descriptor instead.
func (*GetRecordReq) Descriptor() ([]byte, []int) {
	return file_db_ops_proto_rawDescGZIP(), []int{10}
}

func (x *GetRecordReq) GetBusinessId() string {
	if x != nil {
		return x.BusinessId
	}
	return ""
}

func (x *GetRecordReq) GetPos() uint32 {
	if x != nil {
		return x.Pos
	}
	return 0
}

func (x *GetRecordReq) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetRecordReq) GetQueryFields() []string {
	if x != nil {
		return x.QueryFields
	}
	return nil
}

func (x *GetRecordReq) GetWhereCondition() *WhereStruct {
	if x != nil {
		return x.WhereCondition
	}
	return nil
}

func (x *GetRecordReq) GetSorts() []*ColumnSort {
	if x != nil {
		return x.Sorts
	}
	return nil
}

type GetRecordRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrMsg      string         `protobuf:"bytes,1,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`                // 失败信息
	ErrCode     uint32         `protobuf:"varint,2,opt,name=err_code,json=errCode,proto3" json:"err_code,omitempty"`            // 错误码
	Nums        uint64         `protobuf:"varint,3,opt,name=nums,proto3" json:"nums,omitempty"`                                 // 总数
	QueryFields []string       `protobuf:"bytes,4,rep,name=query_fields,json=queryFields,proto3" json:"query_fields,omitempty"` // 查询的字段
	Records     []*QueryRecord `protobuf:"bytes,5,rep,name=records,proto3" json:"records,omitempty"`                            // 返回的具体记录
}

func (x *GetRecordRes) Reset() {
	*x = GetRecordRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_ops_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecordRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecordRes) ProtoMessage() {}

func (x *GetRecordRes) ProtoReflect() protoreflect.Message {
	mi := &file_db_ops_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecordRes.ProtoReflect.Descriptor instead.
func (*GetRecordRes) Descriptor() ([]byte, []int) {
	return file_db_ops_proto_rawDescGZIP(), []int{11}
}

func (x *GetRecordRes) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

func (x *GetRecordRes) GetErrCode() uint32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *GetRecordRes) GetNums() uint64 {
	if x != nil {
		return x.Nums
	}
	return 0
}

func (x *GetRecordRes) GetQueryFields() []string {
	if x != nil {
		return x.QueryFields
	}
	return nil
}

func (x *GetRecordRes) GetRecords() []*QueryRecord {
	if x != nil {
		return x.Records
	}
	return nil
}

// 全部成功或全部失败
type Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrMsg  string `protobuf:"bytes,1,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`     // 失败信息
	ErrCode uint32 `protobuf:"varint,2,opt,name=err_code,json=errCode,proto3" json:"err_code,omitempty"` // 错误码
}

func (x *Res) Reset() {
	*x = Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_ops_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Res) ProtoMessage() {}

func (x *Res) ProtoReflect() protoreflect.Message {
	mi := &file_db_ops_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Res.ProtoReflect.Descriptor instead.
func (*Res) Descriptor() ([]byte, []int) {
	return file_db_ops_proto_rawDescGZIP(), []int{12}
}

func (x *Res) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

func (x *Res) GetErrCode() uint32 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

var File_db_ops_proto protoreflect.FileDescriptor

var file_db_ops_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x64, 0x62, 0x5f, 0x6f, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x64, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x73, 0x22, 0x2f, 0x0a,
	0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x23,
	0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x36, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x2c, 0x0a,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x64, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x73, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x8c, 0x01, 0x0a, 0x09,
	0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x07, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x64, 0x62,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x73, 0x2e, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x2d, 0x0a, 0x04, 0x6d,
	0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x64, 0x62, 0x5f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x73, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74,
	0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x80, 0x01, 0x0a, 0x0a, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x39, 0x0a, 0x0d, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x64, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x73,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x12, 0x37, 0x0a, 0x0c, 0x77, 0x68, 0x65, 0x72, 0x65, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x62, 0x5f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x73, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x52, 0x0b, 0x77, 0x68, 0x65, 0x72, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x5d, 0x0a,
	0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64, 0x62, 0x5f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x65, 0x0a, 0x09,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x0c, 0x77, 0x68,
	0x65, 0x72, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x64, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x73,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x0b, 0x77, 0x68, 0x65, 0x72, 0x65, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x22, 0x3f, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x12, 0x17, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x72, 0x72,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x65, 0x72, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x22, 0x41, 0x0a, 0x0a, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x53, 0x6f,
	0x72, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x22, 0x4b, 0x0a, 0x0b, 0x57, 0x68, 0x65, 0x72, 0x65,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6d, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x74, 0x65, 0x22, 0xf0, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x12, 0x43, 0x0a, 0x0f, 0x77, 0x68, 0x65, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x62, 0x5f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x73, 0x2e, 0x57, 0x68, 0x65, 0x72, 0x65,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0e, 0x77, 0x68, 0x65, 0x72, 0x65, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x05, 0x73, 0x6f, 0x72, 0x74, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5f, 0x6f, 0x70, 0x73, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x53, 0x6f, 0x72, 0x74,
	0x52, 0x05, 0x73, 0x6f, 0x72, 0x74, 0x73, 0x22, 0xaf, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x5f,
	0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73,
	0x67, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x75, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x73,
	0x12, 0x21, 0x0a, 0x0c, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x5f, 0x6f, 0x70, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x39, 0x0a, 0x03, 0x52, 0x65, 0x73,
	0x12, 0x17, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x72, 0x72,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x65, 0x72, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x2a, 0x42, 0x0a, 0x0a, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x4d, 0x6f,
	0x64, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x1b,
	0x0a, 0x17, 0x4f, 0x4e, 0x5f, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x4b,
	0x45, 0x59, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x52,
	0x45, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x10, 0x02, 0x32, 0xfd, 0x01, 0x0a, 0x0c, 0x44, 0x62, 0x4f,
	0x70, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x49, 0x6e, 0x73,
	0x65, 0x72, 0x74, 0x12, 0x18, 0x2e, 0x64, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f,
	0x6f, 0x70, 0x73, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e,
	0x64, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x73, 0x2e, 0x52, 0x65,
	0x73, 0x12, 0x36, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x64, 0x62,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x64, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5f, 0x6f, 0x70, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x64, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f,
	0x6f, 0x70, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e,
	0x64, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x73, 0x2e, 0x52, 0x65,
	0x73, 0x12, 0x45, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x1b,
	0x2e, 0x64, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x64, 0x62,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x64, 0x62, 0x5f, 0x6f, 0x70, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62,
	0x3b, 0x64, 0x62, 0x5f, 0x6f, 0x70, 0x73, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_db_ops_proto_rawDescOnce sync.Once
	file_db_ops_proto_rawDescData = file_db_ops_proto_rawDesc
)

func file_db_ops_proto_rawDescGZIP() []byte {
	file_db_ops_proto_rawDescOnce.Do(func() {
		file_db_ops_proto_rawDescData = protoimpl.X.CompressGZIP(file_db_ops_proto_rawDescData)
	})
	return file_db_ops_proto_rawDescData
}

var file_db_ops_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_db_ops_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_db_ops_proto_goTypes = []interface{}{
	(InsertMode)(0),      // 0: db_common_ops.InsertMode
	(*Field)(nil),        // 1: db_common_ops.Field
	(*QueryRecord)(nil),  // 2: db_common_ops.QueryRecord
	(*Record)(nil),       // 3: db_common_ops.Record
	(*InsertReq)(nil),    // 4: db_common_ops.InsertReq
	(*UpdateItem)(nil),   // 5: db_common_ops.UpdateItem
	(*UpdateReq)(nil),    // 6: db_common_ops.UpdateReq
	(*DeleteReq)(nil),    // 7: db_common_ops.DeleteReq
	(*DeleteRes)(nil),    // 8: db_common_ops.DeleteRes
	(*ColumnSort)(nil),   // 9: db_common_ops.ColumnSort
	(*WhereStruct)(nil),  // 10: db_common_ops.WhereStruct
	(*GetRecordReq)(nil), // 11: db_common_ops.GetRecordReq
	(*GetRecordRes)(nil), // 12: db_common_ops.GetRecordRes
	(*Res)(nil),          // 13: db_common_ops.Res
}
var file_db_ops_proto_depIdxs = []int32{
	1,  // 0: db_common_ops.Record.fields:type_name -> db_common_ops.Field
	3,  // 1: db_common_ops.InsertReq.records:type_name -> db_common_ops.Record
	0,  // 2: db_common_ops.InsertReq.mode:type_name -> db_common_ops.InsertMode
	1,  // 3: db_common_ops.UpdateItem.update_fields:type_name -> db_common_ops.Field
	1,  // 4: db_common_ops.UpdateItem.where_fields:type_name -> db_common_ops.Field
	5,  // 5: db_common_ops.UpdateReq.items:type_name -> db_common_ops.UpdateItem
	1,  // 6: db_common_ops.DeleteReq.where_fields:type_name -> db_common_ops.Field
	10, // 7: db_common_ops.GetRecordReq.where_condition:type_name -> db_common_ops.WhereStruct
	9,  // 8: db_common_ops.GetRecordReq.sorts:type_name -> db_common_ops.ColumnSort
	2,  // 9: db_common_ops.GetRecordRes.records:type_name -> db_common_ops.QueryRecord
	4,  // 10: db_common_ops.DbOpsService.Insert:input_type -> db_common_ops.InsertReq
	6,  // 11: db_common_ops.DbOpsService.Update:input_type -> db_common_ops.UpdateReq
	7,  // 12: db_common_ops.DbOpsService.Delete:input_type -> db_common_ops.DeleteReq
	11, // 13: db_common_ops.DbOpsService.GetRecord:input_type -> db_common_ops.GetRecordReq
	13, // 14: db_common_ops.DbOpsService.Insert:output_type -> db_common_ops.Res
	13, // 15: db_common_ops.DbOpsService.Update:output_type -> db_common_ops.Res
	13, // 16: db_common_ops.DbOpsService.Delete:output_type -> db_common_ops.Res
	12, // 17: db_common_ops.DbOpsService.GetRecord:output_type -> db_common_ops.GetRecordRes
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_db_ops_proto_init() }
func file_db_ops_proto_init() {
	if File_db_ops_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_db_ops_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Field); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_db_ops_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRecord); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_db_ops_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_db_ops_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_db_ops_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_db_ops_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_db_ops_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_db_ops_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_db_ops_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ColumnSort); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_db_ops_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WhereStruct); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_db_ops_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecordReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_db_ops_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecordRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_db_ops_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Res); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_db_ops_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_db_ops_proto_goTypes,
		DependencyIndexes: file_db_ops_proto_depIdxs,
		EnumInfos:         file_db_ops_proto_enumTypes,
		MessageInfos:      file_db_ops_proto_msgTypes,
	}.Build()
	File_db_ops_proto = out.File
	file_db_ops_proto_rawDesc = nil
	file_db_ops_proto_goTypes = nil
	file_db_ops_proto_depIdxs = nil
}
