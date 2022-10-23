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
// source: WinterCampActivityDetailInfo.proto

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

type WinterCampActivityDetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unk2700_FBMHFJHDJNB []*Unk2700_MBIDJDLLBNM `protobuf:"bytes,9,rep,name=Unk2700_FBMHFJHDJNB,json=Unk2700FBMHFJHDJNB,proto3" json:"Unk2700_FBMHFJHDJNB,omitempty"`
	BattleInfo          *Unk2700_DIEGJDEIDKO   `protobuf:"bytes,10,opt,name=battle_info,json=battleInfo,proto3" json:"battle_info,omitempty"`
	Unk2700_GALHBPGEGNL []uint32               `protobuf:"varint,8,rep,packed,name=Unk2700_GALHBPGEGNL,json=Unk2700GALHBPGEGNL,proto3" json:"Unk2700_GALHBPGEGNL,omitempty"`
	Unk2700_DKCGOPBHJHA []uint32               `protobuf:"varint,14,rep,packed,name=Unk2700_DKCGOPBHJHA,json=Unk2700DKCGOPBHJHA,proto3" json:"Unk2700_DKCGOPBHJHA,omitempty"`
	Unk2700_OOBOCEALLBE []uint32               `protobuf:"varint,6,rep,packed,name=Unk2700_OOBOCEALLBE,json=Unk2700OOBOCEALLBE,proto3" json:"Unk2700_OOBOCEALLBE,omitempty"`
	IsContentClosed     bool                   `protobuf:"varint,15,opt,name=is_content_closed,json=isContentClosed,proto3" json:"is_content_closed,omitempty"`
	ExploreInfo         *Unk2700_DIEGJDEIDKO   `protobuf:"bytes,11,opt,name=explore_info,json=exploreInfo,proto3" json:"explore_info,omitempty"`
	Unk2700_CFENLEBIKGG []*ItemParam           `protobuf:"bytes,2,rep,name=Unk2700_CFENLEBIKGG,json=Unk2700CFENLEBIKGG,proto3" json:"Unk2700_CFENLEBIKGG,omitempty"`
}

func (x *WinterCampActivityDetailInfo) Reset() {
	*x = WinterCampActivityDetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WinterCampActivityDetailInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WinterCampActivityDetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WinterCampActivityDetailInfo) ProtoMessage() {}

func (x *WinterCampActivityDetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_WinterCampActivityDetailInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WinterCampActivityDetailInfo.ProtoReflect.Descriptor instead.
func (*WinterCampActivityDetailInfo) Descriptor() ([]byte, []int) {
	return file_WinterCampActivityDetailInfo_proto_rawDescGZIP(), []int{0}
}

func (x *WinterCampActivityDetailInfo) GetUnk2700_FBMHFJHDJNB() []*Unk2700_MBIDJDLLBNM {
	if x != nil {
		return x.Unk2700_FBMHFJHDJNB
	}
	return nil
}

func (x *WinterCampActivityDetailInfo) GetBattleInfo() *Unk2700_DIEGJDEIDKO {
	if x != nil {
		return x.BattleInfo
	}
	return nil
}

func (x *WinterCampActivityDetailInfo) GetUnk2700_GALHBPGEGNL() []uint32 {
	if x != nil {
		return x.Unk2700_GALHBPGEGNL
	}
	return nil
}

func (x *WinterCampActivityDetailInfo) GetUnk2700_DKCGOPBHJHA() []uint32 {
	if x != nil {
		return x.Unk2700_DKCGOPBHJHA
	}
	return nil
}

func (x *WinterCampActivityDetailInfo) GetUnk2700_OOBOCEALLBE() []uint32 {
	if x != nil {
		return x.Unk2700_OOBOCEALLBE
	}
	return nil
}

func (x *WinterCampActivityDetailInfo) GetIsContentClosed() bool {
	if x != nil {
		return x.IsContentClosed
	}
	return false
}

func (x *WinterCampActivityDetailInfo) GetExploreInfo() *Unk2700_DIEGJDEIDKO {
	if x != nil {
		return x.ExploreInfo
	}
	return nil
}

func (x *WinterCampActivityDetailInfo) GetUnk2700_CFENLEBIKGG() []*ItemParam {
	if x != nil {
		return x.Unk2700_CFENLEBIKGG
	}
	return nil
}

var File_WinterCampActivityDetailInfo_proto protoreflect.FileDescriptor

var file_WinterCampActivityDetailInfo_proto_rawDesc = []byte{
	0x0a, 0x22, 0x57, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x43, 0x61, 0x6d, 0x70, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x44,
	0x49, 0x45, 0x47, 0x4a, 0x44, 0x45, 0x49, 0x44, 0x4b, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4d, 0x42, 0x49, 0x44, 0x4a, 0x44,
	0x4c, 0x4c, 0x42, 0x4e, 0x4d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x03, 0x0a, 0x1c,
	0x57, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x43, 0x61, 0x6d, 0x70, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x45, 0x0a, 0x13,
	0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x46, 0x42, 0x4d, 0x48, 0x46, 0x4a, 0x48, 0x44,
	0x4a, 0x4e, 0x42, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x55, 0x6e, 0x6b, 0x32,
	0x37, 0x30, 0x30, 0x5f, 0x4d, 0x42, 0x49, 0x44, 0x4a, 0x44, 0x4c, 0x4c, 0x42, 0x4e, 0x4d, 0x52,
	0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x46, 0x42, 0x4d, 0x48, 0x46, 0x4a, 0x48, 0x44,
	0x4a, 0x4e, 0x42, 0x12, 0x35, 0x0a, 0x0b, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x55, 0x6e, 0x6b, 0x32, 0x37,
	0x30, 0x30, 0x5f, 0x44, 0x49, 0x45, 0x47, 0x4a, 0x44, 0x45, 0x49, 0x44, 0x4b, 0x4f, 0x52, 0x0a,
	0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e,
	0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x47, 0x41, 0x4c, 0x48, 0x42, 0x50, 0x47, 0x45, 0x47, 0x4e,
	0x4c, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30,
	0x47, 0x41, 0x4c, 0x48, 0x42, 0x50, 0x47, 0x45, 0x47, 0x4e, 0x4c, 0x12, 0x2f, 0x0a, 0x13, 0x55,
	0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x44, 0x4b, 0x43, 0x47, 0x4f, 0x50, 0x42, 0x48, 0x4a,
	0x48, 0x41, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30,
	0x30, 0x44, 0x4b, 0x43, 0x47, 0x4f, 0x50, 0x42, 0x48, 0x4a, 0x48, 0x41, 0x12, 0x2f, 0x0a, 0x13,
	0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4f, 0x4f, 0x42, 0x4f, 0x43, 0x45, 0x41, 0x4c,
	0x4c, 0x42, 0x45, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37,
	0x30, 0x30, 0x4f, 0x4f, 0x42, 0x4f, 0x43, 0x45, 0x41, 0x4c, 0x4c, 0x42, 0x45, 0x12, 0x2a, 0x0a,
	0x11, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x6f, 0x73,
	0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x12, 0x37, 0x0a, 0x0c, 0x65, 0x78, 0x70,
	0x6c, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x44, 0x49, 0x45, 0x47, 0x4a, 0x44,
	0x45, 0x49, 0x44, 0x4b, 0x4f, 0x52, 0x0b, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x3b, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x43, 0x46,
	0x45, 0x4e, 0x4c, 0x45, 0x42, 0x49, 0x4b, 0x47, 0x47, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x12, 0x55, 0x6e, 0x6b,
	0x32, 0x37, 0x30, 0x30, 0x43, 0x46, 0x45, 0x4e, 0x4c, 0x45, 0x42, 0x49, 0x4b, 0x47, 0x47, 0x42,
	0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_WinterCampActivityDetailInfo_proto_rawDescOnce sync.Once
	file_WinterCampActivityDetailInfo_proto_rawDescData = file_WinterCampActivityDetailInfo_proto_rawDesc
)

func file_WinterCampActivityDetailInfo_proto_rawDescGZIP() []byte {
	file_WinterCampActivityDetailInfo_proto_rawDescOnce.Do(func() {
		file_WinterCampActivityDetailInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_WinterCampActivityDetailInfo_proto_rawDescData)
	})
	return file_WinterCampActivityDetailInfo_proto_rawDescData
}

var file_WinterCampActivityDetailInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_WinterCampActivityDetailInfo_proto_goTypes = []interface{}{
	(*WinterCampActivityDetailInfo)(nil), // 0: WinterCampActivityDetailInfo
	(*Unk2700_MBIDJDLLBNM)(nil),          // 1: Unk2700_MBIDJDLLBNM
	(*Unk2700_DIEGJDEIDKO)(nil),          // 2: Unk2700_DIEGJDEIDKO
	(*ItemParam)(nil),                    // 3: ItemParam
}
var file_WinterCampActivityDetailInfo_proto_depIdxs = []int32{
	1, // 0: WinterCampActivityDetailInfo.Unk2700_FBMHFJHDJNB:type_name -> Unk2700_MBIDJDLLBNM
	2, // 1: WinterCampActivityDetailInfo.battle_info:type_name -> Unk2700_DIEGJDEIDKO
	2, // 2: WinterCampActivityDetailInfo.explore_info:type_name -> Unk2700_DIEGJDEIDKO
	3, // 3: WinterCampActivityDetailInfo.Unk2700_CFENLEBIKGG:type_name -> ItemParam
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_WinterCampActivityDetailInfo_proto_init() }
func file_WinterCampActivityDetailInfo_proto_init() {
	if File_WinterCampActivityDetailInfo_proto != nil {
		return
	}
	file_ItemParam_proto_init()
	file_Unk2700_DIEGJDEIDKO_proto_init()
	file_Unk2700_MBIDJDLLBNM_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_WinterCampActivityDetailInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WinterCampActivityDetailInfo); i {
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
			RawDescriptor: file_WinterCampActivityDetailInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WinterCampActivityDetailInfo_proto_goTypes,
		DependencyIndexes: file_WinterCampActivityDetailInfo_proto_depIdxs,
		MessageInfos:      file_WinterCampActivityDetailInfo_proto_msgTypes,
	}.Build()
	File_WinterCampActivityDetailInfo_proto = out.File
	file_WinterCampActivityDetailInfo_proto_rawDesc = nil
	file_WinterCampActivityDetailInfo_proto_goTypes = nil
	file_WinterCampActivityDetailInfo_proto_depIdxs = nil
}
