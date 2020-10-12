// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: node_and_port_scan.proto

package proto

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

type NodeScanStartNeoMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeScanTimeout int64 `protobuf:"varint,1,opt,name=NodeScanTimeout,proto3" json:"NodeScanTimeout,omitempty"`
	PortScanTimeout int64 `protobuf:"varint,2,opt,name=PortScanTimeout,proto3" json:"PortScanTimeout,omitempty"`
}

func (x *NodeScanStartNeoMessage) Reset() {
	*x = NodeScanStartNeoMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_and_port_scan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeScanStartNeoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeScanStartNeoMessage) ProtoMessage() {}

func (x *NodeScanStartNeoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_node_and_port_scan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeScanStartNeoMessage.ProtoReflect.Descriptor instead.
func (*NodeScanStartNeoMessage) Descriptor() ([]byte, []int) {
	return file_node_and_port_scan_proto_rawDescGZIP(), []int{0}
}

func (x *NodeScanStartNeoMessage) GetNodeScanTimeout() int64 {
	if x != nil {
		return x.NodeScanTimeout
	}
	return 0
}

func (x *NodeScanStartNeoMessage) GetPortScanTimeout() int64 {
	if x != nil {
		return x.PortScanTimeout
	}
	return 0
}

type NodeScanReferenceNeoMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CreatedAt string `protobuf:"bytes,2,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	Done      bool   `protobuf:"varint,3,opt,name=Done,proto3" json:"Done,omitempty"`
}

func (x *NodeScanReferenceNeoMessage) Reset() {
	*x = NodeScanReferenceNeoMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_and_port_scan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeScanReferenceNeoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeScanReferenceNeoMessage) ProtoMessage() {}

func (x *NodeScanReferenceNeoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_node_and_port_scan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeScanReferenceNeoMessage.ProtoReflect.Descriptor instead.
func (*NodeScanReferenceNeoMessage) Descriptor() ([]byte, []int) {
	return file_node_and_port_scan_proto_rawDescGZIP(), []int{1}
}

func (x *NodeScanReferenceNeoMessage) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *NodeScanReferenceNeoMessage) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *NodeScanReferenceNeoMessage) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

var File_node_and_port_scan_proto protoreflect.FileDescriptor

var file_node_and_port_scan_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x73, 0x63, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x63, 0x6f, 0x6d, 0x2e,
	0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e,
	0x6c, 0x69, 0x77, 0x61, 0x73, 0x63, 0x2e, 0x6e, 0x65, 0x6f, 0x22, 0x6d, 0x0a, 0x17, 0x4e, 0x6f,
	0x64, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4e, 0x65, 0x6f, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x63, 0x61,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f,
	0x4e, 0x6f, 0x64, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12,
	0x28, 0x0a, 0x0f, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x63, 0x61, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x63,
	0x61, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x5f, 0x0a, 0x1b, 0x4e, 0x6f, 0x64,
	0x65, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x65,
	0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x6f, 0x6e, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x44, 0x6f, 0x6e, 0x65, 0x32, 0xa3, 0x01, 0x0a, 0x19, 0x4e,
	0x6f, 0x64, 0x65, 0x41, 0x6e, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x63, 0x61, 0x6e, 0x4e, 0x65,
	0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x85, 0x01, 0x0a, 0x0d, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x12, 0x37, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78,
	0x2e, 0x6c, 0x69, 0x77, 0x61, 0x73, 0x63, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x53, 0x63, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4e, 0x65, 0x6f, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x3b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e,
	0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x6c, 0x69, 0x77, 0x61, 0x73, 0x63,
	0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x65, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x6f, 0x6a, 0x6e, 0x74, 0x66, 0x78, 0x2f, 0x6c, 0x69, 0x77, 0x61, 0x73, 0x63, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_node_and_port_scan_proto_rawDescOnce sync.Once
	file_node_and_port_scan_proto_rawDescData = file_node_and_port_scan_proto_rawDesc
)

func file_node_and_port_scan_proto_rawDescGZIP() []byte {
	file_node_and_port_scan_proto_rawDescOnce.Do(func() {
		file_node_and_port_scan_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_and_port_scan_proto_rawDescData)
	})
	return file_node_and_port_scan_proto_rawDescData
}

var file_node_and_port_scan_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_node_and_port_scan_proto_goTypes = []interface{}{
	(*NodeScanStartNeoMessage)(nil),     // 0: com.pojtinger.felix.liwasc.neo.NodeScanStartNeoMessage
	(*NodeScanReferenceNeoMessage)(nil), // 1: com.pojtinger.felix.liwasc.neo.NodeScanReferenceNeoMessage
}
var file_node_and_port_scan_proto_depIdxs = []int32{
	0, // 0: com.pojtinger.felix.liwasc.neo.NodeAndPortScanNeoService.StartNodeScan:input_type -> com.pojtinger.felix.liwasc.neo.NodeScanStartNeoMessage
	1, // 1: com.pojtinger.felix.liwasc.neo.NodeAndPortScanNeoService.StartNodeScan:output_type -> com.pojtinger.felix.liwasc.neo.NodeScanReferenceNeoMessage
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_node_and_port_scan_proto_init() }
func file_node_and_port_scan_proto_init() {
	if File_node_and_port_scan_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_node_and_port_scan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeScanStartNeoMessage); i {
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
		file_node_and_port_scan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeScanReferenceNeoMessage); i {
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
			RawDescriptor: file_node_and_port_scan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_node_and_port_scan_proto_goTypes,
		DependencyIndexes: file_node_and_port_scan_proto_depIdxs,
		MessageInfos:      file_node_and_port_scan_proto_msgTypes,
	}.Build()
	File_node_and_port_scan_proto = out.File
	file_node_and_port_scan_proto_rawDesc = nil
	file_node_and_port_scan_proto_goTypes = nil
	file_node_and_port_scan_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NodeAndPortScanNeoServiceClient is the client API for NodeAndPortScanNeoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeAndPortScanNeoServiceClient interface {
	StartNodeScan(ctx context.Context, in *NodeScanStartNeoMessage, opts ...grpc.CallOption) (*NodeScanReferenceNeoMessage, error)
}

type nodeAndPortScanNeoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeAndPortScanNeoServiceClient(cc grpc.ClientConnInterface) NodeAndPortScanNeoServiceClient {
	return &nodeAndPortScanNeoServiceClient{cc}
}

func (c *nodeAndPortScanNeoServiceClient) StartNodeScan(ctx context.Context, in *NodeScanStartNeoMessage, opts ...grpc.CallOption) (*NodeScanReferenceNeoMessage, error) {
	out := new(NodeScanReferenceNeoMessage)
	err := c.cc.Invoke(ctx, "/com.pojtinger.felix.liwasc.neo.NodeAndPortScanNeoService/StartNodeScan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeAndPortScanNeoServiceServer is the server API for NodeAndPortScanNeoService service.
type NodeAndPortScanNeoServiceServer interface {
	StartNodeScan(context.Context, *NodeScanStartNeoMessage) (*NodeScanReferenceNeoMessage, error)
}

// UnimplementedNodeAndPortScanNeoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNodeAndPortScanNeoServiceServer struct {
}

func (*UnimplementedNodeAndPortScanNeoServiceServer) StartNodeScan(context.Context, *NodeScanStartNeoMessage) (*NodeScanReferenceNeoMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartNodeScan not implemented")
}

func RegisterNodeAndPortScanNeoServiceServer(s *grpc.Server, srv NodeAndPortScanNeoServiceServer) {
	s.RegisterService(&_NodeAndPortScanNeoService_serviceDesc, srv)
}

func _NodeAndPortScanNeoService_StartNodeScan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeScanStartNeoMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAndPortScanNeoServiceServer).StartNodeScan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.pojtinger.felix.liwasc.neo.NodeAndPortScanNeoService/StartNodeScan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAndPortScanNeoServiceServer).StartNodeScan(ctx, req.(*NodeScanStartNeoMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeAndPortScanNeoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.pojtinger.felix.liwasc.neo.NodeAndPortScanNeoService",
	HandlerType: (*NodeAndPortScanNeoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartNodeScan",
			Handler:    _NodeAndPortScanNeoService_StartNodeScan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node_and_port_scan.proto",
}
