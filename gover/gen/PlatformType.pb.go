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
// source: PlatformType.proto

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

type PlatformType int32

const (
	PlatformType_PLATFORM_TYPE_EDITOR              PlatformType = 0
	PlatformType_PLATFORM_TYPE_IOS                 PlatformType = 1
	PlatformType_PLATFORM_TYPE_ANDROID             PlatformType = 2
	PlatformType_PLATFORM_TYPE_PC                  PlatformType = 3
	PlatformType_PLATFORM_TYPE_PS4                 PlatformType = 4
	PlatformType_PLATFORM_TYPE_SERVER              PlatformType = 5
	PlatformType_PLATFORM_TYPE_CLOUD_ANDROID       PlatformType = 6
	PlatformType_PLATFORM_TYPE_CLOUD_IOS           PlatformType = 7
	PlatformType_PLATFORM_TYPE_PS5                 PlatformType = 8
	PlatformType_PLATFORM_TYPE_CLOUD_WEB           PlatformType = 9
	PlatformType_PLATFORM_TYPE_CLOUD_TV            PlatformType = 10
	PlatformType_PLATFORM_TYPE_Unk2700_IBBEKBJLMAJ PlatformType = 11
	PlatformType_PLATFORM_TYPE_Unk2700_BCEICMDNIIG PlatformType = 12
	PlatformType_PLATFORM_TYPE_Unk2800_EFNGHFNPMKM PlatformType = 13
	PlatformType_PLATFORM_TYPE_Unk2800_FNFHGPABLFB PlatformType = 14
)

// Enum value maps for PlatformType.
var (
	PlatformType_name = map[int32]string{
		0:  "PLATFORM_TYPE_EDITOR",
		1:  "PLATFORM_TYPE_IOS",
		2:  "PLATFORM_TYPE_ANDROID",
		3:  "PLATFORM_TYPE_PC",
		4:  "PLATFORM_TYPE_PS4",
		5:  "PLATFORM_TYPE_SERVER",
		6:  "PLATFORM_TYPE_CLOUD_ANDROID",
		7:  "PLATFORM_TYPE_CLOUD_IOS",
		8:  "PLATFORM_TYPE_PS5",
		9:  "PLATFORM_TYPE_CLOUD_WEB",
		10: "PLATFORM_TYPE_CLOUD_TV",
		11: "PLATFORM_TYPE_Unk2700_IBBEKBJLMAJ",
		12: "PLATFORM_TYPE_Unk2700_BCEICMDNIIG",
		13: "PLATFORM_TYPE_Unk2800_EFNGHFNPMKM",
		14: "PLATFORM_TYPE_Unk2800_FNFHGPABLFB",
	}
	PlatformType_value = map[string]int32{
		"PLATFORM_TYPE_EDITOR":              0,
		"PLATFORM_TYPE_IOS":                 1,
		"PLATFORM_TYPE_ANDROID":             2,
		"PLATFORM_TYPE_PC":                  3,
		"PLATFORM_TYPE_PS4":                 4,
		"PLATFORM_TYPE_SERVER":              5,
		"PLATFORM_TYPE_CLOUD_ANDROID":       6,
		"PLATFORM_TYPE_CLOUD_IOS":           7,
		"PLATFORM_TYPE_PS5":                 8,
		"PLATFORM_TYPE_CLOUD_WEB":           9,
		"PLATFORM_TYPE_CLOUD_TV":            10,
		"PLATFORM_TYPE_Unk2700_IBBEKBJLMAJ": 11,
		"PLATFORM_TYPE_Unk2700_BCEICMDNIIG": 12,
		"PLATFORM_TYPE_Unk2800_EFNGHFNPMKM": 13,
		"PLATFORM_TYPE_Unk2800_FNFHGPABLFB": 14,
	}
)

func (x PlatformType) Enum() *PlatformType {
	p := new(PlatformType)
	*p = x
	return p
}

func (x PlatformType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlatformType) Descriptor() protoreflect.EnumDescriptor {
	return file_PlatformType_proto_enumTypes[0].Descriptor()
}

func (PlatformType) Type() protoreflect.EnumType {
	return &file_PlatformType_proto_enumTypes[0]
}

func (x PlatformType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlatformType.Descriptor instead.
func (PlatformType) EnumDescriptor() ([]byte, []int) {
	return file_PlatformType_proto_rawDescGZIP(), []int{0}
}

var File_PlatformType_proto protoreflect.FileDescriptor

var file_PlatformType_proto_rawDesc = []byte{
	0x0a, 0x12, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xcb, 0x03, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52,
	0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x44, 0x49, 0x54, 0x4f, 0x52, 0x10, 0x00, 0x12,
	0x15, 0x0a, 0x11, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x49, 0x4f, 0x53, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f,
	0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x10,
	0x02, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x50, 0x43, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x4c, 0x41, 0x54, 0x46,
	0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x53, 0x34, 0x10, 0x04, 0x12, 0x18,
	0x0a, 0x14, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x05, 0x12, 0x1f, 0x0a, 0x1b, 0x50, 0x4c, 0x41, 0x54,
	0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f,
	0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x10, 0x06, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x4c, 0x41,
	0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4c, 0x4f, 0x55, 0x44,
	0x5f, 0x49, 0x4f, 0x53, 0x10, 0x07, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f,
	0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x53, 0x35, 0x10, 0x08, 0x12, 0x1b, 0x0a,
	0x17, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43,
	0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x57, 0x45, 0x42, 0x10, 0x09, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x4c,
	0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4c, 0x4f, 0x55,
	0x44, 0x5f, 0x54, 0x56, 0x10, 0x0a, 0x12, 0x25, 0x0a, 0x21, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f,
	0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f,
	0x49, 0x42, 0x42, 0x45, 0x4b, 0x42, 0x4a, 0x4c, 0x4d, 0x41, 0x4a, 0x10, 0x0b, 0x12, 0x25, 0x0a,
	0x21, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x42, 0x43, 0x45, 0x49, 0x43, 0x4d, 0x44, 0x4e, 0x49,
	0x49, 0x47, 0x10, 0x0c, 0x12, 0x25, 0x0a, 0x21, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x5f, 0x45, 0x46,
	0x4e, 0x47, 0x48, 0x46, 0x4e, 0x50, 0x4d, 0x4b, 0x4d, 0x10, 0x0d, 0x12, 0x25, 0x0a, 0x21, 0x50,
	0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x6e, 0x6b,
	0x32, 0x38, 0x30, 0x30, 0x5f, 0x46, 0x4e, 0x46, 0x48, 0x47, 0x50, 0x41, 0x42, 0x4c, 0x46, 0x42,
	0x10, 0x0e, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_PlatformType_proto_rawDescOnce sync.Once
	file_PlatformType_proto_rawDescData = file_PlatformType_proto_rawDesc
)

func file_PlatformType_proto_rawDescGZIP() []byte {
	file_PlatformType_proto_rawDescOnce.Do(func() {
		file_PlatformType_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlatformType_proto_rawDescData)
	})
	return file_PlatformType_proto_rawDescData
}

var file_PlatformType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_PlatformType_proto_goTypes = []interface{}{
	(PlatformType)(0), // 0: PlatformType
}
var file_PlatformType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_PlatformType_proto_init() }
func file_PlatformType_proto_init() {
	if File_PlatformType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_PlatformType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlatformType_proto_goTypes,
		DependencyIndexes: file_PlatformType_proto_depIdxs,
		EnumInfos:         file_PlatformType_proto_enumTypes,
	}.Build()
	File_PlatformType_proto = out.File
	file_PlatformType_proto_rawDesc = nil
	file_PlatformType_proto_goTypes = nil
	file_PlatformType_proto_depIdxs = nil
}
