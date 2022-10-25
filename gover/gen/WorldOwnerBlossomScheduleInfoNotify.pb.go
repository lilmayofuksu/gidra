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
// source: WorldOwnerBlossomScheduleInfoNotify.proto

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

// CmdId: 2707
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type WorldOwnerBlossomScheduleInfoNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScheduleInfo *BlossomScheduleInfo `protobuf:"bytes,3,opt,name=schedule_info,json=scheduleInfo,proto3" json:"schedule_info,omitempty"`
}

func (x *WorldOwnerBlossomScheduleInfoNotify) Reset() {
	*x = WorldOwnerBlossomScheduleInfoNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WorldOwnerBlossomScheduleInfoNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldOwnerBlossomScheduleInfoNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldOwnerBlossomScheduleInfoNotify) ProtoMessage() {}

func (x *WorldOwnerBlossomScheduleInfoNotify) ProtoReflect() protoreflect.Message {
	mi := &file_WorldOwnerBlossomScheduleInfoNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldOwnerBlossomScheduleInfoNotify.ProtoReflect.Descriptor instead.
func (*WorldOwnerBlossomScheduleInfoNotify) Descriptor() ([]byte, []int) {
	return file_WorldOwnerBlossomScheduleInfoNotify_proto_rawDescGZIP(), []int{0}
}

func (x *WorldOwnerBlossomScheduleInfoNotify) GetScheduleInfo() *BlossomScheduleInfo {
	if x != nil {
		return x.ScheduleInfo
	}
	return nil
}

var File_WorldOwnerBlossomScheduleInfoNotify_proto protoreflect.FileDescriptor

var file_WorldOwnerBlossomScheduleInfoNotify_proto_rawDesc = []byte{
	0x0a, 0x29, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x42, 0x6c, 0x6f, 0x73,
	0x73, 0x6f, 0x6d, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x42, 0x6c, 0x6f,
	0x73, 0x73, 0x6f, 0x6d, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a, 0x23, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x4f,
	0x77, 0x6e, 0x65, 0x72, 0x42, 0x6c, 0x6f, 0x73, 0x73, 0x6f, 0x6d, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x39, 0x0a,
	0x0d, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x42, 0x6c, 0x6f, 0x73, 0x73, 0x6f, 0x6d, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_WorldOwnerBlossomScheduleInfoNotify_proto_rawDescOnce sync.Once
	file_WorldOwnerBlossomScheduleInfoNotify_proto_rawDescData = file_WorldOwnerBlossomScheduleInfoNotify_proto_rawDesc
)

func file_WorldOwnerBlossomScheduleInfoNotify_proto_rawDescGZIP() []byte {
	file_WorldOwnerBlossomScheduleInfoNotify_proto_rawDescOnce.Do(func() {
		file_WorldOwnerBlossomScheduleInfoNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_WorldOwnerBlossomScheduleInfoNotify_proto_rawDescData)
	})
	return file_WorldOwnerBlossomScheduleInfoNotify_proto_rawDescData
}

var file_WorldOwnerBlossomScheduleInfoNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_WorldOwnerBlossomScheduleInfoNotify_proto_goTypes = []interface{}{
	(*WorldOwnerBlossomScheduleInfoNotify)(nil), // 0: WorldOwnerBlossomScheduleInfoNotify
	(*BlossomScheduleInfo)(nil),                 // 1: BlossomScheduleInfo
}
var file_WorldOwnerBlossomScheduleInfoNotify_proto_depIdxs = []int32{
	1, // 0: WorldOwnerBlossomScheduleInfoNotify.schedule_info:type_name -> BlossomScheduleInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_WorldOwnerBlossomScheduleInfoNotify_proto_init() }
func file_WorldOwnerBlossomScheduleInfoNotify_proto_init() {
	if File_WorldOwnerBlossomScheduleInfoNotify_proto != nil {
		return
	}
	file_BlossomScheduleInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_WorldOwnerBlossomScheduleInfoNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorldOwnerBlossomScheduleInfoNotify); i {
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
			RawDescriptor: file_WorldOwnerBlossomScheduleInfoNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WorldOwnerBlossomScheduleInfoNotify_proto_goTypes,
		DependencyIndexes: file_WorldOwnerBlossomScheduleInfoNotify_proto_depIdxs,
		MessageInfos:      file_WorldOwnerBlossomScheduleInfoNotify_proto_msgTypes,
	}.Build()
	File_WorldOwnerBlossomScheduleInfoNotify_proto = out.File
	file_WorldOwnerBlossomScheduleInfoNotify_proto_rawDesc = nil
	file_WorldOwnerBlossomScheduleInfoNotify_proto_goTypes = nil
	file_WorldOwnerBlossomScheduleInfoNotify_proto_depIdxs = nil
}
