// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: internal/microservices/auth/AuthPB/authPB.proto

package authpb

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

type UserS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *UserS) Reset() {
	*x = UserS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_microservices_auth_AuthPB_authPB_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserS) ProtoMessage() {}

func (x *UserS) ProtoReflect() protoreflect.Message {
	mi := &file_internal_microservices_auth_AuthPB_authPB_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserS.ProtoReflect.Descriptor instead.
func (*UserS) Descriptor() ([]byte, []int) {
	return file_internal_microservices_auth_AuthPB_authPB_proto_rawDescGZIP(), []int{0}
}

func (x *UserS) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserS) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *UserId) Reset() {
	*x = UserId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_microservices_auth_AuthPB_authPB_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserId) ProtoMessage() {}

func (x *UserId) ProtoReflect() protoreflect.Message {
	mi := &file_internal_microservices_auth_AuthPB_authPB_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserId.ProtoReflect.Descriptor instead.
func (*UserId) Descriptor() ([]byte, []int) {
	return file_internal_microservices_auth_AuthPB_authPB_proto_rawDescGZIP(), []int{1}
}

func (x *UserId) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type ResponseUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
}

func (x *ResponseUser) Reset() {
	*x = ResponseUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_microservices_auth_AuthPB_authPB_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseUser) ProtoMessage() {}

func (x *ResponseUser) ProtoReflect() protoreflect.Message {
	mi := &file_internal_microservices_auth_AuthPB_authPB_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseUser.ProtoReflect.Descriptor instead.
func (*ResponseUser) Descriptor() ([]byte, []int) {
	return file_internal_microservices_auth_AuthPB_authPB_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseUser) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ResponseUser) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

var File_internal_microservices_auth_AuthPB_authPB_proto protoreflect.FileDescriptor

var file_internal_microservices_auth_AuthPB_authPB_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x41, 0x75,
	0x74, 0x68, 0x50, 0x42, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x50, 0x42, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x61, 0x75, 0x74, 0x68, 0x70, 0x62, 0x22, 0x3f, 0x0a, 0x05, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x20, 0x0a, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x0c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xac, 0x01, 0x0a, 0x10, 0x41, 0x75, 0x74,
	0x68, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a,
	0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x1a, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x70,
	0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00,
	0x12, 0x2e, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0d, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x1a, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x70,
	0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00,
	0x12, 0x35, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12,
	0x0e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a,
	0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x3b, 0x61, 0x75,
	0x74, 0x68, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_microservices_auth_AuthPB_authPB_proto_rawDescOnce sync.Once
	file_internal_microservices_auth_AuthPB_authPB_proto_rawDescData = file_internal_microservices_auth_AuthPB_authPB_proto_rawDesc
)

func file_internal_microservices_auth_AuthPB_authPB_proto_rawDescGZIP() []byte {
	file_internal_microservices_auth_AuthPB_authPB_proto_rawDescOnce.Do(func() {
		file_internal_microservices_auth_AuthPB_authPB_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_microservices_auth_AuthPB_authPB_proto_rawDescData)
	})
	return file_internal_microservices_auth_AuthPB_authPB_proto_rawDescData
}

var file_internal_microservices_auth_AuthPB_authPB_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internal_microservices_auth_AuthPB_authPB_proto_goTypes = []interface{}{
	(*UserS)(nil),        // 0: authpb.UserS
	(*UserId)(nil),       // 1: authpb.UserId
	(*ResponseUser)(nil), // 2: authpb.ResponseUser
}
var file_internal_microservices_auth_AuthPB_authPB_proto_depIdxs = []int32{
	0, // 0: authpb.AuthMicroservice.Register:input_type -> authpb.UserS
	0, // 1: authpb.AuthMicroservice.Login:input_type -> authpb.UserS
	1, // 2: authpb.AuthMicroservice.GetUserById:input_type -> authpb.UserId
	2, // 3: authpb.AuthMicroservice.Register:output_type -> authpb.ResponseUser
	2, // 4: authpb.AuthMicroservice.Login:output_type -> authpb.ResponseUser
	2, // 5: authpb.AuthMicroservice.GetUserById:output_type -> authpb.ResponseUser
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_microservices_auth_AuthPB_authPB_proto_init() }
func file_internal_microservices_auth_AuthPB_authPB_proto_init() {
	if File_internal_microservices_auth_AuthPB_authPB_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_microservices_auth_AuthPB_authPB_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserS); i {
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
		file_internal_microservices_auth_AuthPB_authPB_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserId); i {
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
		file_internal_microservices_auth_AuthPB_authPB_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseUser); i {
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
			RawDescriptor: file_internal_microservices_auth_AuthPB_authPB_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_microservices_auth_AuthPB_authPB_proto_goTypes,
		DependencyIndexes: file_internal_microservices_auth_AuthPB_authPB_proto_depIdxs,
		MessageInfos:      file_internal_microservices_auth_AuthPB_authPB_proto_msgTypes,
	}.Build()
	File_internal_microservices_auth_AuthPB_authPB_proto = out.File
	file_internal_microservices_auth_AuthPB_authPB_proto_rawDesc = nil
	file_internal_microservices_auth_AuthPB_authPB_proto_goTypes = nil
	file_internal_microservices_auth_AuthPB_authPB_proto_depIdxs = nil
}