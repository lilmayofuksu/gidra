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
// source: SumoSwitchTeamRsp.proto

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

// CmdId: 8525
// EnetChannelId: 0
// EnetIsReliable: true
type SumoSwitchTeamRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NextValidSwitchTime uint32             `protobuf:"varint,7,opt,name=next_valid_switch_time,json=nextValidSwitchTime,proto3" json:"next_valid_switch_time,omitempty"`
	DungeonTeamList     []*SumoDungeonTeam `protobuf:"bytes,10,rep,name=dungeon_team_list,json=dungeonTeamList,proto3" json:"dungeon_team_list,omitempty"`
	ActivityId          uint32             `protobuf:"varint,6,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"`
	Retcode             int32              `protobuf:"varint,14,opt,name=retcode,proto3" json:"retcode,omitempty"`
	CurTeamIndex        uint32             `protobuf:"varint,11,opt,name=cur_team_index,json=curTeamIndex,proto3" json:"cur_team_index,omitempty"`
	StageId             uint32             `protobuf:"varint,5,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"`
}

func (x *SumoSwitchTeamRsp) Reset() {
	*x = SumoSwitchTeamRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SumoSwitchTeamRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumoSwitchTeamRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumoSwitchTeamRsp) ProtoMessage() {}

func (x *SumoSwitchTeamRsp) ProtoReflect() protoreflect.Message {
	mi := &file_SumoSwitchTeamRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumoSwitchTeamRsp.ProtoReflect.Descriptor instead.
func (*SumoSwitchTeamRsp) Descriptor() ([]byte, []int) {
	return file_SumoSwitchTeamRsp_proto_rawDescGZIP(), []int{0}
}

func (x *SumoSwitchTeamRsp) GetNextValidSwitchTime() uint32 {
	if x != nil {
		return x.NextValidSwitchTime
	}
	return 0
}

func (x *SumoSwitchTeamRsp) GetDungeonTeamList() []*SumoDungeonTeam {
	if x != nil {
		return x.DungeonTeamList
	}
	return nil
}

func (x *SumoSwitchTeamRsp) GetActivityId() uint32 {
	if x != nil {
		return x.ActivityId
	}
	return 0
}

func (x *SumoSwitchTeamRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *SumoSwitchTeamRsp) GetCurTeamIndex() uint32 {
	if x != nil {
		return x.CurTeamIndex
	}
	return 0
}

func (x *SumoSwitchTeamRsp) GetStageId() uint32 {
	if x != nil {
		return x.StageId
	}
	return 0
}

var File_SumoSwitchTeamRsp_proto protoreflect.FileDescriptor

var file_SumoSwitchTeamRsp_proto_rawDesc = []byte{
	0x0a, 0x17, 0x53, 0x75, 0x6d, 0x6f, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x54, 0x65, 0x61, 0x6d,
	0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x53, 0x75, 0x6d, 0x6f, 0x44,
	0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x82, 0x02, 0x0a, 0x11, 0x53, 0x75, 0x6d, 0x6f, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x54,
	0x65, 0x61, 0x6d, 0x52, 0x73, 0x70, 0x12, 0x33, 0x0a, 0x16, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x5f, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x6e, 0x65, 0x78, 0x74, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x11, 0x64,
	0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x53, 0x75, 0x6d, 0x6f, 0x44, 0x75, 0x6e,
	0x67, 0x65, 0x6f, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x0f, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f,
	0x6e, 0x54, 0x65, 0x61, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x75, 0x72, 0x5f, 0x74, 0x65, 0x61, 0x6d,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x75,
	0x72, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x49, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SumoSwitchTeamRsp_proto_rawDescOnce sync.Once
	file_SumoSwitchTeamRsp_proto_rawDescData = file_SumoSwitchTeamRsp_proto_rawDesc
)

func file_SumoSwitchTeamRsp_proto_rawDescGZIP() []byte {
	file_SumoSwitchTeamRsp_proto_rawDescOnce.Do(func() {
		file_SumoSwitchTeamRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_SumoSwitchTeamRsp_proto_rawDescData)
	})
	return file_SumoSwitchTeamRsp_proto_rawDescData
}

var file_SumoSwitchTeamRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SumoSwitchTeamRsp_proto_goTypes = []interface{}{
	(*SumoSwitchTeamRsp)(nil), // 0: SumoSwitchTeamRsp
	(*SumoDungeonTeam)(nil),   // 1: SumoDungeonTeam
}
var file_SumoSwitchTeamRsp_proto_depIdxs = []int32{
	1, // 0: SumoSwitchTeamRsp.dungeon_team_list:type_name -> SumoDungeonTeam
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SumoSwitchTeamRsp_proto_init() }
func file_SumoSwitchTeamRsp_proto_init() {
	if File_SumoSwitchTeamRsp_proto != nil {
		return
	}
	file_SumoDungeonTeam_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SumoSwitchTeamRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumoSwitchTeamRsp); i {
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
			RawDescriptor: file_SumoSwitchTeamRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SumoSwitchTeamRsp_proto_goTypes,
		DependencyIndexes: file_SumoSwitchTeamRsp_proto_depIdxs,
		MessageInfos:      file_SumoSwitchTeamRsp_proto_msgTypes,
	}.Build()
	File_SumoSwitchTeamRsp_proto = out.File
	file_SumoSwitchTeamRsp_proto_rawDesc = nil
	file_SumoSwitchTeamRsp_proto_goTypes = nil
	file_SumoSwitchTeamRsp_proto_depIdxs = nil
}
