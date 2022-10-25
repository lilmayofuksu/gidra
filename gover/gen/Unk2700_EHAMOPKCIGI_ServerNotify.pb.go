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
// source: Unk2700_EHAMOPKCIGI_ServerNotify.proto

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

// CmdId: 4805
// EnetChannelId: 0
// EnetIsReliable: true
type Unk2700_EHAMOPKCIGI_ServerNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GalleryId           uint32               `protobuf:"varint,11,opt,name=gallery_id,json=galleryId,proto3" json:"gallery_id,omitempty"`
	SettleInfo          *Unk2700_KNGDOIDOFFB `protobuf:"bytes,12,opt,name=settle_info,json=settleInfo,proto3" json:"settle_info,omitempty"`
	Unk2700_HAOPLFPOLFM uint32               `protobuf:"varint,7,opt,name=Unk2700_HAOPLFPOLFM,json=Unk2700HAOPLFPOLFM,proto3" json:"Unk2700_HAOPLFPOLFM,omitempty"`
	IsNewRecord         bool                 `protobuf:"varint,2,opt,name=is_new_record,json=isNewRecord,proto3" json:"is_new_record,omitempty"`
}

func (x *Unk2700_EHAMOPKCIGI_ServerNotify) Reset() {
	*x = Unk2700_EHAMOPKCIGI_ServerNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Unk2700_EHAMOPKCIGI_ServerNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unk2700_EHAMOPKCIGI_ServerNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unk2700_EHAMOPKCIGI_ServerNotify) ProtoMessage() {}

func (x *Unk2700_EHAMOPKCIGI_ServerNotify) ProtoReflect() protoreflect.Message {
	mi := &file_Unk2700_EHAMOPKCIGI_ServerNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unk2700_EHAMOPKCIGI_ServerNotify.ProtoReflect.Descriptor instead.
func (*Unk2700_EHAMOPKCIGI_ServerNotify) Descriptor() ([]byte, []int) {
	return file_Unk2700_EHAMOPKCIGI_ServerNotify_proto_rawDescGZIP(), []int{0}
}

func (x *Unk2700_EHAMOPKCIGI_ServerNotify) GetGalleryId() uint32 {
	if x != nil {
		return x.GalleryId
	}
	return 0
}

func (x *Unk2700_EHAMOPKCIGI_ServerNotify) GetSettleInfo() *Unk2700_KNGDOIDOFFB {
	if x != nil {
		return x.SettleInfo
	}
	return nil
}

func (x *Unk2700_EHAMOPKCIGI_ServerNotify) GetUnk2700_HAOPLFPOLFM() uint32 {
	if x != nil {
		return x.Unk2700_HAOPLFPOLFM
	}
	return 0
}

func (x *Unk2700_EHAMOPKCIGI_ServerNotify) GetIsNewRecord() bool {
	if x != nil {
		return x.IsNewRecord
	}
	return false
}

var File_Unk2700_EHAMOPKCIGI_ServerNotify_proto protoreflect.FileDescriptor

var file_Unk2700_EHAMOPKCIGI_ServerNotify_proto_rawDesc = []byte{
	0x0a, 0x26, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x45, 0x48, 0x41, 0x4d, 0x4f, 0x50,
	0x4b, 0x43, 0x49, 0x47, 0x49, 0x5f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30,
	0x30, 0x5f, 0x4b, 0x4e, 0x47, 0x44, 0x4f, 0x49, 0x44, 0x4f, 0x46, 0x46, 0x42, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x01, 0x0a, 0x20, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f,
	0x45, 0x48, 0x41, 0x4d, 0x4f, 0x50, 0x4b, 0x43, 0x49, 0x47, 0x49, 0x5f, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x61, 0x6c, 0x6c,
	0x65, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x67, 0x61,
	0x6c, 0x6c, 0x65, 0x72, 0x79, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x0b, 0x73, 0x65, 0x74, 0x74, 0x6c,
	0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x55,
	0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4b, 0x4e, 0x47, 0x44, 0x4f, 0x49, 0x44, 0x4f, 0x46,
	0x46, 0x42, 0x52, 0x0a, 0x73, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2f,
	0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x48, 0x41, 0x4f, 0x50, 0x4c, 0x46,
	0x50, 0x4f, 0x4c, 0x46, 0x4d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b,
	0x32, 0x37, 0x30, 0x30, 0x48, 0x41, 0x4f, 0x50, 0x4c, 0x46, 0x50, 0x4f, 0x4c, 0x46, 0x4d, 0x12,
	0x22, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_Unk2700_EHAMOPKCIGI_ServerNotify_proto_rawDescOnce sync.Once
	file_Unk2700_EHAMOPKCIGI_ServerNotify_proto_rawDescData = file_Unk2700_EHAMOPKCIGI_ServerNotify_proto_rawDesc
)

func file_Unk2700_EHAMOPKCIGI_ServerNotify_proto_rawDescGZIP() []byte {
	file_Unk2700_EHAMOPKCIGI_ServerNotify_proto_rawDescOnce.Do(func() {
		file_Unk2700_EHAMOPKCIGI_ServerNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_Unk2700_EHAMOPKCIGI_ServerNotify_proto_rawDescData)
	})
	return file_Unk2700_EHAMOPKCIGI_ServerNotify_proto_rawDescData
}

var file_Unk2700_EHAMOPKCIGI_ServerNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Unk2700_EHAMOPKCIGI_ServerNotify_proto_goTypes = []interface{}{
	(*Unk2700_EHAMOPKCIGI_ServerNotify)(nil), // 0: Unk2700_EHAMOPKCIGI_ServerNotify
	(*Unk2700_KNGDOIDOFFB)(nil),              // 1: Unk2700_KNGDOIDOFFB
}
var file_Unk2700_EHAMOPKCIGI_ServerNotify_proto_depIdxs = []int32{
	1, // 0: Unk2700_EHAMOPKCIGI_ServerNotify.settle_info:type_name -> Unk2700_KNGDOIDOFFB
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Unk2700_EHAMOPKCIGI_ServerNotify_proto_init() }
func file_Unk2700_EHAMOPKCIGI_ServerNotify_proto_init() {
	if File_Unk2700_EHAMOPKCIGI_ServerNotify_proto != nil {
		return
	}
	file_Unk2700_KNGDOIDOFFB_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Unk2700_EHAMOPKCIGI_ServerNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unk2700_EHAMOPKCIGI_ServerNotify); i {
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
			RawDescriptor: file_Unk2700_EHAMOPKCIGI_ServerNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Unk2700_EHAMOPKCIGI_ServerNotify_proto_goTypes,
		DependencyIndexes: file_Unk2700_EHAMOPKCIGI_ServerNotify_proto_depIdxs,
		MessageInfos:      file_Unk2700_EHAMOPKCIGI_ServerNotify_proto_msgTypes,
	}.Build()
	File_Unk2700_EHAMOPKCIGI_ServerNotify_proto = out.File
	file_Unk2700_EHAMOPKCIGI_ServerNotify_proto_rawDesc = nil
	file_Unk2700_EHAMOPKCIGI_ServerNotify_proto_goTypes = nil
	file_Unk2700_EHAMOPKCIGI_ServerNotify_proto_depIdxs = nil
}
