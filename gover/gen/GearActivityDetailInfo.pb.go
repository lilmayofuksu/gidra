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
// source: GearActivityDetailInfo.proto

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

type GearActivityDetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unk2800_GBAPCBPMHNJ []*Unk2800_BPOJIIDEADD `protobuf:"bytes,14,rep,name=Unk2800_GBAPCBPMHNJ,json=Unk2800GBAPCBPMHNJ,proto3" json:"Unk2800_GBAPCBPMHNJ,omitempty"`
	Unk2800_IHEHGOBCINC *Unk2800_JIPMJPAKIKE   `protobuf:"bytes,8,opt,name=Unk2800_IHEHGOBCINC,json=Unk2800IHEHGOBCINC,proto3" json:"Unk2800_IHEHGOBCINC,omitempty"`
}

func (x *GearActivityDetailInfo) Reset() {
	*x = GearActivityDetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GearActivityDetailInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GearActivityDetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GearActivityDetailInfo) ProtoMessage() {}

func (x *GearActivityDetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_GearActivityDetailInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GearActivityDetailInfo.ProtoReflect.Descriptor instead.
func (*GearActivityDetailInfo) Descriptor() ([]byte, []int) {
	return file_GearActivityDetailInfo_proto_rawDescGZIP(), []int{0}
}

func (x *GearActivityDetailInfo) GetUnk2800_GBAPCBPMHNJ() []*Unk2800_BPOJIIDEADD {
	if x != nil {
		return x.Unk2800_GBAPCBPMHNJ
	}
	return nil
}

func (x *GearActivityDetailInfo) GetUnk2800_IHEHGOBCINC() *Unk2800_JIPMJPAKIKE {
	if x != nil {
		return x.Unk2800_IHEHGOBCINC
	}
	return nil
}

var File_GearActivityDetailInfo_proto protoreflect.FileDescriptor

var file_GearActivityDetailInfo_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x47, 0x65, 0x61, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x5f, 0x42, 0x50, 0x4f, 0x4a, 0x49, 0x49, 0x44, 0x45,
	0x41, 0x44, 0x44, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x38,
	0x30, 0x30, 0x5f, 0x4a, 0x49, 0x50, 0x4d, 0x4a, 0x50, 0x41, 0x4b, 0x49, 0x4b, 0x45, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x61, 0x72, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x45, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x5f, 0x47, 0x42, 0x41, 0x50, 0x43,
	0x42, 0x50, 0x4d, 0x48, 0x4e, 0x4a, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x55,
	0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x5f, 0x42, 0x50, 0x4f, 0x4a, 0x49, 0x49, 0x44, 0x45, 0x41,
	0x44, 0x44, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x47, 0x42, 0x41, 0x50, 0x43,
	0x42, 0x50, 0x4d, 0x48, 0x4e, 0x4a, 0x12, 0x45, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30,
	0x30, 0x5f, 0x49, 0x48, 0x45, 0x48, 0x47, 0x4f, 0x42, 0x43, 0x49, 0x4e, 0x43, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x5f, 0x4a, 0x49,
	0x50, 0x4d, 0x4a, 0x50, 0x41, 0x4b, 0x49, 0x4b, 0x45, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x38,
	0x30, 0x30, 0x49, 0x48, 0x45, 0x48, 0x47, 0x4f, 0x42, 0x43, 0x49, 0x4e, 0x43, 0x42, 0x06, 0x5a,
	0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GearActivityDetailInfo_proto_rawDescOnce sync.Once
	file_GearActivityDetailInfo_proto_rawDescData = file_GearActivityDetailInfo_proto_rawDesc
)

func file_GearActivityDetailInfo_proto_rawDescGZIP() []byte {
	file_GearActivityDetailInfo_proto_rawDescOnce.Do(func() {
		file_GearActivityDetailInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_GearActivityDetailInfo_proto_rawDescData)
	})
	return file_GearActivityDetailInfo_proto_rawDescData
}

var file_GearActivityDetailInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GearActivityDetailInfo_proto_goTypes = []interface{}{
	(*GearActivityDetailInfo)(nil), // 0: GearActivityDetailInfo
	(*Unk2800_BPOJIIDEADD)(nil),    // 1: Unk2800_BPOJIIDEADD
	(*Unk2800_JIPMJPAKIKE)(nil),    // 2: Unk2800_JIPMJPAKIKE
}
var file_GearActivityDetailInfo_proto_depIdxs = []int32{
	1, // 0: GearActivityDetailInfo.Unk2800_GBAPCBPMHNJ:type_name -> Unk2800_BPOJIIDEADD
	2, // 1: GearActivityDetailInfo.Unk2800_IHEHGOBCINC:type_name -> Unk2800_JIPMJPAKIKE
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_GearActivityDetailInfo_proto_init() }
func file_GearActivityDetailInfo_proto_init() {
	if File_GearActivityDetailInfo_proto != nil {
		return
	}
	file_Unk2800_BPOJIIDEADD_proto_init()
	file_Unk2800_JIPMJPAKIKE_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GearActivityDetailInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GearActivityDetailInfo); i {
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
			RawDescriptor: file_GearActivityDetailInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GearActivityDetailInfo_proto_goTypes,
		DependencyIndexes: file_GearActivityDetailInfo_proto_depIdxs,
		MessageInfos:      file_GearActivityDetailInfo_proto_msgTypes,
	}.Build()
	File_GearActivityDetailInfo_proto = out.File
	file_GearActivityDetailInfo_proto_rawDesc = nil
	file_GearActivityDetailInfo_proto_goTypes = nil
	file_GearActivityDetailInfo_proto_depIdxs = nil
}
