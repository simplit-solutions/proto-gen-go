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
// source: commons/gps/gps.proto

package gps

import (
	proto "github.com/golang/protobuf/proto"
	global "gitlab.com/simpliroute/proto-gen-go/visit-properties/src/commons/global"
	date "google.golang.org/genproto/googleapis/type/date"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

type ActivityType int32

const (
	ActivityType_ACTIVITY_TYPE_NOT_SET ActivityType = 0
	ActivityType_STILL                 ActivityType = 1
	ActivityType_ON_FOOT               ActivityType = 2
	ActivityType_RUNNING               ActivityType = 3
	ActivityType_ON_BICYCLE            ActivityType = 4
	ActivityType_IN_VEHICLE            ActivityType = 5
	ActivityType_UNKNOWN               ActivityType = 6
	ActivityType_WALKING               ActivityType = 7
	ActivityType_TILTING               ActivityType = 8
)

// Enum value maps for ActivityType.
var (
	ActivityType_name = map[int32]string{
		0: "ACTIVITY_TYPE_NOT_SET",
		1: "STILL",
		2: "ON_FOOT",
		3: "RUNNING",
		4: "ON_BICYCLE",
		5: "IN_VEHICLE",
		6: "UNKNOWN",
		7: "WALKING",
		8: "TILTING",
	}
	ActivityType_value = map[string]int32{
		"ACTIVITY_TYPE_NOT_SET": 0,
		"STILL":                 1,
		"ON_FOOT":               2,
		"RUNNING":               3,
		"ON_BICYCLE":            4,
		"IN_VEHICLE":            5,
		"UNKNOWN":               6,
		"WALKING":               7,
		"TILTING":               8,
	}
)

func (x ActivityType) Enum() *ActivityType {
	p := new(ActivityType)
	*p = x
	return p
}

func (x ActivityType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActivityType) Descriptor() protoreflect.EnumDescriptor {
	return file_commons_gps_gps_proto_enumTypes[0].Descriptor()
}

func (ActivityType) Type() protoreflect.EnumType {
	return &file_commons_gps_gps_proto_enumTypes[0]
}

func (x ActivityType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActivityType.Descriptor instead.
func (ActivityType) EnumDescriptor() ([]byte, []int) {
	return file_commons_gps_gps_proto_rawDescGZIP(), []int{0}
}

type OwnerType int32

const (
	OwnerType_OWNER_TYPE_NOT_SET OwnerType = 0
	OwnerType_VEHICLE            OwnerType = 1
	OwnerType_DRIVER             OwnerType = 2
)

// Enum value maps for OwnerType.
var (
	OwnerType_name = map[int32]string{
		0: "OWNER_TYPE_NOT_SET",
		1: "VEHICLE",
		2: "DRIVER",
	}
	OwnerType_value = map[string]int32{
		"OWNER_TYPE_NOT_SET": 0,
		"VEHICLE":            1,
		"DRIVER":             2,
	}
)

func (x OwnerType) Enum() *OwnerType {
	p := new(OwnerType)
	*p = x
	return p
}

func (x OwnerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OwnerType) Descriptor() protoreflect.EnumDescriptor {
	return file_commons_gps_gps_proto_enumTypes[1].Descriptor()
}

func (OwnerType) Type() protoreflect.EnumType {
	return &file_commons_gps_gps_proto_enumTypes[1]
}

func (x OwnerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OwnerType.Descriptor instead.
func (OwnerType) EnumDescriptor() ([]byte, []int) {
	return file_commons_gps_gps_proto_rawDescGZIP(), []int{1}
}

type TrackingLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OwnerType OwnerType `protobuf:"varint,2,opt,name=owner_type,json=ownerType,proto3,enum=commons.gps.OwnerType" json:"owner_type,omitempty"`
	OwnerId   uint32    `protobuf:"varint,3,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	// in the range [1, Infinity); the closer to 1, the more accurate the reported coordinates; not a radius in meters
	Accuracy *wrapperspb.FloatValue `protobuf:"bytes,4,opt,name=accuracy,proto3" json:"accuracy,omitempty"`
	// arbitrary JSON metadata from the GPS device, stored as-is and returned to the client on request
	Alert *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=alert,proto3" json:"alert,omitempty"`
	// source device of tracking data
	ProviderName *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=provider_name,json=providerName,proto3" json:"provider_name,omitempty"`
	BatteryLevel *wrapperspb.FloatValue  `protobuf:"bytes,7,opt,name=battery_level,json=batteryLevel,proto3" json:"battery_level,omitempty"` // [0,1]; % of battery remaining
	ActivityType ActivityType            `protobuf:"varint,8,opt,name=activity_type,json=activityType,proto3,enum=commons.gps.ActivityType" json:"activity_type,omitempty"`
	// when the GPS measurment was taken
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Location  *global.Location       `protobuf:"bytes,10,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *TrackingLocation) Reset() {
	*x = TrackingLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commons_gps_gps_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackingLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackingLocation) ProtoMessage() {}

func (x *TrackingLocation) ProtoReflect() protoreflect.Message {
	mi := &file_commons_gps_gps_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackingLocation.ProtoReflect.Descriptor instead.
func (*TrackingLocation) Descriptor() ([]byte, []int) {
	return file_commons_gps_gps_proto_rawDescGZIP(), []int{0}
}

func (x *TrackingLocation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TrackingLocation) GetOwnerType() OwnerType {
	if x != nil {
		return x.OwnerType
	}
	return OwnerType_OWNER_TYPE_NOT_SET
}

func (x *TrackingLocation) GetOwnerId() uint32 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *TrackingLocation) GetAccuracy() *wrapperspb.FloatValue {
	if x != nil {
		return x.Accuracy
	}
	return nil
}

func (x *TrackingLocation) GetAlert() *wrapperspb.StringValue {
	if x != nil {
		return x.Alert
	}
	return nil
}

func (x *TrackingLocation) GetProviderName() *wrapperspb.StringValue {
	if x != nil {
		return x.ProviderName
	}
	return nil
}

func (x *TrackingLocation) GetBatteryLevel() *wrapperspb.FloatValue {
	if x != nil {
		return x.BatteryLevel
	}
	return nil
}

func (x *TrackingLocation) GetActivityType() ActivityType {
	if x != nil {
		return x.ActivityType
	}
	return ActivityType_ACTIVITY_TYPE_NOT_SET
}

func (x *TrackingLocation) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *TrackingLocation) GetLocation() *global.Location {
	if x != nil {
		return x.Location
	}
	return nil
}

type LatestTrackingStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date           *date.Date        `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	LatestLocation *TrackingLocation `protobuf:"bytes,2,opt,name=latest_location,json=latestLocation,proto3" json:"latest_location,omitempty"`
}

func (x *LatestTrackingStatus) Reset() {
	*x = LatestTrackingStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commons_gps_gps_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LatestTrackingStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LatestTrackingStatus) ProtoMessage() {}

func (x *LatestTrackingStatus) ProtoReflect() protoreflect.Message {
	mi := &file_commons_gps_gps_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LatestTrackingStatus.ProtoReflect.Descriptor instead.
func (*LatestTrackingStatus) Descriptor() ([]byte, []int) {
	return file_commons_gps_gps_proto_rawDescGZIP(), []int{1}
}

func (x *LatestTrackingStatus) GetDate() *date.Date {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *LatestTrackingStatus) GetLatestLocation() *TrackingLocation {
	if x != nil {
		return x.LatestLocation
	}
	return nil
}

var File_commons_gps_gps_proto protoreflect.FileDescriptor

var file_commons_gps_gps_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x70, 0x73, 0x2f, 0x67, 0x70,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x67, 0x70, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x04, 0x0a, 0x10,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x35, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x67,
	0x70, 0x73, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x37, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x08, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x12, 0x32, 0x0a, 0x05, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x12,
	0x41, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0d, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x5f, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x3e, 0x0a, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x70, 0x73, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x34,
	0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x16, 0xc2, 0xf3, 0x18, 0x12, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69,
	0x6e, 0x67, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xa1, 0x01, 0x0a,
	0x14, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x46, 0x0a, 0x0f,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x67, 0x70, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x1a, 0xc2, 0xf3, 0x18, 0x16, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74,
	0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2a, 0x95, 0x01, 0x0a, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x53, 0x54, 0x49, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x4e, 0x5f, 0x46, 0x4f,
	0x4f, 0x54, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10,
	0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x4e, 0x5f, 0x42, 0x49, 0x43, 0x59, 0x43, 0x4c, 0x45, 0x10,
	0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4e, 0x5f, 0x56, 0x45, 0x48, 0x49, 0x43, 0x4c, 0x45, 0x10,
	0x05, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x06, 0x12, 0x0b,
	0x0a, 0x07, 0x57, 0x41, 0x4c, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0x07, 0x12, 0x0b, 0x0a, 0x07, 0x54,
	0x49, 0x4c, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x08, 0x2a, 0x3c, 0x0a, 0x09, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x56, 0x45, 0x48, 0x49, 0x43, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x52,
	0x49, 0x56, 0x45, 0x52, 0x10, 0x02, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67,
	0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x70, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commons_gps_gps_proto_rawDescOnce sync.Once
	file_commons_gps_gps_proto_rawDescData = file_commons_gps_gps_proto_rawDesc
)

func file_commons_gps_gps_proto_rawDescGZIP() []byte {
	file_commons_gps_gps_proto_rawDescOnce.Do(func() {
		file_commons_gps_gps_proto_rawDescData = protoimpl.X.CompressGZIP(file_commons_gps_gps_proto_rawDescData)
	})
	return file_commons_gps_gps_proto_rawDescData
}

var file_commons_gps_gps_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_commons_gps_gps_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_commons_gps_gps_proto_goTypes = []interface{}{
	(ActivityType)(0),              // 0: commons.gps.ActivityType
	(OwnerType)(0),                 // 1: commons.gps.OwnerType
	(*TrackingLocation)(nil),       // 2: commons.gps.TrackingLocation
	(*LatestTrackingStatus)(nil),   // 3: commons.gps.LatestTrackingStatus
	(*wrapperspb.FloatValue)(nil),  // 4: google.protobuf.FloatValue
	(*wrapperspb.StringValue)(nil), // 5: google.protobuf.StringValue
	(*timestamppb.Timestamp)(nil),  // 6: google.protobuf.Timestamp
	(*global.Location)(nil),        // 7: commons.global.Location
	(*date.Date)(nil),              // 8: google.type.Date
}
var file_commons_gps_gps_proto_depIdxs = []int32{
	1,  // 0: commons.gps.TrackingLocation.owner_type:type_name -> commons.gps.OwnerType
	4,  // 1: commons.gps.TrackingLocation.accuracy:type_name -> google.protobuf.FloatValue
	5,  // 2: commons.gps.TrackingLocation.alert:type_name -> google.protobuf.StringValue
	5,  // 3: commons.gps.TrackingLocation.provider_name:type_name -> google.protobuf.StringValue
	4,  // 4: commons.gps.TrackingLocation.battery_level:type_name -> google.protobuf.FloatValue
	0,  // 5: commons.gps.TrackingLocation.activity_type:type_name -> commons.gps.ActivityType
	6,  // 6: commons.gps.TrackingLocation.timestamp:type_name -> google.protobuf.Timestamp
	7,  // 7: commons.gps.TrackingLocation.location:type_name -> commons.global.Location
	8,  // 8: commons.gps.LatestTrackingStatus.date:type_name -> google.type.Date
	2,  // 9: commons.gps.LatestTrackingStatus.latest_location:type_name -> commons.gps.TrackingLocation
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_commons_gps_gps_proto_init() }
func file_commons_gps_gps_proto_init() {
	if File_commons_gps_gps_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commons_gps_gps_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackingLocation); i {
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
		file_commons_gps_gps_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LatestTrackingStatus); i {
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
			RawDescriptor: file_commons_gps_gps_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_commons_gps_gps_proto_goTypes,
		DependencyIndexes: file_commons_gps_gps_proto_depIdxs,
		EnumInfos:         file_commons_gps_gps_proto_enumTypes,
		MessageInfos:      file_commons_gps_gps_proto_msgTypes,
	}.Build()
	File_commons_gps_gps_proto = out.File
	file_commons_gps_gps_proto_rawDesc = nil
	file_commons_gps_gps_proto_goTypes = nil
	file_commons_gps_gps_proto_depIdxs = nil
}
