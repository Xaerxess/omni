// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.24.4
// source: omni/oidc/oidc.proto

package oidc

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthenticateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Auth Request ID.
	AuthRequestId string `protobuf:"bytes,1,opt,name=auth_request_id,json=authRequestId,proto3" json:"auth_request_id,omitempty"`
}

func (x *AuthenticateRequest) Reset() {
	*x = AuthenticateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omni_oidc_oidc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateRequest) ProtoMessage() {}

func (x *AuthenticateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_omni_oidc_oidc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateRequest.ProtoReflect.Descriptor instead.
func (*AuthenticateRequest) Descriptor() ([]byte, []int) {
	return file_omni_oidc_oidc_proto_rawDescGZIP(), []int{0}
}

func (x *AuthenticateRequest) GetAuthRequestId() string {
	if x != nil {
		return x.AuthRequestId
	}
	return ""
}

type AuthenticateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// URL to redirect the user to.
	RedirectUrl string `protobuf:"bytes,1,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
}

func (x *AuthenticateResponse) Reset() {
	*x = AuthenticateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omni_oidc_oidc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateResponse) ProtoMessage() {}

func (x *AuthenticateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_omni_oidc_oidc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateResponse.ProtoReflect.Descriptor instead.
func (*AuthenticateResponse) Descriptor() ([]byte, []int) {
	return file_omni_oidc_oidc_proto_rawDescGZIP(), []int{1}
}

func (x *AuthenticateResponse) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

var File_omni_oidc_oidc_proto protoreflect.FileDescriptor

var file_omni_oidc_oidc_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6f, 0x6d, 0x6e, 0x69, 0x2f, 0x6f, 0x69, 0x64, 0x63, 0x2f, 0x6f, 0x69, 0x64, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6f, 0x69, 0x64, 0x63, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x13, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x26, 0x0a, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0x39, 0x0a, 0x14, 0x41, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x55, 0x72, 0x6c, 0x32, 0x54, 0x0a, 0x0b, 0x4f, 0x49, 0x44, 0x43, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x12, 0x19, 0x2e, 0x6f, 0x69, 0x64, 0x63, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x6f, 0x69, 0x64, 0x63, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x64, 0x65, 0x72, 0x6f, 0x6c, 0x61,
	0x62, 0x73, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6f, 0x6d, 0x6e, 0x69, 0x2f, 0x6f, 0x69, 0x64, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_omni_oidc_oidc_proto_rawDescOnce sync.Once
	file_omni_oidc_oidc_proto_rawDescData = file_omni_oidc_oidc_proto_rawDesc
)

func file_omni_oidc_oidc_proto_rawDescGZIP() []byte {
	file_omni_oidc_oidc_proto_rawDescOnce.Do(func() {
		file_omni_oidc_oidc_proto_rawDescData = protoimpl.X.CompressGZIP(file_omni_oidc_oidc_proto_rawDescData)
	})
	return file_omni_oidc_oidc_proto_rawDescData
}

var file_omni_oidc_oidc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_omni_oidc_oidc_proto_goTypes = []interface{}{
	(*AuthenticateRequest)(nil),  // 0: oidc.AuthenticateRequest
	(*AuthenticateResponse)(nil), // 1: oidc.AuthenticateResponse
}
var file_omni_oidc_oidc_proto_depIdxs = []int32{
	0, // 0: oidc.OIDCService.Authenticate:input_type -> oidc.AuthenticateRequest
	1, // 1: oidc.OIDCService.Authenticate:output_type -> oidc.AuthenticateResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_omni_oidc_oidc_proto_init() }
func file_omni_oidc_oidc_proto_init() {
	if File_omni_oidc_oidc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_omni_oidc_oidc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateRequest); i {
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
		file_omni_oidc_oidc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateResponse); i {
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
			RawDescriptor: file_omni_oidc_oidc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_omni_oidc_oidc_proto_goTypes,
		DependencyIndexes: file_omni_oidc_oidc_proto_depIdxs,
		MessageInfos:      file_omni_oidc_oidc_proto_msgTypes,
	}.Build()
	File_omni_oidc_oidc_proto = out.File
	file_omni_oidc_oidc_proto_rawDesc = nil
	file_omni_oidc_oidc_proto_goTypes = nil
	file_omni_oidc_oidc_proto_depIdxs = nil
}
