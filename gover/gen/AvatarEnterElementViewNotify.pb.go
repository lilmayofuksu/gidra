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
// source: AvatarEnterElementViewNotify.proto

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

// CmdId: 334
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type AvatarEnterElementViewNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsTriggerd     bool   `protobuf:"varint,3,opt,name=is_triggerd,json=isTriggerd,proto3" json:"is_triggerd,omitempty"`
	AvatarEntityId uint32 `protobuf:"varint,12,opt,name=avatar_entity_id,json=avatarEntityId,proto3" json:"avatar_entity_id,omitempty"`
}

func (x *AvatarEnterElementViewNotify) Reset() {
	*x = AvatarEnterElementViewNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AvatarEnterElementViewNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarEnterElementViewNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarEnterElementViewNotify) ProtoMessage() {}

func (x *AvatarEnterElementViewNotify) ProtoReflect() protoreflect.Message {
	mi := &file_AvatarEnterElementViewNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarEnterElementViewNotify.ProtoReflect.Descriptor instead.
func (*AvatarEnterElementViewNotify) Descriptor() ([]byte, []int) {
	return file_AvatarEnterElementViewNotify_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarEnterElementViewNotify) GetIsTriggerd() bool {
	if x != nil {
		return x.IsTriggerd
	}
	return false
}

func (x *AvatarEnterElementViewNotify) GetAvatarEntityId() uint32 {
	if x != nil {
		return x.AvatarEntityId
	}
	return 0
}

var File_AvatarEnterElementViewNotify_proto protoreflect.FileDescriptor

var file_AvatarEnterElementViewNotify_proto_rawDesc = []byte{
	0x0a, 0x22, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x45, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x56, 0x69, 0x65, 0x77, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x1c, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x6e,
	0x74, 0x65, 0x72, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x69, 0x65, 0x77, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x74, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x54, 0x72, 0x69,
	0x67, 0x67, 0x65, 0x72, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0e, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x42,
	0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AvatarEnterElementViewNotify_proto_rawDescOnce sync.Once
	file_AvatarEnterElementViewNotify_proto_rawDescData = file_AvatarEnterElementViewNotify_proto_rawDesc
)

func file_AvatarEnterElementViewNotify_proto_rawDescGZIP() []byte {
	file_AvatarEnterElementViewNotify_proto_rawDescOnce.Do(func() {
		file_AvatarEnterElementViewNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_AvatarEnterElementViewNotify_proto_rawDescData)
	})
	return file_AvatarEnterElementViewNotify_proto_rawDescData
}

var file_AvatarEnterElementViewNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AvatarEnterElementViewNotify_proto_goTypes = []interface{}{
	(*AvatarEnterElementViewNotify)(nil), // 0: AvatarEnterElementViewNotify
}
var file_AvatarEnterElementViewNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AvatarEnterElementViewNotify_proto_init() }
func file_AvatarEnterElementViewNotify_proto_init() {
	if File_AvatarEnterElementViewNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AvatarEnterElementViewNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarEnterElementViewNotify); i {
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
			RawDescriptor: file_AvatarEnterElementViewNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AvatarEnterElementViewNotify_proto_goTypes,
		DependencyIndexes: file_AvatarEnterElementViewNotify_proto_depIdxs,
		MessageInfos:      file_AvatarEnterElementViewNotify_proto_msgTypes,
	}.Build()
	File_AvatarEnterElementViewNotify_proto = out.File
	file_AvatarEnterElementViewNotify_proto_rawDesc = nil
	file_AvatarEnterElementViewNotify_proto_goTypes = nil
	file_AvatarEnterElementViewNotify_proto_depIdxs = nil
}
