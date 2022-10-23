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
// source: PlayerEnterSceneNotify.proto

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

// CmdId: 272
// EnetChannelId: 0
// EnetIsReliable: true
type PlayerEnterSceneNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrevSceneId            uint32    `protobuf:"varint,6,opt,name=prev_scene_id,json=prevSceneId,proto3" json:"prev_scene_id,omitempty"`
	DungeonId              uint32    `protobuf:"varint,12,opt,name=dungeon_id,json=dungeonId,proto3" json:"dungeon_id,omitempty"`
	IsSkipUi               bool      `protobuf:"varint,1732,opt,name=is_skip_ui,json=isSkipUi,proto3" json:"is_skip_ui,omitempty"`
	SceneId                uint32    `protobuf:"varint,15,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	Type                   EnterType `protobuf:"varint,13,opt,name=type,proto3,enum=EnterType" json:"type,omitempty"`
	SceneBeginTime         uint64    `protobuf:"varint,14,opt,name=scene_begin_time,json=sceneBeginTime,proto3" json:"scene_begin_time,omitempty"`
	WorldLevel             uint32    `protobuf:"varint,11,opt,name=world_level,json=worldLevel,proto3" json:"world_level,omitempty"`
	WorldType              uint32    `protobuf:"varint,1490,opt,name=world_type,json=worldType,proto3" json:"world_type,omitempty"`
	TargetUid              uint32    `protobuf:"varint,4,opt,name=target_uid,json=targetUid,proto3" json:"target_uid,omitempty"`
	IsFirstLoginEnterScene bool      `protobuf:"varint,3,opt,name=is_first_login_enter_scene,json=isFirstLoginEnterScene,proto3" json:"is_first_login_enter_scene,omitempty"`
	SceneTagIdList         []uint32  `protobuf:"varint,5,rep,packed,name=scene_tag_id_list,json=sceneTagIdList,proto3" json:"scene_tag_id_list,omitempty"`
	SceneTransaction       string    `protobuf:"bytes,1842,opt,name=scene_transaction,json=sceneTransaction,proto3" json:"scene_transaction,omitempty"`
	PrevPos                *Vector   `protobuf:"bytes,8,opt,name=prev_pos,json=prevPos,proto3" json:"prev_pos,omitempty"`
	EnterReason            uint32    `protobuf:"varint,1828,opt,name=enter_reason,json=enterReason,proto3" json:"enter_reason,omitempty"`
	Pos                    *Vector   `protobuf:"bytes,7,opt,name=pos,proto3" json:"pos,omitempty"`
	EnterSceneToken        uint32    `protobuf:"varint,2,opt,name=enter_scene_token,json=enterSceneToken,proto3" json:"enter_scene_token,omitempty"`
}

func (x *PlayerEnterSceneNotify) Reset() {
	*x = PlayerEnterSceneNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerEnterSceneNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerEnterSceneNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerEnterSceneNotify) ProtoMessage() {}

func (x *PlayerEnterSceneNotify) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerEnterSceneNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerEnterSceneNotify.ProtoReflect.Descriptor instead.
func (*PlayerEnterSceneNotify) Descriptor() ([]byte, []int) {
	return file_PlayerEnterSceneNotify_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerEnterSceneNotify) GetPrevSceneId() uint32 {
	if x != nil {
		return x.PrevSceneId
	}
	return 0
}

func (x *PlayerEnterSceneNotify) GetDungeonId() uint32 {
	if x != nil {
		return x.DungeonId
	}
	return 0
}

func (x *PlayerEnterSceneNotify) GetIsSkipUi() bool {
	if x != nil {
		return x.IsSkipUi
	}
	return false
}

func (x *PlayerEnterSceneNotify) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

func (x *PlayerEnterSceneNotify) GetType() EnterType {
	if x != nil {
		return x.Type
	}
	return EnterType_ENTER_TYPE_NONE
}

func (x *PlayerEnterSceneNotify) GetSceneBeginTime() uint64 {
	if x != nil {
		return x.SceneBeginTime
	}
	return 0
}

func (x *PlayerEnterSceneNotify) GetWorldLevel() uint32 {
	if x != nil {
		return x.WorldLevel
	}
	return 0
}

func (x *PlayerEnterSceneNotify) GetWorldType() uint32 {
	if x != nil {
		return x.WorldType
	}
	return 0
}

func (x *PlayerEnterSceneNotify) GetTargetUid() uint32 {
	if x != nil {
		return x.TargetUid
	}
	return 0
}

func (x *PlayerEnterSceneNotify) GetIsFirstLoginEnterScene() bool {
	if x != nil {
		return x.IsFirstLoginEnterScene
	}
	return false
}

func (x *PlayerEnterSceneNotify) GetSceneTagIdList() []uint32 {
	if x != nil {
		return x.SceneTagIdList
	}
	return nil
}

func (x *PlayerEnterSceneNotify) GetSceneTransaction() string {
	if x != nil {
		return x.SceneTransaction
	}
	return ""
}

func (x *PlayerEnterSceneNotify) GetPrevPos() *Vector {
	if x != nil {
		return x.PrevPos
	}
	return nil
}

func (x *PlayerEnterSceneNotify) GetEnterReason() uint32 {
	if x != nil {
		return x.EnterReason
	}
	return 0
}

func (x *PlayerEnterSceneNotify) GetPos() *Vector {
	if x != nil {
		return x.Pos
	}
	return nil
}

func (x *PlayerEnterSceneNotify) GetEnterSceneToken() uint32 {
	if x != nil {
		return x.EnterSceneToken
	}
	return 0
}

var File_PlayerEnterSceneNotify_proto protoreflect.FileDescriptor

var file_PlayerEnterSceneNotify_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x63, 0x65,
	0x6e, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f,
	0x45, 0x6e, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x04,
	0x0a, 0x16, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x63, 0x65,
	0x6e, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x22, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x76,
	0x5f, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x70, 0x72, 0x65, 0x76, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69,
	0x73, 0x5f, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x75, 0x69, 0x18, 0xc4, 0x0d, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x73, 0x53, 0x6b, 0x69, 0x70, 0x55, 0x69, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63,
	0x65, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x63,
	0x65, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x62,
	0x65, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0e, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0xd2,
	0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x69, 0x64, 0x12,
	0x3a, 0x0a, 0x1a, 0x69, 0x73, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x5f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x16, 0x69, 0x73, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x12, 0x29, 0x0a, 0x11, 0x73,
	0x63, 0x65, 0x6e, 0x65, 0x5f, 0x74, 0x61, 0x67, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0e, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x54, 0x61, 0x67,
	0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0xb2, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x70, 0x6f, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52,
	0x07, 0x70, 0x72, 0x65, 0x76, 0x50, 0x6f, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0xa4, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x03,
	0x70, 0x6f, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x5f, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0f, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_PlayerEnterSceneNotify_proto_rawDescOnce sync.Once
	file_PlayerEnterSceneNotify_proto_rawDescData = file_PlayerEnterSceneNotify_proto_rawDesc
)

func file_PlayerEnterSceneNotify_proto_rawDescGZIP() []byte {
	file_PlayerEnterSceneNotify_proto_rawDescOnce.Do(func() {
		file_PlayerEnterSceneNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerEnterSceneNotify_proto_rawDescData)
	})
	return file_PlayerEnterSceneNotify_proto_rawDescData
}

var file_PlayerEnterSceneNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerEnterSceneNotify_proto_goTypes = []interface{}{
	(*PlayerEnterSceneNotify)(nil), // 0: PlayerEnterSceneNotify
	(EnterType)(0),                 // 1: EnterType
	(*Vector)(nil),                 // 2: Vector
}
var file_PlayerEnterSceneNotify_proto_depIdxs = []int32{
	1, // 0: PlayerEnterSceneNotify.type:type_name -> EnterType
	2, // 1: PlayerEnterSceneNotify.prev_pos:type_name -> Vector
	2, // 2: PlayerEnterSceneNotify.pos:type_name -> Vector
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_PlayerEnterSceneNotify_proto_init() }
func file_PlayerEnterSceneNotify_proto_init() {
	if File_PlayerEnterSceneNotify_proto != nil {
		return
	}
	file_EnterType_proto_init()
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerEnterSceneNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerEnterSceneNotify); i {
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
			RawDescriptor: file_PlayerEnterSceneNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerEnterSceneNotify_proto_goTypes,
		DependencyIndexes: file_PlayerEnterSceneNotify_proto_depIdxs,
		MessageInfos:      file_PlayerEnterSceneNotify_proto_msgTypes,
	}.Build()
	File_PlayerEnterSceneNotify_proto = out.File
	file_PlayerEnterSceneNotify_proto_rawDesc = nil
	file_PlayerEnterSceneNotify_proto_goTypes = nil
	file_PlayerEnterSceneNotify_proto_depIdxs = nil
}
