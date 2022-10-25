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
// source: PlantFlowerGiveFriendFlowerRsp.proto

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

// CmdId: 8386
// EnetChannelId: 0
// EnetIsReliable: true
type PlantFlowerGiveFriendFlowerRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LimitFlowerList []uint32 `protobuf:"varint,5,rep,packed,name=limit_flower_list,json=limitFlowerList,proto3" json:"limit_flower_list,omitempty"`
	Retcode         int32    `protobuf:"varint,3,opt,name=retcode,proto3" json:"retcode,omitempty"`
	ScheduleId      uint32   `protobuf:"varint,14,opt,name=schedule_id,json=scheduleId,proto3" json:"schedule_id,omitempty"`
}

func (x *PlantFlowerGiveFriendFlowerRsp) Reset() {
	*x = PlantFlowerGiveFriendFlowerRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlantFlowerGiveFriendFlowerRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlantFlowerGiveFriendFlowerRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlantFlowerGiveFriendFlowerRsp) ProtoMessage() {}

func (x *PlantFlowerGiveFriendFlowerRsp) ProtoReflect() protoreflect.Message {
	mi := &file_PlantFlowerGiveFriendFlowerRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlantFlowerGiveFriendFlowerRsp.ProtoReflect.Descriptor instead.
func (*PlantFlowerGiveFriendFlowerRsp) Descriptor() ([]byte, []int) {
	return file_PlantFlowerGiveFriendFlowerRsp_proto_rawDescGZIP(), []int{0}
}

func (x *PlantFlowerGiveFriendFlowerRsp) GetLimitFlowerList() []uint32 {
	if x != nil {
		return x.LimitFlowerList
	}
	return nil
}

func (x *PlantFlowerGiveFriendFlowerRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *PlantFlowerGiveFriendFlowerRsp) GetScheduleId() uint32 {
	if x != nil {
		return x.ScheduleId
	}
	return 0
}

var File_PlantFlowerGiveFriendFlowerRsp_proto protoreflect.FileDescriptor

var file_PlantFlowerGiveFriendFlowerRsp_proto_rawDesc = []byte{
	0x0a, 0x24, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x47, 0x69, 0x76,
	0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x46, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x73, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x1e, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x46, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x47, 0x69, 0x76, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x46, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x11, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0d, 0x52, 0x0f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x46, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64,
	0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlantFlowerGiveFriendFlowerRsp_proto_rawDescOnce sync.Once
	file_PlantFlowerGiveFriendFlowerRsp_proto_rawDescData = file_PlantFlowerGiveFriendFlowerRsp_proto_rawDesc
)

func file_PlantFlowerGiveFriendFlowerRsp_proto_rawDescGZIP() []byte {
	file_PlantFlowerGiveFriendFlowerRsp_proto_rawDescOnce.Do(func() {
		file_PlantFlowerGiveFriendFlowerRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlantFlowerGiveFriendFlowerRsp_proto_rawDescData)
	})
	return file_PlantFlowerGiveFriendFlowerRsp_proto_rawDescData
}

var file_PlantFlowerGiveFriendFlowerRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlantFlowerGiveFriendFlowerRsp_proto_goTypes = []interface{}{
	(*PlantFlowerGiveFriendFlowerRsp)(nil), // 0: PlantFlowerGiveFriendFlowerRsp
}
var file_PlantFlowerGiveFriendFlowerRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_PlantFlowerGiveFriendFlowerRsp_proto_init() }
func file_PlantFlowerGiveFriendFlowerRsp_proto_init() {
	if File_PlantFlowerGiveFriendFlowerRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_PlantFlowerGiveFriendFlowerRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlantFlowerGiveFriendFlowerRsp); i {
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
			RawDescriptor: file_PlantFlowerGiveFriendFlowerRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlantFlowerGiveFriendFlowerRsp_proto_goTypes,
		DependencyIndexes: file_PlantFlowerGiveFriendFlowerRsp_proto_depIdxs,
		MessageInfos:      file_PlantFlowerGiveFriendFlowerRsp_proto_msgTypes,
	}.Build()
	File_PlantFlowerGiveFriendFlowerRsp_proto = out.File
	file_PlantFlowerGiveFriendFlowerRsp_proto_rawDesc = nil
	file_PlantFlowerGiveFriendFlowerRsp_proto_goTypes = nil
	file_PlantFlowerGiveFriendFlowerRsp_proto_depIdxs = nil
}
