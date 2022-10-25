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
// source: Unk2700_JOEPIGNPDGH.proto

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

type Unk2700_JOEPIGNPDGH int32

const (
	Unk2700_JOEPIGNPDGH_Unk2700_JOEPIGNPDGH_Unk2700_GIGONJIGKBM Unk2700_JOEPIGNPDGH = 0
	Unk2700_JOEPIGNPDGH_Unk2700_JOEPIGNPDGH_Unk2700_AEKNMJMKIPN Unk2700_JOEPIGNPDGH = 1
	Unk2700_JOEPIGNPDGH_Unk2700_JOEPIGNPDGH_Unk2700_LKCIHNNHIFO Unk2700_JOEPIGNPDGH = 2
	Unk2700_JOEPIGNPDGH_Unk2700_JOEPIGNPDGH_Unk2700_EPAPGLMBAEB Unk2700_JOEPIGNPDGH = 3
)

// Enum value maps for Unk2700_JOEPIGNPDGH.
var (
	Unk2700_JOEPIGNPDGH_name = map[int32]string{
		0: "Unk2700_JOEPIGNPDGH_Unk2700_GIGONJIGKBM",
		1: "Unk2700_JOEPIGNPDGH_Unk2700_AEKNMJMKIPN",
		2: "Unk2700_JOEPIGNPDGH_Unk2700_LKCIHNNHIFO",
		3: "Unk2700_JOEPIGNPDGH_Unk2700_EPAPGLMBAEB",
	}
	Unk2700_JOEPIGNPDGH_value = map[string]int32{
		"Unk2700_JOEPIGNPDGH_Unk2700_GIGONJIGKBM": 0,
		"Unk2700_JOEPIGNPDGH_Unk2700_AEKNMJMKIPN": 1,
		"Unk2700_JOEPIGNPDGH_Unk2700_LKCIHNNHIFO": 2,
		"Unk2700_JOEPIGNPDGH_Unk2700_EPAPGLMBAEB": 3,
	}
)

func (x Unk2700_JOEPIGNPDGH) Enum() *Unk2700_JOEPIGNPDGH {
	p := new(Unk2700_JOEPIGNPDGH)
	*p = x
	return p
}

func (x Unk2700_JOEPIGNPDGH) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Unk2700_JOEPIGNPDGH) Descriptor() protoreflect.EnumDescriptor {
	return file_Unk2700_JOEPIGNPDGH_proto_enumTypes[0].Descriptor()
}

func (Unk2700_JOEPIGNPDGH) Type() protoreflect.EnumType {
	return &file_Unk2700_JOEPIGNPDGH_proto_enumTypes[0]
}

func (x Unk2700_JOEPIGNPDGH) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Unk2700_JOEPIGNPDGH.Descriptor instead.
func (Unk2700_JOEPIGNPDGH) EnumDescriptor() ([]byte, []int) {
	return file_Unk2700_JOEPIGNPDGH_proto_rawDescGZIP(), []int{0}
}

var File_Unk2700_JOEPIGNPDGH_proto protoreflect.FileDescriptor

var file_Unk2700_JOEPIGNPDGH_proto_rawDesc = []byte{
	0x0a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4a, 0x4f, 0x45, 0x50, 0x49, 0x47,
	0x4e, 0x50, 0x44, 0x47, 0x48, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xc9, 0x01, 0x0a, 0x13,
	0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4a, 0x4f, 0x45, 0x50, 0x49, 0x47, 0x4e, 0x50,
	0x44, 0x47, 0x48, 0x12, 0x2b, 0x0a, 0x27, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4a,
	0x4f, 0x45, 0x50, 0x49, 0x47, 0x4e, 0x50, 0x44, 0x47, 0x48, 0x5f, 0x55, 0x6e, 0x6b, 0x32, 0x37,
	0x30, 0x30, 0x5f, 0x47, 0x49, 0x47, 0x4f, 0x4e, 0x4a, 0x49, 0x47, 0x4b, 0x42, 0x4d, 0x10, 0x00,
	0x12, 0x2b, 0x0a, 0x27, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4a, 0x4f, 0x45, 0x50,
	0x49, 0x47, 0x4e, 0x50, 0x44, 0x47, 0x48, 0x5f, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f,
	0x41, 0x45, 0x4b, 0x4e, 0x4d, 0x4a, 0x4d, 0x4b, 0x49, 0x50, 0x4e, 0x10, 0x01, 0x12, 0x2b, 0x0a,
	0x27, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4a, 0x4f, 0x45, 0x50, 0x49, 0x47, 0x4e,
	0x50, 0x44, 0x47, 0x48, 0x5f, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4c, 0x4b, 0x43,
	0x49, 0x48, 0x4e, 0x4e, 0x48, 0x49, 0x46, 0x4f, 0x10, 0x02, 0x12, 0x2b, 0x0a, 0x27, 0x55, 0x6e,
	0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4a, 0x4f, 0x45, 0x50, 0x49, 0x47, 0x4e, 0x50, 0x44, 0x47,
	0x48, 0x5f, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x45, 0x50, 0x41, 0x50, 0x47, 0x4c,
	0x4d, 0x42, 0x41, 0x45, 0x42, 0x10, 0x03, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Unk2700_JOEPIGNPDGH_proto_rawDescOnce sync.Once
	file_Unk2700_JOEPIGNPDGH_proto_rawDescData = file_Unk2700_JOEPIGNPDGH_proto_rawDesc
)

func file_Unk2700_JOEPIGNPDGH_proto_rawDescGZIP() []byte {
	file_Unk2700_JOEPIGNPDGH_proto_rawDescOnce.Do(func() {
		file_Unk2700_JOEPIGNPDGH_proto_rawDescData = protoimpl.X.CompressGZIP(file_Unk2700_JOEPIGNPDGH_proto_rawDescData)
	})
	return file_Unk2700_JOEPIGNPDGH_proto_rawDescData
}

var file_Unk2700_JOEPIGNPDGH_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Unk2700_JOEPIGNPDGH_proto_goTypes = []interface{}{
	(Unk2700_JOEPIGNPDGH)(0), // 0: Unk2700_JOEPIGNPDGH
}
var file_Unk2700_JOEPIGNPDGH_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Unk2700_JOEPIGNPDGH_proto_init() }
func file_Unk2700_JOEPIGNPDGH_proto_init() {
	if File_Unk2700_JOEPIGNPDGH_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Unk2700_JOEPIGNPDGH_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Unk2700_JOEPIGNPDGH_proto_goTypes,
		DependencyIndexes: file_Unk2700_JOEPIGNPDGH_proto_depIdxs,
		EnumInfos:         file_Unk2700_JOEPIGNPDGH_proto_enumTypes,
	}.Build()
	File_Unk2700_JOEPIGNPDGH_proto = out.File
	file_Unk2700_JOEPIGNPDGH_proto_rawDesc = nil
	file_Unk2700_JOEPIGNPDGH_proto_goTypes = nil
	file_Unk2700_JOEPIGNPDGH_proto_depIdxs = nil
}
