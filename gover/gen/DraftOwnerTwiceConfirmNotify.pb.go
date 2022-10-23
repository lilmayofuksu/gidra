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
// source: DraftOwnerTwiceConfirmNotify.proto

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

// CmdId: 5499
// EnetChannelId: 0
// EnetIsReliable: true
type DraftOwnerTwiceConfirmNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TwiceConfirmDeadlineTime uint32 `protobuf:"varint,15,opt,name=twice_confirm_deadline_time,json=twiceConfirmDeadlineTime,proto3" json:"twice_confirm_deadline_time,omitempty"`
	DraftId                  uint32 `protobuf:"varint,14,opt,name=draft_id,json=draftId,proto3" json:"draft_id,omitempty"`
}

func (x *DraftOwnerTwiceConfirmNotify) Reset() {
	*x = DraftOwnerTwiceConfirmNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DraftOwnerTwiceConfirmNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DraftOwnerTwiceConfirmNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DraftOwnerTwiceConfirmNotify) ProtoMessage() {}

func (x *DraftOwnerTwiceConfirmNotify) ProtoReflect() protoreflect.Message {
	mi := &file_DraftOwnerTwiceConfirmNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DraftOwnerTwiceConfirmNotify.ProtoReflect.Descriptor instead.
func (*DraftOwnerTwiceConfirmNotify) Descriptor() ([]byte, []int) {
	return file_DraftOwnerTwiceConfirmNotify_proto_rawDescGZIP(), []int{0}
}

func (x *DraftOwnerTwiceConfirmNotify) GetTwiceConfirmDeadlineTime() uint32 {
	if x != nil {
		return x.TwiceConfirmDeadlineTime
	}
	return 0
}

func (x *DraftOwnerTwiceConfirmNotify) GetDraftId() uint32 {
	if x != nil {
		return x.DraftId
	}
	return 0
}

var File_DraftOwnerTwiceConfirmNotify_proto protoreflect.FileDescriptor

var file_DraftOwnerTwiceConfirmNotify_proto_rawDesc = []byte{
	0x0a, 0x22, 0x44, 0x72, 0x61, 0x66, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x54, 0x77, 0x69, 0x63,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x78, 0x0a, 0x1c, 0x44, 0x72, 0x61, 0x66, 0x74, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x54, 0x77, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x12, 0x3d, 0x0a, 0x1b, 0x74, 0x77, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x5f, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x18, 0x74, 0x77, 0x69, 0x63, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x64, 0x72, 0x61, 0x66, 0x74, 0x49, 0x64, 0x42, 0x06,
	0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DraftOwnerTwiceConfirmNotify_proto_rawDescOnce sync.Once
	file_DraftOwnerTwiceConfirmNotify_proto_rawDescData = file_DraftOwnerTwiceConfirmNotify_proto_rawDesc
)

func file_DraftOwnerTwiceConfirmNotify_proto_rawDescGZIP() []byte {
	file_DraftOwnerTwiceConfirmNotify_proto_rawDescOnce.Do(func() {
		file_DraftOwnerTwiceConfirmNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_DraftOwnerTwiceConfirmNotify_proto_rawDescData)
	})
	return file_DraftOwnerTwiceConfirmNotify_proto_rawDescData
}

var file_DraftOwnerTwiceConfirmNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DraftOwnerTwiceConfirmNotify_proto_goTypes = []interface{}{
	(*DraftOwnerTwiceConfirmNotify)(nil), // 0: DraftOwnerTwiceConfirmNotify
}
var file_DraftOwnerTwiceConfirmNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_DraftOwnerTwiceConfirmNotify_proto_init() }
func file_DraftOwnerTwiceConfirmNotify_proto_init() {
	if File_DraftOwnerTwiceConfirmNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_DraftOwnerTwiceConfirmNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DraftOwnerTwiceConfirmNotify); i {
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
			RawDescriptor: file_DraftOwnerTwiceConfirmNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DraftOwnerTwiceConfirmNotify_proto_goTypes,
		DependencyIndexes: file_DraftOwnerTwiceConfirmNotify_proto_depIdxs,
		MessageInfos:      file_DraftOwnerTwiceConfirmNotify_proto_msgTypes,
	}.Build()
	File_DraftOwnerTwiceConfirmNotify_proto = out.File
	file_DraftOwnerTwiceConfirmNotify_proto_rawDesc = nil
	file_DraftOwnerTwiceConfirmNotify_proto_goTypes = nil
	file_DraftOwnerTwiceConfirmNotify_proto_depIdxs = nil
}
