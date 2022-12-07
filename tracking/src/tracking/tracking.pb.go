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
// source: tracking/tracking.proto

package tracking

import (
	proto "github.com/golang/protobuf/proto"
	global "gitlab.com/simpliroute/proto-gen-go/tracking/src/commons/global"
	gps "gitlab.com/simpliroute/proto-gen-go/tracking/src/commons/gps"
	date "google.golang.org/genproto/googleapis/type/date"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type StoreTrackingLocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location *gps.TrackingLocation `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *StoreTrackingLocationRequest) Reset() {
	*x = StoreTrackingLocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracking_tracking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreTrackingLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreTrackingLocationRequest) ProtoMessage() {}

func (x *StoreTrackingLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tracking_tracking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreTrackingLocationRequest.ProtoReflect.Descriptor instead.
func (*StoreTrackingLocationRequest) Descriptor() ([]byte, []int) {
	return file_tracking_tracking_proto_rawDescGZIP(), []int{0}
}

func (x *StoreTrackingLocationRequest) GetLocation() *gps.TrackingLocation {
	if x != nil {
		return x.Location
	}
	return nil
}

type GetTrackingLocationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Period    *global.Period `protobuf:"bytes,1,opt,name=period,proto3" json:"period,omitempty"`
	OwnerType gps.OwnerType  `protobuf:"varint,2,opt,name=owner_type,json=ownerType,proto3,enum=commons.gps.OwnerType" json:"owner_type,omitempty"`
	OwnerIds  []int32        `protobuf:"varint,3,rep,packed,name=owner_ids,json=ownerIds,proto3" json:"owner_ids,omitempty"`
}

func (x *GetTrackingLocationsRequest) Reset() {
	*x = GetTrackingLocationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracking_tracking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTrackingLocationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTrackingLocationsRequest) ProtoMessage() {}

func (x *GetTrackingLocationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tracking_tracking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTrackingLocationsRequest.ProtoReflect.Descriptor instead.
func (*GetTrackingLocationsRequest) Descriptor() ([]byte, []int) {
	return file_tracking_tracking_proto_rawDescGZIP(), []int{1}
}

func (x *GetTrackingLocationsRequest) GetPeriod() *global.Period {
	if x != nil {
		return x.Period
	}
	return nil
}

func (x *GetTrackingLocationsRequest) GetOwnerType() gps.OwnerType {
	if x != nil {
		return x.OwnerType
	}
	return gps.OwnerType_OWNER_TYPE_NOT_SET
}

func (x *GetTrackingLocationsRequest) GetOwnerIds() []int32 {
	if x != nil {
		return x.OwnerIds
	}
	return nil
}

type GetTrackingLocationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Locations []*gps.TrackingLocation `protobuf:"bytes,1,rep,name=locations,proto3" json:"locations,omitempty"`
}

func (x *GetTrackingLocationsResponse) Reset() {
	*x = GetTrackingLocationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracking_tracking_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTrackingLocationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTrackingLocationsResponse) ProtoMessage() {}

func (x *GetTrackingLocationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tracking_tracking_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTrackingLocationsResponse.ProtoReflect.Descriptor instead.
func (*GetTrackingLocationsResponse) Descriptor() ([]byte, []int) {
	return file_tracking_tracking_proto_rawDescGZIP(), []int{2}
}

func (x *GetTrackingLocationsResponse) GetLocations() []*gps.TrackingLocation {
	if x != nil {
		return x.Locations
	}
	return nil
}

type GetLatestTrackingStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date      *date.Date    `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	OwnerType gps.OwnerType `protobuf:"varint,2,opt,name=owner_type,json=ownerType,proto3,enum=commons.gps.OwnerType" json:"owner_type,omitempty"`
	OwnerIds  []int32       `protobuf:"varint,3,rep,packed,name=owner_ids,json=ownerIds,proto3" json:"owner_ids,omitempty"`
}

func (x *GetLatestTrackingStatusRequest) Reset() {
	*x = GetLatestTrackingStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracking_tracking_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLatestTrackingStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLatestTrackingStatusRequest) ProtoMessage() {}

func (x *GetLatestTrackingStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tracking_tracking_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLatestTrackingStatusRequest.ProtoReflect.Descriptor instead.
func (*GetLatestTrackingStatusRequest) Descriptor() ([]byte, []int) {
	return file_tracking_tracking_proto_rawDescGZIP(), []int{3}
}

func (x *GetLatestTrackingStatusRequest) GetDate() *date.Date {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *GetLatestTrackingStatusRequest) GetOwnerType() gps.OwnerType {
	if x != nil {
		return x.OwnerType
	}
	return gps.OwnerType_OWNER_TYPE_NOT_SET
}

func (x *GetLatestTrackingStatusRequest) GetOwnerIds() []int32 {
	if x != nil {
		return x.OwnerIds
	}
	return nil
}

type GetLatestTrackingStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status []*gps.LatestTrackingStatus `protobuf:"bytes,1,rep,name=status,proto3" json:"status,omitempty"`
}

func (x *GetLatestTrackingStatusResponse) Reset() {
	*x = GetLatestTrackingStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracking_tracking_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLatestTrackingStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLatestTrackingStatusResponse) ProtoMessage() {}

func (x *GetLatestTrackingStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tracking_tracking_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLatestTrackingStatusResponse.ProtoReflect.Descriptor instead.
func (*GetLatestTrackingStatusResponse) Descriptor() ([]byte, []int) {
	return file_tracking_tracking_proto_rawDescGZIP(), []int{4}
}

func (x *GetLatestTrackingStatusResponse) GetStatus() []*gps.LatestTrackingStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_tracking_tracking_proto protoreflect.FileDescriptor

var file_tracking_tracking_proto_rawDesc = []byte{
	0x0a, 0x17, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x73, 0x69, 0x6d, 0x70, 0x6c,
	0x69, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x1a,
	0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x70,
	0x73, 0x2f, 0x67, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x1c, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x70, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x63,
	0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa1, 0x01, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x54, 0x72,
	0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x06,
	0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x35, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x70, 0x73, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0x5b, 0x0a, 0x1c, 0x47, 0x65,
	0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x09, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x70, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x63,
	0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x9b, 0x01, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x4c,
	0x61, 0x74, 0x65, 0x73, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x35, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x67, 0x70, 0x73, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0x5c, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65,
	0x73, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x67, 0x70, 0x73, 0x2e, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x32, 0xa5, 0x02, 0x0a, 0x17, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x7f, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x31, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x69,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x73, 0x69, 0x6d,
	0x70, 0x6c, 0x69, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x88, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x54, 0x72,
	0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x34, 0x2e, 0x73,
	0x69, 0x6d, 0x70, 0x6c, 0x69, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x35, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74,
	0x65, 0x73, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x51, 0x0a, 0x21, 0x63,
	0x6f, 0x6d, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67,
	0x42, 0x08, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x48, 0x01, 0x5a, 0x20, 0x67, 0x69,
	0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tracking_tracking_proto_rawDescOnce sync.Once
	file_tracking_tracking_proto_rawDescData = file_tracking_tracking_proto_rawDesc
)

func file_tracking_tracking_proto_rawDescGZIP() []byte {
	file_tracking_tracking_proto_rawDescOnce.Do(func() {
		file_tracking_tracking_proto_rawDescData = protoimpl.X.CompressGZIP(file_tracking_tracking_proto_rawDescData)
	})
	return file_tracking_tracking_proto_rawDescData
}

var file_tracking_tracking_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_tracking_tracking_proto_goTypes = []interface{}{
	(*StoreTrackingLocationRequest)(nil),    // 0: simpliroute.tracking.StoreTrackingLocationRequest
	(*GetTrackingLocationsRequest)(nil),     // 1: simpliroute.tracking.GetTrackingLocationsRequest
	(*GetTrackingLocationsResponse)(nil),    // 2: simpliroute.tracking.GetTrackingLocationsResponse
	(*GetLatestTrackingStatusRequest)(nil),  // 3: simpliroute.tracking.GetLatestTrackingStatusRequest
	(*GetLatestTrackingStatusResponse)(nil), // 4: simpliroute.tracking.GetLatestTrackingStatusResponse
	(*gps.TrackingLocation)(nil),            // 5: commons.gps.TrackingLocation
	(*global.Period)(nil),                   // 6: commons.global.Period
	(gps.OwnerType)(0),                      // 7: commons.gps.OwnerType
	(*date.Date)(nil),                       // 8: google.type.Date
	(*gps.LatestTrackingStatus)(nil),        // 9: commons.gps.LatestTrackingStatus
}
var file_tracking_tracking_proto_depIdxs = []int32{
	5, // 0: simpliroute.tracking.StoreTrackingLocationRequest.location:type_name -> commons.gps.TrackingLocation
	6, // 1: simpliroute.tracking.GetTrackingLocationsRequest.period:type_name -> commons.global.Period
	7, // 2: simpliroute.tracking.GetTrackingLocationsRequest.owner_type:type_name -> commons.gps.OwnerType
	5, // 3: simpliroute.tracking.GetTrackingLocationsResponse.locations:type_name -> commons.gps.TrackingLocation
	8, // 4: simpliroute.tracking.GetLatestTrackingStatusRequest.date:type_name -> google.type.Date
	7, // 5: simpliroute.tracking.GetLatestTrackingStatusRequest.owner_type:type_name -> commons.gps.OwnerType
	9, // 6: simpliroute.tracking.GetLatestTrackingStatusResponse.status:type_name -> commons.gps.LatestTrackingStatus
	1, // 7: simpliroute.tracking.TrackingQueryingService.GetTrackingLocations:input_type -> simpliroute.tracking.GetTrackingLocationsRequest
	3, // 8: simpliroute.tracking.TrackingQueryingService.GetLatestTrackingStatus:input_type -> simpliroute.tracking.GetLatestTrackingStatusRequest
	2, // 9: simpliroute.tracking.TrackingQueryingService.GetTrackingLocations:output_type -> simpliroute.tracking.GetTrackingLocationsResponse
	4, // 10: simpliroute.tracking.TrackingQueryingService.GetLatestTrackingStatus:output_type -> simpliroute.tracking.GetLatestTrackingStatusResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_tracking_tracking_proto_init() }
func file_tracking_tracking_proto_init() {
	if File_tracking_tracking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tracking_tracking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreTrackingLocationRequest); i {
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
		file_tracking_tracking_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTrackingLocationsRequest); i {
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
		file_tracking_tracking_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTrackingLocationsResponse); i {
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
		file_tracking_tracking_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLatestTrackingStatusRequest); i {
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
		file_tracking_tracking_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLatestTrackingStatusResponse); i {
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
			RawDescriptor: file_tracking_tracking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tracking_tracking_proto_goTypes,
		DependencyIndexes: file_tracking_tracking_proto_depIdxs,
		MessageInfos:      file_tracking_tracking_proto_msgTypes,
	}.Build()
	File_tracking_tracking_proto = out.File
	file_tracking_tracking_proto_rawDesc = nil
	file_tracking_tracking_proto_goTypes = nil
	file_tracking_tracking_proto_depIdxs = nil
}
