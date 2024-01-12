// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: sending.proto

package sending

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

type SendEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *SendEmailRequest) Reset() {
	*x = SendEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sending_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailRequest) ProtoMessage() {}

func (x *SendEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sending_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailRequest.ProtoReflect.Descriptor instead.
func (*SendEmailRequest) Descriptor() ([]byte, []int) {
	return file_sending_proto_rawDescGZIP(), []int{0}
}

func (x *SendEmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type SendEmailReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendEmailReply) Reset() {
	*x = SendEmailReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sending_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailReply) ProtoMessage() {}

func (x *SendEmailReply) ProtoReflect() protoreflect.Message {
	mi := &file_sending_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailReply.ProtoReflect.Descriptor instead.
func (*SendEmailReply) Descriptor() ([]byte, []int) {
	return file_sending_proto_rawDescGZIP(), []int{1}
}

func (x *SendEmailReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_sending_proto protoreflect.FileDescriptor

var file_sending_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x28, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x22, 0x2a, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x49,
	0x0a, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x19, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x18, 0x5a, 0x16, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sending_proto_rawDescOnce sync.Once
	file_sending_proto_rawDescData = file_sending_proto_rawDesc
)

func file_sending_proto_rawDescGZIP() []byte {
	file_sending_proto_rawDescOnce.Do(func() {
		file_sending_proto_rawDescData = protoimpl.X.CompressGZIP(file_sending_proto_rawDescData)
	})
	return file_sending_proto_rawDescData
}

var file_sending_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sending_proto_goTypes = []interface{}{
	(*SendEmailRequest)(nil), // 0: sending.SendEmailRequest
	(*SendEmailReply)(nil),   // 1: sending.SendEmailReply
}
var file_sending_proto_depIdxs = []int32{
	0, // 0: sending.Sender.SendEmail:input_type -> sending.SendEmailRequest
	1, // 1: sending.Sender.SendEmail:output_type -> sending.SendEmailReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sending_proto_init() }
func file_sending_proto_init() {
	if File_sending_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sending_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailRequest); i {
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
		file_sending_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailReply); i {
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
			RawDescriptor: file_sending_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sending_proto_goTypes,
		DependencyIndexes: file_sending_proto_depIdxs,
		MessageInfos:      file_sending_proto_msgTypes,
	}.Build()
	File_sending_proto = out.File
	file_sending_proto_rawDesc = nil
	file_sending_proto_goTypes = nil
	file_sending_proto_depIdxs = nil
}
