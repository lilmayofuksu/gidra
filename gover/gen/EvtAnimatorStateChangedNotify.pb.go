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
// source: EvtAnimatorStateChangedNotify.proto

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

// CmdId: 331
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type EvtAnimatorStateChangedNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ForwardType                 ForwardType                  `protobuf:"varint,3,opt,name=forward_type,json=forwardType,proto3,enum=ForwardType" json:"forward_type,omitempty"`
	EvtAnimatorStateChangedInfo *EvtAnimatorStateChangedInfo `protobuf:"bytes,10,opt,name=evt_animator_state_changed_info,json=evtAnimatorStateChangedInfo,proto3" json:"evt_animator_state_changed_info,omitempty"`
}

func (x *EvtAnimatorStateChangedNotify) Reset() {
	*x = EvtAnimatorStateChangedNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EvtAnimatorStateChangedNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvtAnimatorStateChangedNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvtAnimatorStateChangedNotify) ProtoMessage() {}

func (x *EvtAnimatorStateChangedNotify) ProtoReflect() protoreflect.Message {
	mi := &file_EvtAnimatorStateChangedNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvtAnimatorStateChangedNotify.ProtoReflect.Descriptor instead.
func (*EvtAnimatorStateChangedNotify) Descriptor() ([]byte, []int) {
	return file_EvtAnimatorStateChangedNotify_proto_rawDescGZIP(), []int{0}
}

func (x *EvtAnimatorStateChangedNotify) GetForwardType() ForwardType {
	if x != nil {
		return x.ForwardType
	}
	return ForwardType_FORWARD_TYPE_LOCAL
}

func (x *EvtAnimatorStateChangedNotify) GetEvtAnimatorStateChangedInfo() *EvtAnimatorStateChangedInfo {
	if x != nil {
		return x.EvtAnimatorStateChangedInfo
	}
	return nil
}

var File_EvtAnimatorStateChangedNotify_proto protoreflect.FileDescriptor

var file_EvtAnimatorStateChangedNotify_proto_rawDesc = []byte{
	0x0a, 0x23, 0x45, 0x76, 0x74, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x45, 0x76, 0x74, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x74,
	0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x01, 0x0a, 0x1d,
	0x45, 0x76, 0x74, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x2f, 0x0a,
	0x0c, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0b, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x62,
	0x0a, 0x1f, 0x65, 0x76, 0x74, 0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x45, 0x76, 0x74, 0x41, 0x6e, 0x69,
	0x6d, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x1b, 0x65, 0x76, 0x74, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x74,
	0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_EvtAnimatorStateChangedNotify_proto_rawDescOnce sync.Once
	file_EvtAnimatorStateChangedNotify_proto_rawDescData = file_EvtAnimatorStateChangedNotify_proto_rawDesc
)

func file_EvtAnimatorStateChangedNotify_proto_rawDescGZIP() []byte {
	file_EvtAnimatorStateChangedNotify_proto_rawDescOnce.Do(func() {
		file_EvtAnimatorStateChangedNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_EvtAnimatorStateChangedNotify_proto_rawDescData)
	})
	return file_EvtAnimatorStateChangedNotify_proto_rawDescData
}

var file_EvtAnimatorStateChangedNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EvtAnimatorStateChangedNotify_proto_goTypes = []interface{}{
	(*EvtAnimatorStateChangedNotify)(nil), // 0: EvtAnimatorStateChangedNotify
	(ForwardType)(0),                      // 1: ForwardType
	(*EvtAnimatorStateChangedInfo)(nil),   // 2: EvtAnimatorStateChangedInfo
}
var file_EvtAnimatorStateChangedNotify_proto_depIdxs = []int32{
	1, // 0: EvtAnimatorStateChangedNotify.forward_type:type_name -> ForwardType
	2, // 1: EvtAnimatorStateChangedNotify.evt_animator_state_changed_info:type_name -> EvtAnimatorStateChangedInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_EvtAnimatorStateChangedNotify_proto_init() }
func file_EvtAnimatorStateChangedNotify_proto_init() {
	if File_EvtAnimatorStateChangedNotify_proto != nil {
		return
	}
	file_EvtAnimatorStateChangedInfo_proto_init()
	file_ForwardType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EvtAnimatorStateChangedNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvtAnimatorStateChangedNotify); i {
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
			RawDescriptor: file_EvtAnimatorStateChangedNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EvtAnimatorStateChangedNotify_proto_goTypes,
		DependencyIndexes: file_EvtAnimatorStateChangedNotify_proto_depIdxs,
		MessageInfos:      file_EvtAnimatorStateChangedNotify_proto_msgTypes,
	}.Build()
	File_EvtAnimatorStateChangedNotify_proto = out.File
	file_EvtAnimatorStateChangedNotify_proto_rawDesc = nil
	file_EvtAnimatorStateChangedNotify_proto_goTypes = nil
	file_EvtAnimatorStateChangedNotify_proto_depIdxs = nil
}
