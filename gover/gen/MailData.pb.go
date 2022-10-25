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
// source: MailData.proto

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

type MailData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MailId              uint32              `protobuf:"varint,1,opt,name=mail_id,json=mailId,proto3" json:"mail_id,omitempty"`
	MailTextContent     *MailTextContent    `protobuf:"bytes,4,opt,name=mail_text_content,json=mailTextContent,proto3" json:"mail_text_content,omitempty"`
	ItemList            []*MailItem         `protobuf:"bytes,7,rep,name=item_list,json=itemList,proto3" json:"item_list,omitempty"`
	SendTime            uint32              `protobuf:"varint,8,opt,name=send_time,json=sendTime,proto3" json:"send_time,omitempty"`
	ExpireTime          uint32              `protobuf:"varint,9,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
	Importance          uint32              `protobuf:"varint,10,opt,name=importance,proto3" json:"importance,omitempty"`
	IsRead              bool                `protobuf:"varint,11,opt,name=is_read,json=isRead,proto3" json:"is_read,omitempty"`
	IsAttachmentGot     bool                `protobuf:"varint,12,opt,name=is_attachment_got,json=isAttachmentGot,proto3" json:"is_attachment_got,omitempty"`
	ConfigId            uint32              `protobuf:"varint,13,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
	ArgumentList        []string            `protobuf:"bytes,14,rep,name=argument_list,json=argumentList,proto3" json:"argument_list,omitempty"`
	Unk2700_NDPPGJKJOMH Unk2700_CBJEDMGOBPL `protobuf:"varint,15,opt,name=Unk2700_NDPPGJKJOMH,json=Unk2700NDPPGJKJOMH,proto3,enum=Unk2700_CBJEDMGOBPL" json:"Unk2700_NDPPGJKJOMH,omitempty"`
}

func (x *MailData) Reset() {
	*x = MailData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MailData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailData) ProtoMessage() {}

func (x *MailData) ProtoReflect() protoreflect.Message {
	mi := &file_MailData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailData.ProtoReflect.Descriptor instead.
func (*MailData) Descriptor() ([]byte, []int) {
	return file_MailData_proto_rawDescGZIP(), []int{0}
}

func (x *MailData) GetMailId() uint32 {
	if x != nil {
		return x.MailId
	}
	return 0
}

func (x *MailData) GetMailTextContent() *MailTextContent {
	if x != nil {
		return x.MailTextContent
	}
	return nil
}

func (x *MailData) GetItemList() []*MailItem {
	if x != nil {
		return x.ItemList
	}
	return nil
}

func (x *MailData) GetSendTime() uint32 {
	if x != nil {
		return x.SendTime
	}
	return 0
}

func (x *MailData) GetExpireTime() uint32 {
	if x != nil {
		return x.ExpireTime
	}
	return 0
}

func (x *MailData) GetImportance() uint32 {
	if x != nil {
		return x.Importance
	}
	return 0
}

func (x *MailData) GetIsRead() bool {
	if x != nil {
		return x.IsRead
	}
	return false
}

func (x *MailData) GetIsAttachmentGot() bool {
	if x != nil {
		return x.IsAttachmentGot
	}
	return false
}

func (x *MailData) GetConfigId() uint32 {
	if x != nil {
		return x.ConfigId
	}
	return 0
}

func (x *MailData) GetArgumentList() []string {
	if x != nil {
		return x.ArgumentList
	}
	return nil
}

func (x *MailData) GetUnk2700_NDPPGJKJOMH() Unk2700_CBJEDMGOBPL {
	if x != nil {
		return x.Unk2700_NDPPGJKJOMH
	}
	return Unk2700_CBJEDMGOBPL_Unk2700_CBJEDMGOBPL_Unk2700_MBLDLJOKLBL
}

var File_MailData_proto protoreflect.FileDescriptor

var file_MailData_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x4d, 0x61, 0x69, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0e, 0x4d, 0x61, 0x69, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x4d, 0x61, 0x69, 0x6c, 0x54, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30,
	0x5f, 0x43, 0x42, 0x4a, 0x45, 0x44, 0x4d, 0x47, 0x4f, 0x42, 0x50, 0x4c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb5, 0x03, 0x0a, 0x08, 0x4d, 0x61, 0x69, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x17, 0x0a, 0x07, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x11, 0x6d, 0x61, 0x69, 0x6c,
	0x5f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x54, 0x65, 0x78, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x0f, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x65, 0x78, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4d, 0x61, 0x69, 0x6c,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x69, 0x73, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69,
	0x73, 0x52, 0x65, 0x61, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x73, 0x5f, 0x61, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x67, 0x6f, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0f, 0x69, 0x73, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x6f,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x0e, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4e,
	0x44, 0x50, 0x50, 0x47, 0x4a, 0x4b, 0x4a, 0x4f, 0x4d, 0x48, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x14, 0x2e, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x43, 0x42, 0x4a, 0x45, 0x44,
	0x4d, 0x47, 0x4f, 0x42, 0x50, 0x4c, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x4e,
	0x44, 0x50, 0x50, 0x47, 0x4a, 0x4b, 0x4a, 0x4f, 0x4d, 0x48, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67,
	0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MailData_proto_rawDescOnce sync.Once
	file_MailData_proto_rawDescData = file_MailData_proto_rawDesc
)

func file_MailData_proto_rawDescGZIP() []byte {
	file_MailData_proto_rawDescOnce.Do(func() {
		file_MailData_proto_rawDescData = protoimpl.X.CompressGZIP(file_MailData_proto_rawDescData)
	})
	return file_MailData_proto_rawDescData
}

var file_MailData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MailData_proto_goTypes = []interface{}{
	(*MailData)(nil),         // 0: MailData
	(*MailTextContent)(nil),  // 1: MailTextContent
	(*MailItem)(nil),         // 2: MailItem
	(Unk2700_CBJEDMGOBPL)(0), // 3: Unk2700_CBJEDMGOBPL
}
var file_MailData_proto_depIdxs = []int32{
	1, // 0: MailData.mail_text_content:type_name -> MailTextContent
	2, // 1: MailData.item_list:type_name -> MailItem
	3, // 2: MailData.Unk2700_NDPPGJKJOMH:type_name -> Unk2700_CBJEDMGOBPL
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_MailData_proto_init() }
func file_MailData_proto_init() {
	if File_MailData_proto != nil {
		return
	}
	file_MailItem_proto_init()
	file_MailTextContent_proto_init()
	file_Unk2700_CBJEDMGOBPL_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_MailData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailData); i {
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
			RawDescriptor: file_MailData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MailData_proto_goTypes,
		DependencyIndexes: file_MailData_proto_depIdxs,
		MessageInfos:      file_MailData_proto_msgTypes,
	}.Build()
	File_MailData_proto = out.File
	file_MailData_proto_rawDesc = nil
	file_MailData_proto_goTypes = nil
	file_MailData_proto_depIdxs = nil
}
