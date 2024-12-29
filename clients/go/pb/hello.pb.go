// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v4.25.1
// source: hello.proto

package pb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// message 对应生成代码的 struct
// 定义客户端请求的数据格式
// @validator=HelloReq
type HelloReq struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// [修饰符] 类型 字段名 = 标识符;
	// @inject_tag: json:"id" validate:"required,min=1"
	Id            int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id" validate:"required,min=1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HelloReq) Reset() {
	*x = HelloReq{}
	mi := &file_hello_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HelloReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReq) ProtoMessage() {}

func (x *HelloReq) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReq.ProtoReflect.Descriptor instead.
func (*HelloReq) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{0}
}

func (x *HelloReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 定义服务端响应的数据格式
type HelloReply struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// @inject_tag: json:"name"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	// @inject_tag: json:"message"
	Message       string `protobuf:"bytes,2,opt,name=message,proto3" json:"message"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	mi := &file_hello_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// @validator=InfoReq
type InfoReq struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 主要用于grpc validator参数校验
	// @inject_tag: json:"name" validate:"required,min=1"
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" validate:"required,min=1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InfoReq) Reset() {
	*x = InfoReq{}
	mi := &file_hello_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoReq) ProtoMessage() {}

func (x *InfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoReq.ProtoReflect.Descriptor instead.
func (*InfoReq) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{2}
}

func (x *InfoReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// InfoReply info reply
type InfoReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Address       string                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InfoReply) Reset() {
	*x = InfoReply{}
	mi := &file_hello_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoReply) ProtoMessage() {}

func (x *InfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoReply.ProtoReflect.Descriptor instead.
func (*InfoReply) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{3}
}

func (x *InfoReply) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *InfoReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UserEntity struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserEntity) Reset() {
	*x = UserEntity{}
	mi := &file_hello_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserEntity) ProtoMessage() {}

func (x *UserEntity) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserEntity.ProtoReflect.Descriptor instead.
func (*UserEntity) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{4}
}

func (x *UserEntity) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserEntity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type BatchUsersReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ids           []int64                `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BatchUsersReq) Reset() {
	*x = BatchUsersReq{}
	mi := &file_hello_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchUsersReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchUsersReq) ProtoMessage() {}

func (x *BatchUsersReq) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchUsersReq.ProtoReflect.Descriptor instead.
func (*BatchUsersReq) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{5}
}

func (x *BatchUsersReq) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type BatchUsersReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Users         []*UserEntity          `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BatchUsersReply) Reset() {
	*x = BatchUsersReply{}
	mi := &file_hello_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchUsersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchUsersReply) ProtoMessage() {}

func (x *BatchUsersReply) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchUsersReply.ProtoReflect.Descriptor instead.
func (*BatchUsersReply) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{6}
}

func (x *BatchUsersReply) GetUsers() []*UserEntity {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_hello_proto protoreflect.FileDescriptor

var file_hello_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x41,
	0x70, 0x70, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1a, 0x0a, 0x08, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3a, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x1d, 0x0a, 0x07, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x3f, 0x0a, 0x09, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x30, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x21, 0x0a, 0x0d, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x43, 0x0a, 0x0f, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x41, 0x70, 0x70,
	0x2e, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x32, 0xa1, 0x02,
	0x0a, 0x0e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x56, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x18, 0x2e, 0x41,
	0x70, 0x70, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x41, 0x70, 0x70, 0x2e, 0x47, 0x72, 0x70,
	0x63, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x61, 0x79, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x53, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x17, 0x2e, 0x41, 0x70, 0x70, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x41, 0x70, 0x70, 0x2e,
	0x47, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x76,
	0x31, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x62, 0x0a,
	0x0a, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1d, 0x2e, 0x41, 0x70,
	0x70, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x41, 0x70, 0x70,
	0x2e, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_hello_proto_rawDescOnce sync.Once
	file_hello_proto_rawDescData = file_hello_proto_rawDesc
)

func file_hello_proto_rawDescGZIP() []byte {
	file_hello_proto_rawDescOnce.Do(func() {
		file_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_hello_proto_rawDescData)
	})
	return file_hello_proto_rawDescData
}

var file_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_hello_proto_goTypes = []any{
	(*HelloReq)(nil),        // 0: App.Grpc.Hello.HelloReq
	(*HelloReply)(nil),      // 1: App.Grpc.Hello.HelloReply
	(*InfoReq)(nil),         // 2: App.Grpc.Hello.InfoReq
	(*InfoReply)(nil),       // 3: App.Grpc.Hello.InfoReply
	(*UserEntity)(nil),      // 4: App.Grpc.Hello.UserEntity
	(*BatchUsersReq)(nil),   // 5: App.Grpc.Hello.BatchUsersReq
	(*BatchUsersReply)(nil), // 6: App.Grpc.Hello.BatchUsersReply
}
var file_hello_proto_depIdxs = []int32{
	4, // 0: App.Grpc.Hello.BatchUsersReply.users:type_name -> App.Grpc.Hello.UserEntity
	0, // 1: App.Grpc.Hello.GreeterService.SayHello:input_type -> App.Grpc.Hello.HelloReq
	2, // 2: App.Grpc.Hello.GreeterService.Info:input_type -> App.Grpc.Hello.InfoReq
	5, // 3: App.Grpc.Hello.GreeterService.BatchUsers:input_type -> App.Grpc.Hello.BatchUsersReq
	1, // 4: App.Grpc.Hello.GreeterService.SayHello:output_type -> App.Grpc.Hello.HelloReply
	3, // 5: App.Grpc.Hello.GreeterService.Info:output_type -> App.Grpc.Hello.InfoReply
	6, // 6: App.Grpc.Hello.GreeterService.BatchUsers:output_type -> App.Grpc.Hello.BatchUsersReply
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_hello_proto_init() }
func file_hello_proto_init() {
	if File_hello_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hello_proto_goTypes,
		DependencyIndexes: file_hello_proto_depIdxs,
		MessageInfos:      file_hello_proto_msgTypes,
	}.Build()
	File_hello_proto = out.File
	file_hello_proto_rawDesc = nil
	file_hello_proto_goTypes = nil
	file_hello_proto_depIdxs = nil
}
