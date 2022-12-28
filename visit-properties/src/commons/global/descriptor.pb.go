//
// (C) Copyright SimpliRoute.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: commons/global/descriptor.proto

package global

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// types of DB mappers
type DBFieldMapperType int32

const (
	// unset
	DBFieldMapperType_NO_TYPE DBFieldMapperType = 0
	// primary key
	DBFieldMapperType_PK DBFieldMapperType = 1
	// creation time
	DBFieldMapperType_CREATION_TIME DBFieldMapperType = 2
	// update time
	DBFieldMapperType_UPDATE_TIME DBFieldMapperType = 3
)

// Enum value maps for DBFieldMapperType.
var (
	DBFieldMapperType_name = map[int32]string{
		0: "NO_TYPE",
		1: "PK",
		2: "CREATION_TIME",
		3: "UPDATE_TIME",
	}
	DBFieldMapperType_value = map[string]int32{
		"NO_TYPE":       0,
		"PK":            1,
		"CREATION_TIME": 2,
		"UPDATE_TIME":   3,
	}
)

func (x DBFieldMapperType) Enum() *DBFieldMapperType {
	p := new(DBFieldMapperType)
	*p = x
	return p
}

func (x DBFieldMapperType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DBFieldMapperType) Descriptor() protoreflect.EnumDescriptor {
	return file_commons_global_descriptor_proto_enumTypes[0].Descriptor()
}

func (DBFieldMapperType) Type() protoreflect.EnumType {
	return &file_commons_global_descriptor_proto_enumTypes[0]
}

func (x DBFieldMapperType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DBFieldMapperType.Descriptor instead.
func (DBFieldMapperType) EnumDescriptor() ([]byte, []int) {
	return file_commons_global_descriptor_proto_rawDescGZIP(), []int{0}
}

var file_commons_global_descriptor_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         51000,
		Name:          "commons.global.descriptor.table_name",
		Tag:           "bytes,51000,opt,name=table_name",
		Filename:      "commons/global/descriptor.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*DBFieldMapperType)(nil),
		Field:         51001,
		Name:          "commons.global.descriptor.db_mapper",
		Tag:           "varint,51001,opt,name=db_mapper,enum=commons.global.descriptor.DBFieldMapperType",
		Filename:      "commons/global/descriptor.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional string table_name = 51000;
	E_TableName = &file_commons_global_descriptor_proto_extTypes[0]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional commons.global.descriptor.DBFieldMapperType db_mapper = 51001;
	E_DbMapper = &file_commons_global_descriptor_proto_extTypes[1]
)

var File_commons_global_descriptor_proto protoreflect.FileDescriptor

var file_commons_global_descriptor_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x2e, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x4c,
	0x0a, 0x11, 0x44, 0x42, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x4f, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00,
	0x12, 0x06, 0x0a, 0x02, 0x50, 0x4b, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x52, 0x45, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x55,
	0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x03, 0x3a, 0x40, 0x0a, 0x0a,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb8, 0x8e, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x3a, 0x6a,
	0x0a, 0x09, 0x64, 0x62, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb9, 0x8e, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x2e, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x44,
	0x42, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x08, 0x64, 0x62, 0x4d, 0x61, 0x70, 0x70, 0x65, 0x72, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69,
	0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67,
	0x6f, 0x2f, 0x76, 0x69, 0x73, 0x69, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69,
	0x65, 0x73, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x67,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commons_global_descriptor_proto_rawDescOnce sync.Once
	file_commons_global_descriptor_proto_rawDescData = file_commons_global_descriptor_proto_rawDesc
)

func file_commons_global_descriptor_proto_rawDescGZIP() []byte {
	file_commons_global_descriptor_proto_rawDescOnce.Do(func() {
		file_commons_global_descriptor_proto_rawDescData = protoimpl.X.CompressGZIP(file_commons_global_descriptor_proto_rawDescData)
	})
	return file_commons_global_descriptor_proto_rawDescData
}

var file_commons_global_descriptor_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_commons_global_descriptor_proto_goTypes = []interface{}{
	(DBFieldMapperType)(0),              // 0: commons.global.descriptor.DBFieldMapperType
	(*descriptorpb.MessageOptions)(nil), // 1: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),   // 2: google.protobuf.FieldOptions
}
var file_commons_global_descriptor_proto_depIdxs = []int32{
	1, // 0: commons.global.descriptor.table_name:extendee -> google.protobuf.MessageOptions
	2, // 1: commons.global.descriptor.db_mapper:extendee -> google.protobuf.FieldOptions
	0, // 2: commons.global.descriptor.db_mapper:type_name -> commons.global.descriptor.DBFieldMapperType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	2, // [2:3] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_commons_global_descriptor_proto_init() }
func file_commons_global_descriptor_proto_init() {
	if File_commons_global_descriptor_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_commons_global_descriptor_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_commons_global_descriptor_proto_goTypes,
		DependencyIndexes: file_commons_global_descriptor_proto_depIdxs,
		EnumInfos:         file_commons_global_descriptor_proto_enumTypes,
		ExtensionInfos:    file_commons_global_descriptor_proto_extTypes,
	}.Build()
	File_commons_global_descriptor_proto = out.File
	file_commons_global_descriptor_proto_rawDesc = nil
	file_commons_global_descriptor_proto_goTypes = nil
	file_commons_global_descriptor_proto_depIdxs = nil
}
