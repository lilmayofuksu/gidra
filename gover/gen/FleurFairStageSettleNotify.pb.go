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
// source: FleurFairStageSettleNotify.proto

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

// CmdId: 5356
// EnetChannelId: 0
// EnetIsReliable: true
type FleurFairStageSettleNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StageType uint32 `protobuf:"varint,10,opt,name=stage_type,json=stageType,proto3" json:"stage_type,omitempty"`
	// Types that are assignable to Detail:
	//
	//	*FleurFairStageSettleNotify_GallerySettleInfo
	//	*FleurFairStageSettleNotify_BossSettleInfo
	Detail isFleurFairStageSettleNotify_Detail `protobuf_oneof:"detail"`
}

func (x *FleurFairStageSettleNotify) Reset() {
	*x = FleurFairStageSettleNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FleurFairStageSettleNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FleurFairStageSettleNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FleurFairStageSettleNotify) ProtoMessage() {}

func (x *FleurFairStageSettleNotify) ProtoReflect() protoreflect.Message {
	mi := &file_FleurFairStageSettleNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FleurFairStageSettleNotify.ProtoReflect.Descriptor instead.
func (*FleurFairStageSettleNotify) Descriptor() ([]byte, []int) {
	return file_FleurFairStageSettleNotify_proto_rawDescGZIP(), []int{0}
}

func (x *FleurFairStageSettleNotify) GetStageType() uint32 {
	if x != nil {
		return x.StageType
	}
	return 0
}

func (m *FleurFairStageSettleNotify) GetDetail() isFleurFairStageSettleNotify_Detail {
	if m != nil {
		return m.Detail
	}
	return nil
}

func (x *FleurFairStageSettleNotify) GetGallerySettleInfo() *FleurFairGallerySettleInfo {
	if x, ok := x.GetDetail().(*FleurFairStageSettleNotify_GallerySettleInfo); ok {
		return x.GallerySettleInfo
	}
	return nil
}

func (x *FleurFairStageSettleNotify) GetBossSettleInfo() *FleurFairBossSettleInfo {
	if x, ok := x.GetDetail().(*FleurFairStageSettleNotify_BossSettleInfo); ok {
		return x.BossSettleInfo
	}
	return nil
}

type isFleurFairStageSettleNotify_Detail interface {
	isFleurFairStageSettleNotify_Detail()
}

type FleurFairStageSettleNotify_GallerySettleInfo struct {
	GallerySettleInfo *FleurFairGallerySettleInfo `protobuf:"bytes,13,opt,name=gallery_settle_info,json=gallerySettleInfo,proto3,oneof"`
}

type FleurFairStageSettleNotify_BossSettleInfo struct {
	BossSettleInfo *FleurFairBossSettleInfo `protobuf:"bytes,14,opt,name=boss_settle_info,json=bossSettleInfo,proto3,oneof"`
}

func (*FleurFairStageSettleNotify_GallerySettleInfo) isFleurFairStageSettleNotify_Detail() {}

func (*FleurFairStageSettleNotify_BossSettleInfo) isFleurFairStageSettleNotify_Detail() {}

var File_FleurFairStageSettleNotify_proto protoreflect.FileDescriptor

var file_FleurFairStageSettleNotify_proto_rawDesc = []byte{
	0x0a, 0x20, 0x46, 0x6c, 0x65, 0x75, 0x72, 0x46, 0x61, 0x69, 0x72, 0x53, 0x74, 0x61, 0x67, 0x65,
	0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x46, 0x6c, 0x65, 0x75, 0x72, 0x46, 0x61, 0x69, 0x72, 0x42, 0x6f, 0x73,
	0x73, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x46, 0x6c, 0x65, 0x75, 0x72, 0x46, 0x61, 0x69, 0x72, 0x47, 0x61, 0x6c, 0x6c,
	0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xda, 0x01, 0x0a, 0x1a, 0x46, 0x6c, 0x65, 0x75, 0x72, 0x46, 0x61, 0x69,
	0x72, 0x53, 0x74, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x74, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x4d, 0x0a, 0x13, 0x67, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x74,
	0x74, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x46, 0x6c, 0x65, 0x75, 0x72, 0x46, 0x61, 0x69, 0x72, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72,
	0x79, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x11, 0x67,
	0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x44, 0x0a, 0x10, 0x62, 0x6f, 0x73, 0x73, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x46, 0x6c, 0x65,
	0x75, 0x72, 0x46, 0x61, 0x69, 0x72, 0x42, 0x6f, 0x73, 0x73, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x0e, 0x62, 0x6f, 0x73, 0x73, 0x53, 0x65, 0x74, 0x74,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x08, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FleurFairStageSettleNotify_proto_rawDescOnce sync.Once
	file_FleurFairStageSettleNotify_proto_rawDescData = file_FleurFairStageSettleNotify_proto_rawDesc
)

func file_FleurFairStageSettleNotify_proto_rawDescGZIP() []byte {
	file_FleurFairStageSettleNotify_proto_rawDescOnce.Do(func() {
		file_FleurFairStageSettleNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_FleurFairStageSettleNotify_proto_rawDescData)
	})
	return file_FleurFairStageSettleNotify_proto_rawDescData
}

var file_FleurFairStageSettleNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FleurFairStageSettleNotify_proto_goTypes = []interface{}{
	(*FleurFairStageSettleNotify)(nil), // 0: FleurFairStageSettleNotify
	(*FleurFairGallerySettleInfo)(nil), // 1: FleurFairGallerySettleInfo
	(*FleurFairBossSettleInfo)(nil),    // 2: FleurFairBossSettleInfo
}
var file_FleurFairStageSettleNotify_proto_depIdxs = []int32{
	1, // 0: FleurFairStageSettleNotify.gallery_settle_info:type_name -> FleurFairGallerySettleInfo
	2, // 1: FleurFairStageSettleNotify.boss_settle_info:type_name -> FleurFairBossSettleInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_FleurFairStageSettleNotify_proto_init() }
func file_FleurFairStageSettleNotify_proto_init() {
	if File_FleurFairStageSettleNotify_proto != nil {
		return
	}
	file_FleurFairBossSettleInfo_proto_init()
	file_FleurFairGallerySettleInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FleurFairStageSettleNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FleurFairStageSettleNotify); i {
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
	file_FleurFairStageSettleNotify_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*FleurFairStageSettleNotify_GallerySettleInfo)(nil),
		(*FleurFairStageSettleNotify_BossSettleInfo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_FleurFairStageSettleNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FleurFairStageSettleNotify_proto_goTypes,
		DependencyIndexes: file_FleurFairStageSettleNotify_proto_depIdxs,
		MessageInfos:      file_FleurFairStageSettleNotify_proto_msgTypes,
	}.Build()
	File_FleurFairStageSettleNotify_proto = out.File
	file_FleurFairStageSettleNotify_proto_rawDesc = nil
	file_FleurFairStageSettleNotify_proto_goTypes = nil
	file_FleurFairStageSettleNotify_proto_depIdxs = nil
}
