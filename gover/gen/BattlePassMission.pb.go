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
// source: BattlePassMission.proto

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

type BattlePassMission_MissionStatus int32

const (
	BattlePassMission_MISSION_STATUS_INVALID     BattlePassMission_MissionStatus = 0
	BattlePassMission_MISSION_STATUS_UNFINISHED  BattlePassMission_MissionStatus = 1
	BattlePassMission_MISSION_STATUS_FINISHED    BattlePassMission_MissionStatus = 2
	BattlePassMission_MISSION_STATUS_POINT_TAKEN BattlePassMission_MissionStatus = 3
)

// Enum value maps for BattlePassMission_MissionStatus.
var (
	BattlePassMission_MissionStatus_name = map[int32]string{
		0: "MISSION_STATUS_INVALID",
		1: "MISSION_STATUS_UNFINISHED",
		2: "MISSION_STATUS_FINISHED",
		3: "MISSION_STATUS_POINT_TAKEN",
	}
	BattlePassMission_MissionStatus_value = map[string]int32{
		"MISSION_STATUS_INVALID":     0,
		"MISSION_STATUS_UNFINISHED":  1,
		"MISSION_STATUS_FINISHED":    2,
		"MISSION_STATUS_POINT_TAKEN": 3,
	}
)

func (x BattlePassMission_MissionStatus) Enum() *BattlePassMission_MissionStatus {
	p := new(BattlePassMission_MissionStatus)
	*p = x
	return p
}

func (x BattlePassMission_MissionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BattlePassMission_MissionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_BattlePassMission_proto_enumTypes[0].Descriptor()
}

func (BattlePassMission_MissionStatus) Type() protoreflect.EnumType {
	return &file_BattlePassMission_proto_enumTypes[0]
}

func (x BattlePassMission_MissionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BattlePassMission_MissionStatus.Descriptor instead.
func (BattlePassMission_MissionStatus) EnumDescriptor() ([]byte, []int) {
	return file_BattlePassMission_proto_rawDescGZIP(), []int{0, 0}
}

type BattlePassMission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurProgress           uint32                          `protobuf:"varint,13,opt,name=cur_progress,json=curProgress,proto3" json:"cur_progress,omitempty"`
	MissionStatus         BattlePassMission_MissionStatus `protobuf:"varint,15,opt,name=mission_status,json=missionStatus,proto3,enum=BattlePassMission_MissionStatus" json:"mission_status,omitempty"`
	MissionId             uint32                          `protobuf:"varint,11,opt,name=mission_id,json=missionId,proto3" json:"mission_id,omitempty"`
	RewardBattlePassPoint uint32                          `protobuf:"varint,3,opt,name=reward_battle_pass_point,json=rewardBattlePassPoint,proto3" json:"reward_battle_pass_point,omitempty"`
	MissionType           uint32                          `protobuf:"varint,12,opt,name=mission_type,json=missionType,proto3" json:"mission_type,omitempty"`
	TotalProgress         uint32                          `protobuf:"varint,6,opt,name=total_progress,json=totalProgress,proto3" json:"total_progress,omitempty"`
}

func (x *BattlePassMission) Reset() {
	*x = BattlePassMission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BattlePassMission_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BattlePassMission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BattlePassMission) ProtoMessage() {}

func (x *BattlePassMission) ProtoReflect() protoreflect.Message {
	mi := &file_BattlePassMission_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BattlePassMission.ProtoReflect.Descriptor instead.
func (*BattlePassMission) Descriptor() ([]byte, []int) {
	return file_BattlePassMission_proto_rawDescGZIP(), []int{0}
}

func (x *BattlePassMission) GetCurProgress() uint32 {
	if x != nil {
		return x.CurProgress
	}
	return 0
}

func (x *BattlePassMission) GetMissionStatus() BattlePassMission_MissionStatus {
	if x != nil {
		return x.MissionStatus
	}
	return BattlePassMission_MISSION_STATUS_INVALID
}

func (x *BattlePassMission) GetMissionId() uint32 {
	if x != nil {
		return x.MissionId
	}
	return 0
}

func (x *BattlePassMission) GetRewardBattlePassPoint() uint32 {
	if x != nil {
		return x.RewardBattlePassPoint
	}
	return 0
}

func (x *BattlePassMission) GetMissionType() uint32 {
	if x != nil {
		return x.MissionType
	}
	return 0
}

func (x *BattlePassMission) GetTotalProgress() uint32 {
	if x != nil {
		return x.TotalProgress
	}
	return 0
}

var File_BattlePassMission_proto protoreflect.FileDescriptor

var file_BattlePassMission_proto_rawDesc = []byte{
	0x0a, 0x17, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x50, 0x61, 0x73, 0x73, 0x4d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x03, 0x0a, 0x11, 0x42, 0x61,
	0x74, 0x74, 0x6c, 0x65, 0x50, 0x61, 0x73, 0x73, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x47, 0x0a, 0x0e, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x42, 0x61, 0x74,
	0x74, 0x6c, 0x65, 0x50, 0x61, 0x73, 0x73, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0d, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x18, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x5f, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x73, 0x73,
	0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x50, 0x61, 0x73, 0x73, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x22, 0x87, 0x01,
	0x0a, 0x0d, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1a, 0x0a, 0x16, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x4d,
	0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e,
	0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x49,
	0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x49, 0x4e,
	0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x02, 0x12, 0x1e, 0x0a, 0x1a, 0x4d, 0x49, 0x53, 0x53, 0x49,
	0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x5f,
	0x54, 0x41, 0x4b, 0x45, 0x4e, 0x10, 0x03, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BattlePassMission_proto_rawDescOnce sync.Once
	file_BattlePassMission_proto_rawDescData = file_BattlePassMission_proto_rawDesc
)

func file_BattlePassMission_proto_rawDescGZIP() []byte {
	file_BattlePassMission_proto_rawDescOnce.Do(func() {
		file_BattlePassMission_proto_rawDescData = protoimpl.X.CompressGZIP(file_BattlePassMission_proto_rawDescData)
	})
	return file_BattlePassMission_proto_rawDescData
}

var file_BattlePassMission_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_BattlePassMission_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BattlePassMission_proto_goTypes = []interface{}{
	(BattlePassMission_MissionStatus)(0), // 0: BattlePassMission.MissionStatus
	(*BattlePassMission)(nil),            // 1: BattlePassMission
}
var file_BattlePassMission_proto_depIdxs = []int32{
	0, // 0: BattlePassMission.mission_status:type_name -> BattlePassMission.MissionStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_BattlePassMission_proto_init() }
func file_BattlePassMission_proto_init() {
	if File_BattlePassMission_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_BattlePassMission_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BattlePassMission); i {
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
			RawDescriptor: file_BattlePassMission_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BattlePassMission_proto_goTypes,
		DependencyIndexes: file_BattlePassMission_proto_depIdxs,
		EnumInfos:         file_BattlePassMission_proto_enumTypes,
		MessageInfos:      file_BattlePassMission_proto_msgTypes,
	}.Build()
	File_BattlePassMission_proto = out.File
	file_BattlePassMission_proto_rawDesc = nil
	file_BattlePassMission_proto_goTypes = nil
	file_BattlePassMission_proto_depIdxs = nil
}
