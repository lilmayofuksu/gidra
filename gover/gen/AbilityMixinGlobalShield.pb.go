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
// source: AbilityMixinGlobalShield.proto

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

type AbilityMixinGlobalShield struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsCreateEffect   bool    `protobuf:"varint,4,opt,name=is_create_effect,json=isCreateEffect,proto3" json:"is_create_effect,omitempty"`
	SubShield        float32 `protobuf:"fixed32,7,opt,name=sub_shield,json=subShield,proto3" json:"sub_shield,omitempty"`
	HeightOffset     float32 `protobuf:"fixed32,5,opt,name=height_offset,json=heightOffset,proto3" json:"height_offset,omitempty"`
	AvatarId         uint32  `protobuf:"varint,11,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id,omitempty"`
	MaxShield        float32 `protobuf:"fixed32,10,opt,name=max_shield,json=maxShield,proto3" json:"max_shield,omitempty"`
	ShieldEffectName string  `protobuf:"bytes,2,opt,name=shield_effect_name,json=shieldEffectName,proto3" json:"shield_effect_name,omitempty"`
}

func (x *AbilityMixinGlobalShield) Reset() {
	*x = AbilityMixinGlobalShield{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AbilityMixinGlobalShield_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbilityMixinGlobalShield) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbilityMixinGlobalShield) ProtoMessage() {}

func (x *AbilityMixinGlobalShield) ProtoReflect() protoreflect.Message {
	mi := &file_AbilityMixinGlobalShield_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbilityMixinGlobalShield.ProtoReflect.Descriptor instead.
func (*AbilityMixinGlobalShield) Descriptor() ([]byte, []int) {
	return file_AbilityMixinGlobalShield_proto_rawDescGZIP(), []int{0}
}

func (x *AbilityMixinGlobalShield) GetIsCreateEffect() bool {
	if x != nil {
		return x.IsCreateEffect
	}
	return false
}

func (x *AbilityMixinGlobalShield) GetSubShield() float32 {
	if x != nil {
		return x.SubShield
	}
	return 0
}

func (x *AbilityMixinGlobalShield) GetHeightOffset() float32 {
	if x != nil {
		return x.HeightOffset
	}
	return 0
}

func (x *AbilityMixinGlobalShield) GetAvatarId() uint32 {
	if x != nil {
		return x.AvatarId
	}
	return 0
}

func (x *AbilityMixinGlobalShield) GetMaxShield() float32 {
	if x != nil {
		return x.MaxShield
	}
	return 0
}

func (x *AbilityMixinGlobalShield) GetShieldEffectName() string {
	if x != nil {
		return x.ShieldEffectName
	}
	return ""
}

var File_AbilityMixinGlobalShield_proto protoreflect.FileDescriptor

var file_AbilityMixinGlobalShield_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4d, 0x69, 0x78, 0x69, 0x6e, 0x47, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x53, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xf2, 0x01, 0x0a, 0x18, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4d, 0x69, 0x78, 0x69,
	0x6e, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x53, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x28, 0x0a,
	0x10, 0x69, 0x73, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x65, 0x66, 0x66, 0x65, 0x63,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x5f, 0x73,
	0x68, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x73, 0x75, 0x62,
	0x53, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x5f,
	0x73, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6d, 0x61,
	0x78, 0x53, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x68, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x45, 0x66, 0x66, 0x65, 0x63,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AbilityMixinGlobalShield_proto_rawDescOnce sync.Once
	file_AbilityMixinGlobalShield_proto_rawDescData = file_AbilityMixinGlobalShield_proto_rawDesc
)

func file_AbilityMixinGlobalShield_proto_rawDescGZIP() []byte {
	file_AbilityMixinGlobalShield_proto_rawDescOnce.Do(func() {
		file_AbilityMixinGlobalShield_proto_rawDescData = protoimpl.X.CompressGZIP(file_AbilityMixinGlobalShield_proto_rawDescData)
	})
	return file_AbilityMixinGlobalShield_proto_rawDescData
}

var file_AbilityMixinGlobalShield_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AbilityMixinGlobalShield_proto_goTypes = []interface{}{
	(*AbilityMixinGlobalShield)(nil), // 0: AbilityMixinGlobalShield
}
var file_AbilityMixinGlobalShield_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AbilityMixinGlobalShield_proto_init() }
func file_AbilityMixinGlobalShield_proto_init() {
	if File_AbilityMixinGlobalShield_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AbilityMixinGlobalShield_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbilityMixinGlobalShield); i {
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
			RawDescriptor: file_AbilityMixinGlobalShield_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AbilityMixinGlobalShield_proto_goTypes,
		DependencyIndexes: file_AbilityMixinGlobalShield_proto_depIdxs,
		MessageInfos:      file_AbilityMixinGlobalShield_proto_msgTypes,
	}.Build()
	File_AbilityMixinGlobalShield_proto = out.File
	file_AbilityMixinGlobalShield_proto_rawDesc = nil
	file_AbilityMixinGlobalShield_proto_goTypes = nil
	file_AbilityMixinGlobalShield_proto_depIdxs = nil
}
