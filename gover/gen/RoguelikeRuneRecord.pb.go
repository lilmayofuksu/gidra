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
// source: RoguelikeRuneRecord.proto

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

type RoguelikeRuneRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeftCount uint32 `protobuf:"varint,14,opt,name=left_count,json=leftCount,proto3" json:"left_count,omitempty"`
	RuneId    uint32 `protobuf:"varint,6,opt,name=rune_id,json=runeId,proto3" json:"rune_id,omitempty"`
	MaxCount  uint32 `protobuf:"varint,4,opt,name=max_count,json=maxCount,proto3" json:"max_count,omitempty"`
}

func (x *RoguelikeRuneRecord) Reset() {
	*x = RoguelikeRuneRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RoguelikeRuneRecord_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoguelikeRuneRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoguelikeRuneRecord) ProtoMessage() {}

func (x *RoguelikeRuneRecord) ProtoReflect() protoreflect.Message {
	mi := &file_RoguelikeRuneRecord_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoguelikeRuneRecord.ProtoReflect.Descriptor instead.
func (*RoguelikeRuneRecord) Descriptor() ([]byte, []int) {
	return file_RoguelikeRuneRecord_proto_rawDescGZIP(), []int{0}
}

func (x *RoguelikeRuneRecord) GetLeftCount() uint32 {
	if x != nil {
		return x.LeftCount
	}
	return 0
}

func (x *RoguelikeRuneRecord) GetRuneId() uint32 {
	if x != nil {
		return x.RuneId
	}
	return 0
}

func (x *RoguelikeRuneRecord) GetMaxCount() uint32 {
	if x != nil {
		return x.MaxCount
	}
	return 0
}

var File_RoguelikeRuneRecord_proto protoreflect.FileDescriptor

var file_RoguelikeRuneRecord_proto_rawDesc = []byte{
	0x0a, 0x19, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x6c, 0x69, 0x6b, 0x65, 0x52, 0x75, 0x6e, 0x65, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x13, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x6c, 0x69, 0x6b, 0x65, 0x52, 0x75, 0x6e, 0x65, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x65, 0x66, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6c, 0x65, 0x66, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x72, 0x75, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61,
	0x78, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6d,
	0x61, 0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RoguelikeRuneRecord_proto_rawDescOnce sync.Once
	file_RoguelikeRuneRecord_proto_rawDescData = file_RoguelikeRuneRecord_proto_rawDesc
)

func file_RoguelikeRuneRecord_proto_rawDescGZIP() []byte {
	file_RoguelikeRuneRecord_proto_rawDescOnce.Do(func() {
		file_RoguelikeRuneRecord_proto_rawDescData = protoimpl.X.CompressGZIP(file_RoguelikeRuneRecord_proto_rawDescData)
	})
	return file_RoguelikeRuneRecord_proto_rawDescData
}

var file_RoguelikeRuneRecord_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RoguelikeRuneRecord_proto_goTypes = []interface{}{
	(*RoguelikeRuneRecord)(nil), // 0: RoguelikeRuneRecord
}
var file_RoguelikeRuneRecord_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RoguelikeRuneRecord_proto_init() }
func file_RoguelikeRuneRecord_proto_init() {
	if File_RoguelikeRuneRecord_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_RoguelikeRuneRecord_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoguelikeRuneRecord); i {
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
			RawDescriptor: file_RoguelikeRuneRecord_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RoguelikeRuneRecord_proto_goTypes,
		DependencyIndexes: file_RoguelikeRuneRecord_proto_depIdxs,
		MessageInfos:      file_RoguelikeRuneRecord_proto_msgTypes,
	}.Build()
	File_RoguelikeRuneRecord_proto = out.File
	file_RoguelikeRuneRecord_proto_rawDesc = nil
	file_RoguelikeRuneRecord_proto_goTypes = nil
	file_RoguelikeRuneRecord_proto_depIdxs = nil
}
