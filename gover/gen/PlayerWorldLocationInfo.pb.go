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
// source: PlayerWorldLocationInfo.proto

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

type PlayerWorldLocationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SceneId   uint32              `protobuf:"varint,1,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	PlayerLoc *PlayerLocationInfo `protobuf:"bytes,12,opt,name=player_loc,json=playerLoc,proto3" json:"player_loc,omitempty"`
}

func (x *PlayerWorldLocationInfo) Reset() {
	*x = PlayerWorldLocationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerWorldLocationInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerWorldLocationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerWorldLocationInfo) ProtoMessage() {}

func (x *PlayerWorldLocationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerWorldLocationInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerWorldLocationInfo.ProtoReflect.Descriptor instead.
func (*PlayerWorldLocationInfo) Descriptor() ([]byte, []int) {
	return file_PlayerWorldLocationInfo_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerWorldLocationInfo) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

func (x *PlayerWorldLocationInfo) GetPlayerLoc() *PlayerLocationInfo {
	if x != nil {
		return x.PlayerLoc
	}
	return nil
}

var File_PlayerWorldLocationInfo_proto protoreflect.FileDescriptor

var file_PlayerWorldLocationInfo_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x18, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x17, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x64, 0x12,
	0x32, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x63, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x4c, 0x6f, 0x63, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_PlayerWorldLocationInfo_proto_rawDescOnce sync.Once
	file_PlayerWorldLocationInfo_proto_rawDescData = file_PlayerWorldLocationInfo_proto_rawDesc
)

func file_PlayerWorldLocationInfo_proto_rawDescGZIP() []byte {
	file_PlayerWorldLocationInfo_proto_rawDescOnce.Do(func() {
		file_PlayerWorldLocationInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerWorldLocationInfo_proto_rawDescData)
	})
	return file_PlayerWorldLocationInfo_proto_rawDescData
}

var file_PlayerWorldLocationInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerWorldLocationInfo_proto_goTypes = []interface{}{
	(*PlayerWorldLocationInfo)(nil), // 0: PlayerWorldLocationInfo
	(*PlayerLocationInfo)(nil),      // 1: PlayerLocationInfo
}
var file_PlayerWorldLocationInfo_proto_depIdxs = []int32{
	1, // 0: PlayerWorldLocationInfo.player_loc:type_name -> PlayerLocationInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PlayerWorldLocationInfo_proto_init() }
func file_PlayerWorldLocationInfo_proto_init() {
	if File_PlayerWorldLocationInfo_proto != nil {
		return
	}
	file_PlayerLocationInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerWorldLocationInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerWorldLocationInfo); i {
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
			RawDescriptor: file_PlayerWorldLocationInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerWorldLocationInfo_proto_goTypes,
		DependencyIndexes: file_PlayerWorldLocationInfo_proto_depIdxs,
		MessageInfos:      file_PlayerWorldLocationInfo_proto_msgTypes,
	}.Build()
	File_PlayerWorldLocationInfo_proto = out.File
	file_PlayerWorldLocationInfo_proto_rawDesc = nil
	file_PlayerWorldLocationInfo_proto_goTypes = nil
	file_PlayerWorldLocationInfo_proto_depIdxs = nil
}
