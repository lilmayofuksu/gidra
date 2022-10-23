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
// source: GetDungeonEntryExploreConditionRsp.proto

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

// CmdId: 3269
// EnetChannelId: 0
// EnetIsReliable: true
type GetDungeonEntryExploreConditionRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DungeonEntryCond *DungeonEntryCond `protobuf:"bytes,5,opt,name=dungeon_entry_cond,json=dungeonEntryCond,proto3" json:"dungeon_entry_cond,omitempty"`
	Retcode          int32             `protobuf:"varint,3,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *GetDungeonEntryExploreConditionRsp) Reset() {
	*x = GetDungeonEntryExploreConditionRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetDungeonEntryExploreConditionRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDungeonEntryExploreConditionRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDungeonEntryExploreConditionRsp) ProtoMessage() {}

func (x *GetDungeonEntryExploreConditionRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetDungeonEntryExploreConditionRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDungeonEntryExploreConditionRsp.ProtoReflect.Descriptor instead.
func (*GetDungeonEntryExploreConditionRsp) Descriptor() ([]byte, []int) {
	return file_GetDungeonEntryExploreConditionRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetDungeonEntryExploreConditionRsp) GetDungeonEntryCond() *DungeonEntryCond {
	if x != nil {
		return x.DungeonEntryCond
	}
	return nil
}

func (x *GetDungeonEntryExploreConditionRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_GetDungeonEntryExploreConditionRsp_proto protoreflect.FileDescriptor

var file_GetDungeonEntryExploreConditionRsp_proto_rawDesc = []byte{
	0x0a, 0x28, 0x47, 0x65, 0x74, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x44, 0x75, 0x6e, 0x67,
	0x65, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x3f, 0x0a, 0x12, 0x64, 0x75, 0x6e, 0x67,
	0x65, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x64, 0x52, 0x10, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63,
	0x6f, 0x64, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_GetDungeonEntryExploreConditionRsp_proto_rawDescOnce sync.Once
	file_GetDungeonEntryExploreConditionRsp_proto_rawDescData = file_GetDungeonEntryExploreConditionRsp_proto_rawDesc
)

func file_GetDungeonEntryExploreConditionRsp_proto_rawDescGZIP() []byte {
	file_GetDungeonEntryExploreConditionRsp_proto_rawDescOnce.Do(func() {
		file_GetDungeonEntryExploreConditionRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetDungeonEntryExploreConditionRsp_proto_rawDescData)
	})
	return file_GetDungeonEntryExploreConditionRsp_proto_rawDescData
}

var file_GetDungeonEntryExploreConditionRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetDungeonEntryExploreConditionRsp_proto_goTypes = []interface{}{
	(*GetDungeonEntryExploreConditionRsp)(nil), // 0: GetDungeonEntryExploreConditionRsp
	(*DungeonEntryCond)(nil),                   // 1: DungeonEntryCond
}
var file_GetDungeonEntryExploreConditionRsp_proto_depIdxs = []int32{
	1, // 0: GetDungeonEntryExploreConditionRsp.dungeon_entry_cond:type_name -> DungeonEntryCond
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GetDungeonEntryExploreConditionRsp_proto_init() }
func file_GetDungeonEntryExploreConditionRsp_proto_init() {
	if File_GetDungeonEntryExploreConditionRsp_proto != nil {
		return
	}
	file_DungeonEntryCond_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetDungeonEntryExploreConditionRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDungeonEntryExploreConditionRsp); i {
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
			RawDescriptor: file_GetDungeonEntryExploreConditionRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetDungeonEntryExploreConditionRsp_proto_goTypes,
		DependencyIndexes: file_GetDungeonEntryExploreConditionRsp_proto_depIdxs,
		MessageInfos:      file_GetDungeonEntryExploreConditionRsp_proto_msgTypes,
	}.Build()
	File_GetDungeonEntryExploreConditionRsp_proto = out.File
	file_GetDungeonEntryExploreConditionRsp_proto_rawDesc = nil
	file_GetDungeonEntryExploreConditionRsp_proto_goTypes = nil
	file_GetDungeonEntryExploreConditionRsp_proto_depIdxs = nil
}
