//
// Copyright © 2024 honeysense.com All rights reserved.
// Author: sunrui
// Date: 2024-03-26 10:46:24

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: pkg/rpc/proto/chat/chat_card.proto

package chat

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

// 名片
type ChatCard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 姓名
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// 职位
	Position string `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
	// 公司名称
	CompanyName string `protobuf:"bytes,3,opt,name=company_name,json=companyName,proto3" json:"company_name,omitempty"`
	// 手机号码
	MobileNumber string `protobuf:"bytes,4,opt,name=mobile_number,json=mobileNumber,proto3" json:"mobile_number,omitempty"`
	// 办公电话
	OfficePhone string `protobuf:"bytes,5,opt,name=office_phone,json=officePhone,proto3" json:"office_phone,omitempty"`
	// 电子邮箱
	Email string `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	// 地址
	Address string `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	// 个人简介
	Bio string `protobuf:"bytes,8,opt,name=bio,proto3" json:"bio,omitempty"`
	// 社交媒体账号列表
	SocialMediaAccounts []*SocialMedia `protobuf:"bytes,9,rep,name=social_media_accounts,json=socialMediaAccounts,proto3" json:"social_media_accounts,omitempty"`
	// 自定义标签
	CustomLabels map[string]string `protobuf:"bytes,10,rep,name=custom_labels,json=customLabels,proto3" json:"custom_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ChatCard) Reset() {
	*x = ChatCard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rpc_proto_chat_chat_card_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatCard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatCard) ProtoMessage() {}

func (x *ChatCard) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_proto_chat_chat_card_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatCard.ProtoReflect.Descriptor instead.
func (*ChatCard) Descriptor() ([]byte, []int) {
	return file_pkg_rpc_proto_chat_chat_card_proto_rawDescGZIP(), []int{0}
}

func (x *ChatCard) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChatCard) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *ChatCard) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *ChatCard) GetMobileNumber() string {
	if x != nil {
		return x.MobileNumber
	}
	return ""
}

func (x *ChatCard) GetOfficePhone() string {
	if x != nil {
		return x.OfficePhone
	}
	return ""
}

func (x *ChatCard) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ChatCard) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ChatCard) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *ChatCard) GetSocialMediaAccounts() []*SocialMedia {
	if x != nil {
		return x.SocialMediaAccounts
	}
	return nil
}

func (x *ChatCard) GetCustomLabels() map[string]string {
	if x != nil {
		return x.CustomLabels
	}
	return nil
}

// 定义社交媒体账号消息类型
type SocialMedia struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 社交媒体平台类型（如：LinkedIn、Twitter、Facebook）
	Platform string `protobuf:"bytes,1,opt,name=platform,proto3" json:"platform,omitempty"`
	// 用户名或者链接
	UsernameOrLink string `protobuf:"bytes,2,opt,name=username_or_link,json=usernameOrLink,proto3" json:"username_or_link,omitempty"`
}

func (x *SocialMedia) Reset() {
	*x = SocialMedia{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rpc_proto_chat_chat_card_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocialMedia) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocialMedia) ProtoMessage() {}

func (x *SocialMedia) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rpc_proto_chat_chat_card_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocialMedia.ProtoReflect.Descriptor instead.
func (*SocialMedia) Descriptor() ([]byte, []int) {
	return file_pkg_rpc_proto_chat_chat_card_proto_rawDescGZIP(), []int{1}
}

func (x *SocialMedia) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *SocialMedia) GetUsernameOrLink() string {
	if x != nil {
		return x.UsernameOrLink
	}
	return ""
}

var File_pkg_rpc_proto_chat_chat_card_proto protoreflect.FileDescriptor

var file_pkg_rpc_proto_chat_chat_card_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x22, 0xd2, 0x03, 0x0a, 0x08, 0x43, 0x68, 0x61,
	0x74, 0x43, 0x61, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x21, 0x0a,
	0x0c, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62,
	0x69, 0x6f, 0x12, 0x53, 0x0a, 0x15, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4d, 0x65, 0x64,
	0x69, 0x61, 0x52, 0x13, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x53, 0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e,
	0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x43, 0x61, 0x72, 0x64, 0x2e, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x3f, 0x0a, 0x11,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x53, 0x0a,
	0x0b, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x28, 0x0a, 0x10, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x6f, 0x72, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x4f, 0x72, 0x4c, 0x69,
	0x6e, 0x6b, 0x42, 0x14, 0x5a, 0x12, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_rpc_proto_chat_chat_card_proto_rawDescOnce sync.Once
	file_pkg_rpc_proto_chat_chat_card_proto_rawDescData = file_pkg_rpc_proto_chat_chat_card_proto_rawDesc
)

func file_pkg_rpc_proto_chat_chat_card_proto_rawDescGZIP() []byte {
	file_pkg_rpc_proto_chat_chat_card_proto_rawDescOnce.Do(func() {
		file_pkg_rpc_proto_chat_chat_card_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_rpc_proto_chat_chat_card_proto_rawDescData)
	})
	return file_pkg_rpc_proto_chat_chat_card_proto_rawDescData
}

var file_pkg_rpc_proto_chat_chat_card_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_rpc_proto_chat_chat_card_proto_goTypes = []interface{}{
	(*ChatCard)(nil),    // 0: pkg.rpc.proto.chat.ChatCard
	(*SocialMedia)(nil), // 1: pkg.rpc.proto.chat.SocialMedia
	nil,                 // 2: pkg.rpc.proto.chat.ChatCard.CustomLabelsEntry
}
var file_pkg_rpc_proto_chat_chat_card_proto_depIdxs = []int32{
	1, // 0: pkg.rpc.proto.chat.ChatCard.social_media_accounts:type_name -> pkg.rpc.proto.chat.SocialMedia
	2, // 1: pkg.rpc.proto.chat.ChatCard.custom_labels:type_name -> pkg.rpc.proto.chat.ChatCard.CustomLabelsEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_rpc_proto_chat_chat_card_proto_init() }
func file_pkg_rpc_proto_chat_chat_card_proto_init() {
	if File_pkg_rpc_proto_chat_chat_card_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_rpc_proto_chat_chat_card_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatCard); i {
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
		file_pkg_rpc_proto_chat_chat_card_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocialMedia); i {
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
			RawDescriptor: file_pkg_rpc_proto_chat_chat_card_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_rpc_proto_chat_chat_card_proto_goTypes,
		DependencyIndexes: file_pkg_rpc_proto_chat_chat_card_proto_depIdxs,
		MessageInfos:      file_pkg_rpc_proto_chat_chat_card_proto_msgTypes,
	}.Build()
	File_pkg_rpc_proto_chat_chat_card_proto = out.File
	file_pkg_rpc_proto_chat_chat_card_proto_rawDesc = nil
	file_pkg_rpc_proto_chat_chat_card_proto_goTypes = nil
	file_pkg_rpc_proto_chat_chat_card_proto_depIdxs = nil
}
