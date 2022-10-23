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
// source: RegionInfo.proto

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

type RegionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GateserverIp               string            `protobuf:"bytes,1,opt,name=gateserver_ip,json=gateserverIp,proto3" json:"gateserver_ip,omitempty"`
	GateserverPort             uint32            `protobuf:"varint,2,opt,name=gateserver_port,json=gateserverPort,proto3" json:"gateserver_port,omitempty"`
	PayCallbackUrl             string            `protobuf:"bytes,3,opt,name=pay_callback_url,json=payCallbackUrl,proto3" json:"pay_callback_url,omitempty"`
	AreaType                   string            `protobuf:"bytes,7,opt,name=area_type,json=areaType,proto3" json:"area_type,omitempty"`
	ResourceUrl                string            `protobuf:"bytes,8,opt,name=resource_url,json=resourceUrl,proto3" json:"resource_url,omitempty"`
	DataUrl                    string            `protobuf:"bytes,9,opt,name=data_url,json=dataUrl,proto3" json:"data_url,omitempty"`
	FeedbackUrl                string            `protobuf:"bytes,10,opt,name=feedback_url,json=feedbackUrl,proto3" json:"feedback_url,omitempty"`
	BulletinUrl                string            `protobuf:"bytes,11,opt,name=bulletin_url,json=bulletinUrl,proto3" json:"bulletin_url,omitempty"`
	ResourceUrlBak             string            `protobuf:"bytes,12,opt,name=resource_url_bak,json=resourceUrlBak,proto3" json:"resource_url_bak,omitempty"`
	DataUrlBak                 string            `protobuf:"bytes,13,opt,name=data_url_bak,json=dataUrlBak,proto3" json:"data_url_bak,omitempty"`
	ClientDataVersion          uint32            `protobuf:"varint,14,opt,name=client_data_version,json=clientDataVersion,proto3" json:"client_data_version,omitempty"`
	HandbookUrl                string            `protobuf:"bytes,16,opt,name=handbook_url,json=handbookUrl,proto3" json:"handbook_url,omitempty"`
	ClientSilenceDataVersion   uint32            `protobuf:"varint,18,opt,name=client_silence_data_version,json=clientSilenceDataVersion,proto3" json:"client_silence_data_version,omitempty"`
	ClientDataMd5              string            `protobuf:"bytes,19,opt,name=client_data_md5,json=clientDataMd5,proto3" json:"client_data_md5,omitempty"`
	ClientSilenceDataMd5       string            `protobuf:"bytes,20,opt,name=client_silence_data_md5,json=clientSilenceDataMd5,proto3" json:"client_silence_data_md5,omitempty"`
	ResVersionConfig           *ResVersionConfig `protobuf:"bytes,22,opt,name=res_version_config,json=resVersionConfig,proto3" json:"res_version_config,omitempty"`
	SecretKey                  []byte            `protobuf:"bytes,23,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	OfficialCommunityUrl       string            `protobuf:"bytes,24,opt,name=official_community_url,json=officialCommunityUrl,proto3" json:"official_community_url,omitempty"`
	ClientVersionSuffix        string            `protobuf:"bytes,26,opt,name=client_version_suffix,json=clientVersionSuffix,proto3" json:"client_version_suffix,omitempty"`
	ClientSilenceVersionSuffix string            `protobuf:"bytes,27,opt,name=client_silence_version_suffix,json=clientSilenceVersionSuffix,proto3" json:"client_silence_version_suffix,omitempty"`
	UseGateserverDomainName    bool              `protobuf:"varint,28,opt,name=use_gateserver_domain_name,json=useGateserverDomainName,proto3" json:"use_gateserver_domain_name,omitempty"`
	GateserverDomainName       string            `protobuf:"bytes,29,opt,name=gateserver_domain_name,json=gateserverDomainName,proto3" json:"gateserver_domain_name,omitempty"`
	UserCenterUrl              string            `protobuf:"bytes,30,opt,name=user_center_url,json=userCenterUrl,proto3" json:"user_center_url,omitempty"`
	AccountBindUrl             string            `protobuf:"bytes,31,opt,name=account_bind_url,json=accountBindUrl,proto3" json:"account_bind_url,omitempty"`
	CdkeyUrl                   string            `protobuf:"bytes,32,opt,name=cdkey_url,json=cdkeyUrl,proto3" json:"cdkey_url,omitempty"`
	PrivacyPolicyUrl           string            `protobuf:"bytes,33,opt,name=privacy_policy_url,json=privacyPolicyUrl,proto3" json:"privacy_policy_url,omitempty"`
	NextResourceUrl            string            `protobuf:"bytes,34,opt,name=next_resource_url,json=nextResourceUrl,proto3" json:"next_resource_url,omitempty"`
	NextResVersionConfig       *ResVersionConfig `protobuf:"bytes,35,opt,name=next_res_version_config,json=nextResVersionConfig,proto3" json:"next_res_version_config,omitempty"`
}

func (x *RegionInfo) Reset() {
	*x = RegionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RegionInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegionInfo) ProtoMessage() {}

func (x *RegionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_RegionInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegionInfo.ProtoReflect.Descriptor instead.
func (*RegionInfo) Descriptor() ([]byte, []int) {
	return file_RegionInfo_proto_rawDescGZIP(), []int{0}
}

func (x *RegionInfo) GetGateserverIp() string {
	if x != nil {
		return x.GateserverIp
	}
	return ""
}

func (x *RegionInfo) GetGateserverPort() uint32 {
	if x != nil {
		return x.GateserverPort
	}
	return 0
}

func (x *RegionInfo) GetPayCallbackUrl() string {
	if x != nil {
		return x.PayCallbackUrl
	}
	return ""
}

func (x *RegionInfo) GetAreaType() string {
	if x != nil {
		return x.AreaType
	}
	return ""
}

func (x *RegionInfo) GetResourceUrl() string {
	if x != nil {
		return x.ResourceUrl
	}
	return ""
}

func (x *RegionInfo) GetDataUrl() string {
	if x != nil {
		return x.DataUrl
	}
	return ""
}

func (x *RegionInfo) GetFeedbackUrl() string {
	if x != nil {
		return x.FeedbackUrl
	}
	return ""
}

func (x *RegionInfo) GetBulletinUrl() string {
	if x != nil {
		return x.BulletinUrl
	}
	return ""
}

func (x *RegionInfo) GetResourceUrlBak() string {
	if x != nil {
		return x.ResourceUrlBak
	}
	return ""
}

func (x *RegionInfo) GetDataUrlBak() string {
	if x != nil {
		return x.DataUrlBak
	}
	return ""
}

func (x *RegionInfo) GetClientDataVersion() uint32 {
	if x != nil {
		return x.ClientDataVersion
	}
	return 0
}

func (x *RegionInfo) GetHandbookUrl() string {
	if x != nil {
		return x.HandbookUrl
	}
	return ""
}

func (x *RegionInfo) GetClientSilenceDataVersion() uint32 {
	if x != nil {
		return x.ClientSilenceDataVersion
	}
	return 0
}

func (x *RegionInfo) GetClientDataMd5() string {
	if x != nil {
		return x.ClientDataMd5
	}
	return ""
}

func (x *RegionInfo) GetClientSilenceDataMd5() string {
	if x != nil {
		return x.ClientSilenceDataMd5
	}
	return ""
}

func (x *RegionInfo) GetResVersionConfig() *ResVersionConfig {
	if x != nil {
		return x.ResVersionConfig
	}
	return nil
}

func (x *RegionInfo) GetSecretKey() []byte {
	if x != nil {
		return x.SecretKey
	}
	return nil
}

func (x *RegionInfo) GetOfficialCommunityUrl() string {
	if x != nil {
		return x.OfficialCommunityUrl
	}
	return ""
}

func (x *RegionInfo) GetClientVersionSuffix() string {
	if x != nil {
		return x.ClientVersionSuffix
	}
	return ""
}

func (x *RegionInfo) GetClientSilenceVersionSuffix() string {
	if x != nil {
		return x.ClientSilenceVersionSuffix
	}
	return ""
}

func (x *RegionInfo) GetUseGateserverDomainName() bool {
	if x != nil {
		return x.UseGateserverDomainName
	}
	return false
}

func (x *RegionInfo) GetGateserverDomainName() string {
	if x != nil {
		return x.GateserverDomainName
	}
	return ""
}

func (x *RegionInfo) GetUserCenterUrl() string {
	if x != nil {
		return x.UserCenterUrl
	}
	return ""
}

func (x *RegionInfo) GetAccountBindUrl() string {
	if x != nil {
		return x.AccountBindUrl
	}
	return ""
}

func (x *RegionInfo) GetCdkeyUrl() string {
	if x != nil {
		return x.CdkeyUrl
	}
	return ""
}

func (x *RegionInfo) GetPrivacyPolicyUrl() string {
	if x != nil {
		return x.PrivacyPolicyUrl
	}
	return ""
}

func (x *RegionInfo) GetNextResourceUrl() string {
	if x != nil {
		return x.NextResourceUrl
	}
	return ""
}

func (x *RegionInfo) GetNextResVersionConfig() *ResVersionConfig {
	if x != nil {
		return x.NextResVersionConfig
	}
	return nil
}

var File_RegionInfo_proto protoreflect.FileDescriptor

var file_RegionInfo_proto_rawDesc = []byte{
	0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x52, 0x65, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf5, 0x09, 0x0a, 0x0a, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a, 0x0d, 0x67, 0x61, 0x74,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x67, 0x61, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x70, 0x12, 0x27,
	0x0a, 0x0f, 0x67, 0x61, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x67, 0x61, 0x74, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x61, 0x79, 0x5f, 0x63,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x70, 0x61, 0x79, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x55, 0x72,
	0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x72, 0x65, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x72,
	0x6c, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x55, 0x72, 0x6c, 0x12, 0x21, 0x0a, 0x0c,
	0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x55, 0x72, 0x6c, 0x12,
	0x21, 0x0a, 0x0c, 0x62, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x69, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x69, 0x6e, 0x55,
	0x72, 0x6c, 0x12, 0x28, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75,
	0x72, 0x6c, 0x5f, 0x62, 0x61, 0x6b, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x42, 0x61, 0x6b, 0x12, 0x20, 0x0a, 0x0c,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x75, 0x72, 0x6c, 0x5f, 0x62, 0x61, 0x6b, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x55, 0x72, 0x6c, 0x42, 0x61, 0x6b, 0x12, 0x2e,
	0x0a, 0x13, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21,
	0x0a, 0x0c, 0x68, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x68, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x6f, 0x6b, 0x55, 0x72,
	0x6c, 0x12, 0x3d, 0x0a, 0x1b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x6c, 0x65,
	0x6e, 0x63, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x18, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x69,
	0x6c, 0x65, 0x6e, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x26, 0x0a, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x6d, 0x64, 0x35, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x64, 0x35, 0x12, 0x35, 0x0a, 0x17, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x69, 0x6c, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x6d, 0x64, 0x35, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x53, 0x69, 0x6c, 0x65, 0x6e, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x64, 0x35, 0x12,
	0x3f, 0x0a, 0x12, 0x72, 0x65, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x52, 0x65,
	0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x10,
	0x72, 0x65, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x17,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12,
	0x34, 0x0a, 0x16, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x14, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x32, 0x0a, 0x15, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x1a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x53, 0x75, 0x66, 0x66, 0x69, 0x78, 0x12, 0x41, 0x0a, 0x1d, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x6c, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x1a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x6c, 0x65, 0x6e, 0x63, 0x65, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x66, 0x66, 0x69, 0x78, 0x12, 0x3b, 0x0a, 0x1a,
	0x75, 0x73, 0x65, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x17, 0x75, 0x73, 0x65, 0x47, 0x61, 0x74, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x67, 0x61, 0x74,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x67, 0x61, 0x74, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x26, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x43, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x1f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x69, 0x6e, 0x64, 0x55, 0x72,
	0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x64, 0x6b, 0x65, 0x79, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x20,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x64, 0x6b, 0x65, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x2c,
	0x0a, 0x12, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x21, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x63, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x2a, 0x0a, 0x11,
	0x6e, 0x65, 0x78, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x22, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x48, 0x0a, 0x17, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x72, 0x65, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x23, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x52, 0x65, 0x73, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x14, 0x6e, 0x65,
	0x78, 0x74, 0x52, 0x65, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_RegionInfo_proto_rawDescOnce sync.Once
	file_RegionInfo_proto_rawDescData = file_RegionInfo_proto_rawDesc
)

func file_RegionInfo_proto_rawDescGZIP() []byte {
	file_RegionInfo_proto_rawDescOnce.Do(func() {
		file_RegionInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_RegionInfo_proto_rawDescData)
	})
	return file_RegionInfo_proto_rawDescData
}

var file_RegionInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RegionInfo_proto_goTypes = []interface{}{
	(*RegionInfo)(nil),       // 0: RegionInfo
	(*ResVersionConfig)(nil), // 1: ResVersionConfig
}
var file_RegionInfo_proto_depIdxs = []int32{
	1, // 0: RegionInfo.res_version_config:type_name -> ResVersionConfig
	1, // 1: RegionInfo.next_res_version_config:type_name -> ResVersionConfig
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_RegionInfo_proto_init() }
func file_RegionInfo_proto_init() {
	if File_RegionInfo_proto != nil {
		return
	}
	file_ResVersionConfig_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RegionInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegionInfo); i {
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
			RawDescriptor: file_RegionInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RegionInfo_proto_goTypes,
		DependencyIndexes: file_RegionInfo_proto_depIdxs,
		MessageInfos:      file_RegionInfo_proto_msgTypes,
	}.Build()
	File_RegionInfo_proto = out.File
	file_RegionInfo_proto_rawDesc = nil
	file_RegionInfo_proto_goTypes = nil
	file_RegionInfo_proto_depIdxs = nil
}
