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
// source: GadgetBornType.proto

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

type GadgetBornType int32

const (
	GadgetBornType_GADGET_BORN_TYPE_NONE        GadgetBornType = 0
	GadgetBornType_GADGET_BORN_TYPE_IN_AIR      GadgetBornType = 1
	GadgetBornType_GADGET_BORN_TYPE_PLAYER      GadgetBornType = 2
	GadgetBornType_GADGET_BORN_TYPE_MONSTER_HIT GadgetBornType = 3
	GadgetBornType_GADGET_BORN_TYPE_MONSTER_DIE GadgetBornType = 4
	GadgetBornType_GADGET_BORN_TYPE_GADGET      GadgetBornType = 5
	GadgetBornType_GADGET_BORN_TYPE_GROUND      GadgetBornType = 6
)

// Enum value maps for GadgetBornType.
var (
	GadgetBornType_name = map[int32]string{
		0: "GADGET_BORN_TYPE_NONE",
		1: "GADGET_BORN_TYPE_IN_AIR",
		2: "GADGET_BORN_TYPE_PLAYER",
		3: "GADGET_BORN_TYPE_MONSTER_HIT",
		4: "GADGET_BORN_TYPE_MONSTER_DIE",
		5: "GADGET_BORN_TYPE_GADGET",
		6: "GADGET_BORN_TYPE_GROUND",
	}
	GadgetBornType_value = map[string]int32{
		"GADGET_BORN_TYPE_NONE":        0,
		"GADGET_BORN_TYPE_IN_AIR":      1,
		"GADGET_BORN_TYPE_PLAYER":      2,
		"GADGET_BORN_TYPE_MONSTER_HIT": 3,
		"GADGET_BORN_TYPE_MONSTER_DIE": 4,
		"GADGET_BORN_TYPE_GADGET":      5,
		"GADGET_BORN_TYPE_GROUND":      6,
	}
)

func (x GadgetBornType) Enum() *GadgetBornType {
	p := new(GadgetBornType)
	*p = x
	return p
}

func (x GadgetBornType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GadgetBornType) Descriptor() protoreflect.EnumDescriptor {
	return file_GadgetBornType_proto_enumTypes[0].Descriptor()
}

func (GadgetBornType) Type() protoreflect.EnumType {
	return &file_GadgetBornType_proto_enumTypes[0]
}

func (x GadgetBornType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GadgetBornType.Descriptor instead.
func (GadgetBornType) EnumDescriptor() ([]byte, []int) {
	return file_GadgetBornType_proto_rawDescGZIP(), []int{0}
}

var File_GadgetBornType_proto protoreflect.FileDescriptor

var file_GadgetBornType_proto_rawDesc = []byte{
	0x0a, 0x14, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x42, 0x6f, 0x72, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xe3, 0x01, 0x0a, 0x0e, 0x47, 0x61, 0x64, 0x67, 0x65,
	0x74, 0x42, 0x6f, 0x72, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x47, 0x41, 0x44,
	0x47, 0x45, 0x54, 0x5f, 0x42, 0x4f, 0x52, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f,
	0x4e, 0x45, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x47, 0x41, 0x44, 0x47, 0x45, 0x54, 0x5f, 0x42,
	0x4f, 0x52, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x5f, 0x41, 0x49, 0x52, 0x10,
	0x01, 0x12, 0x1b, 0x0a, 0x17, 0x47, 0x41, 0x44, 0x47, 0x45, 0x54, 0x5f, 0x42, 0x4f, 0x52, 0x4e,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x10, 0x02, 0x12, 0x20,
	0x0a, 0x1c, 0x47, 0x41, 0x44, 0x47, 0x45, 0x54, 0x5f, 0x42, 0x4f, 0x52, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4d, 0x4f, 0x4e, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x48, 0x49, 0x54, 0x10, 0x03,
	0x12, 0x20, 0x0a, 0x1c, 0x47, 0x41, 0x44, 0x47, 0x45, 0x54, 0x5f, 0x42, 0x4f, 0x52, 0x4e, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x4f, 0x4e, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x44, 0x49, 0x45,
	0x10, 0x04, 0x12, 0x1b, 0x0a, 0x17, 0x47, 0x41, 0x44, 0x47, 0x45, 0x54, 0x5f, 0x42, 0x4f, 0x52,
	0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x41, 0x44, 0x47, 0x45, 0x54, 0x10, 0x05, 0x12,
	0x1b, 0x0a, 0x17, 0x47, 0x41, 0x44, 0x47, 0x45, 0x54, 0x5f, 0x42, 0x4f, 0x52, 0x4e, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x06, 0x42, 0x06, 0x5a, 0x04,
	0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GadgetBornType_proto_rawDescOnce sync.Once
	file_GadgetBornType_proto_rawDescData = file_GadgetBornType_proto_rawDesc
)

func file_GadgetBornType_proto_rawDescGZIP() []byte {
	file_GadgetBornType_proto_rawDescOnce.Do(func() {
		file_GadgetBornType_proto_rawDescData = protoimpl.X.CompressGZIP(file_GadgetBornType_proto_rawDescData)
	})
	return file_GadgetBornType_proto_rawDescData
}

var file_GadgetBornType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_GadgetBornType_proto_goTypes = []interface{}{
	(GadgetBornType)(0), // 0: GadgetBornType
}
var file_GadgetBornType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GadgetBornType_proto_init() }
func file_GadgetBornType_proto_init() {
	if File_GadgetBornType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_GadgetBornType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GadgetBornType_proto_goTypes,
		DependencyIndexes: file_GadgetBornType_proto_depIdxs,
		EnumInfos:         file_GadgetBornType_proto_enumTypes,
	}.Build()
	File_GadgetBornType_proto = out.File
	file_GadgetBornType_proto_rawDesc = nil
	file_GadgetBornType_proto_goTypes = nil
	file_GadgetBornType_proto_depIdxs = nil
}
