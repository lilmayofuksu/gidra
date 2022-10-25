// Sorapointa - A server software re-implementation for a certain anime game,
// and avoid sorapointa. Copyright (C) 2022  Sorapointa Team
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: OpActivityInfo.proto

package gen

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

type OpActivityInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActivityId  uint32 `protobuf:"varint,2,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"`
	EndTime     uint32 `protobuf:"varint,6,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	BeginTime   uint32 `protobuf:"varint,5,opt,name=begin_time,json=beginTime,proto3" json:"begin_time,omitempty"`
	IsHasChange bool   `protobuf:"varint,1,opt,name=is_has_change,json=isHasChange,proto3" json:"is_has_change,omitempty"`
	ScheduleId  uint32 `protobuf:"varint,13,opt,name=schedule_id,json=scheduleId,proto3" json:"schedule_id,omitempty"`
	// Types that are assignable to Detail:
	//
	//	*OpActivityInfo_BonusInfo
	Detail isOpActivityInfo_Detail `protobuf_oneof:"detail"`
}

func (x *OpActivityInfo) Reset() {
	*x = OpActivityInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OpActivityInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpActivityInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpActivityInfo) ProtoMessage() {}

func (x *OpActivityInfo) ProtoReflect() protoreflect.Message {
	mi := &file_OpActivityInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpActivityInfo.ProtoReflect.Descriptor instead.
func (*OpActivityInfo) Descriptor() ([]byte, []int) {
	return file_OpActivityInfo_proto_rawDescGZIP(), []int{0}
}

func (x *OpActivityInfo) GetActivityId() uint32 {
	if x != nil {
		return x.ActivityId
	}
	return 0
}

func (x *OpActivityInfo) GetEndTime() uint32 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *OpActivityInfo) GetBeginTime() uint32 {
	if x != nil {
		return x.BeginTime
	}
	return 0
}

func (x *OpActivityInfo) GetIsHasChange() bool {
	if x != nil {
		return x.IsHasChange
	}
	return false
}

func (x *OpActivityInfo) GetScheduleId() uint32 {
	if x != nil {
		return x.ScheduleId
	}
	return 0
}

func (m *OpActivityInfo) GetDetail() isOpActivityInfo_Detail {
	if m != nil {
		return m.Detail
	}
	return nil
}

func (x *OpActivityInfo) GetBonusInfo() *BonusOpActivityInfo {
	if x, ok := x.GetDetail().(*OpActivityInfo_BonusInfo); ok {
		return x.BonusInfo
	}
	return nil
}

type isOpActivityInfo_Detail interface {
	isOpActivityInfo_Detail()
}

type OpActivityInfo_BonusInfo struct {
	BonusInfo *BonusOpActivityInfo `protobuf:"bytes,12,opt,name=bonus_info,json=bonusInfo,proto3,oneof"`
}

func (*OpActivityInfo_BonusInfo) isOpActivityInfo_Detail() {}

var File_OpActivityInfo_proto protoreflect.FileDescriptor

var file_OpActivityInfo_proto_rawDesc = []byte{
	0x0a, 0x14, 0x4f, 0x70, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x4f, 0x70, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xf1, 0x01, 0x0a, 0x0e, 0x4f, 0x70, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x22, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x68, 0x61, 0x73, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x48, 0x61, 0x73, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x0a, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x42, 0x6f, 0x6e, 0x75, 0x73,
	0x4f, 0x70, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00,
	0x52, 0x09, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x08, 0x0a, 0x06, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_OpActivityInfo_proto_rawDescOnce sync.Once
	file_OpActivityInfo_proto_rawDescData = file_OpActivityInfo_proto_rawDesc
)

func file_OpActivityInfo_proto_rawDescGZIP() []byte {
	file_OpActivityInfo_proto_rawDescOnce.Do(func() {
		file_OpActivityInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_OpActivityInfo_proto_rawDescData)
	})
	return file_OpActivityInfo_proto_rawDescData
}

var file_OpActivityInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_OpActivityInfo_proto_goTypes = []interface{}{
	(*OpActivityInfo)(nil),      // 0: OpActivityInfo
	(*BonusOpActivityInfo)(nil), // 1: BonusOpActivityInfo
}
var file_OpActivityInfo_proto_depIdxs = []int32{
	1, // 0: OpActivityInfo.bonus_info:type_name -> BonusOpActivityInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_OpActivityInfo_proto_init() }
func file_OpActivityInfo_proto_init() {
	if File_OpActivityInfo_proto != nil {
		return
	}
	file_BonusOpActivityInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_OpActivityInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpActivityInfo); i {
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
	file_OpActivityInfo_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*OpActivityInfo_BonusInfo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_OpActivityInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_OpActivityInfo_proto_goTypes,
		DependencyIndexes: file_OpActivityInfo_proto_depIdxs,
		MessageInfos:      file_OpActivityInfo_proto_msgTypes,
	}.Build()
	File_OpActivityInfo_proto = out.File
	file_OpActivityInfo_proto_rawDesc = nil
	file_OpActivityInfo_proto_goTypes = nil
	file_OpActivityInfo_proto_depIdxs = nil
}
