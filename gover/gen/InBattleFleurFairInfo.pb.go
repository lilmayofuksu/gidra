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
// source: InBattleFleurFairInfo.proto

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

type InBattleFleurFairInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GalleryIdList          []uint32 `protobuf:"varint,5,rep,packed,name=gallery_id_list,json=galleryIdList,proto3" json:"gallery_id_list,omitempty"`
	GalleryStageIndex      uint32   `protobuf:"varint,6,opt,name=gallery_stage_index,json=galleryStageIndex,proto3" json:"gallery_stage_index,omitempty"`
	PreviewStageIndex      uint32   `protobuf:"varint,8,opt,name=preview_stage_index,json=previewStageIndex,proto3" json:"preview_stage_index,omitempty"`
	AbilityGroupIdList     []uint32 `protobuf:"varint,2,rep,packed,name=ability_group_id_list,json=abilityGroupIdList,proto3" json:"ability_group_id_list,omitempty"`
	PreviewDisplayDuration uint32   `protobuf:"varint,12,opt,name=preview_display_duration,json=previewDisplayDuration,proto3" json:"preview_display_duration,omitempty"`
}

func (x *InBattleFleurFairInfo) Reset() {
	*x = InBattleFleurFairInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_InBattleFleurFairInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InBattleFleurFairInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InBattleFleurFairInfo) ProtoMessage() {}

func (x *InBattleFleurFairInfo) ProtoReflect() protoreflect.Message {
	mi := &file_InBattleFleurFairInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InBattleFleurFairInfo.ProtoReflect.Descriptor instead.
func (*InBattleFleurFairInfo) Descriptor() ([]byte, []int) {
	return file_InBattleFleurFairInfo_proto_rawDescGZIP(), []int{0}
}

func (x *InBattleFleurFairInfo) GetGalleryIdList() []uint32 {
	if x != nil {
		return x.GalleryIdList
	}
	return nil
}

func (x *InBattleFleurFairInfo) GetGalleryStageIndex() uint32 {
	if x != nil {
		return x.GalleryStageIndex
	}
	return 0
}

func (x *InBattleFleurFairInfo) GetPreviewStageIndex() uint32 {
	if x != nil {
		return x.PreviewStageIndex
	}
	return 0
}

func (x *InBattleFleurFairInfo) GetAbilityGroupIdList() []uint32 {
	if x != nil {
		return x.AbilityGroupIdList
	}
	return nil
}

func (x *InBattleFleurFairInfo) GetPreviewDisplayDuration() uint32 {
	if x != nil {
		return x.PreviewDisplayDuration
	}
	return 0
}

var File_InBattleFleurFairInfo_proto protoreflect.FileDescriptor

var file_InBattleFleurFairInfo_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x49, 0x6e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x46, 0x6c, 0x65, 0x75, 0x72, 0x46,
	0x61, 0x69, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x02,
	0x0a, 0x15, 0x49, 0x6e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x46, 0x6c, 0x65, 0x75, 0x72, 0x46,
	0x61, 0x69, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26, 0x0a, 0x0f, 0x67, 0x61, 0x6c, 0x6c, 0x65,
	0x72, 0x79, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x0d, 0x67, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x2e, 0x0a, 0x13, 0x67, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x67, 0x61,
	0x6c, 0x6c, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x2e, 0x0a, 0x13, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x70, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x31, 0x0a, 0x15, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x12,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x38, 0x0a, 0x18, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x44, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x06, 0x5a, 0x04,
	0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_InBattleFleurFairInfo_proto_rawDescOnce sync.Once
	file_InBattleFleurFairInfo_proto_rawDescData = file_InBattleFleurFairInfo_proto_rawDesc
)

func file_InBattleFleurFairInfo_proto_rawDescGZIP() []byte {
	file_InBattleFleurFairInfo_proto_rawDescOnce.Do(func() {
		file_InBattleFleurFairInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_InBattleFleurFairInfo_proto_rawDescData)
	})
	return file_InBattleFleurFairInfo_proto_rawDescData
}

var file_InBattleFleurFairInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_InBattleFleurFairInfo_proto_goTypes = []interface{}{
	(*InBattleFleurFairInfo)(nil), // 0: InBattleFleurFairInfo
}
var file_InBattleFleurFairInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_InBattleFleurFairInfo_proto_init() }
func file_InBattleFleurFairInfo_proto_init() {
	if File_InBattleFleurFairInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_InBattleFleurFairInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InBattleFleurFairInfo); i {
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
			RawDescriptor: file_InBattleFleurFairInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_InBattleFleurFairInfo_proto_goTypes,
		DependencyIndexes: file_InBattleFleurFairInfo_proto_depIdxs,
		MessageInfos:      file_InBattleFleurFairInfo_proto_msgTypes,
	}.Build()
	File_InBattleFleurFairInfo_proto = out.File
	file_InBattleFleurFairInfo_proto_rawDesc = nil
	file_InBattleFleurFairInfo_proto_goTypes = nil
	file_InBattleFleurFairInfo_proto_depIdxs = nil
}
