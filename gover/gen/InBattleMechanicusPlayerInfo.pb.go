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
// source: InBattleMechanicusPlayerInfo.proto

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

type InBattleMechanicusPlayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PickCardId      uint32                            `protobuf:"varint,5,opt,name=pick_card_id,json=pickCardId,proto3" json:"pick_card_id,omitempty"`
	Uid             uint32                            `protobuf:"varint,14,opt,name=uid,proto3" json:"uid,omitempty"`
	BuildingList    []*InBattleMechanicusBuildingInfo `protobuf:"bytes,4,rep,name=building_list,json=buildingList,proto3" json:"building_list,omitempty"`
	IsCardConfirmed bool                              `protobuf:"varint,13,opt,name=is_card_confirmed,json=isCardConfirmed,proto3" json:"is_card_confirmed,omitempty"`
	BuildingPoints  uint32                            `protobuf:"varint,3,opt,name=building_points,json=buildingPoints,proto3" json:"building_points,omitempty"`
}

func (x *InBattleMechanicusPlayerInfo) Reset() {
	*x = InBattleMechanicusPlayerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_InBattleMechanicusPlayerInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InBattleMechanicusPlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InBattleMechanicusPlayerInfo) ProtoMessage() {}

func (x *InBattleMechanicusPlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_InBattleMechanicusPlayerInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InBattleMechanicusPlayerInfo.ProtoReflect.Descriptor instead.
func (*InBattleMechanicusPlayerInfo) Descriptor() ([]byte, []int) {
	return file_InBattleMechanicusPlayerInfo_proto_rawDescGZIP(), []int{0}
}

func (x *InBattleMechanicusPlayerInfo) GetPickCardId() uint32 {
	if x != nil {
		return x.PickCardId
	}
	return 0
}

func (x *InBattleMechanicusPlayerInfo) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *InBattleMechanicusPlayerInfo) GetBuildingList() []*InBattleMechanicusBuildingInfo {
	if x != nil {
		return x.BuildingList
	}
	return nil
}

func (x *InBattleMechanicusPlayerInfo) GetIsCardConfirmed() bool {
	if x != nil {
		return x.IsCardConfirmed
	}
	return false
}

func (x *InBattleMechanicusPlayerInfo) GetBuildingPoints() uint32 {
	if x != nil {
		return x.BuildingPoints
	}
	return 0
}

var File_InBattleMechanicusPlayerInfo_proto protoreflect.FileDescriptor

var file_InBattleMechanicusPlayerInfo_proto_rawDesc = []byte{
	0x0a, 0x22, 0x49, 0x6e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e,
	0x69, 0x63, 0x75, 0x73, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x49, 0x6e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x4d, 0x65,
	0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xed, 0x01, 0x0a, 0x1c, 0x49,
	0x6e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75,
	0x73, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0c, 0x70,
	0x69, 0x63, 0x6b, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x70, 0x69, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12,
	0x44, 0x0a, 0x0d, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x49, 0x6e, 0x42, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e,
	0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x73, 0x5f, 0x63, 0x61, 0x72, 0x64,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0f, 0x69, 0x73, 0x43, 0x61, 0x72, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65,
	0x64, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67,
	0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_InBattleMechanicusPlayerInfo_proto_rawDescOnce sync.Once
	file_InBattleMechanicusPlayerInfo_proto_rawDescData = file_InBattleMechanicusPlayerInfo_proto_rawDesc
)

func file_InBattleMechanicusPlayerInfo_proto_rawDescGZIP() []byte {
	file_InBattleMechanicusPlayerInfo_proto_rawDescOnce.Do(func() {
		file_InBattleMechanicusPlayerInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_InBattleMechanicusPlayerInfo_proto_rawDescData)
	})
	return file_InBattleMechanicusPlayerInfo_proto_rawDescData
}

var file_InBattleMechanicusPlayerInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_InBattleMechanicusPlayerInfo_proto_goTypes = []interface{}{
	(*InBattleMechanicusPlayerInfo)(nil),   // 0: InBattleMechanicusPlayerInfo
	(*InBattleMechanicusBuildingInfo)(nil), // 1: InBattleMechanicusBuildingInfo
}
var file_InBattleMechanicusPlayerInfo_proto_depIdxs = []int32{
	1, // 0: InBattleMechanicusPlayerInfo.building_list:type_name -> InBattleMechanicusBuildingInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_InBattleMechanicusPlayerInfo_proto_init() }
func file_InBattleMechanicusPlayerInfo_proto_init() {
	if File_InBattleMechanicusPlayerInfo_proto != nil {
		return
	}
	file_InBattleMechanicusBuildingInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_InBattleMechanicusPlayerInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InBattleMechanicusPlayerInfo); i {
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
			RawDescriptor: file_InBattleMechanicusPlayerInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_InBattleMechanicusPlayerInfo_proto_goTypes,
		DependencyIndexes: file_InBattleMechanicusPlayerInfo_proto_depIdxs,
		MessageInfos:      file_InBattleMechanicusPlayerInfo_proto_msgTypes,
	}.Build()
	File_InBattleMechanicusPlayerInfo_proto = out.File
	file_InBattleMechanicusPlayerInfo_proto_rawDesc = nil
	file_InBattleMechanicusPlayerInfo_proto_goTypes = nil
	file_InBattleMechanicusPlayerInfo_proto_depIdxs = nil
}
