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
// source: AuthorityChange.proto

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

type AuthorityChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityAuthorityInfo *EntityAuthorityInfo `protobuf:"bytes,5,opt,name=entity_authority_info,json=entityAuthorityInfo,proto3" json:"entity_authority_info,omitempty"`
	AuthorityPeerId     uint32               `protobuf:"varint,3,opt,name=authority_peer_id,json=authorityPeerId,proto3" json:"authority_peer_id,omitempty"`
	EntityId            uint32               `protobuf:"varint,13,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
}

func (x *AuthorityChange) Reset() {
	*x = AuthorityChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AuthorityChange_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorityChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorityChange) ProtoMessage() {}

func (x *AuthorityChange) ProtoReflect() protoreflect.Message {
	mi := &file_AuthorityChange_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorityChange.ProtoReflect.Descriptor instead.
func (*AuthorityChange) Descriptor() ([]byte, []int) {
	return file_AuthorityChange_proto_rawDescGZIP(), []int{0}
}

func (x *AuthorityChange) GetEntityAuthorityInfo() *EntityAuthorityInfo {
	if x != nil {
		return x.EntityAuthorityInfo
	}
	return nil
}

func (x *AuthorityChange) GetAuthorityPeerId() uint32 {
	if x != nil {
		return x.AuthorityPeerId
	}
	return 0
}

func (x *AuthorityChange) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

var File_AuthorityChange_proto protoreflect.FileDescriptor

var file_AuthorityChange_proto_rawDesc = []byte{
	0x0a, 0x15, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa4, 0x01, 0x0a, 0x0f, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x48, 0x0a, 0x15, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x13, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x2a, 0x0a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x65,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x50, 0x65, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AuthorityChange_proto_rawDescOnce sync.Once
	file_AuthorityChange_proto_rawDescData = file_AuthorityChange_proto_rawDesc
)

func file_AuthorityChange_proto_rawDescGZIP() []byte {
	file_AuthorityChange_proto_rawDescOnce.Do(func() {
		file_AuthorityChange_proto_rawDescData = protoimpl.X.CompressGZIP(file_AuthorityChange_proto_rawDescData)
	})
	return file_AuthorityChange_proto_rawDescData
}

var file_AuthorityChange_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AuthorityChange_proto_goTypes = []interface{}{
	(*AuthorityChange)(nil),     // 0: AuthorityChange
	(*EntityAuthorityInfo)(nil), // 1: EntityAuthorityInfo
}
var file_AuthorityChange_proto_depIdxs = []int32{
	1, // 0: AuthorityChange.entity_authority_info:type_name -> EntityAuthorityInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_AuthorityChange_proto_init() }
func file_AuthorityChange_proto_init() {
	if File_AuthorityChange_proto != nil {
		return
	}
	file_EntityAuthorityInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AuthorityChange_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorityChange); i {
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
			RawDescriptor: file_AuthorityChange_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AuthorityChange_proto_goTypes,
		DependencyIndexes: file_AuthorityChange_proto_depIdxs,
		MessageInfos:      file_AuthorityChange_proto_msgTypes,
	}.Build()
	File_AuthorityChange_proto = out.File
	file_AuthorityChange_proto_rawDesc = nil
	file_AuthorityChange_proto_goTypes = nil
	file_AuthorityChange_proto_depIdxs = nil
}
