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
// source: Unk3000_GCBMILHPIKA.proto

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

// CmdId: 4659
// EnetChannelId: 0
// EnetIsReliable: true
type Unk3000_GCBMILHPIKA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode             int32                                      `protobuf:"varint,10,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Unk3000_EBIEGNHLMFP []*Unk3000_GCBMILHPIKA_Unk3000_PPGINNAFPIF `protobuf:"bytes,5,rep,name=Unk3000_EBIEGNHLMFP,json=Unk3000EBIEGNHLMFP,proto3" json:"Unk3000_EBIEGNHLMFP,omitempty"`
}

func (x *Unk3000_GCBMILHPIKA) Reset() {
	*x = Unk3000_GCBMILHPIKA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Unk3000_GCBMILHPIKA_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unk3000_GCBMILHPIKA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unk3000_GCBMILHPIKA) ProtoMessage() {}

func (x *Unk3000_GCBMILHPIKA) ProtoReflect() protoreflect.Message {
	mi := &file_Unk3000_GCBMILHPIKA_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unk3000_GCBMILHPIKA.ProtoReflect.Descriptor instead.
func (*Unk3000_GCBMILHPIKA) Descriptor() ([]byte, []int) {
	return file_Unk3000_GCBMILHPIKA_proto_rawDescGZIP(), []int{0}
}

func (x *Unk3000_GCBMILHPIKA) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *Unk3000_GCBMILHPIKA) GetUnk3000_EBIEGNHLMFP() []*Unk3000_GCBMILHPIKA_Unk3000_PPGINNAFPIF {
	if x != nil {
		return x.Unk3000_EBIEGNHLMFP
	}
	return nil
}

type Unk3000_GCBMILHPIKA_Unk3000_PPGINNAFPIF struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unk3000_CLMLONOEHLB uint32 `protobuf:"varint,7,opt,name=Unk3000_CLMLONOEHLB,json=Unk3000CLMLONOEHLB,proto3" json:"Unk3000_CLMLONOEHLB,omitempty"`
	Unk3000_HCAJDIBHKDG uint32 `protobuf:"varint,12,opt,name=Unk3000_HCAJDIBHKDG,json=Unk3000HCAJDIBHKDG,proto3" json:"Unk3000_HCAJDIBHKDG,omitempty"`
	NextRefreshTime     uint32 `protobuf:"varint,14,opt,name=next_refresh_time,json=nextRefreshTime,proto3" json:"next_refresh_time,omitempty"`
	Unk3000_LOFNFMJFGNB uint32 `protobuf:"varint,2,opt,name=Unk3000_LOFNFMJFGNB,json=Unk3000LOFNFMJFGNB,proto3" json:"Unk3000_LOFNFMJFGNB,omitempty"`
}

func (x *Unk3000_GCBMILHPIKA_Unk3000_PPGINNAFPIF) Reset() {
	*x = Unk3000_GCBMILHPIKA_Unk3000_PPGINNAFPIF{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Unk3000_GCBMILHPIKA_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unk3000_GCBMILHPIKA_Unk3000_PPGINNAFPIF) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unk3000_GCBMILHPIKA_Unk3000_PPGINNAFPIF) ProtoMessage() {}

func (x *Unk3000_GCBMILHPIKA_Unk3000_PPGINNAFPIF) ProtoReflect() protoreflect.Message {
	mi := &file_Unk3000_GCBMILHPIKA_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unk3000_GCBMILHPIKA_Unk3000_PPGINNAFPIF.ProtoReflect.Descriptor instead.
func (*Unk3000_GCBMILHPIKA_Unk3000_PPGINNAFPIF) Descriptor() ([]byte, []int) {
	return file_Unk3000_GCBMILHPIKA_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Unk3000_GCBMILHPIKA_Unk3000_PPGINNAFPIF) GetUnk3000_CLMLONOEHLB() uint32 {
	if x != nil {
		return x.Unk3000_CLMLONOEHLB
	}
	return 0
}

func (x *Unk3000_GCBMILHPIKA_Unk3000_PPGINNAFPIF) GetUnk3000_HCAJDIBHKDG() uint32 {
	if x != nil {
		return x.Unk3000_HCAJDIBHKDG
	}
	return 0
}

func (x *Unk3000_GCBMILHPIKA_Unk3000_PPGINNAFPIF) GetNextRefreshTime() uint32 {
	if x != nil {
		return x.NextRefreshTime
	}
	return 0
}

func (x *Unk3000_GCBMILHPIKA_Unk3000_PPGINNAFPIF) GetUnk3000_LOFNFMJFGNB() uint32 {
	if x != nil {
		return x.Unk3000_LOFNFMJFGNB
	}
	return 0
}

var File_Unk3000_GCBMILHPIKA_proto protoreflect.FileDescriptor

var file_Unk3000_GCBMILHPIKA_proto_rawDesc = []byte{
	0x0a, 0x19, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x47, 0x43, 0x42, 0x4d, 0x49, 0x4c,
	0x48, 0x50, 0x49, 0x4b, 0x41, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x02, 0x0a, 0x13,
	0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x47, 0x43, 0x42, 0x4d, 0x49, 0x4c, 0x48, 0x50,
	0x49, 0x4b, 0x41, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x59, 0x0a,
	0x13, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x45, 0x42, 0x49, 0x45, 0x47, 0x4e, 0x48,
	0x4c, 0x4d, 0x46, 0x50, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x55, 0x6e, 0x6b,
	0x33, 0x30, 0x30, 0x30, 0x5f, 0x47, 0x43, 0x42, 0x4d, 0x49, 0x4c, 0x48, 0x50, 0x49, 0x4b, 0x41,
	0x2e, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x50, 0x50, 0x47, 0x49, 0x4e, 0x4e, 0x41,
	0x46, 0x50, 0x49, 0x46, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x45, 0x42, 0x49,
	0x45, 0x47, 0x4e, 0x48, 0x4c, 0x4d, 0x46, 0x50, 0x1a, 0xd4, 0x01, 0x0a, 0x13, 0x55, 0x6e, 0x6b,
	0x33, 0x30, 0x30, 0x30, 0x5f, 0x50, 0x50, 0x47, 0x49, 0x4e, 0x4e, 0x41, 0x46, 0x50, 0x49, 0x46,
	0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x43, 0x4c, 0x4d, 0x4c,
	0x4f, 0x4e, 0x4f, 0x45, 0x48, 0x4c, 0x42, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55,
	0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x43, 0x4c, 0x4d, 0x4c, 0x4f, 0x4e, 0x4f, 0x45, 0x48, 0x4c,
	0x42, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x48, 0x43, 0x41,
	0x4a, 0x44, 0x49, 0x42, 0x48, 0x4b, 0x44, 0x47, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12,
	0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x48, 0x43, 0x41, 0x4a, 0x44, 0x49, 0x42, 0x48, 0x4b,
	0x44, 0x47, 0x12, 0x2a, 0x0a, 0x11, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6e,
	0x65, 0x78, 0x74, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2f,
	0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x4c, 0x4f, 0x46, 0x4e, 0x46, 0x4d,
	0x4a, 0x46, 0x47, 0x4e, 0x42, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b,
	0x33, 0x30, 0x30, 0x30, 0x4c, 0x4f, 0x46, 0x4e, 0x46, 0x4d, 0x4a, 0x46, 0x47, 0x4e, 0x42, 0x42,
	0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Unk3000_GCBMILHPIKA_proto_rawDescOnce sync.Once
	file_Unk3000_GCBMILHPIKA_proto_rawDescData = file_Unk3000_GCBMILHPIKA_proto_rawDesc
)

func file_Unk3000_GCBMILHPIKA_proto_rawDescGZIP() []byte {
	file_Unk3000_GCBMILHPIKA_proto_rawDescOnce.Do(func() {
		file_Unk3000_GCBMILHPIKA_proto_rawDescData = protoimpl.X.CompressGZIP(file_Unk3000_GCBMILHPIKA_proto_rawDescData)
	})
	return file_Unk3000_GCBMILHPIKA_proto_rawDescData
}

var file_Unk3000_GCBMILHPIKA_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Unk3000_GCBMILHPIKA_proto_goTypes = []interface{}{
	(*Unk3000_GCBMILHPIKA)(nil),                     // 0: Unk3000_GCBMILHPIKA
	(*Unk3000_GCBMILHPIKA_Unk3000_PPGINNAFPIF)(nil), // 1: Unk3000_GCBMILHPIKA.Unk3000_PPGINNAFPIF
}
var file_Unk3000_GCBMILHPIKA_proto_depIdxs = []int32{
	1, // 0: Unk3000_GCBMILHPIKA.Unk3000_EBIEGNHLMFP:type_name -> Unk3000_GCBMILHPIKA.Unk3000_PPGINNAFPIF
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Unk3000_GCBMILHPIKA_proto_init() }
func file_Unk3000_GCBMILHPIKA_proto_init() {
	if File_Unk3000_GCBMILHPIKA_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Unk3000_GCBMILHPIKA_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unk3000_GCBMILHPIKA); i {
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
		file_Unk3000_GCBMILHPIKA_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unk3000_GCBMILHPIKA_Unk3000_PPGINNAFPIF); i {
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
			RawDescriptor: file_Unk3000_GCBMILHPIKA_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Unk3000_GCBMILHPIKA_proto_goTypes,
		DependencyIndexes: file_Unk3000_GCBMILHPIKA_proto_depIdxs,
		MessageInfos:      file_Unk3000_GCBMILHPIKA_proto_msgTypes,
	}.Build()
	File_Unk3000_GCBMILHPIKA_proto = out.File
	file_Unk3000_GCBMILHPIKA_proto_rawDesc = nil
	file_Unk3000_GCBMILHPIKA_proto_goTypes = nil
	file_Unk3000_GCBMILHPIKA_proto_depIdxs = nil
}
