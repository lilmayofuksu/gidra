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
// source: SummerTimeSprintBoatSettleNotify.proto

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

// CmdId: 8651
// EnetChannelId: 0
// EnetIsReliable: true
type SummerTimeSprintBoatSettleNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalNum    uint32 `protobuf:"varint,13,opt,name=total_num,json=totalNum,proto3" json:"total_num,omitempty"`
	GroupId     uint32 `protobuf:"varint,12,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	IsSuccess   bool   `protobuf:"varint,15,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	CollectNum  uint32 `protobuf:"varint,6,opt,name=collect_num,json=collectNum,proto3" json:"collect_num,omitempty"`
	LeftTime    uint32 `protobuf:"varint,8,opt,name=left_time,json=leftTime,proto3" json:"left_time,omitempty"`
	MedalLevel  uint32 `protobuf:"varint,2,opt,name=medal_level,json=medalLevel,proto3" json:"medal_level,omitempty"`
	Score       uint32 `protobuf:"varint,10,opt,name=score,proto3" json:"score,omitempty"`
	IsNewRecord bool   `protobuf:"varint,7,opt,name=is_new_record,json=isNewRecord,proto3" json:"is_new_record,omitempty"`
}

func (x *SummerTimeSprintBoatSettleNotify) Reset() {
	*x = SummerTimeSprintBoatSettleNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SummerTimeSprintBoatSettleNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SummerTimeSprintBoatSettleNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SummerTimeSprintBoatSettleNotify) ProtoMessage() {}

func (x *SummerTimeSprintBoatSettleNotify) ProtoReflect() protoreflect.Message {
	mi := &file_SummerTimeSprintBoatSettleNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SummerTimeSprintBoatSettleNotify.ProtoReflect.Descriptor instead.
func (*SummerTimeSprintBoatSettleNotify) Descriptor() ([]byte, []int) {
	return file_SummerTimeSprintBoatSettleNotify_proto_rawDescGZIP(), []int{0}
}

func (x *SummerTimeSprintBoatSettleNotify) GetTotalNum() uint32 {
	if x != nil {
		return x.TotalNum
	}
	return 0
}

func (x *SummerTimeSprintBoatSettleNotify) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *SummerTimeSprintBoatSettleNotify) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *SummerTimeSprintBoatSettleNotify) GetCollectNum() uint32 {
	if x != nil {
		return x.CollectNum
	}
	return 0
}

func (x *SummerTimeSprintBoatSettleNotify) GetLeftTime() uint32 {
	if x != nil {
		return x.LeftTime
	}
	return 0
}

func (x *SummerTimeSprintBoatSettleNotify) GetMedalLevel() uint32 {
	if x != nil {
		return x.MedalLevel
	}
	return 0
}

func (x *SummerTimeSprintBoatSettleNotify) GetScore() uint32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *SummerTimeSprintBoatSettleNotify) GetIsNewRecord() bool {
	if x != nil {
		return x.IsNewRecord
	}
	return false
}

var File_SummerTimeSprintBoatSettleNotify_proto protoreflect.FileDescriptor

var file_SummerTimeSprintBoatSettleNotify_proto_rawDesc = []byte{
	0x0a, 0x26, 0x53, 0x75, 0x6d, 0x6d, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x70, 0x72, 0x69,
	0x6e, 0x74, 0x42, 0x6f, 0x61, 0x74, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x02, 0x0a, 0x20, 0x53, 0x75, 0x6d,
	0x6d, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x42, 0x6f, 0x61,
	0x74, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1b, 0x0a,
	0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x5f,
	0x6e, 0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x66, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6c, 0x65, 0x66, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x64, 0x61, 0x6c, 0x5f, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6d, 0x65, 0x64, 0x61, 0x6c, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x69, 0x73, 0x5f,
	0x6e, 0x65, 0x77, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x69, 0x73, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x42, 0x06, 0x5a,
	0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SummerTimeSprintBoatSettleNotify_proto_rawDescOnce sync.Once
	file_SummerTimeSprintBoatSettleNotify_proto_rawDescData = file_SummerTimeSprintBoatSettleNotify_proto_rawDesc
)

func file_SummerTimeSprintBoatSettleNotify_proto_rawDescGZIP() []byte {
	file_SummerTimeSprintBoatSettleNotify_proto_rawDescOnce.Do(func() {
		file_SummerTimeSprintBoatSettleNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_SummerTimeSprintBoatSettleNotify_proto_rawDescData)
	})
	return file_SummerTimeSprintBoatSettleNotify_proto_rawDescData
}

var file_SummerTimeSprintBoatSettleNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SummerTimeSprintBoatSettleNotify_proto_goTypes = []interface{}{
	(*SummerTimeSprintBoatSettleNotify)(nil), // 0: SummerTimeSprintBoatSettleNotify
}
var file_SummerTimeSprintBoatSettleNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SummerTimeSprintBoatSettleNotify_proto_init() }
func file_SummerTimeSprintBoatSettleNotify_proto_init() {
	if File_SummerTimeSprintBoatSettleNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SummerTimeSprintBoatSettleNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SummerTimeSprintBoatSettleNotify); i {
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
			RawDescriptor: file_SummerTimeSprintBoatSettleNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SummerTimeSprintBoatSettleNotify_proto_goTypes,
		DependencyIndexes: file_SummerTimeSprintBoatSettleNotify_proto_depIdxs,
		MessageInfos:      file_SummerTimeSprintBoatSettleNotify_proto_msgTypes,
	}.Build()
	File_SummerTimeSprintBoatSettleNotify_proto = out.File
	file_SummerTimeSprintBoatSettleNotify_proto_rawDesc = nil
	file_SummerTimeSprintBoatSettleNotify_proto_goTypes = nil
	file_SummerTimeSprintBoatSettleNotify_proto_depIdxs = nil
}
