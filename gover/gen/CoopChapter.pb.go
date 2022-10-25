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
// source: CoopChapter.proto

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

type CoopChapter_State int32

const (
	CoopChapter_STATE_CLOSE         CoopChapter_State = 0
	CoopChapter_STATE_COND_NOT_MEET CoopChapter_State = 1
	CoopChapter_STATE_COND_MEET     CoopChapter_State = 2
	CoopChapter_STATE_ACCEPT        CoopChapter_State = 3
)

// Enum value maps for CoopChapter_State.
var (
	CoopChapter_State_name = map[int32]string{
		0: "STATE_CLOSE",
		1: "STATE_COND_NOT_MEET",
		2: "STATE_COND_MEET",
		3: "STATE_ACCEPT",
	}
	CoopChapter_State_value = map[string]int32{
		"STATE_CLOSE":         0,
		"STATE_COND_NOT_MEET": 1,
		"STATE_COND_MEET":     2,
		"STATE_ACCEPT":        3,
	}
)

func (x CoopChapter_State) Enum() *CoopChapter_State {
	p := new(CoopChapter_State)
	*p = x
	return p
}

func (x CoopChapter_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CoopChapter_State) Descriptor() protoreflect.EnumDescriptor {
	return file_CoopChapter_proto_enumTypes[0].Descriptor()
}

func (CoopChapter_State) Type() protoreflect.EnumType {
	return &file_CoopChapter_proto_enumTypes[0]
}

func (x CoopChapter_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CoopChapter_State.Descriptor instead.
func (CoopChapter_State) EnumDescriptor() ([]byte, []int) {
	return file_CoopChapter_proto_rawDescGZIP(), []int{0, 0}
}

type CoopChapter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoopCgList       []*CoopCg         `protobuf:"bytes,1,rep,name=coop_cg_list,json=coopCgList,proto3" json:"coop_cg_list,omitempty"`
	Id               uint32            `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	CoopPointList    []*CoopPoint      `protobuf:"bytes,11,rep,name=coop_point_list,json=coopPointList,proto3" json:"coop_point_list,omitempty"`
	FinishDialogList []uint32          `protobuf:"varint,10,rep,packed,name=finish_dialog_list,json=finishDialogList,proto3" json:"finish_dialog_list,omitempty"`
	FinishedEndCount uint32            `protobuf:"varint,14,opt,name=finished_end_count,json=finishedEndCount,proto3" json:"finished_end_count,omitempty"`
	TotalEndCount    uint32            `protobuf:"varint,7,opt,name=total_end_count,json=totalEndCount,proto3" json:"total_end_count,omitempty"`
	CoopRewardList   []*CoopReward     `protobuf:"bytes,5,rep,name=coop_reward_list,json=coopRewardList,proto3" json:"coop_reward_list,omitempty"`
	LockReasonList   []uint32          `protobuf:"varint,12,rep,packed,name=lock_reason_list,json=lockReasonList,proto3" json:"lock_reason_list,omitempty"`
	State            CoopChapter_State `protobuf:"varint,4,opt,name=state,proto3,enum=CoopChapter_State" json:"state,omitempty"`
	SeenEndingMap    map[uint32]uint32 `protobuf:"bytes,9,rep,name=seen_ending_map,json=seenEndingMap,proto3" json:"seen_ending_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *CoopChapter) Reset() {
	*x = CoopChapter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CoopChapter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoopChapter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoopChapter) ProtoMessage() {}

func (x *CoopChapter) ProtoReflect() protoreflect.Message {
	mi := &file_CoopChapter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoopChapter.ProtoReflect.Descriptor instead.
func (*CoopChapter) Descriptor() ([]byte, []int) {
	return file_CoopChapter_proto_rawDescGZIP(), []int{0}
}

func (x *CoopChapter) GetCoopCgList() []*CoopCg {
	if x != nil {
		return x.CoopCgList
	}
	return nil
}

func (x *CoopChapter) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CoopChapter) GetCoopPointList() []*CoopPoint {
	if x != nil {
		return x.CoopPointList
	}
	return nil
}

func (x *CoopChapter) GetFinishDialogList() []uint32 {
	if x != nil {
		return x.FinishDialogList
	}
	return nil
}

func (x *CoopChapter) GetFinishedEndCount() uint32 {
	if x != nil {
		return x.FinishedEndCount
	}
	return 0
}

func (x *CoopChapter) GetTotalEndCount() uint32 {
	if x != nil {
		return x.TotalEndCount
	}
	return 0
}

func (x *CoopChapter) GetCoopRewardList() []*CoopReward {
	if x != nil {
		return x.CoopRewardList
	}
	return nil
}

func (x *CoopChapter) GetLockReasonList() []uint32 {
	if x != nil {
		return x.LockReasonList
	}
	return nil
}

func (x *CoopChapter) GetState() CoopChapter_State {
	if x != nil {
		return x.State
	}
	return CoopChapter_STATE_CLOSE
}

func (x *CoopChapter) GetSeenEndingMap() map[uint32]uint32 {
	if x != nil {
		return x.SeenEndingMap
	}
	return nil
}

var File_CoopChapter_proto protoreflect.FileDescriptor

var file_CoopChapter_proto_rawDesc = []byte{
	0x0a, 0x11, 0x43, 0x6f, 0x6f, 0x70, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x43, 0x6f, 0x6f, 0x70, 0x43, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0f, 0x43, 0x6f, 0x6f, 0x70, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x10, 0x43, 0x6f, 0x6f, 0x70, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x04, 0x0a, 0x0b, 0x43, 0x6f, 0x6f, 0x70, 0x43, 0x68, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x0c, 0x63, 0x6f, 0x6f, 0x70, 0x5f, 0x63, 0x67, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x43, 0x6f, 0x6f,
	0x70, 0x43, 0x67, 0x52, 0x0a, 0x63, 0x6f, 0x6f, 0x70, 0x43, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x32, 0x0a, 0x0f, 0x63, 0x6f, 0x6f, 0x70, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x43, 0x6f, 0x6f, 0x70, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0d, 0x63, 0x6f, 0x6f, 0x70, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x64, 0x69,
	0x61, 0x6c, 0x6f, 0x67, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x10, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x2c, 0x0a, 0x12, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x65, 0x6e,
	0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x66,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x45, 0x6e, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x26, 0x0a, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45,
	0x6e, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x35, 0x0a, 0x10, 0x63, 0x6f, 0x6f, 0x70, 0x5f,
	0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x43, 0x6f, 0x6f, 0x70, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x0e,
	0x63, 0x6f, 0x6f, 0x70, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x28,
	0x0a, 0x10, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0e, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x43, 0x6f, 0x6f, 0x70, 0x43, 0x68,
	0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x47, 0x0a, 0x0f, 0x73, 0x65, 0x65, 0x6e, 0x5f, 0x65, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x43, 0x6f,
	0x6f, 0x70, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x65, 0x6e, 0x45, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x73, 0x65,
	0x65, 0x6e, 0x45, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x1a, 0x40, 0x0a, 0x12, 0x53,
	0x65, 0x65, 0x6e, 0x45, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x58, 0x0a,
	0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x43, 0x4c, 0x4f, 0x53, 0x45, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4d, 0x45, 0x45, 0x54, 0x10, 0x01,
	0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x44, 0x5f, 0x4d,
	0x45, 0x45, 0x54, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x41,
	0x43, 0x43, 0x45, 0x50, 0x54, 0x10, 0x03, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CoopChapter_proto_rawDescOnce sync.Once
	file_CoopChapter_proto_rawDescData = file_CoopChapter_proto_rawDesc
)

func file_CoopChapter_proto_rawDescGZIP() []byte {
	file_CoopChapter_proto_rawDescOnce.Do(func() {
		file_CoopChapter_proto_rawDescData = protoimpl.X.CompressGZIP(file_CoopChapter_proto_rawDescData)
	})
	return file_CoopChapter_proto_rawDescData
}

var file_CoopChapter_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_CoopChapter_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_CoopChapter_proto_goTypes = []interface{}{
	(CoopChapter_State)(0), // 0: CoopChapter.State
	(*CoopChapter)(nil),    // 1: CoopChapter
	nil,                    // 2: CoopChapter.SeenEndingMapEntry
	(*CoopCg)(nil),         // 3: CoopCg
	(*CoopPoint)(nil),      // 4: CoopPoint
	(*CoopReward)(nil),     // 5: CoopReward
}
var file_CoopChapter_proto_depIdxs = []int32{
	3, // 0: CoopChapter.coop_cg_list:type_name -> CoopCg
	4, // 1: CoopChapter.coop_point_list:type_name -> CoopPoint
	5, // 2: CoopChapter.coop_reward_list:type_name -> CoopReward
	0, // 3: CoopChapter.state:type_name -> CoopChapter.State
	2, // 4: CoopChapter.seen_ending_map:type_name -> CoopChapter.SeenEndingMapEntry
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_CoopChapter_proto_init() }
func file_CoopChapter_proto_init() {
	if File_CoopChapter_proto != nil {
		return
	}
	file_CoopCg_proto_init()
	file_CoopPoint_proto_init()
	file_CoopReward_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CoopChapter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoopChapter); i {
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
			RawDescriptor: file_CoopChapter_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CoopChapter_proto_goTypes,
		DependencyIndexes: file_CoopChapter_proto_depIdxs,
		EnumInfos:         file_CoopChapter_proto_enumTypes,
		MessageInfos:      file_CoopChapter_proto_msgTypes,
	}.Build()
	File_CoopChapter_proto = out.File
	file_CoopChapter_proto_rawDesc = nil
	file_CoopChapter_proto_goTypes = nil
	file_CoopChapter_proto_depIdxs = nil
}
