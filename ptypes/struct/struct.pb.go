// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/golang/protobuf/ptypes/struct/struct.proto

package structpb

import (
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoregistry "github.com/golang/protobuf/v2/reflect/protoregistry"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	known "github.com/golang/protobuf/v2/types/known"
	sync "sync"
)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

// Symbols defined in public import of google/protobuf/struct.proto

type NullValue = known.NullValue

const NullValue_NULL_VALUE = known.NullValue_NULL_VALUE

var NullValue_name = known.NullValue_name
var NullValue_value = known.NullValue_value

type Struct = known.Struct
type Value = known.Value
type Value_NullValue = known.Value_NullValue
type Value_NumberValue = known.Value_NumberValue
type Value_StringValue = known.Value_StringValue
type Value_BoolValue = known.Value_BoolValue
type Value_StructValue = known.Value_StructValue
type Value_ListValue = known.Value_ListValue
type ListValue = known.ListValue

var File_github_com_golang_protobuf_ptypes_struct_struct_proto protoreflect.FileDescriptor

var xxx_File_github_com_golang_protobuf_ptypes_struct_struct_proto_rawDesc = []byte{
	0x0a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x70, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x3b, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x70, 0x62, 0x50, 0x00, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	xxx_File_github_com_golang_protobuf_ptypes_struct_struct_proto_rawDesc_once sync.Once
	xxx_File_github_com_golang_protobuf_ptypes_struct_struct_proto_rawDesc_data = xxx_File_github_com_golang_protobuf_ptypes_struct_struct_proto_rawDesc
)

func xxx_File_github_com_golang_protobuf_ptypes_struct_struct_proto_rawDescGZIP() []byte {
	xxx_File_github_com_golang_protobuf_ptypes_struct_struct_proto_rawDesc_once.Do(func() {
		xxx_File_github_com_golang_protobuf_ptypes_struct_struct_proto_rawDesc_data = protoimpl.X.CompressGZIP(xxx_File_github_com_golang_protobuf_ptypes_struct_struct_proto_rawDesc_data)
	})
	return xxx_File_github_com_golang_protobuf_ptypes_struct_struct_proto_rawDesc_data
}

var xxx_File_github_com_golang_protobuf_ptypes_struct_struct_proto_goTypes = []interface{}{}
var xxx_File_github_com_golang_protobuf_ptypes_struct_struct_proto_depIdxs = []int32{}

func init() { xxx_File_github_com_golang_protobuf_ptypes_struct_struct_proto_init() }
func xxx_File_github_com_golang_protobuf_ptypes_struct_struct_proto_init() {
	if File_github_com_golang_protobuf_ptypes_struct_struct_proto != nil {
		return
	}
	File_github_com_golang_protobuf_ptypes_struct_struct_proto = protoimpl.FileBuilder{
		RawDescriptor:     xxx_File_github_com_golang_protobuf_ptypes_struct_struct_proto_rawDesc,
		GoTypes:           xxx_File_github_com_golang_protobuf_ptypes_struct_struct_proto_goTypes,
		DependencyIndexes: xxx_File_github_com_golang_protobuf_ptypes_struct_struct_proto_depIdxs,
		FilesRegistry:     protoregistry.GlobalFiles,
		TypesRegistry:     protoregistry.GlobalTypes,
	}.Init()
	xxx_File_github_com_golang_protobuf_ptypes_struct_struct_proto_rawDesc = nil
	xxx_File_github_com_golang_protobuf_ptypes_struct_struct_proto_goTypes = nil
	xxx_File_github_com_golang_protobuf_ptypes_struct_struct_proto_depIdxs = nil
}
