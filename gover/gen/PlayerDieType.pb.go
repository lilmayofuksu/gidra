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
// source: PlayerDieType.proto

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

type PlayerDieType int32

const (
	PlayerDieType_PLAYER_DIE_TYPE_NONE            PlayerDieType = 0
	PlayerDieType_PLAYER_DIE_TYPE_KILL_BY_MONSTER PlayerDieType = 1
	PlayerDieType_PLAYER_DIE_TYPE_KILL_BY_GEAR    PlayerDieType = 2
	PlayerDieType_PLAYER_DIE_TYPE_FALL            PlayerDieType = 3
	PlayerDieType_PLAYER_DIE_TYPE_DRAWN           PlayerDieType = 4
	PlayerDieType_PLAYER_DIE_TYPE_ABYSS           PlayerDieType = 5
	PlayerDieType_PLAYER_DIE_TYPE_GM              PlayerDieType = 6
	PlayerDieType_PLAYER_DIE_TYPE_CLIMATE_COLD    PlayerDieType = 7
	PlayerDieType_PLAYER_DIE_TYPE_STORM_LIGHTING  PlayerDieType = 8
)

// Enum value maps for PlayerDieType.
var (
	PlayerDieType_name = map[int32]string{
		0: "PLAYER_DIE_TYPE_NONE",
		1: "PLAYER_DIE_TYPE_KILL_BY_MONSTER",
		2: "PLAYER_DIE_TYPE_KILL_BY_GEAR",
		3: "PLAYER_DIE_TYPE_FALL",
		4: "PLAYER_DIE_TYPE_DRAWN",
		5: "PLAYER_DIE_TYPE_ABYSS",
		6: "PLAYER_DIE_TYPE_GM",
		7: "PLAYER_DIE_TYPE_CLIMATE_COLD",
		8: "PLAYER_DIE_TYPE_STORM_LIGHTING",
	}
	PlayerDieType_value = map[string]int32{
		"PLAYER_DIE_TYPE_NONE":            0,
		"PLAYER_DIE_TYPE_KILL_BY_MONSTER": 1,
		"PLAYER_DIE_TYPE_KILL_BY_GEAR":    2,
		"PLAYER_DIE_TYPE_FALL":            3,
		"PLAYER_DIE_TYPE_DRAWN":           4,
		"PLAYER_DIE_TYPE_ABYSS":           5,
		"PLAYER_DIE_TYPE_GM":              6,
		"PLAYER_DIE_TYPE_CLIMATE_COLD":    7,
		"PLAYER_DIE_TYPE_STORM_LIGHTING":  8,
	}
)

func (x PlayerDieType) Enum() *PlayerDieType {
	p := new(PlayerDieType)
	*p = x
	return p
}

func (x PlayerDieType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerDieType) Descriptor() protoreflect.EnumDescriptor {
	return file_PlayerDieType_proto_enumTypes[0].Descriptor()
}

func (PlayerDieType) Type() protoreflect.EnumType {
	return &file_PlayerDieType_proto_enumTypes[0]
}

func (x PlayerDieType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerDieType.Descriptor instead.
func (PlayerDieType) EnumDescriptor() ([]byte, []int) {
	return file_PlayerDieType_proto_rawDescGZIP(), []int{0}
}

var File_PlayerDieType_proto protoreflect.FileDescriptor

var file_PlayerDieType_proto_rawDesc = []byte{
	0x0a, 0x13, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x69, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x9e, 0x02, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x44, 0x69, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x4c, 0x41, 0x59, 0x45,
	0x52, 0x5f, 0x44, 0x49, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10,
	0x00, 0x12, 0x23, 0x0a, 0x1f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x44, 0x49, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4b, 0x49, 0x4c, 0x4c, 0x5f, 0x42, 0x59, 0x5f, 0x4d, 0x4f, 0x4e,
	0x53, 0x54, 0x45, 0x52, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52,
	0x5f, 0x44, 0x49, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4b, 0x49, 0x4c, 0x4c, 0x5f, 0x42,
	0x59, 0x5f, 0x47, 0x45, 0x41, 0x52, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x4c, 0x41, 0x59,
	0x45, 0x52, 0x5f, 0x44, 0x49, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x41, 0x4c, 0x4c,
	0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x44, 0x49, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x52, 0x41, 0x57, 0x4e, 0x10, 0x04, 0x12, 0x19, 0x0a,
	0x15, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x44, 0x49, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x41, 0x42, 0x59, 0x53, 0x53, 0x10, 0x05, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x4c, 0x41, 0x59,
	0x45, 0x52, 0x5f, 0x44, 0x49, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x4d, 0x10, 0x06,
	0x12, 0x20, 0x0a, 0x1c, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x44, 0x49, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x43, 0x4c, 0x49, 0x4d, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x4f, 0x4c, 0x44,
	0x10, 0x07, 0x12, 0x22, 0x0a, 0x1e, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x44, 0x49, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x4d, 0x5f, 0x4c, 0x49, 0x47, 0x48,
	0x54, 0x49, 0x4e, 0x47, 0x10, 0x08, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerDieType_proto_rawDescOnce sync.Once
	file_PlayerDieType_proto_rawDescData = file_PlayerDieType_proto_rawDesc
)

func file_PlayerDieType_proto_rawDescGZIP() []byte {
	file_PlayerDieType_proto_rawDescOnce.Do(func() {
		file_PlayerDieType_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerDieType_proto_rawDescData)
	})
	return file_PlayerDieType_proto_rawDescData
}

var file_PlayerDieType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_PlayerDieType_proto_goTypes = []interface{}{
	(PlayerDieType)(0), // 0: PlayerDieType
}
var file_PlayerDieType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_PlayerDieType_proto_init() }
func file_PlayerDieType_proto_init() {
	if File_PlayerDieType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_PlayerDieType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerDieType_proto_goTypes,
		DependencyIndexes: file_PlayerDieType_proto_depIdxs,
		EnumInfos:         file_PlayerDieType_proto_enumTypes,
	}.Build()
	File_PlayerDieType_proto = out.File
	file_PlayerDieType_proto_rawDesc = nil
	file_PlayerDieType_proto_goTypes = nil
	file_PlayerDieType_proto_depIdxs = nil
}
