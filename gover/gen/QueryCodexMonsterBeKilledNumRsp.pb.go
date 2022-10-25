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
// source: QueryCodexMonsterBeKilledNumRsp.proto

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

// CmdId: 4209
// EnetChannelId: 0
// EnetIsReliable: true
type QueryCodexMonsterBeKilledNumRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CodexIdList         []uint32 `protobuf:"varint,4,rep,packed,name=codex_id_list,json=codexIdList,proto3" json:"codex_id_list,omitempty"`
	Unk2700_MKOBMGGPNMI []uint32 `protobuf:"varint,6,rep,packed,name=Unk2700_MKOBMGGPNMI,json=Unk2700MKOBMGGPNMI,proto3" json:"Unk2700_MKOBMGGPNMI,omitempty"`
	BeKilledNumList     []uint32 `protobuf:"varint,12,rep,packed,name=be_killed_num_list,json=beKilledNumList,proto3" json:"be_killed_num_list,omitempty"`
	Retcode             int32    `protobuf:"varint,5,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *QueryCodexMonsterBeKilledNumRsp) Reset() {
	*x = QueryCodexMonsterBeKilledNumRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_QueryCodexMonsterBeKilledNumRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryCodexMonsterBeKilledNumRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryCodexMonsterBeKilledNumRsp) ProtoMessage() {}

func (x *QueryCodexMonsterBeKilledNumRsp) ProtoReflect() protoreflect.Message {
	mi := &file_QueryCodexMonsterBeKilledNumRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryCodexMonsterBeKilledNumRsp.ProtoReflect.Descriptor instead.
func (*QueryCodexMonsterBeKilledNumRsp) Descriptor() ([]byte, []int) {
	return file_QueryCodexMonsterBeKilledNumRsp_proto_rawDescGZIP(), []int{0}
}

func (x *QueryCodexMonsterBeKilledNumRsp) GetCodexIdList() []uint32 {
	if x != nil {
		return x.CodexIdList
	}
	return nil
}

func (x *QueryCodexMonsterBeKilledNumRsp) GetUnk2700_MKOBMGGPNMI() []uint32 {
	if x != nil {
		return x.Unk2700_MKOBMGGPNMI
	}
	return nil
}

func (x *QueryCodexMonsterBeKilledNumRsp) GetBeKilledNumList() []uint32 {
	if x != nil {
		return x.BeKilledNumList
	}
	return nil
}

func (x *QueryCodexMonsterBeKilledNumRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_QueryCodexMonsterBeKilledNumRsp_proto protoreflect.FileDescriptor

var file_QueryCodexMonsterBeKilledNumRsp_proto_rawDesc = []byte{
	0x0a, 0x25, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x78, 0x4d, 0x6f, 0x6e, 0x73,
	0x74, 0x65, 0x72, 0x42, 0x65, 0x4b, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x4e, 0x75, 0x6d, 0x52, 0x73,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x01, 0x0a, 0x1f, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x78, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x42, 0x65, 0x4b,
	0x69, 0x6c, 0x6c, 0x65, 0x64, 0x4e, 0x75, 0x6d, 0x52, 0x73, 0x70, 0x12, 0x22, 0x0a, 0x0d, 0x63,
	0x6f, 0x64, 0x65, 0x78, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x0b, 0x63, 0x6f, 0x64, 0x65, 0x78, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4d, 0x4b, 0x4f, 0x42, 0x4d,
	0x47, 0x47, 0x50, 0x4e, 0x4d, 0x49, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e,
	0x6b, 0x32, 0x37, 0x30, 0x30, 0x4d, 0x4b, 0x4f, 0x42, 0x4d, 0x47, 0x47, 0x50, 0x4e, 0x4d, 0x49,
	0x12, 0x2b, 0x0a, 0x12, 0x62, 0x65, 0x5f, 0x6b, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x5f, 0x6e, 0x75,
	0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0f, 0x62, 0x65,
	0x4b, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x4e, 0x75, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_QueryCodexMonsterBeKilledNumRsp_proto_rawDescOnce sync.Once
	file_QueryCodexMonsterBeKilledNumRsp_proto_rawDescData = file_QueryCodexMonsterBeKilledNumRsp_proto_rawDesc
)

func file_QueryCodexMonsterBeKilledNumRsp_proto_rawDescGZIP() []byte {
	file_QueryCodexMonsterBeKilledNumRsp_proto_rawDescOnce.Do(func() {
		file_QueryCodexMonsterBeKilledNumRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_QueryCodexMonsterBeKilledNumRsp_proto_rawDescData)
	})
	return file_QueryCodexMonsterBeKilledNumRsp_proto_rawDescData
}

var file_QueryCodexMonsterBeKilledNumRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_QueryCodexMonsterBeKilledNumRsp_proto_goTypes = []interface{}{
	(*QueryCodexMonsterBeKilledNumRsp)(nil), // 0: QueryCodexMonsterBeKilledNumRsp
}
var file_QueryCodexMonsterBeKilledNumRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_QueryCodexMonsterBeKilledNumRsp_proto_init() }
func file_QueryCodexMonsterBeKilledNumRsp_proto_init() {
	if File_QueryCodexMonsterBeKilledNumRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_QueryCodexMonsterBeKilledNumRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryCodexMonsterBeKilledNumRsp); i {
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
			RawDescriptor: file_QueryCodexMonsterBeKilledNumRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_QueryCodexMonsterBeKilledNumRsp_proto_goTypes,
		DependencyIndexes: file_QueryCodexMonsterBeKilledNumRsp_proto_depIdxs,
		MessageInfos:      file_QueryCodexMonsterBeKilledNumRsp_proto_msgTypes,
	}.Build()
	File_QueryCodexMonsterBeKilledNumRsp_proto = out.File
	file_QueryCodexMonsterBeKilledNumRsp_proto_rawDesc = nil
	file_QueryCodexMonsterBeKilledNumRsp_proto_goTypes = nil
	file_QueryCodexMonsterBeKilledNumRsp_proto_depIdxs = nil
}
