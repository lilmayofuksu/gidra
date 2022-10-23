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
// source: GalleryStageType.proto

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

type GalleryStageType int32

const (
	GalleryStageType_GALLERY_STAGE_TYPE_NONE     GalleryStageType = 0
	GalleryStageType_GALLERY_STAGE_TYPE_PRESTART GalleryStageType = 1
	GalleryStageType_GALLERY_STAGE_TYPE_START    GalleryStageType = 2
)

// Enum value maps for GalleryStageType.
var (
	GalleryStageType_name = map[int32]string{
		0: "GALLERY_STAGE_TYPE_NONE",
		1: "GALLERY_STAGE_TYPE_PRESTART",
		2: "GALLERY_STAGE_TYPE_START",
	}
	GalleryStageType_value = map[string]int32{
		"GALLERY_STAGE_TYPE_NONE":     0,
		"GALLERY_STAGE_TYPE_PRESTART": 1,
		"GALLERY_STAGE_TYPE_START":    2,
	}
)

func (x GalleryStageType) Enum() *GalleryStageType {
	p := new(GalleryStageType)
	*p = x
	return p
}

func (x GalleryStageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GalleryStageType) Descriptor() protoreflect.EnumDescriptor {
	return file_GalleryStageType_proto_enumTypes[0].Descriptor()
}

func (GalleryStageType) Type() protoreflect.EnumType {
	return &file_GalleryStageType_proto_enumTypes[0]
}

func (x GalleryStageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GalleryStageType.Descriptor instead.
func (GalleryStageType) EnumDescriptor() ([]byte, []int) {
	return file_GalleryStageType_proto_rawDescGZIP(), []int{0}
}

var File_GalleryStageType_proto protoreflect.FileDescriptor

var file_GalleryStageType_proto_rawDesc = []byte{
	0x0a, 0x16, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x6e, 0x0a, 0x10, 0x47, 0x61, 0x6c, 0x6c,
	0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17,
	0x47, 0x41, 0x4c, 0x4c, 0x45, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x1b, 0x47, 0x41, 0x4c,
	0x4c, 0x45, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x50, 0x52, 0x45, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x47, 0x41,
	0x4c, 0x4c, 0x45, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x02, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GalleryStageType_proto_rawDescOnce sync.Once
	file_GalleryStageType_proto_rawDescData = file_GalleryStageType_proto_rawDesc
)

func file_GalleryStageType_proto_rawDescGZIP() []byte {
	file_GalleryStageType_proto_rawDescOnce.Do(func() {
		file_GalleryStageType_proto_rawDescData = protoimpl.X.CompressGZIP(file_GalleryStageType_proto_rawDescData)
	})
	return file_GalleryStageType_proto_rawDescData
}

var file_GalleryStageType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_GalleryStageType_proto_goTypes = []interface{}{
	(GalleryStageType)(0), // 0: GalleryStageType
}
var file_GalleryStageType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GalleryStageType_proto_init() }
func file_GalleryStageType_proto_init() {
	if File_GalleryStageType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_GalleryStageType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GalleryStageType_proto_goTypes,
		DependencyIndexes: file_GalleryStageType_proto_depIdxs,
		EnumInfos:         file_GalleryStageType_proto_enumTypes,
	}.Build()
	File_GalleryStageType_proto = out.File
	file_GalleryStageType_proto_rawDesc = nil
	file_GalleryStageType_proto_goTypes = nil
	file_GalleryStageType_proto_depIdxs = nil
}
