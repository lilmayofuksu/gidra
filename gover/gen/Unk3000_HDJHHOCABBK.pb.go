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
// source: Unk3000_HDJHHOCABBK.proto

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

type Unk3000_HDJHHOCABBK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsDone              bool    `protobuf:"varint,12,opt,name=is_done,json=isDone,proto3" json:"is_done,omitempty"`
	Unk3000_LIHPABKOAIP uint32  `protobuf:"varint,6,opt,name=Unk3000_LIHPABKOAIP,json=Unk3000LIHPABKOAIP,proto3" json:"Unk3000_LIHPABKOAIP,omitempty"`
	Unk3000_AEGHMLLEOJF uint32  `protobuf:"varint,10,opt,name=Unk3000_AEGHMLLEOJF,json=Unk3000AEGHMLLEOJF,proto3" json:"Unk3000_AEGHMLLEOJF,omitempty"`
	RegionRadius        float32 `protobuf:"fixed32,7,opt,name=region_radius,json=regionRadius,proto3" json:"region_radius,omitempty"`
	IsOpen              bool    `protobuf:"varint,9,opt,name=is_open,json=isOpen,proto3" json:"is_open,omitempty"`
	OpenTime            uint32  `protobuf:"varint,8,opt,name=open_time,json=openTime,proto3" json:"open_time,omitempty"`
	RegionCenterPos     *Vector `protobuf:"bytes,11,opt,name=region_center_pos,json=regionCenterPos,proto3" json:"region_center_pos,omitempty"`
	SceneId             uint32  `protobuf:"varint,13,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	Unk3000_KNNPMAMOCOM uint32  `protobuf:"varint,15,opt,name=Unk3000_KNNPMAMOCOM,json=Unk3000KNNPMAMOCOM,proto3" json:"Unk3000_KNNPMAMOCOM,omitempty"`
	RegionId            uint32  `protobuf:"varint,1,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
}

func (x *Unk3000_HDJHHOCABBK) Reset() {
	*x = Unk3000_HDJHHOCABBK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Unk3000_HDJHHOCABBK_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unk3000_HDJHHOCABBK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unk3000_HDJHHOCABBK) ProtoMessage() {}

func (x *Unk3000_HDJHHOCABBK) ProtoReflect() protoreflect.Message {
	mi := &file_Unk3000_HDJHHOCABBK_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unk3000_HDJHHOCABBK.ProtoReflect.Descriptor instead.
func (*Unk3000_HDJHHOCABBK) Descriptor() ([]byte, []int) {
	return file_Unk3000_HDJHHOCABBK_proto_rawDescGZIP(), []int{0}
}

func (x *Unk3000_HDJHHOCABBK) GetIsDone() bool {
	if x != nil {
		return x.IsDone
	}
	return false
}

func (x *Unk3000_HDJHHOCABBK) GetUnk3000_LIHPABKOAIP() uint32 {
	if x != nil {
		return x.Unk3000_LIHPABKOAIP
	}
	return 0
}

func (x *Unk3000_HDJHHOCABBK) GetUnk3000_AEGHMLLEOJF() uint32 {
	if x != nil {
		return x.Unk3000_AEGHMLLEOJF
	}
	return 0
}

func (x *Unk3000_HDJHHOCABBK) GetRegionRadius() float32 {
	if x != nil {
		return x.RegionRadius
	}
	return 0
}

func (x *Unk3000_HDJHHOCABBK) GetIsOpen() bool {
	if x != nil {
		return x.IsOpen
	}
	return false
}

func (x *Unk3000_HDJHHOCABBK) GetOpenTime() uint32 {
	if x != nil {
		return x.OpenTime
	}
	return 0
}

func (x *Unk3000_HDJHHOCABBK) GetRegionCenterPos() *Vector {
	if x != nil {
		return x.RegionCenterPos
	}
	return nil
}

func (x *Unk3000_HDJHHOCABBK) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

func (x *Unk3000_HDJHHOCABBK) GetUnk3000_KNNPMAMOCOM() uint32 {
	if x != nil {
		return x.Unk3000_KNNPMAMOCOM
	}
	return 0
}

func (x *Unk3000_HDJHHOCABBK) GetRegionId() uint32 {
	if x != nil {
		return x.RegionId
	}
	return 0
}

var File_Unk3000_HDJHHOCABBK_proto protoreflect.FileDescriptor

var file_Unk3000_HDJHHOCABBK_proto_rawDesc = []byte{
	0x0a, 0x19, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x48, 0x44, 0x4a, 0x48, 0x48, 0x4f,
	0x43, 0x41, 0x42, 0x42, 0x4b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x03, 0x0a, 0x13, 0x55, 0x6e,
	0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x48, 0x44, 0x4a, 0x48, 0x48, 0x4f, 0x43, 0x41, 0x42, 0x42,
	0x4b, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x44, 0x6f, 0x6e, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e,
	0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x4c, 0x49, 0x48, 0x50, 0x41, 0x42, 0x4b, 0x4f, 0x41, 0x49,
	0x50, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30,
	0x4c, 0x49, 0x48, 0x50, 0x41, 0x42, 0x4b, 0x4f, 0x41, 0x49, 0x50, 0x12, 0x2f, 0x0a, 0x13, 0x55,
	0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x41, 0x45, 0x47, 0x48, 0x4d, 0x4c, 0x4c, 0x45, 0x4f,
	0x4a, 0x46, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30,
	0x30, 0x41, 0x45, 0x47, 0x48, 0x4d, 0x4c, 0x4c, 0x45, 0x4f, 0x4a, 0x46, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x64, 0x69, 0x75,
	0x73, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x70,
	0x65, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6f,
	0x70, 0x65, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x11, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0f, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x12, 0x19, 0x0a, 0x08,
	0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x73, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x30,
	0x30, 0x30, 0x5f, 0x4b, 0x4e, 0x4e, 0x50, 0x4d, 0x41, 0x4d, 0x4f, 0x43, 0x4f, 0x4d, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x4b, 0x4e, 0x4e,
	0x50, 0x4d, 0x41, 0x4d, 0x4f, 0x43, 0x4f, 0x4d, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Unk3000_HDJHHOCABBK_proto_rawDescOnce sync.Once
	file_Unk3000_HDJHHOCABBK_proto_rawDescData = file_Unk3000_HDJHHOCABBK_proto_rawDesc
)

func file_Unk3000_HDJHHOCABBK_proto_rawDescGZIP() []byte {
	file_Unk3000_HDJHHOCABBK_proto_rawDescOnce.Do(func() {
		file_Unk3000_HDJHHOCABBK_proto_rawDescData = protoimpl.X.CompressGZIP(file_Unk3000_HDJHHOCABBK_proto_rawDescData)
	})
	return file_Unk3000_HDJHHOCABBK_proto_rawDescData
}

var file_Unk3000_HDJHHOCABBK_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Unk3000_HDJHHOCABBK_proto_goTypes = []interface{}{
	(*Unk3000_HDJHHOCABBK)(nil), // 0: Unk3000_HDJHHOCABBK
	(*Vector)(nil),              // 1: Vector
}
var file_Unk3000_HDJHHOCABBK_proto_depIdxs = []int32{
	1, // 0: Unk3000_HDJHHOCABBK.region_center_pos:type_name -> Vector
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Unk3000_HDJHHOCABBK_proto_init() }
func file_Unk3000_HDJHHOCABBK_proto_init() {
	if File_Unk3000_HDJHHOCABBK_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Unk3000_HDJHHOCABBK_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unk3000_HDJHHOCABBK); i {
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
			RawDescriptor: file_Unk3000_HDJHHOCABBK_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Unk3000_HDJHHOCABBK_proto_goTypes,
		DependencyIndexes: file_Unk3000_HDJHHOCABBK_proto_depIdxs,
		MessageInfos:      file_Unk3000_HDJHHOCABBK_proto_msgTypes,
	}.Build()
	File_Unk3000_HDJHHOCABBK_proto = out.File
	file_Unk3000_HDJHHOCABBK_proto_rawDesc = nil
	file_Unk3000_HDJHHOCABBK_proto_goTypes = nil
	file_Unk3000_HDJHHOCABBK_proto_depIdxs = nil
}
