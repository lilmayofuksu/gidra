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
// source: Unk2800_CHEDEMEDPPM.proto

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

// CmdId: 5565
// EnetChannelId: 0
// EnetIsReliable: true
type Unk2800_CHEDEMEDPPM struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PointId             uint32 `protobuf:"varint,7,opt,name=point_id,json=pointId,proto3" json:"point_id,omitempty"`
	Coin                uint32 `protobuf:"varint,15,opt,name=coin,proto3" json:"coin,omitempty"`
	Unk2800_EOFOECJJMLJ uint32 `protobuf:"varint,3,opt,name=Unk2800_EOFOECJJMLJ,json=Unk2800EOFOECJJMLJ,proto3" json:"Unk2800_EOFOECJJMLJ,omitempty"`
	Unk2800_BAEEDEAADIA uint32 `protobuf:"varint,13,opt,name=Unk2800_BAEEDEAADIA,json=Unk2800BAEEDEAADIA,proto3" json:"Unk2800_BAEEDEAADIA,omitempty"`
}

func (x *Unk2800_CHEDEMEDPPM) Reset() {
	*x = Unk2800_CHEDEMEDPPM{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Unk2800_CHEDEMEDPPM_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unk2800_CHEDEMEDPPM) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unk2800_CHEDEMEDPPM) ProtoMessage() {}

func (x *Unk2800_CHEDEMEDPPM) ProtoReflect() protoreflect.Message {
	mi := &file_Unk2800_CHEDEMEDPPM_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unk2800_CHEDEMEDPPM.ProtoReflect.Descriptor instead.
func (*Unk2800_CHEDEMEDPPM) Descriptor() ([]byte, []int) {
	return file_Unk2800_CHEDEMEDPPM_proto_rawDescGZIP(), []int{0}
}

func (x *Unk2800_CHEDEMEDPPM) GetPointId() uint32 {
	if x != nil {
		return x.PointId
	}
	return 0
}

func (x *Unk2800_CHEDEMEDPPM) GetCoin() uint32 {
	if x != nil {
		return x.Coin
	}
	return 0
}

func (x *Unk2800_CHEDEMEDPPM) GetUnk2800_EOFOECJJMLJ() uint32 {
	if x != nil {
		return x.Unk2800_EOFOECJJMLJ
	}
	return 0
}

func (x *Unk2800_CHEDEMEDPPM) GetUnk2800_BAEEDEAADIA() uint32 {
	if x != nil {
		return x.Unk2800_BAEEDEAADIA
	}
	return 0
}

var File_Unk2800_CHEDEMEDPPM_proto protoreflect.FileDescriptor

var file_Unk2800_CHEDEMEDPPM_proto_rawDesc = []byte{
	0x0a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x5f, 0x43, 0x48, 0x45, 0x44, 0x45, 0x4d,
	0x45, 0x44, 0x50, 0x50, 0x4d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x13,
	0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x5f, 0x43, 0x48, 0x45, 0x44, 0x45, 0x4d, 0x45, 0x44,
	0x50, 0x50, 0x4d, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x69, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f,
	0x69, 0x6e, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x5f, 0x45, 0x4f,
	0x46, 0x4f, 0x45, 0x43, 0x4a, 0x4a, 0x4d, 0x4c, 0x4a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x12, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x45, 0x4f, 0x46, 0x4f, 0x45, 0x43, 0x4a, 0x4a,
	0x4d, 0x4c, 0x4a, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x5f, 0x42,
	0x41, 0x45, 0x45, 0x44, 0x45, 0x41, 0x41, 0x44, 0x49, 0x41, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x42, 0x41, 0x45, 0x45, 0x44, 0x45, 0x41,
	0x41, 0x44, 0x49, 0x41, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Unk2800_CHEDEMEDPPM_proto_rawDescOnce sync.Once
	file_Unk2800_CHEDEMEDPPM_proto_rawDescData = file_Unk2800_CHEDEMEDPPM_proto_rawDesc
)

func file_Unk2800_CHEDEMEDPPM_proto_rawDescGZIP() []byte {
	file_Unk2800_CHEDEMEDPPM_proto_rawDescOnce.Do(func() {
		file_Unk2800_CHEDEMEDPPM_proto_rawDescData = protoimpl.X.CompressGZIP(file_Unk2800_CHEDEMEDPPM_proto_rawDescData)
	})
	return file_Unk2800_CHEDEMEDPPM_proto_rawDescData
}

var file_Unk2800_CHEDEMEDPPM_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Unk2800_CHEDEMEDPPM_proto_goTypes = []interface{}{
	(*Unk2800_CHEDEMEDPPM)(nil), // 0: Unk2800_CHEDEMEDPPM
}
var file_Unk2800_CHEDEMEDPPM_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Unk2800_CHEDEMEDPPM_proto_init() }
func file_Unk2800_CHEDEMEDPPM_proto_init() {
	if File_Unk2800_CHEDEMEDPPM_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Unk2800_CHEDEMEDPPM_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unk2800_CHEDEMEDPPM); i {
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
			RawDescriptor: file_Unk2800_CHEDEMEDPPM_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Unk2800_CHEDEMEDPPM_proto_goTypes,
		DependencyIndexes: file_Unk2800_CHEDEMEDPPM_proto_depIdxs,
		MessageInfos:      file_Unk2800_CHEDEMEDPPM_proto_msgTypes,
	}.Build()
	File_Unk2800_CHEDEMEDPPM_proto = out.File
	file_Unk2800_CHEDEMEDPPM_proto_rawDesc = nil
	file_Unk2800_CHEDEMEDPPM_proto_goTypes = nil
	file_Unk2800_CHEDEMEDPPM_proto_depIdxs = nil
}
