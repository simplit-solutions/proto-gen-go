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
// source: indexer/elastic-search-indexer.proto

package indexer

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type OperationType int32

const (
	OperationType_OPERATION_TYPE_NOT_SET OperationType = 0
	OperationType_SET                    OperationType = 1
	OperationType_UPDATE                 OperationType = 2
)

// Enum value maps for OperationType.
var (
	OperationType_name = map[int32]string{
		0: "OPERATION_TYPE_NOT_SET",
		1: "SET",
		2: "UPDATE",
	}
	OperationType_value = map[string]int32{
		"OPERATION_TYPE_NOT_SET": 0,
		"SET":                    1,
		"UPDATE":                 2,
	}
)

func (x OperationType) Enum() *OperationType {
	p := new(OperationType)
	*p = x
	return p
}

func (x OperationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationType) Descriptor() protoreflect.EnumDescriptor {
	return file_indexer_elastic_search_indexer_proto_enumTypes[0].Descriptor()
}

func (OperationType) Type() protoreflect.EnumType {
	return &file_indexer_elastic_search_indexer_proto_enumTypes[0]
}

func (x OperationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationType.Descriptor instead.
func (OperationType) EnumDescriptor() ([]byte, []int) {
	return file_indexer_elastic_search_indexer_proto_rawDescGZIP(), []int{0}
}

type IndexDocumentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperationType OperationType        `protobuf:"varint,1,opt,name=operation_type,json=operationType,proto3,enum=simpliroute.indexer.OperationType" json:"operation_type,omitempty"`
	Index         string               `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
	Documents     []*IndexableDocument `protobuf:"bytes,3,rep,name=documents,proto3" json:"documents,omitempty"`
}

func (x *IndexDocumentsRequest) Reset() {
	*x = IndexDocumentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indexer_elastic_search_indexer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexDocumentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexDocumentsRequest) ProtoMessage() {}

func (x *IndexDocumentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_indexer_elastic_search_indexer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexDocumentsRequest.ProtoReflect.Descriptor instead.
func (*IndexDocumentsRequest) Descriptor() ([]byte, []int) {
	return file_indexer_elastic_search_indexer_proto_rawDescGZIP(), []int{0}
}

func (x *IndexDocumentsRequest) GetOperationType() OperationType {
	if x != nil {
		return x.OperationType
	}
	return OperationType_OPERATION_TYPE_NOT_SET
}

func (x *IndexDocumentsRequest) GetIndex() string {
	if x != nil {
		return x.Index
	}
	return ""
}

func (x *IndexDocumentsRequest) GetDocuments() []*IndexableDocument {
	if x != nil {
		return x.Documents
	}
	return nil
}

type IndexableDocument struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Supports the same operation on multiple IDs
	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	// Json serialized document contents
	Contents string `protobuf:"bytes,2,opt,name=contents,proto3" json:"contents,omitempty"`
}

func (x *IndexableDocument) Reset() {
	*x = IndexableDocument{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indexer_elastic_search_indexer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexableDocument) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexableDocument) ProtoMessage() {}

func (x *IndexableDocument) ProtoReflect() protoreflect.Message {
	mi := &file_indexer_elastic_search_indexer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexableDocument.ProtoReflect.Descriptor instead.
func (*IndexableDocument) Descriptor() ([]byte, []int) {
	return file_indexer_elastic_search_indexer_proto_rawDescGZIP(), []int{1}
}

func (x *IndexableDocument) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *IndexableDocument) GetContents() string {
	if x != nil {
		return x.Contents
	}
	return ""
}

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Healthcheck endpoint for load balancer
	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indexer_elastic_search_indexer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_indexer_elastic_search_indexer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_indexer_elastic_search_indexer_proto_rawDescGZIP(), []int{2}
}

func (x *Ping) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

var File_indexer_elastic_search_indexer_proto protoreflect.FileDescriptor

var file_indexer_elastic_search_indexer_proto_rawDesc = []byte{
	0x0a, 0x24, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2f, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69,
	0x63, 0x2d, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2d, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x01, 0x0a, 0x15, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x49, 0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x73, 0x69, 0x6d,
	0x70, 0x6c, 0x69, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x44, 0x0a, 0x09, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x09,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x41, 0x0a, 0x11, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x1a, 0x0a, 0x04,
	0x50, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x2a, 0x40, 0x0a, 0x0d, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x4f, 0x50, 0x45,
	0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x45, 0x54, 0x10, 0x01, 0x12, 0x0a,
	0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x02, 0x32, 0x50, 0x0a, 0x0e, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x19, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x42, 0x45, 0x0a, 0x17,
	0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x42, 0x07, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72,
	0x48, 0x01, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_indexer_elastic_search_indexer_proto_rawDescOnce sync.Once
	file_indexer_elastic_search_indexer_proto_rawDescData = file_indexer_elastic_search_indexer_proto_rawDesc
)

func file_indexer_elastic_search_indexer_proto_rawDescGZIP() []byte {
	file_indexer_elastic_search_indexer_proto_rawDescOnce.Do(func() {
		file_indexer_elastic_search_indexer_proto_rawDescData = protoimpl.X.CompressGZIP(file_indexer_elastic_search_indexer_proto_rawDescData)
	})
	return file_indexer_elastic_search_indexer_proto_rawDescData
}

var file_indexer_elastic_search_indexer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_indexer_elastic_search_indexer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_indexer_elastic_search_indexer_proto_goTypes = []interface{}{
	(OperationType)(0),            // 0: simpliroute.indexer.OperationType
	(*IndexDocumentsRequest)(nil), // 1: simpliroute.indexer.IndexDocumentsRequest
	(*IndexableDocument)(nil),     // 2: simpliroute.indexer.IndexableDocument
	(*Ping)(nil),                  // 3: simpliroute.indexer.Ping
	(*emptypb.Empty)(nil),         // 4: google.protobuf.Empty
}
var file_indexer_elastic_search_indexer_proto_depIdxs = []int32{
	0, // 0: simpliroute.indexer.IndexDocumentsRequest.operation_type:type_name -> simpliroute.indexer.OperationType
	2, // 1: simpliroute.indexer.IndexDocumentsRequest.documents:type_name -> simpliroute.indexer.IndexableDocument
	4, // 2: simpliroute.indexer.IndexerService.GetPing:input_type -> google.protobuf.Empty
	3, // 3: simpliroute.indexer.IndexerService.GetPing:output_type -> simpliroute.indexer.Ping
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_indexer_elastic_search_indexer_proto_init() }
func file_indexer_elastic_search_indexer_proto_init() {
	if File_indexer_elastic_search_indexer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_indexer_elastic_search_indexer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexDocumentsRequest); i {
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
		file_indexer_elastic_search_indexer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexableDocument); i {
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
		file_indexer_elastic_search_indexer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
			RawDescriptor: file_indexer_elastic_search_indexer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_indexer_elastic_search_indexer_proto_goTypes,
		DependencyIndexes: file_indexer_elastic_search_indexer_proto_depIdxs,
		EnumInfos:         file_indexer_elastic_search_indexer_proto_enumTypes,
		MessageInfos:      file_indexer_elastic_search_indexer_proto_msgTypes,
	}.Build()
	File_indexer_elastic_search_indexer_proto = out.File
	file_indexer_elastic_search_indexer_proto_rawDesc = nil
	file_indexer_elastic_search_indexer_proto_goTypes = nil
	file_indexer_elastic_search_indexer_proto_depIdxs = nil
}
