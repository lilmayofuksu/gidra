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
// source: LunaRiteTakeSacrificeRewardReq.proto

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

// CmdId: 8045
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type LunaRiteTakeSacrificeRewardReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AreaId uint32 `protobuf:"varint,11,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	Index  uint32 `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *LunaRiteTakeSacrificeRewardReq) Reset() {
	*x = LunaRiteTakeSacrificeRewardReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LunaRiteTakeSacrificeRewardReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LunaRiteTakeSacrificeRewardReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LunaRiteTakeSacrificeRewardReq) ProtoMessage() {}

func (x *LunaRiteTakeSacrificeRewardReq) ProtoReflect() protoreflect.Message {
	mi := &file_LunaRiteTakeSacrificeRewardReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LunaRiteTakeSacrificeRewardReq.ProtoReflect.Descriptor instead.
func (*LunaRiteTakeSacrificeRewardReq) Descriptor() ([]byte, []int) {
	return file_LunaRiteTakeSacrificeRewardReq_proto_rawDescGZIP(), []int{0}
}

func (x *LunaRiteTakeSacrificeRewardReq) GetAreaId() uint32 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *LunaRiteTakeSacrificeRewardReq) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

var File_LunaRiteTakeSacrificeRewardReq_proto protoreflect.FileDescriptor

var file_LunaRiteTakeSacrificeRewardReq_proto_rawDesc = []byte{
	0x0a, 0x24, 0x4c, 0x75, 0x6e, 0x61, 0x52, 0x69, 0x74, 0x65, 0x54, 0x61, 0x6b, 0x65, 0x53, 0x61,
	0x63, 0x72, 0x69, 0x66, 0x69, 0x63, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x1e, 0x4c, 0x75, 0x6e, 0x61, 0x52, 0x69,
	0x74, 0x65, 0x54, 0x61, 0x6b, 0x65, 0x53, 0x61, 0x63, 0x72, 0x69, 0x66, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x72, 0x65, 0x61,
	0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x72, 0x65, 0x61, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_LunaRiteTakeSacrificeRewardReq_proto_rawDescOnce sync.Once
	file_LunaRiteTakeSacrificeRewardReq_proto_rawDescData = file_LunaRiteTakeSacrificeRewardReq_proto_rawDesc
)

func file_LunaRiteTakeSacrificeRewardReq_proto_rawDescGZIP() []byte {
	file_LunaRiteTakeSacrificeRewardReq_proto_rawDescOnce.Do(func() {
		file_LunaRiteTakeSacrificeRewardReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_LunaRiteTakeSacrificeRewardReq_proto_rawDescData)
	})
	return file_LunaRiteTakeSacrificeRewardReq_proto_rawDescData
}

var file_LunaRiteTakeSacrificeRewardReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_LunaRiteTakeSacrificeRewardReq_proto_goTypes = []interface{}{
	(*LunaRiteTakeSacrificeRewardReq)(nil), // 0: LunaRiteTakeSacrificeRewardReq
}
var file_LunaRiteTakeSacrificeRewardReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_LunaRiteTakeSacrificeRewardReq_proto_init() }
func file_LunaRiteTakeSacrificeRewardReq_proto_init() {
	if File_LunaRiteTakeSacrificeRewardReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_LunaRiteTakeSacrificeRewardReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LunaRiteTakeSacrificeRewardReq); i {
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
			RawDescriptor: file_LunaRiteTakeSacrificeRewardReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LunaRiteTakeSacrificeRewardReq_proto_goTypes,
		DependencyIndexes: file_LunaRiteTakeSacrificeRewardReq_proto_depIdxs,
		MessageInfos:      file_LunaRiteTakeSacrificeRewardReq_proto_msgTypes,
	}.Build()
	File_LunaRiteTakeSacrificeRewardReq_proto = out.File
	file_LunaRiteTakeSacrificeRewardReq_proto_rawDesc = nil
	file_LunaRiteTakeSacrificeRewardReq_proto_goTypes = nil
	file_LunaRiteTakeSacrificeRewardReq_proto_depIdxs = nil
}
