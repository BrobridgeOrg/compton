// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.2
// source: types/record/record.proto

package record

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DataType int32

const (
	DataType_BOOLEAN DataType = 0
	DataType_BINARY  DataType = 1
	DataType_STRING  DataType = 2
	DataType_UINT64  DataType = 3
	DataType_INT64   DataType = 4
	DataType_FLOAT64 DataType = 5
	DataType_ARRAY   DataType = 6
	DataType_MAP     DataType = 7
	DataType_NULL    DataType = 8
	DataType_TIME    DataType = 9
)

// Enum value maps for DataType.
var (
	DataType_name = map[int32]string{
		0: "BOOLEAN",
		1: "BINARY",
		2: "STRING",
		3: "UINT64",
		4: "INT64",
		5: "FLOAT64",
		6: "ARRAY",
		7: "MAP",
		8: "NULL",
		9: "TIME",
	}
	DataType_value = map[string]int32{
		"BOOLEAN": 0,
		"BINARY":  1,
		"STRING":  2,
		"UINT64":  3,
		"INT64":   4,
		"FLOAT64": 5,
		"ARRAY":   6,
		"MAP":     7,
		"NULL":    8,
		"TIME":    9,
	}
)

func (x DataType) Enum() *DataType {
	p := new(DataType)
	*p = x
	return p
}

func (x DataType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataType) Descriptor() protoreflect.EnumDescriptor {
	return file_types_record_record_proto_enumTypes[0].Descriptor()
}

func (DataType) Type() protoreflect.EnumType {
	return &file_types_record_record_proto_enumTypes[0]
}

func (x DataType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataType.Descriptor instead.
func (DataType) EnumDescriptor() ([]byte, []int) {
	return file_types_record_record_proto_rawDescGZIP(), []int{0}
}

type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta    *structpb.Struct `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Payload *Value           `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_record_record_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_types_record_record_proto_msgTypes[0]
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
	return file_types_record_record_proto_rawDescGZIP(), []int{0}
}

func (x *Record) GetMeta() *structpb.Struct {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Record) GetPayload() *Value {
	if x != nil {
		return x.Payload
	}
	return nil
}

type Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value *Value `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Field) Reset() {
	*x = Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_record_record_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field) ProtoMessage() {}

func (x *Field) ProtoReflect() protoreflect.Message {
	mi := &file_types_record_record_proto_msgTypes[1]
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
	return file_types_record_record_proto_rawDescGZIP(), []int{1}
}

func (x *Field) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Field) GetValue() *Value {
	if x != nil {
		return x.Value
	}
	return nil
}

type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      DataType               `protobuf:"varint,1,opt,name=type,proto3,enum=compton.types.record.DataType" json:"type,omitempty"`
	Value     []byte                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Map       *MapValue              `protobuf:"bytes,3,opt,name=map,proto3" json:"map,omitempty"`
	Array     *ArrayValue            `protobuf:"bytes,4,opt,name=array,proto3" json:"array,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_record_record_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_types_record_record_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_types_record_record_proto_rawDescGZIP(), []int{2}
}

func (x *Value) GetType() DataType {
	if x != nil {
		return x.Type
	}
	return DataType_BOOLEAN
}

func (x *Value) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Value) GetMap() *MapValue {
	if x != nil {
		return x.Map
	}
	return nil
}

func (x *Value) GetArray() *ArrayValue {
	if x != nil {
		return x.Array
	}
	return nil
}

func (x *Value) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type MapValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []*Field `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *MapValue) Reset() {
	*x = MapValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_record_record_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapValue) ProtoMessage() {}

func (x *MapValue) ProtoReflect() protoreflect.Message {
	mi := &file_types_record_record_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapValue.ProtoReflect.Descriptor instead.
func (*MapValue) Descriptor() ([]byte, []int) {
	return file_types_record_record_proto_rawDescGZIP(), []int{3}
}

func (x *MapValue) GetFields() []*Field {
	if x != nil {
		return x.Fields
	}
	return nil
}

type ArrayValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Elements []*Value `protobuf:"bytes,1,rep,name=elements,proto3" json:"elements,omitempty"`
}

func (x *ArrayValue) Reset() {
	*x = ArrayValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_record_record_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArrayValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArrayValue) ProtoMessage() {}

func (x *ArrayValue) ProtoReflect() protoreflect.Message {
	mi := &file_types_record_record_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArrayValue.ProtoReflect.Descriptor instead.
func (*ArrayValue) Descriptor() ([]byte, []int) {
	return file_types_record_record_proto_rawDescGZIP(), []int{4}
}

func (x *ArrayValue) GetElements() []*Value {
	if x != nil {
		return x.Elements
	}
	return nil
}

var File_types_record_record_proto protoreflect.FileDescriptor

var file_types_record_record_proto_rawDesc = []byte{
	0x0a, 0x19, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2f, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x63, 0x6f, 0x6d,
	0x70, 0x74, 0x6f, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x6c, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x2b, 0x0a, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x35, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x74,
	0x6f, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x4e,
	0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x74, 0x6f, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xf5,
	0x01, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x74, 0x6f, 0x6e,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x30, 0x0a, 0x03, 0x6d, 0x61, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x74, 0x6f, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x4d, 0x61, 0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x03, 0x6d, 0x61, 0x70, 0x12, 0x36, 0x0a, 0x05, 0x61, 0x72, 0x72, 0x61, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x74, 0x6f, 0x6e, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x41, 0x72, 0x72, 0x61, 0x79,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x61, 0x72, 0x72, 0x61, 0x79, 0x12, 0x38, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x3f, 0x0a, 0x08, 0x4d, 0x61, 0x70, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x74, 0x6f, 0x6e, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x45, 0x0a, 0x0a, 0x41, 0x72, 0x72, 0x61, 0x79,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x74, 0x6f,
	0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2a, 0x7b,
	0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x4f,
	0x4f, 0x4c, 0x45, 0x41, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x49, 0x4e, 0x41, 0x52,
	0x59, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12,
	0x0a, 0x0a, 0x06, 0x55, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x49,
	0x4e, 0x54, 0x36, 0x34, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x36,
	0x34, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x52, 0x52, 0x41, 0x59, 0x10, 0x06, 0x12, 0x07,
	0x0a, 0x03, 0x4d, 0x41, 0x50, 0x10, 0x07, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x55, 0x4c, 0x4c, 0x10,
	0x08, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x09, 0x42, 0x2e, 0x5a, 0x2c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x72, 0x6f, 0x62, 0x72, 0x69,
	0x64, 0x67, 0x65, 0x4f, 0x72, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x74, 0x6f, 0x6e, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_types_record_record_proto_rawDescOnce sync.Once
	file_types_record_record_proto_rawDescData = file_types_record_record_proto_rawDesc
)

func file_types_record_record_proto_rawDescGZIP() []byte {
	file_types_record_record_proto_rawDescOnce.Do(func() {
		file_types_record_record_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_record_record_proto_rawDescData)
	})
	return file_types_record_record_proto_rawDescData
}

var file_types_record_record_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_types_record_record_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_types_record_record_proto_goTypes = []interface{}{
	(DataType)(0),                 // 0: compton.types.record.DataType
	(*Record)(nil),                // 1: compton.types.record.Record
	(*Field)(nil),                 // 2: compton.types.record.Field
	(*Value)(nil),                 // 3: compton.types.record.Value
	(*MapValue)(nil),              // 4: compton.types.record.MapValue
	(*ArrayValue)(nil),            // 5: compton.types.record.ArrayValue
	(*structpb.Struct)(nil),       // 6: google.protobuf.Struct
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_types_record_record_proto_depIdxs = []int32{
	6, // 0: compton.types.record.Record.meta:type_name -> google.protobuf.Struct
	3, // 1: compton.types.record.Record.payload:type_name -> compton.types.record.Value
	3, // 2: compton.types.record.Field.value:type_name -> compton.types.record.Value
	0, // 3: compton.types.record.Value.type:type_name -> compton.types.record.DataType
	4, // 4: compton.types.record.Value.map:type_name -> compton.types.record.MapValue
	5, // 5: compton.types.record.Value.array:type_name -> compton.types.record.ArrayValue
	7, // 6: compton.types.record.Value.timestamp:type_name -> google.protobuf.Timestamp
	2, // 7: compton.types.record.MapValue.fields:type_name -> compton.types.record.Field
	3, // 8: compton.types.record.ArrayValue.elements:type_name -> compton.types.record.Value
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_types_record_record_proto_init() }
func file_types_record_record_proto_init() {
	if File_types_record_record_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_record_record_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_types_record_record_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_types_record_record_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
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
		file_types_record_record_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapValue); i {
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
		file_types_record_record_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArrayValue); i {
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
			RawDescriptor: file_types_record_record_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_record_record_proto_goTypes,
		DependencyIndexes: file_types_record_record_proto_depIdxs,
		EnumInfos:         file_types_record_record_proto_enumTypes,
		MessageInfos:      file_types_record_record_proto_msgTypes,
	}.Build()
	File_types_record_record_proto = out.File
	file_types_record_record_proto_rawDesc = nil
	file_types_record_record_proto_goTypes = nil
	file_types_record_record_proto_depIdxs = nil
}
