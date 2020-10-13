// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: node_and_port_scan.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type NodeScanNeoMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CreatedAt string `protobuf:"bytes,2,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	Done      bool   `protobuf:"varint,3,opt,name=Done,proto3" json:"Done,omitempty"`
}

func (x *NodeScanNeoMessage) Reset() {
	*x = NodeScanNeoMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_and_port_scan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeScanNeoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeScanNeoMessage) ProtoMessage() {}

func (x *NodeScanNeoMessage) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use NodeScanNeoMessage.ProtoReflect.Descriptor instead.
func (*NodeScanNeoMessage) Descriptor() ([]byte, []int) {
	return file_node_and_port_scan_proto_rawDescGZIP(), []int{1}
}

func (x *NodeScanNeoMessage) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *NodeScanNeoMessage) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *NodeScanNeoMessage) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

type NodeNeoMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CreatedAt  string `protobuf:"bytes,2,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	Priority   int64  `protobuf:"varint,3,opt,name=Priority,proto3" json:"Priority,omitempty"`
	MACAddress string `protobuf:"bytes,4,opt,name=MACAddress,proto3" json:"MACAddress,omitempty"`
	IPAddress  string `protobuf:"bytes,5,opt,name=IPAddress,proto3" json:"IPAddress,omitempty"`
	NodeScanID int64  `protobuf:"varint,6,opt,name=NodeScanID,proto3" json:"NodeScanID,omitempty"`
}

func (x *NodeNeoMessage) Reset() {
	*x = NodeNeoMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_and_port_scan_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeNeoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeNeoMessage) ProtoMessage() {}

func (x *NodeNeoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_node_and_port_scan_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeNeoMessage.ProtoReflect.Descriptor instead.
func (*NodeNeoMessage) Descriptor() ([]byte, []int) {
	return file_node_and_port_scan_proto_rawDescGZIP(), []int{2}
}

func (x *NodeNeoMessage) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *NodeNeoMessage) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *NodeNeoMessage) GetPriority() int64 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *NodeNeoMessage) GetMACAddress() string {
	if x != nil {
		return x.MACAddress
	}
	return ""
}

func (x *NodeNeoMessage) GetIPAddress() string {
	if x != nil {
		return x.IPAddress
	}
	return ""
}

func (x *NodeNeoMessage) GetNodeScanID() int64 {
	if x != nil {
		return x.NodeScanID
	}
	return 0
}

type PortScanNeoMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CreatedAt string `protobuf:"bytes,2,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	Done      bool   `protobuf:"varint,3,opt,name=Done,proto3" json:"Done,omitempty"`
	NodeID    int64  `protobuf:"varint,4,opt,name=NodeID,proto3" json:"NodeID,omitempty"`
}

func (x *PortScanNeoMessage) Reset() {
	*x = PortScanNeoMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_and_port_scan_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortScanNeoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortScanNeoMessage) ProtoMessage() {}

func (x *PortScanNeoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_node_and_port_scan_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortScanNeoMessage.ProtoReflect.Descriptor instead.
func (*PortScanNeoMessage) Descriptor() ([]byte, []int) {
	return file_node_and_port_scan_proto_rawDescGZIP(), []int{3}
}

func (x *PortScanNeoMessage) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *PortScanNeoMessage) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *PortScanNeoMessage) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

func (x *PortScanNeoMessage) GetNodeID() int64 {
	if x != nil {
		return x.NodeID
	}
	return 0
}

type PortNeoMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CreatedAt         string `protobuf:"bytes,2,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	Priority          int64  `protobuf:"varint,3,opt,name=Priority,proto3" json:"Priority,omitempty"`
	PortNumber        int64  `protobuf:"varint,4,opt,name=PortNumber,proto3" json:"PortNumber,omitempty"`
	TransportProtocol string `protobuf:"bytes,5,opt,name=TransportProtocol,proto3" json:"TransportProtocol,omitempty"`
	PortScanID        int64  `protobuf:"varint,6,opt,name=PortScanID,proto3" json:"PortScanID,omitempty"`
}

func (x *PortNeoMessage) Reset() {
	*x = PortNeoMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_and_port_scan_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortNeoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortNeoMessage) ProtoMessage() {}

func (x *PortNeoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_node_and_port_scan_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortNeoMessage.ProtoReflect.Descriptor instead.
func (*PortNeoMessage) Descriptor() ([]byte, []int) {
	return file_node_and_port_scan_proto_rawDescGZIP(), []int{4}
}

func (x *PortNeoMessage) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *PortNeoMessage) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *PortNeoMessage) GetPriority() int64 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *PortNeoMessage) GetPortNumber() int64 {
	if x != nil {
		return x.PortNumber
	}
	return 0
}

func (x *PortNeoMessage) GetTransportProtocol() string {
	if x != nil {
		return x.TransportProtocol
	}
	return ""
}

func (x *PortNeoMessage) GetPortScanID() int64 {
	if x != nil {
		return x.PortScanID
	}
	return 0
}

var File_node_and_port_scan_proto protoreflect.FileDescriptor

var file_node_and_port_scan_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x73, 0x63, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x63, 0x6f, 0x6d, 0x2e,
	0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e,
	0x6c, 0x69, 0x77, 0x61, 0x73, 0x63, 0x2e, 0x6e, 0x65, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x17, 0x4e, 0x6f, 0x64, 0x65, 0x53,
	0x63, 0x61, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4e, 0x65, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x4e, 0x6f, 0x64,
	0x65, 0x53, 0x63, 0x61, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x28, 0x0a, 0x0f,
	0x50, 0x6f, 0x72, 0x74, 0x53, 0x63, 0x61, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x63, 0x61, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x56, 0x0a, 0x12, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x63,
	0x61, 0x6e, 0x4e, 0x65, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x6f,
	0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x44, 0x6f, 0x6e, 0x65, 0x22, 0xb8,
	0x01, 0x0a, 0x0e, 0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x65, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x4d,
	0x41, 0x43, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x4d, 0x41, 0x43, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x49,
	0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x4e, 0x6f, 0x64,
	0x65, 0x53, 0x63, 0x61, 0x6e, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x4e,
	0x6f, 0x64, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x49, 0x44, 0x22, 0x6e, 0x0a, 0x12, 0x50, 0x6f, 0x72,
	0x74, 0x53, 0x63, 0x61, 0x6e, 0x4e, 0x65, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x44, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x44, 0x6f, 0x6e,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x22, 0xc8, 0x01, 0x0a, 0x0e, 0x50, 0x6f,
	0x72, 0x74, 0x4e, 0x65, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x50, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x6f, 0x72, 0x74, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x50, 0x6f, 0x72, 0x74,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x63, 0x61, 0x6e,
	0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x63,
	0x61, 0x6e, 0x49, 0x44, 0x32, 0xf1, 0x04, 0x0a, 0x19, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x6e, 0x64,
	0x50, 0x6f, 0x72, 0x74, 0x53, 0x63, 0x61, 0x6e, 0x4e, 0x65, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x7c, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x53,
	0x63, 0x61, 0x6e, 0x12, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e,
	0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x6c, 0x69, 0x77, 0x61, 0x73, 0x63,
	0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x4e, 0x65, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x32, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c,
	0x69, 0x78, 0x2e, 0x6c, 0x69, 0x77, 0x61, 0x73, 0x63, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x4e, 0x65, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x64, 0x0a, 0x14, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x4e,
	0x6f, 0x64, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72,
	0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x6c, 0x69, 0x77, 0x61, 0x73, 0x63, 0x2e, 0x6e, 0x65,
	0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x4e, 0x65, 0x6f, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x30, 0x01, 0x12, 0x78, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x54, 0x6f, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x32, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78,
	0x2e, 0x6c, 0x69, 0x77, 0x61, 0x73, 0x63, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x53, 0x63, 0x61, 0x6e, 0x4e, 0x65, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x2e,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66,
	0x65, 0x6c, 0x69, 0x78, 0x2e, 0x6c, 0x69, 0x77, 0x61, 0x73, 0x63, 0x2e, 0x6e, 0x65, 0x6f, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x65, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x30, 0x01,
	0x12, 0x7c, 0x0a, 0x14, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x50,
	0x6f, 0x72, 0x74, 0x53, 0x63, 0x61, 0x6e, 0x73, 0x12, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70,
	0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x6c,
	0x69, 0x77, 0x61, 0x73, 0x63, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x65,
	0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70,
	0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x6c,
	0x69, 0x77, 0x61, 0x73, 0x63, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x63,
	0x61, 0x6e, 0x4e, 0x65, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x30, 0x01, 0x12, 0x78,
	0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x50, 0x6f, 0x72,
	0x74, 0x73, 0x12, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6a, 0x74, 0x69, 0x6e, 0x67,
	0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x6c, 0x69, 0x77, 0x61, 0x73, 0x63, 0x2e,
	0x6e, 0x65, 0x6f, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x63, 0x61, 0x6e, 0x4e, 0x65, 0x6f, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x6f, 0x6a,
	0x74, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x6c, 0x69, 0x78, 0x2e, 0x6c, 0x69, 0x77,
	0x61, 0x73, 0x63, 0x2e, 0x6e, 0x65, 0x6f, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x4e, 0x65, 0x6f, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x30, 0x01, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f, 0x6a, 0x6e, 0x74, 0x66, 0x78, 0x2f, 0x6c,
	0x69, 0x77, 0x61, 0x73, 0x63, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_node_and_port_scan_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_node_and_port_scan_proto_goTypes = []interface{}{
	(*NodeScanStartNeoMessage)(nil), // 0: com.pojtinger.felix.liwasc.neo.NodeScanStartNeoMessage
	(*NodeScanNeoMessage)(nil),      // 1: com.pojtinger.felix.liwasc.neo.NodeScanNeoMessage
	(*NodeNeoMessage)(nil),          // 2: com.pojtinger.felix.liwasc.neo.NodeNeoMessage
	(*PortScanNeoMessage)(nil),      // 3: com.pojtinger.felix.liwasc.neo.PortScanNeoMessage
	(*PortNeoMessage)(nil),          // 4: com.pojtinger.felix.liwasc.neo.PortNeoMessage
	(*empty.Empty)(nil),             // 5: google.protobuf.Empty
}
var file_node_and_port_scan_proto_depIdxs = []int32{
	0, // 0: com.pojtinger.felix.liwasc.neo.NodeAndPortScanNeoService.StartNodeScan:input_type -> com.pojtinger.felix.liwasc.neo.NodeScanStartNeoMessage
	5, // 1: com.pojtinger.felix.liwasc.neo.NodeAndPortScanNeoService.SubscribeToNodeScans:input_type -> google.protobuf.Empty
	1, // 2: com.pojtinger.felix.liwasc.neo.NodeAndPortScanNeoService.SubscribeToNodes:input_type -> com.pojtinger.felix.liwasc.neo.NodeScanNeoMessage
	2, // 3: com.pojtinger.felix.liwasc.neo.NodeAndPortScanNeoService.SubscribeToPortScans:input_type -> com.pojtinger.felix.liwasc.neo.NodeNeoMessage
	3, // 4: com.pojtinger.felix.liwasc.neo.NodeAndPortScanNeoService.SubscribeToPorts:input_type -> com.pojtinger.felix.liwasc.neo.PortScanNeoMessage
	1, // 5: com.pojtinger.felix.liwasc.neo.NodeAndPortScanNeoService.StartNodeScan:output_type -> com.pojtinger.felix.liwasc.neo.NodeScanNeoMessage
	1, // 6: com.pojtinger.felix.liwasc.neo.NodeAndPortScanNeoService.SubscribeToNodeScans:output_type -> com.pojtinger.felix.liwasc.neo.NodeScanNeoMessage
	2, // 7: com.pojtinger.felix.liwasc.neo.NodeAndPortScanNeoService.SubscribeToNodes:output_type -> com.pojtinger.felix.liwasc.neo.NodeNeoMessage
	3, // 8: com.pojtinger.felix.liwasc.neo.NodeAndPortScanNeoService.SubscribeToPortScans:output_type -> com.pojtinger.felix.liwasc.neo.PortScanNeoMessage
	4, // 9: com.pojtinger.felix.liwasc.neo.NodeAndPortScanNeoService.SubscribeToPorts:output_type -> com.pojtinger.felix.liwasc.neo.PortNeoMessage
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
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
			switch v := v.(*NodeScanNeoMessage); i {
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
		file_node_and_port_scan_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeNeoMessage); i {
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
		file_node_and_port_scan_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortScanNeoMessage); i {
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
		file_node_and_port_scan_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortNeoMessage); i {
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
			NumMessages:   5,
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
	StartNodeScan(ctx context.Context, in *NodeScanStartNeoMessage, opts ...grpc.CallOption) (*NodeScanNeoMessage, error)
	SubscribeToNodeScans(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (NodeAndPortScanNeoService_SubscribeToNodeScansClient, error)
	SubscribeToNodes(ctx context.Context, in *NodeScanNeoMessage, opts ...grpc.CallOption) (NodeAndPortScanNeoService_SubscribeToNodesClient, error)
	SubscribeToPortScans(ctx context.Context, in *NodeNeoMessage, opts ...grpc.CallOption) (NodeAndPortScanNeoService_SubscribeToPortScansClient, error)
	SubscribeToPorts(ctx context.Context, in *PortScanNeoMessage, opts ...grpc.CallOption) (NodeAndPortScanNeoService_SubscribeToPortsClient, error)
}

type nodeAndPortScanNeoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeAndPortScanNeoServiceClient(cc grpc.ClientConnInterface) NodeAndPortScanNeoServiceClient {
	return &nodeAndPortScanNeoServiceClient{cc}
}

func (c *nodeAndPortScanNeoServiceClient) StartNodeScan(ctx context.Context, in *NodeScanStartNeoMessage, opts ...grpc.CallOption) (*NodeScanNeoMessage, error) {
	out := new(NodeScanNeoMessage)
	err := c.cc.Invoke(ctx, "/com.pojtinger.felix.liwasc.neo.NodeAndPortScanNeoService/StartNodeScan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeAndPortScanNeoServiceClient) SubscribeToNodeScans(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (NodeAndPortScanNeoService_SubscribeToNodeScansClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NodeAndPortScanNeoService_serviceDesc.Streams[0], "/com.pojtinger.felix.liwasc.neo.NodeAndPortScanNeoService/SubscribeToNodeScans", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeAndPortScanNeoServiceSubscribeToNodeScansClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeAndPortScanNeoService_SubscribeToNodeScansClient interface {
	Recv() (*NodeScanNeoMessage, error)
	grpc.ClientStream
}

type nodeAndPortScanNeoServiceSubscribeToNodeScansClient struct {
	grpc.ClientStream
}

func (x *nodeAndPortScanNeoServiceSubscribeToNodeScansClient) Recv() (*NodeScanNeoMessage, error) {
	m := new(NodeScanNeoMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeAndPortScanNeoServiceClient) SubscribeToNodes(ctx context.Context, in *NodeScanNeoMessage, opts ...grpc.CallOption) (NodeAndPortScanNeoService_SubscribeToNodesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NodeAndPortScanNeoService_serviceDesc.Streams[1], "/com.pojtinger.felix.liwasc.neo.NodeAndPortScanNeoService/SubscribeToNodes", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeAndPortScanNeoServiceSubscribeToNodesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeAndPortScanNeoService_SubscribeToNodesClient interface {
	Recv() (*NodeNeoMessage, error)
	grpc.ClientStream
}

type nodeAndPortScanNeoServiceSubscribeToNodesClient struct {
	grpc.ClientStream
}

func (x *nodeAndPortScanNeoServiceSubscribeToNodesClient) Recv() (*NodeNeoMessage, error) {
	m := new(NodeNeoMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeAndPortScanNeoServiceClient) SubscribeToPortScans(ctx context.Context, in *NodeNeoMessage, opts ...grpc.CallOption) (NodeAndPortScanNeoService_SubscribeToPortScansClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NodeAndPortScanNeoService_serviceDesc.Streams[2], "/com.pojtinger.felix.liwasc.neo.NodeAndPortScanNeoService/SubscribeToPortScans", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeAndPortScanNeoServiceSubscribeToPortScansClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeAndPortScanNeoService_SubscribeToPortScansClient interface {
	Recv() (*PortScanNeoMessage, error)
	grpc.ClientStream
}

type nodeAndPortScanNeoServiceSubscribeToPortScansClient struct {
	grpc.ClientStream
}

func (x *nodeAndPortScanNeoServiceSubscribeToPortScansClient) Recv() (*PortScanNeoMessage, error) {
	m := new(PortScanNeoMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeAndPortScanNeoServiceClient) SubscribeToPorts(ctx context.Context, in *PortScanNeoMessage, opts ...grpc.CallOption) (NodeAndPortScanNeoService_SubscribeToPortsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NodeAndPortScanNeoService_serviceDesc.Streams[3], "/com.pojtinger.felix.liwasc.neo.NodeAndPortScanNeoService/SubscribeToPorts", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeAndPortScanNeoServiceSubscribeToPortsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeAndPortScanNeoService_SubscribeToPortsClient interface {
	Recv() (*PortNeoMessage, error)
	grpc.ClientStream
}

type nodeAndPortScanNeoServiceSubscribeToPortsClient struct {
	grpc.ClientStream
}

func (x *nodeAndPortScanNeoServiceSubscribeToPortsClient) Recv() (*PortNeoMessage, error) {
	m := new(PortNeoMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NodeAndPortScanNeoServiceServer is the server API for NodeAndPortScanNeoService service.
type NodeAndPortScanNeoServiceServer interface {
	StartNodeScan(context.Context, *NodeScanStartNeoMessage) (*NodeScanNeoMessage, error)
	SubscribeToNodeScans(*empty.Empty, NodeAndPortScanNeoService_SubscribeToNodeScansServer) error
	SubscribeToNodes(*NodeScanNeoMessage, NodeAndPortScanNeoService_SubscribeToNodesServer) error
	SubscribeToPortScans(*NodeNeoMessage, NodeAndPortScanNeoService_SubscribeToPortScansServer) error
	SubscribeToPorts(*PortScanNeoMessage, NodeAndPortScanNeoService_SubscribeToPortsServer) error
}

// UnimplementedNodeAndPortScanNeoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNodeAndPortScanNeoServiceServer struct {
}

func (*UnimplementedNodeAndPortScanNeoServiceServer) StartNodeScan(context.Context, *NodeScanStartNeoMessage) (*NodeScanNeoMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartNodeScan not implemented")
}
func (*UnimplementedNodeAndPortScanNeoServiceServer) SubscribeToNodeScans(*empty.Empty, NodeAndPortScanNeoService_SubscribeToNodeScansServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToNodeScans not implemented")
}
func (*UnimplementedNodeAndPortScanNeoServiceServer) SubscribeToNodes(*NodeScanNeoMessage, NodeAndPortScanNeoService_SubscribeToNodesServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToNodes not implemented")
}
func (*UnimplementedNodeAndPortScanNeoServiceServer) SubscribeToPortScans(*NodeNeoMessage, NodeAndPortScanNeoService_SubscribeToPortScansServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToPortScans not implemented")
}
func (*UnimplementedNodeAndPortScanNeoServiceServer) SubscribeToPorts(*PortScanNeoMessage, NodeAndPortScanNeoService_SubscribeToPortsServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToPorts not implemented")
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

func _NodeAndPortScanNeoService_SubscribeToNodeScans_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeAndPortScanNeoServiceServer).SubscribeToNodeScans(m, &nodeAndPortScanNeoServiceSubscribeToNodeScansServer{stream})
}

type NodeAndPortScanNeoService_SubscribeToNodeScansServer interface {
	Send(*NodeScanNeoMessage) error
	grpc.ServerStream
}

type nodeAndPortScanNeoServiceSubscribeToNodeScansServer struct {
	grpc.ServerStream
}

func (x *nodeAndPortScanNeoServiceSubscribeToNodeScansServer) Send(m *NodeScanNeoMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeAndPortScanNeoService_SubscribeToNodes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NodeScanNeoMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeAndPortScanNeoServiceServer).SubscribeToNodes(m, &nodeAndPortScanNeoServiceSubscribeToNodesServer{stream})
}

type NodeAndPortScanNeoService_SubscribeToNodesServer interface {
	Send(*NodeNeoMessage) error
	grpc.ServerStream
}

type nodeAndPortScanNeoServiceSubscribeToNodesServer struct {
	grpc.ServerStream
}

func (x *nodeAndPortScanNeoServiceSubscribeToNodesServer) Send(m *NodeNeoMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeAndPortScanNeoService_SubscribeToPortScans_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NodeNeoMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeAndPortScanNeoServiceServer).SubscribeToPortScans(m, &nodeAndPortScanNeoServiceSubscribeToPortScansServer{stream})
}

type NodeAndPortScanNeoService_SubscribeToPortScansServer interface {
	Send(*PortScanNeoMessage) error
	grpc.ServerStream
}

type nodeAndPortScanNeoServiceSubscribeToPortScansServer struct {
	grpc.ServerStream
}

func (x *nodeAndPortScanNeoServiceSubscribeToPortScansServer) Send(m *PortScanNeoMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeAndPortScanNeoService_SubscribeToPorts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PortScanNeoMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeAndPortScanNeoServiceServer).SubscribeToPorts(m, &nodeAndPortScanNeoServiceSubscribeToPortsServer{stream})
}

type NodeAndPortScanNeoService_SubscribeToPortsServer interface {
	Send(*PortNeoMessage) error
	grpc.ServerStream
}

type nodeAndPortScanNeoServiceSubscribeToPortsServer struct {
	grpc.ServerStream
}

func (x *nodeAndPortScanNeoServiceSubscribeToPortsServer) Send(m *PortNeoMessage) error {
	return x.ServerStream.SendMsg(m)
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
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeToNodeScans",
			Handler:       _NodeAndPortScanNeoService_SubscribeToNodeScans_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeToNodes",
			Handler:       _NodeAndPortScanNeoService_SubscribeToNodes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeToPortScans",
			Handler:       _NodeAndPortScanNeoService_SubscribeToPortScans_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeToPorts",
			Handler:       _NodeAndPortScanNeoService_SubscribeToPorts_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "node_and_port_scan.proto",
}