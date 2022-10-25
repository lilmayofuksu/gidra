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
// source: PlayerMatchSuccNotify.proto

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

// CmdId: 4179
// EnetChannelId: 0
// EnetIsReliable: true
type PlayerMatchSuccNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MechanicusDifficultLevel uint32            `protobuf:"varint,1,opt,name=mechanicus_difficult_level,json=mechanicusDifficultLevel,proto3" json:"mechanicus_difficult_level,omitempty"`
	DungeonId                uint32            `protobuf:"varint,6,opt,name=dungeon_id,json=dungeonId,proto3" json:"dungeon_id,omitempty"`
	MatchType                MatchType         `protobuf:"varint,5,opt,name=match_type,json=matchType,proto3,enum=MatchType" json:"match_type,omitempty"`
	MpPlayId                 uint32            `protobuf:"varint,15,opt,name=mp_play_id,json=mpPlayId,proto3" json:"mp_play_id,omitempty"`
	GeneralMatchInfo         *GeneralMatchInfo `protobuf:"bytes,7,opt,name=general_match_info,json=generalMatchInfo,proto3" json:"general_match_info,omitempty"`
	HostUid                  uint32            `protobuf:"varint,3,opt,name=host_uid,json=hostUid,proto3" json:"host_uid,omitempty"`
	ConfirmEndTime           uint32            `protobuf:"varint,2,opt,name=confirm_end_time,json=confirmEndTime,proto3" json:"confirm_end_time,omitempty"`
}

func (x *PlayerMatchSuccNotify) Reset() {
	*x = PlayerMatchSuccNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerMatchSuccNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerMatchSuccNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerMatchSuccNotify) ProtoMessage() {}

func (x *PlayerMatchSuccNotify) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerMatchSuccNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerMatchSuccNotify.ProtoReflect.Descriptor instead.
func (*PlayerMatchSuccNotify) Descriptor() ([]byte, []int) {
	return file_PlayerMatchSuccNotify_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerMatchSuccNotify) GetMechanicusDifficultLevel() uint32 {
	if x != nil {
		return x.MechanicusDifficultLevel
	}
	return 0
}

func (x *PlayerMatchSuccNotify) GetDungeonId() uint32 {
	if x != nil {
		return x.DungeonId
	}
	return 0
}

func (x *PlayerMatchSuccNotify) GetMatchType() MatchType {
	if x != nil {
		return x.MatchType
	}
	return MatchType_MATCH_TYPE_NONE
}

func (x *PlayerMatchSuccNotify) GetMpPlayId() uint32 {
	if x != nil {
		return x.MpPlayId
	}
	return 0
}

func (x *PlayerMatchSuccNotify) GetGeneralMatchInfo() *GeneralMatchInfo {
	if x != nil {
		return x.GeneralMatchInfo
	}
	return nil
}

func (x *PlayerMatchSuccNotify) GetHostUid() uint32 {
	if x != nil {
		return x.HostUid
	}
	return 0
}

func (x *PlayerMatchSuccNotify) GetConfirmEndTime() uint32 {
	if x != nil {
		return x.ConfirmEndTime
	}
	return 0
}

var File_PlayerMatchSuccNotify_proto protoreflect.FileDescriptor

var file_PlayerMatchSuccNotify_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x75, 0x63,
	0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x02, 0x0a, 0x15, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x75, 0x63, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x12, 0x3c, 0x0a, 0x1a, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x5f, 0x64,
	0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x18, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73,
	0x44, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1d,
	0x0a, 0x0a, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x29, 0x0a,
	0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0a, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x0a, 0x6d, 0x70, 0x5f, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6d, 0x70,
	0x50, 0x6c, 0x61, 0x79, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x12, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x6c, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x10, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x5f,
	0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x68, 0x6f, 0x73, 0x74, 0x55,
	0x69, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x5f, 0x65, 0x6e,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x06, 0x5a, 0x04,
	0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerMatchSuccNotify_proto_rawDescOnce sync.Once
	file_PlayerMatchSuccNotify_proto_rawDescData = file_PlayerMatchSuccNotify_proto_rawDesc
)

func file_PlayerMatchSuccNotify_proto_rawDescGZIP() []byte {
	file_PlayerMatchSuccNotify_proto_rawDescOnce.Do(func() {
		file_PlayerMatchSuccNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerMatchSuccNotify_proto_rawDescData)
	})
	return file_PlayerMatchSuccNotify_proto_rawDescData
}

var file_PlayerMatchSuccNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerMatchSuccNotify_proto_goTypes = []interface{}{
	(*PlayerMatchSuccNotify)(nil), // 0: PlayerMatchSuccNotify
	(MatchType)(0),                // 1: MatchType
	(*GeneralMatchInfo)(nil),      // 2: GeneralMatchInfo
}
var file_PlayerMatchSuccNotify_proto_depIdxs = []int32{
	1, // 0: PlayerMatchSuccNotify.match_type:type_name -> MatchType
	2, // 1: PlayerMatchSuccNotify.general_match_info:type_name -> GeneralMatchInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_PlayerMatchSuccNotify_proto_init() }
func file_PlayerMatchSuccNotify_proto_init() {
	if File_PlayerMatchSuccNotify_proto != nil {
		return
	}
	file_GeneralMatchInfo_proto_init()
	file_MatchType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerMatchSuccNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerMatchSuccNotify); i {
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
			RawDescriptor: file_PlayerMatchSuccNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerMatchSuccNotify_proto_goTypes,
		DependencyIndexes: file_PlayerMatchSuccNotify_proto_depIdxs,
		MessageInfos:      file_PlayerMatchSuccNotify_proto_msgTypes,
	}.Build()
	File_PlayerMatchSuccNotify_proto = out.File
	file_PlayerMatchSuccNotify_proto_rawDesc = nil
	file_PlayerMatchSuccNotify_proto_goTypes = nil
	file_PlayerMatchSuccNotify_proto_depIdxs = nil
}
