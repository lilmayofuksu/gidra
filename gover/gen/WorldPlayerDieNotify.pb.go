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
// source: WorldPlayerDieNotify.proto

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

// CmdId: 285
// EnetChannelId: 0
// EnetIsReliable: true
type WorldPlayerDieNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DieType          PlayerDieType `protobuf:"varint,12,opt,name=die_type,json=dieType,proto3,enum=PlayerDieType" json:"die_type,omitempty"`
	MurdererEntityId uint32        `protobuf:"varint,15,opt,name=murderer_entity_id,json=murdererEntityId,proto3" json:"murderer_entity_id,omitempty"`
	// Types that are assignable to Entity:
	//
	//	*WorldPlayerDieNotify_MonsterId
	//	*WorldPlayerDieNotify_GadgetId
	Entity isWorldPlayerDieNotify_Entity `protobuf_oneof:"entity"`
}

func (x *WorldPlayerDieNotify) Reset() {
	*x = WorldPlayerDieNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WorldPlayerDieNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldPlayerDieNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldPlayerDieNotify) ProtoMessage() {}

func (x *WorldPlayerDieNotify) ProtoReflect() protoreflect.Message {
	mi := &file_WorldPlayerDieNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldPlayerDieNotify.ProtoReflect.Descriptor instead.
func (*WorldPlayerDieNotify) Descriptor() ([]byte, []int) {
	return file_WorldPlayerDieNotify_proto_rawDescGZIP(), []int{0}
}

func (x *WorldPlayerDieNotify) GetDieType() PlayerDieType {
	if x != nil {
		return x.DieType
	}
	return PlayerDieType_PLAYER_DIE_TYPE_NONE
}

func (x *WorldPlayerDieNotify) GetMurdererEntityId() uint32 {
	if x != nil {
		return x.MurdererEntityId
	}
	return 0
}

func (m *WorldPlayerDieNotify) GetEntity() isWorldPlayerDieNotify_Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (x *WorldPlayerDieNotify) GetMonsterId() uint32 {
	if x, ok := x.GetEntity().(*WorldPlayerDieNotify_MonsterId); ok {
		return x.MonsterId
	}
	return 0
}

func (x *WorldPlayerDieNotify) GetGadgetId() uint32 {
	if x, ok := x.GetEntity().(*WorldPlayerDieNotify_GadgetId); ok {
		return x.GadgetId
	}
	return 0
}

type isWorldPlayerDieNotify_Entity interface {
	isWorldPlayerDieNotify_Entity()
}

type WorldPlayerDieNotify_MonsterId struct {
	MonsterId uint32 `protobuf:"varint,8,opt,name=monster_id,json=monsterId,proto3,oneof"`
}

type WorldPlayerDieNotify_GadgetId struct {
	GadgetId uint32 `protobuf:"varint,4,opt,name=gadget_id,json=gadgetId,proto3,oneof"`
}

func (*WorldPlayerDieNotify_MonsterId) isWorldPlayerDieNotify_Entity() {}

func (*WorldPlayerDieNotify_GadgetId) isWorldPlayerDieNotify_Entity() {}

var File_WorldPlayerDieNotify_proto protoreflect.FileDescriptor

var file_WorldPlayerDieNotify_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x69, 0x65,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x44, 0x69, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb9, 0x01, 0x0a, 0x14, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x44, 0x69, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x29, 0x0a, 0x08, 0x64, 0x69,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x69, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x64, 0x69,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x6d, 0x75, 0x72, 0x64, 0x65, 0x72, 0x65,
	0x72, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x10, 0x6d, 0x75, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0a, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x09, 0x6d, 0x6f, 0x6e, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x09, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x08, 0x67, 0x61, 0x64, 0x67, 0x65,
	0x74, 0x49, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x06, 0x5a,
	0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_WorldPlayerDieNotify_proto_rawDescOnce sync.Once
	file_WorldPlayerDieNotify_proto_rawDescData = file_WorldPlayerDieNotify_proto_rawDesc
)

func file_WorldPlayerDieNotify_proto_rawDescGZIP() []byte {
	file_WorldPlayerDieNotify_proto_rawDescOnce.Do(func() {
		file_WorldPlayerDieNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_WorldPlayerDieNotify_proto_rawDescData)
	})
	return file_WorldPlayerDieNotify_proto_rawDescData
}

var file_WorldPlayerDieNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_WorldPlayerDieNotify_proto_goTypes = []interface{}{
	(*WorldPlayerDieNotify)(nil), // 0: WorldPlayerDieNotify
	(PlayerDieType)(0),           // 1: PlayerDieType
}
var file_WorldPlayerDieNotify_proto_depIdxs = []int32{
	1, // 0: WorldPlayerDieNotify.die_type:type_name -> PlayerDieType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_WorldPlayerDieNotify_proto_init() }
func file_WorldPlayerDieNotify_proto_init() {
	if File_WorldPlayerDieNotify_proto != nil {
		return
	}
	file_PlayerDieType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_WorldPlayerDieNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorldPlayerDieNotify); i {
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
	file_WorldPlayerDieNotify_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*WorldPlayerDieNotify_MonsterId)(nil),
		(*WorldPlayerDieNotify_GadgetId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_WorldPlayerDieNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WorldPlayerDieNotify_proto_goTypes,
		DependencyIndexes: file_WorldPlayerDieNotify_proto_depIdxs,
		MessageInfos:      file_WorldPlayerDieNotify_proto_msgTypes,
	}.Build()
	File_WorldPlayerDieNotify_proto = out.File
	file_WorldPlayerDieNotify_proto_rawDesc = nil
	file_WorldPlayerDieNotify_proto_goTypes = nil
	file_WorldPlayerDieNotify_proto_depIdxs = nil
}
