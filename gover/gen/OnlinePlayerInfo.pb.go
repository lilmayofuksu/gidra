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
// source: OnlinePlayerInfo.proto

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

type OnlinePlayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid                 uint32          `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Nickname            string          `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	PlayerLevel         uint32          `protobuf:"varint,3,opt,name=player_level,json=playerLevel,proto3" json:"player_level,omitempty"`
	AvatarId            uint32          `protobuf:"varint,4,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id,omitempty"`
	MpSettingType       MpSettingType   `protobuf:"varint,5,opt,name=mp_setting_type,json=mpSettingType,proto3,enum=MpSettingType" json:"mp_setting_type,omitempty"`
	CurPlayerNumInWorld uint32          `protobuf:"varint,6,opt,name=cur_player_num_in_world,json=curPlayerNumInWorld,proto3" json:"cur_player_num_in_world,omitempty"`
	WorldLevel          uint32          `protobuf:"varint,7,opt,name=world_level,json=worldLevel,proto3" json:"world_level,omitempty"`
	OnlineId            string          `protobuf:"bytes,8,opt,name=online_id,json=onlineId,proto3" json:"online_id,omitempty"`
	NameCardId          uint32          `protobuf:"varint,9,opt,name=name_card_id,json=nameCardId,proto3" json:"name_card_id,omitempty"`
	BlacklistUidList    []uint32        `protobuf:"varint,10,rep,packed,name=blacklist_uid_list,json=blacklistUidList,proto3" json:"blacklist_uid_list,omitempty"`
	Signature           string          `protobuf:"bytes,11,opt,name=signature,proto3" json:"signature,omitempty"`
	ProfilePicture      *ProfilePicture `protobuf:"bytes,12,opt,name=profile_picture,json=profilePicture,proto3" json:"profile_picture,omitempty"`
	PsnId               string          `protobuf:"bytes,13,opt,name=psn_id,json=psnId,proto3" json:"psn_id,omitempty"`
}

func (x *OnlinePlayerInfo) Reset() {
	*x = OnlinePlayerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OnlinePlayerInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnlinePlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnlinePlayerInfo) ProtoMessage() {}

func (x *OnlinePlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_OnlinePlayerInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnlinePlayerInfo.ProtoReflect.Descriptor instead.
func (*OnlinePlayerInfo) Descriptor() ([]byte, []int) {
	return file_OnlinePlayerInfo_proto_rawDescGZIP(), []int{0}
}

func (x *OnlinePlayerInfo) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *OnlinePlayerInfo) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *OnlinePlayerInfo) GetPlayerLevel() uint32 {
	if x != nil {
		return x.PlayerLevel
	}
	return 0
}

func (x *OnlinePlayerInfo) GetAvatarId() uint32 {
	if x != nil {
		return x.AvatarId
	}
	return 0
}

func (x *OnlinePlayerInfo) GetMpSettingType() MpSettingType {
	if x != nil {
		return x.MpSettingType
	}
	return MpSettingType_MP_SETTING_TYPE_NO_ENTER
}

func (x *OnlinePlayerInfo) GetCurPlayerNumInWorld() uint32 {
	if x != nil {
		return x.CurPlayerNumInWorld
	}
	return 0
}

func (x *OnlinePlayerInfo) GetWorldLevel() uint32 {
	if x != nil {
		return x.WorldLevel
	}
	return 0
}

func (x *OnlinePlayerInfo) GetOnlineId() string {
	if x != nil {
		return x.OnlineId
	}
	return ""
}

func (x *OnlinePlayerInfo) GetNameCardId() uint32 {
	if x != nil {
		return x.NameCardId
	}
	return 0
}

func (x *OnlinePlayerInfo) GetBlacklistUidList() []uint32 {
	if x != nil {
		return x.BlacklistUidList
	}
	return nil
}

func (x *OnlinePlayerInfo) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *OnlinePlayerInfo) GetProfilePicture() *ProfilePicture {
	if x != nil {
		return x.ProfilePicture
	}
	return nil
}

func (x *OnlinePlayerInfo) GetPsnId() string {
	if x != nil {
		return x.PsnId
	}
	return ""
}

var File_OnlinePlayerInfo_proto protoreflect.FileDescriptor

var file_OnlinePlayerInfo_proto_rawDesc = []byte{
	0x0a, 0x16, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x4d, 0x70, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xeb, 0x03, 0x0a, 0x10, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x0f, 0x6d, 0x70, 0x5f, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0e, 0x2e, 0x4d, 0x70, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0d, 0x6d, 0x70, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x34,
	0x0a, 0x17, 0x63, 0x75, 0x72, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d,
	0x5f, 0x69, 0x6e, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x13, 0x63, 0x75, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x49, 0x6e, 0x57,
	0x6f, 0x72, 0x6c, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f,
	0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x43, 0x61,
	0x72, 0x64, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73,
	0x74, 0x5f, 0x75, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x10, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x55, 0x69, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x38, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x69, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x70, 0x73,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x73, 0x6e, 0x49,
	0x64, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_OnlinePlayerInfo_proto_rawDescOnce sync.Once
	file_OnlinePlayerInfo_proto_rawDescData = file_OnlinePlayerInfo_proto_rawDesc
)

func file_OnlinePlayerInfo_proto_rawDescGZIP() []byte {
	file_OnlinePlayerInfo_proto_rawDescOnce.Do(func() {
		file_OnlinePlayerInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_OnlinePlayerInfo_proto_rawDescData)
	})
	return file_OnlinePlayerInfo_proto_rawDescData
}

var file_OnlinePlayerInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_OnlinePlayerInfo_proto_goTypes = []interface{}{
	(*OnlinePlayerInfo)(nil), // 0: OnlinePlayerInfo
	(MpSettingType)(0),       // 1: MpSettingType
	(*ProfilePicture)(nil),   // 2: ProfilePicture
}
var file_OnlinePlayerInfo_proto_depIdxs = []int32{
	1, // 0: OnlinePlayerInfo.mp_setting_type:type_name -> MpSettingType
	2, // 1: OnlinePlayerInfo.profile_picture:type_name -> ProfilePicture
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_OnlinePlayerInfo_proto_init() }
func file_OnlinePlayerInfo_proto_init() {
	if File_OnlinePlayerInfo_proto != nil {
		return
	}
	file_MpSettingType_proto_init()
	file_ProfilePicture_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_OnlinePlayerInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnlinePlayerInfo); i {
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
			RawDescriptor: file_OnlinePlayerInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_OnlinePlayerInfo_proto_goTypes,
		DependencyIndexes: file_OnlinePlayerInfo_proto_depIdxs,
		MessageInfos:      file_OnlinePlayerInfo_proto_msgTypes,
	}.Build()
	File_OnlinePlayerInfo_proto = out.File
	file_OnlinePlayerInfo_proto_rawDesc = nil
	file_OnlinePlayerInfo_proto_goTypes = nil
	file_OnlinePlayerInfo_proto_depIdxs = nil
}
