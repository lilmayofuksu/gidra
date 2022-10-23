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
// source: GadgetPlayUidOpNotify.proto

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

// CmdId: 875
// EnetChannelId: 0
// EnetIsReliable: true
type GadgetPlayUidOpNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityId  uint32   `protobuf:"varint,11,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	UidList   []uint32 `protobuf:"varint,2,rep,packed,name=uid_list,json=uidList,proto3" json:"uid_list,omitempty"`
	PlayType  uint32   `protobuf:"varint,6,opt,name=play_type,json=playType,proto3" json:"play_type,omitempty"`
	ParamStr  string   `protobuf:"bytes,1,opt,name=param_str,json=paramStr,proto3" json:"param_str,omitempty"`
	Op        uint32   `protobuf:"varint,7,opt,name=op,proto3" json:"op,omitempty"`
	ParamList []uint32 `protobuf:"varint,4,rep,packed,name=param_list,json=paramList,proto3" json:"param_list,omitempty"`
}

func (x *GadgetPlayUidOpNotify) Reset() {
	*x = GadgetPlayUidOpNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GadgetPlayUidOpNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GadgetPlayUidOpNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GadgetPlayUidOpNotify) ProtoMessage() {}

func (x *GadgetPlayUidOpNotify) ProtoReflect() protoreflect.Message {
	mi := &file_GadgetPlayUidOpNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GadgetPlayUidOpNotify.ProtoReflect.Descriptor instead.
func (*GadgetPlayUidOpNotify) Descriptor() ([]byte, []int) {
	return file_GadgetPlayUidOpNotify_proto_rawDescGZIP(), []int{0}
}

func (x *GadgetPlayUidOpNotify) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *GadgetPlayUidOpNotify) GetUidList() []uint32 {
	if x != nil {
		return x.UidList
	}
	return nil
}

func (x *GadgetPlayUidOpNotify) GetPlayType() uint32 {
	if x != nil {
		return x.PlayType
	}
	return 0
}

func (x *GadgetPlayUidOpNotify) GetParamStr() string {
	if x != nil {
		return x.ParamStr
	}
	return ""
}

func (x *GadgetPlayUidOpNotify) GetOp() uint32 {
	if x != nil {
		return x.Op
	}
	return 0
}

func (x *GadgetPlayUidOpNotify) GetParamList() []uint32 {
	if x != nil {
		return x.ParamList
	}
	return nil
}

var File_GadgetPlayUidOpNotify_proto protoreflect.FileDescriptor

var file_GadgetPlayUidOpNotify_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x55, 0x69, 0x64, 0x4f,
	0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x01,
	0x0a, 0x15, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x55, 0x69, 0x64, 0x4f,
	0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x07, 0x75, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x53, 0x74, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x70, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x09, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GadgetPlayUidOpNotify_proto_rawDescOnce sync.Once
	file_GadgetPlayUidOpNotify_proto_rawDescData = file_GadgetPlayUidOpNotify_proto_rawDesc
)

func file_GadgetPlayUidOpNotify_proto_rawDescGZIP() []byte {
	file_GadgetPlayUidOpNotify_proto_rawDescOnce.Do(func() {
		file_GadgetPlayUidOpNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_GadgetPlayUidOpNotify_proto_rawDescData)
	})
	return file_GadgetPlayUidOpNotify_proto_rawDescData
}

var file_GadgetPlayUidOpNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GadgetPlayUidOpNotify_proto_goTypes = []interface{}{
	(*GadgetPlayUidOpNotify)(nil), // 0: GadgetPlayUidOpNotify
}
var file_GadgetPlayUidOpNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GadgetPlayUidOpNotify_proto_init() }
func file_GadgetPlayUidOpNotify_proto_init() {
	if File_GadgetPlayUidOpNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GadgetPlayUidOpNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GadgetPlayUidOpNotify); i {
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
			RawDescriptor: file_GadgetPlayUidOpNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GadgetPlayUidOpNotify_proto_goTypes,
		DependencyIndexes: file_GadgetPlayUidOpNotify_proto_depIdxs,
		MessageInfos:      file_GadgetPlayUidOpNotify_proto_msgTypes,
	}.Build()
	File_GadgetPlayUidOpNotify_proto = out.File
	file_GadgetPlayUidOpNotify_proto_rawDesc = nil
	file_GadgetPlayUidOpNotify_proto_goTypes = nil
	file_GadgetPlayUidOpNotify_proto_depIdxs = nil
}
