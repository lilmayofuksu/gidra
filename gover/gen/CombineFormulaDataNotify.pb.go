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
// source: CombineFormulaDataNotify.proto

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

// CmdId: 632
// EnetChannelId: 0
// EnetIsReliable: true
type CombineFormulaDataNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CombineId uint32 `protobuf:"varint,6,opt,name=combine_id,json=combineId,proto3" json:"combine_id,omitempty"`
	IsLocked  bool   `protobuf:"varint,3,opt,name=is_locked,json=isLocked,proto3" json:"is_locked,omitempty"`
}

func (x *CombineFormulaDataNotify) Reset() {
	*x = CombineFormulaDataNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CombineFormulaDataNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CombineFormulaDataNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CombineFormulaDataNotify) ProtoMessage() {}

func (x *CombineFormulaDataNotify) ProtoReflect() protoreflect.Message {
	mi := &file_CombineFormulaDataNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CombineFormulaDataNotify.ProtoReflect.Descriptor instead.
func (*CombineFormulaDataNotify) Descriptor() ([]byte, []int) {
	return file_CombineFormulaDataNotify_proto_rawDescGZIP(), []int{0}
}

func (x *CombineFormulaDataNotify) GetCombineId() uint32 {
	if x != nil {
		return x.CombineId
	}
	return 0
}

func (x *CombineFormulaDataNotify) GetIsLocked() bool {
	if x != nil {
		return x.IsLocked
	}
	return false
}

var File_CombineFormulaDataNotify_proto protoreflect.FileDescriptor

var file_CombineFormulaDataNotify_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61,
	0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x56, 0x0a, 0x18, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x75,
	0x6c, 0x61, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x63, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69,
	0x73, 0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x69, 0x73, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CombineFormulaDataNotify_proto_rawDescOnce sync.Once
	file_CombineFormulaDataNotify_proto_rawDescData = file_CombineFormulaDataNotify_proto_rawDesc
)

func file_CombineFormulaDataNotify_proto_rawDescGZIP() []byte {
	file_CombineFormulaDataNotify_proto_rawDescOnce.Do(func() {
		file_CombineFormulaDataNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_CombineFormulaDataNotify_proto_rawDescData)
	})
	return file_CombineFormulaDataNotify_proto_rawDescData
}

var file_CombineFormulaDataNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CombineFormulaDataNotify_proto_goTypes = []interface{}{
	(*CombineFormulaDataNotify)(nil), // 0: CombineFormulaDataNotify
}
var file_CombineFormulaDataNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_CombineFormulaDataNotify_proto_init() }
func file_CombineFormulaDataNotify_proto_init() {
	if File_CombineFormulaDataNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CombineFormulaDataNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CombineFormulaDataNotify); i {
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
			RawDescriptor: file_CombineFormulaDataNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CombineFormulaDataNotify_proto_goTypes,
		DependencyIndexes: file_CombineFormulaDataNotify_proto_depIdxs,
		MessageInfos:      file_CombineFormulaDataNotify_proto_msgTypes,
	}.Build()
	File_CombineFormulaDataNotify_proto = out.File
	file_CombineFormulaDataNotify_proto_rawDesc = nil
	file_CombineFormulaDataNotify_proto_goTypes = nil
	file_CombineFormulaDataNotify_proto_depIdxs = nil
}
