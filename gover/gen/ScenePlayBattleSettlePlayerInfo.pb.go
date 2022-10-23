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
// source: ScenePlayBattleSettlePlayerInfo.proto

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

type ScenePlayBattleSettlePlayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CardList       []*ExhibitionDisplayInfo `protobuf:"bytes,14,rep,name=card_list,json=cardList,proto3" json:"card_list,omitempty"`
	ProfilePicture *ProfilePicture          `protobuf:"bytes,10,opt,name=profile_picture,json=profilePicture,proto3" json:"profile_picture,omitempty"`
	HeadImage      uint32                   `protobuf:"varint,11,opt,name=head_image,json=headImage,proto3" json:"head_image,omitempty"`
	StatisticId    uint32                   `protobuf:"varint,4,opt,name=statistic_id,json=statisticId,proto3" json:"statistic_id,omitempty"`
	Uid            uint32                   `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Param          int64                    `protobuf:"varint,5,opt,name=param,proto3" json:"param,omitempty"`
	OnlineId       string                   `protobuf:"bytes,12,opt,name=online_id,json=onlineId,proto3" json:"online_id,omitempty"`
	Nickname       string                   `protobuf:"bytes,15,opt,name=nickname,proto3" json:"nickname,omitempty"`
}

func (x *ScenePlayBattleSettlePlayerInfo) Reset() {
	*x = ScenePlayBattleSettlePlayerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ScenePlayBattleSettlePlayerInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScenePlayBattleSettlePlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScenePlayBattleSettlePlayerInfo) ProtoMessage() {}

func (x *ScenePlayBattleSettlePlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ScenePlayBattleSettlePlayerInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScenePlayBattleSettlePlayerInfo.ProtoReflect.Descriptor instead.
func (*ScenePlayBattleSettlePlayerInfo) Descriptor() ([]byte, []int) {
	return file_ScenePlayBattleSettlePlayerInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ScenePlayBattleSettlePlayerInfo) GetCardList() []*ExhibitionDisplayInfo {
	if x != nil {
		return x.CardList
	}
	return nil
}

func (x *ScenePlayBattleSettlePlayerInfo) GetProfilePicture() *ProfilePicture {
	if x != nil {
		return x.ProfilePicture
	}
	return nil
}

func (x *ScenePlayBattleSettlePlayerInfo) GetHeadImage() uint32 {
	if x != nil {
		return x.HeadImage
	}
	return 0
}

func (x *ScenePlayBattleSettlePlayerInfo) GetStatisticId() uint32 {
	if x != nil {
		return x.StatisticId
	}
	return 0
}

func (x *ScenePlayBattleSettlePlayerInfo) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ScenePlayBattleSettlePlayerInfo) GetParam() int64 {
	if x != nil {
		return x.Param
	}
	return 0
}

func (x *ScenePlayBattleSettlePlayerInfo) GetOnlineId() string {
	if x != nil {
		return x.OnlineId
	}
	return ""
}

func (x *ScenePlayBattleSettlePlayerInfo) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

var File_ScenePlayBattleSettlePlayerInfo_proto protoreflect.FileDescriptor

var file_ScenePlayBattleSettlePlayerInfo_proto_rawDesc = []byte{
	0x0a, 0x25, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x42, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x45, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x02, 0x0a, 0x1f, 0x53,
	0x63, 0x65, 0x6e, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x53, 0x65,
	0x74, 0x74, 0x6c, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x33,
	0x0a, 0x09, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x45, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x63, 0x61, 0x72, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70,
	0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x0e, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x68, 0x65, 0x61, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x68, 0x65, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x49, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65,
	0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ScenePlayBattleSettlePlayerInfo_proto_rawDescOnce sync.Once
	file_ScenePlayBattleSettlePlayerInfo_proto_rawDescData = file_ScenePlayBattleSettlePlayerInfo_proto_rawDesc
)

func file_ScenePlayBattleSettlePlayerInfo_proto_rawDescGZIP() []byte {
	file_ScenePlayBattleSettlePlayerInfo_proto_rawDescOnce.Do(func() {
		file_ScenePlayBattleSettlePlayerInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ScenePlayBattleSettlePlayerInfo_proto_rawDescData)
	})
	return file_ScenePlayBattleSettlePlayerInfo_proto_rawDescData
}

var file_ScenePlayBattleSettlePlayerInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ScenePlayBattleSettlePlayerInfo_proto_goTypes = []interface{}{
	(*ScenePlayBattleSettlePlayerInfo)(nil), // 0: ScenePlayBattleSettlePlayerInfo
	(*ExhibitionDisplayInfo)(nil),           // 1: ExhibitionDisplayInfo
	(*ProfilePicture)(nil),                  // 2: ProfilePicture
}
var file_ScenePlayBattleSettlePlayerInfo_proto_depIdxs = []int32{
	1, // 0: ScenePlayBattleSettlePlayerInfo.card_list:type_name -> ExhibitionDisplayInfo
	2, // 1: ScenePlayBattleSettlePlayerInfo.profile_picture:type_name -> ProfilePicture
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ScenePlayBattleSettlePlayerInfo_proto_init() }
func file_ScenePlayBattleSettlePlayerInfo_proto_init() {
	if File_ScenePlayBattleSettlePlayerInfo_proto != nil {
		return
	}
	file_ExhibitionDisplayInfo_proto_init()
	file_ProfilePicture_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ScenePlayBattleSettlePlayerInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScenePlayBattleSettlePlayerInfo); i {
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
			RawDescriptor: file_ScenePlayBattleSettlePlayerInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ScenePlayBattleSettlePlayerInfo_proto_goTypes,
		DependencyIndexes: file_ScenePlayBattleSettlePlayerInfo_proto_depIdxs,
		MessageInfos:      file_ScenePlayBattleSettlePlayerInfo_proto_msgTypes,
	}.Build()
	File_ScenePlayBattleSettlePlayerInfo_proto = out.File
	file_ScenePlayBattleSettlePlayerInfo_proto_rawDesc = nil
	file_ScenePlayBattleSettlePlayerInfo_proto_goTypes = nil
	file_ScenePlayBattleSettlePlayerInfo_proto_depIdxs = nil
}
