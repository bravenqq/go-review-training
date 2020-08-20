// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: fileServer.proto

package file

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ServeContentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *ServeContentRequest) Reset() {
	*x = ServeContentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileServer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServeContentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServeContentRequest) ProtoMessage() {}

func (x *ServeContentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fileServer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServeContentRequest.ProtoReflect.Descriptor instead.
func (*ServeContentRequest) Descriptor() ([]byte, []int) {
	return file_fileServer_proto_rawDescGZIP(), []int{0}
}

func (x *ServeContentRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fileServer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_fileServer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_fileServer_proto_rawDescGZIP(), []int{1}
}

func (x *Chunk) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_fileServer_proto protoreflect.FileDescriptor

var file_fileServer_proto_rawDesc = []byte{
	0x0a, 0x10, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x29, 0x0a, 0x13, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x22, 0x21, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x18, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0x47, 0x0a, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x12, 0x3a, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b,
	0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x00, 0x30, 0x01, 0x42,
	0x08, 0x5a, 0x06, 0x2e, 0x3b, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_fileServer_proto_rawDescOnce sync.Once
	file_fileServer_proto_rawDescData = file_fileServer_proto_rawDesc
)

func file_fileServer_proto_rawDescGZIP() []byte {
	file_fileServer_proto_rawDescOnce.Do(func() {
		file_fileServer_proto_rawDescData = protoimpl.X.CompressGZIP(file_fileServer_proto_rawDescData)
	})
	return file_fileServer_proto_rawDescData
}

var file_fileServer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_fileServer_proto_goTypes = []interface{}{
	(*ServeContentRequest)(nil), // 0: file.ServeContentRequest
	(*Chunk)(nil),               // 1: file.chunk
}
var file_fileServer_proto_depIdxs = []int32{
	0, // 0: file.FileServe.ServeContent:input_type -> file.ServeContentRequest
	1, // 1: file.FileServe.ServeContent:output_type -> file.chunk
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fileServer_proto_init() }
func file_fileServer_proto_init() {
	if File_fileServer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fileServer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServeContentRequest); i {
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
		file_fileServer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
			RawDescriptor: file_fileServer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fileServer_proto_goTypes,
		DependencyIndexes: file_fileServer_proto_depIdxs,
		MessageInfos:      file_fileServer_proto_msgTypes,
	}.Build()
	File_fileServer_proto = out.File
	file_fileServer_proto_rawDesc = nil
	file_fileServer_proto_goTypes = nil
	file_fileServer_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FileServeClient is the client API for FileServe service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileServeClient interface {
	ServeContent(ctx context.Context, in *ServeContentRequest, opts ...grpc.CallOption) (FileServe_ServeContentClient, error)
}

type fileServeClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServeClient(cc grpc.ClientConnInterface) FileServeClient {
	return &fileServeClient{cc}
}

func (c *fileServeClient) ServeContent(ctx context.Context, in *ServeContentRequest, opts ...grpc.CallOption) (FileServe_ServeContentClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FileServe_serviceDesc.Streams[0], "/file.FileServe/ServeContent", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileServeServeContentClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileServe_ServeContentClient interface {
	Recv() (*Chunk, error)
	grpc.ClientStream
}

type fileServeServeContentClient struct {
	grpc.ClientStream
}

func (x *fileServeServeContentClient) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileServeServer is the server API for FileServe service.
type FileServeServer interface {
	ServeContent(*ServeContentRequest, FileServe_ServeContentServer) error
}

// UnimplementedFileServeServer can be embedded to have forward compatible implementations.
type UnimplementedFileServeServer struct {
}

func (*UnimplementedFileServeServer) ServeContent(*ServeContentRequest, FileServe_ServeContentServer) error {
	return status.Errorf(codes.Unimplemented, "method ServeContent not implemented")
}

func RegisterFileServeServer(s *grpc.Server, srv FileServeServer) {
	s.RegisterService(&_FileServe_serviceDesc, srv)
}

func _FileServe_ServeContent_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ServeContentRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileServeServer).ServeContent(m, &fileServeServeContentServer{stream})
}

type FileServe_ServeContentServer interface {
	Send(*Chunk) error
	grpc.ServerStream
}

type fileServeServeContentServer struct {
	grpc.ServerStream
}

func (x *fileServeServeContentServer) Send(m *Chunk) error {
	return x.ServerStream.SendMsg(m)
}

var _FileServe_serviceDesc = grpc.ServiceDesc{
	ServiceName: "file.FileServe",
	HandlerType: (*FileServeServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServeContent",
			Handler:       _FileServe_ServeContent_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "fileServer.proto",
}
