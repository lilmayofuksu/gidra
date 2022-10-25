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
// source: WindFieldDungeonSettleInfo.proto

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

type WindFieldDungeonSettleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unk3100_ABGAMIOBKAB []uint32            `protobuf:"varint,11,rep,packed,name=Unk3100_ABGAMIOBKAB,json=Unk3100ABGAMIOBKAB,proto3" json:"Unk3100_ABGAMIOBKAB,omitempty"`
	Unk3100_MPGPNBOHCMC []uint32            `protobuf:"varint,7,rep,packed,name=Unk3100_MPGPNBOHCMC,json=Unk3100MPGPNBOHCMC,proto3" json:"Unk3100_MPGPNBOHCMC,omitempty"`
	Unk3100_AOFJAJACNAJ Unk3100_HJALLGOLFGL `protobuf:"varint,2,opt,name=Unk3100_AOFJAJACNAJ,json=Unk3100AOFJAJACNAJ,proto3,enum=Unk3100_HJALLGOLFGL" json:"Unk3100_AOFJAJACNAJ,omitempty"`
}

func (x *WindFieldDungeonSettleInfo) Reset() {
	*x = WindFieldDungeonSettleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WindFieldDungeonSettleInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WindFieldDungeonSettleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WindFieldDungeonSettleInfo) ProtoMessage() {}

func (x *WindFieldDungeonSettleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_WindFieldDungeonSettleInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WindFieldDungeonSettleInfo.ProtoReflect.Descriptor instead.
func (*WindFieldDungeonSettleInfo) Descriptor() ([]byte, []int) {
	return file_WindFieldDungeonSettleInfo_proto_rawDescGZIP(), []int{0}
}

func (x *WindFieldDungeonSettleInfo) GetUnk3100_ABGAMIOBKAB() []uint32 {
	if x != nil {
		return x.Unk3100_ABGAMIOBKAB
	}
	return nil
}

func (x *WindFieldDungeonSettleInfo) GetUnk3100_MPGPNBOHCMC() []uint32 {
	if x != nil {
		return x.Unk3100_MPGPNBOHCMC
	}
	return nil
}

func (x *WindFieldDungeonSettleInfo) GetUnk3100_AOFJAJACNAJ() Unk3100_HJALLGOLFGL {
	if x != nil {
		return x.Unk3100_AOFJAJACNAJ
	}
	return Unk3100_HJALLGOLFGL_Unk3100_HJALLGOLFGL_Unk3100_KAADIPNHPAM
}

var File_WindFieldDungeonSettleInfo_proto protoreflect.FileDescriptor

var file_WindFieldDungeonSettleInfo_proto_rawDesc = []byte{
	0x0a, 0x20, 0x57, 0x69, 0x6e, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x75, 0x6e, 0x67, 0x65,
	0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b, 0x33, 0x31, 0x30, 0x30, 0x5f, 0x48, 0x4a, 0x41, 0x4c,
	0x4c, 0x47, 0x4f, 0x4c, 0x46, 0x47, 0x4c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x01,
	0x0a, 0x1a, 0x57, 0x69, 0x6e, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x75, 0x6e, 0x67, 0x65,
	0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2f, 0x0a, 0x13,
	0x55, 0x6e, 0x6b, 0x33, 0x31, 0x30, 0x30, 0x5f, 0x41, 0x42, 0x47, 0x41, 0x4d, 0x49, 0x4f, 0x42,
	0x4b, 0x41, 0x42, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x31,
	0x30, 0x30, 0x41, 0x42, 0x47, 0x41, 0x4d, 0x49, 0x4f, 0x42, 0x4b, 0x41, 0x42, 0x12, 0x2f, 0x0a,
	0x13, 0x55, 0x6e, 0x6b, 0x33, 0x31, 0x30, 0x30, 0x5f, 0x4d, 0x50, 0x47, 0x50, 0x4e, 0x42, 0x4f,
	0x48, 0x43, 0x4d, 0x43, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33,
	0x31, 0x30, 0x30, 0x4d, 0x50, 0x47, 0x50, 0x4e, 0x42, 0x4f, 0x48, 0x43, 0x4d, 0x43, 0x12, 0x45,
	0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x31, 0x30, 0x30, 0x5f, 0x41, 0x4f, 0x46, 0x4a, 0x41, 0x4a,
	0x41, 0x43, 0x4e, 0x41, 0x4a, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x55, 0x6e,
	0x6b, 0x33, 0x31, 0x30, 0x30, 0x5f, 0x48, 0x4a, 0x41, 0x4c, 0x4c, 0x47, 0x4f, 0x4c, 0x46, 0x47,
	0x4c, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x31, 0x30, 0x30, 0x41, 0x4f, 0x46, 0x4a, 0x41, 0x4a,
	0x41, 0x43, 0x4e, 0x41, 0x4a, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_WindFieldDungeonSettleInfo_proto_rawDescOnce sync.Once
	file_WindFieldDungeonSettleInfo_proto_rawDescData = file_WindFieldDungeonSettleInfo_proto_rawDesc
)

func file_WindFieldDungeonSettleInfo_proto_rawDescGZIP() []byte {
	file_WindFieldDungeonSettleInfo_proto_rawDescOnce.Do(func() {
		file_WindFieldDungeonSettleInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_WindFieldDungeonSettleInfo_proto_rawDescData)
	})
	return file_WindFieldDungeonSettleInfo_proto_rawDescData
}

var file_WindFieldDungeonSettleInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_WindFieldDungeonSettleInfo_proto_goTypes = []interface{}{
	(*WindFieldDungeonSettleInfo)(nil), // 0: WindFieldDungeonSettleInfo
	(Unk3100_HJALLGOLFGL)(0),           // 1: Unk3100_HJALLGOLFGL
}
var file_WindFieldDungeonSettleInfo_proto_depIdxs = []int32{
	1, // 0: WindFieldDungeonSettleInfo.Unk3100_AOFJAJACNAJ:type_name -> Unk3100_HJALLGOLFGL
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_WindFieldDungeonSettleInfo_proto_init() }
func file_WindFieldDungeonSettleInfo_proto_init() {
	if File_WindFieldDungeonSettleInfo_proto != nil {
		return
	}
	file_Unk3100_HJALLGOLFGL_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_WindFieldDungeonSettleInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WindFieldDungeonSettleInfo); i {
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
			RawDescriptor: file_WindFieldDungeonSettleInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WindFieldDungeonSettleInfo_proto_goTypes,
		DependencyIndexes: file_WindFieldDungeonSettleInfo_proto_depIdxs,
		MessageInfos:      file_WindFieldDungeonSettleInfo_proto_msgTypes,
	}.Build()
	File_WindFieldDungeonSettleInfo_proto = out.File
	file_WindFieldDungeonSettleInfo_proto_rawDesc = nil
	file_WindFieldDungeonSettleInfo_proto_goTypes = nil
	file_WindFieldDungeonSettleInfo_proto_depIdxs = nil
}
