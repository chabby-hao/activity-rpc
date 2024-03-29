// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/chabby-hao/acitivity-rpc.proto

package activity

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ListAllRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAllRequest) Reset()         { *m = ListAllRequest{} }
func (m *ListAllRequest) String() string { return proto.CompactTextString(m) }
func (*ListAllRequest) ProtoMessage()    {}
func (*ListAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a684c9a0549e7832, []int{0}
}

func (m *ListAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAllRequest.Unmarshal(m, b)
}
func (m *ListAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAllRequest.Marshal(b, m, deterministic)
}
func (m *ListAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAllRequest.Merge(m, src)
}
func (m *ListAllRequest) XXX_Size() int {
	return xxx_messageInfo_ListAllRequest.Size(m)
}
func (m *ListAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListAllRequest proto.InternalMessageInfo

type ListAllResponse struct {
	List                 []*ActivityInfo `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListAllResponse) Reset()         { *m = ListAllResponse{} }
func (m *ListAllResponse) String() string { return proto.CompactTextString(m) }
func (*ListAllResponse) ProtoMessage()    {}
func (*ListAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a684c9a0549e7832, []int{1}
}

func (m *ListAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAllResponse.Unmarshal(m, b)
}
func (m *ListAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAllResponse.Marshal(b, m, deterministic)
}
func (m *ListAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAllResponse.Merge(m, src)
}
func (m *ListAllResponse) XXX_Size() int {
	return xxx_messageInfo_ListAllResponse.Size(m)
}
func (m *ListAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAllResponse proto.InternalMessageInfo

func (m *ListAllResponse) GetList() []*ActivityInfo {
	if m != nil {
		return m.List
	}
	return nil
}

type ActivityInfo struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 int64    `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Level                int64    `protobuf:"varint,3,opt,name=level,proto3" json:"level,omitempty"`
	Status               int64    `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	StartAt              int64    `protobuf:"varint,5,opt,name=startAt,proto3" json:"startAt,omitempty"`
	EndAt                int64    `protobuf:"varint,6,opt,name=endAt,proto3" json:"endAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActivityInfo) Reset()         { *m = ActivityInfo{} }
func (m *ActivityInfo) String() string { return proto.CompactTextString(m) }
func (*ActivityInfo) ProtoMessage()    {}
func (*ActivityInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a684c9a0549e7832, []int{2}
}

func (m *ActivityInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActivityInfo.Unmarshal(m, b)
}
func (m *ActivityInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActivityInfo.Marshal(b, m, deterministic)
}
func (m *ActivityInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivityInfo.Merge(m, src)
}
func (m *ActivityInfo) XXX_Size() int {
	return xxx_messageInfo_ActivityInfo.Size(m)
}
func (m *ActivityInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivityInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ActivityInfo proto.InternalMessageInfo

func (m *ActivityInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ActivityInfo) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ActivityInfo) GetLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *ActivityInfo) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ActivityInfo) GetStartAt() int64 {
	if m != nil {
		return m.StartAt
	}
	return 0
}

func (m *ActivityInfo) GetEndAt() int64 {
	if m != nil {
		return m.EndAt
	}
	return 0
}

type ListByTypeRequest struct {
	Type                 int64    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListByTypeRequest) Reset()         { *m = ListByTypeRequest{} }
func (m *ListByTypeRequest) String() string { return proto.CompactTextString(m) }
func (*ListByTypeRequest) ProtoMessage()    {}
func (*ListByTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a684c9a0549e7832, []int{3}
}

func (m *ListByTypeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListByTypeRequest.Unmarshal(m, b)
}
func (m *ListByTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListByTypeRequest.Marshal(b, m, deterministic)
}
func (m *ListByTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListByTypeRequest.Merge(m, src)
}
func (m *ListByTypeRequest) XXX_Size() int {
	return xxx_messageInfo_ListByTypeRequest.Size(m)
}
func (m *ListByTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListByTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListByTypeRequest proto.InternalMessageInfo

func (m *ListByTypeRequest) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

type ListByTypeResponse struct {
	List                 []*ActivityInfo `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListByTypeResponse) Reset()         { *m = ListByTypeResponse{} }
func (m *ListByTypeResponse) String() string { return proto.CompactTextString(m) }
func (*ListByTypeResponse) ProtoMessage()    {}
func (*ListByTypeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a684c9a0549e7832, []int{4}
}

func (m *ListByTypeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListByTypeResponse.Unmarshal(m, b)
}
func (m *ListByTypeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListByTypeResponse.Marshal(b, m, deterministic)
}
func (m *ListByTypeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListByTypeResponse.Merge(m, src)
}
func (m *ListByTypeResponse) XXX_Size() int {
	return xxx_messageInfo_ListByTypeResponse.Size(m)
}
func (m *ListByTypeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListByTypeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListByTypeResponse proto.InternalMessageInfo

func (m *ListByTypeResponse) GetList() []*ActivityInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*ListAllRequest)(nil), "github.com/chabby-hao/acitivity-rpc.ListAllRequest")
	proto.RegisterType((*ListAllResponse)(nil), "github.com/chabby-hao/acitivity-rpc.ListAllResponse")
	proto.RegisterType((*ActivityInfo)(nil), "github.com/chabby-hao/acitivity-rpc.ActivityInfo")
	proto.RegisterType((*ListByTypeRequest)(nil), "github.com/chabby-hao/acitivity-rpc.ListByTypeRequest")
	proto.RegisterType((*ListByTypeResponse)(nil), "github.com/chabby-hao/acitivity-rpc.ListByTypeResponse")
}

func init() { proto.RegisterFile("github.com/chabby-hao/acitivity-rpc.proto", fileDescriptor_a684c9a0549e7832) }

var fileDescriptor_a684c9a0549e7832 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xd9, 0x26, 0x4d, 0xcb, 0x28, 0x51, 0x07, 0x29, 0x63, 0xf5, 0x50, 0xf6, 0x62, 0xf1,
	0xd0, 0x43, 0x3d, 0x2b, 0xc6, 0x8b, 0x08, 0x9e, 0x82, 0x2f, 0x10, 0xcd, 0x0a, 0x0b, 0x4b, 0x12,
	0xbb, 0xd3, 0x42, 0xde, 0x42, 0xf0, 0x85, 0xa5, 0x9b, 0x4d, 0xd2, 0x40, 0x2f, 0xde, 0xe6, 0xff,
	0x66, 0xf2, 0x67, 0xfe, 0x59, 0x88, 0xb3, 0x4f, 0xd6, 0x3b, 0xcd, 0xf5, 0xaa, 0xda, 0x94, 0x5c,
	0xe2, 0xb4, 0xd5, 0xf2, 0x1c, 0xe2, 0x37, 0x6d, 0x39, 0x31, 0x26, 0x55, 0xdf, 0x5b, 0x65, 0x59,
	0x3e, 0xc0, 0x59, 0x47, 0x6c, 0x55, 0x16, 0x56, 0xe1, 0x1d, 0x84, 0x46, 0x5b, 0x26, 0xb1, 0x08,
	0x96, 0x27, 0xeb, 0xd9, 0xaa, 0x73, 0x4b, 0x7c, 0xf1, 0x5a, 0x7c, 0x95, 0xa9, 0x9b, 0x91, 0x3f,
	0x02, 0x4e, 0x0f, 0x31, 0xc6, 0x30, 0xd2, 0x39, 0x89, 0x85, 0x58, 0x06, 0xe9, 0x48, 0xe7, 0x88,
	0x10, 0x72, 0x5d, 0x29, 0x1a, 0x39, 0xe2, 0x6a, 0xbc, 0x84, 0xb1, 0x51, 0x3b, 0x65, 0x28, 0x70,
	0xb0, 0x11, 0x38, 0x83, 0xc8, 0x72, 0xc6, 0x5b, 0x4b, 0xa1, 0xc3, 0x5e, 0x21, 0xc1, 0xc4, 0x72,
	0xb6, 0xe1, 0x84, 0x69, 0xec, 0x1a, 0xad, 0xdc, 0xfb, 0xa8, 0x22, 0x4f, 0x98, 0xa2, 0xc6, 0xc7,
	0x09, 0x79, 0x0b, 0x17, 0xfb, 0x44, 0xcf, 0xf5, 0x7b, 0x5d, 0x29, 0x1f, 0xb3, 0x5b, 0x43, 0xf4,
	0x6b, 0xc8, 0x27, 0xc0, 0xc3, 0xc1, 0xff, 0xa7, 0x5f, 0xff, 0x0a, 0x98, 0xb6, 0x18, 0x1f, 0x61,
	0xe2, 0x2f, 0x89, 0xd4, 0x7f, 0x35, 0x3c, 0xf7, 0xfc, 0xea, 0x48, 0xc7, 0xff, 0xf8, 0x05, 0xa0,
	0x5f, 0x07, 0xaf, 0x87, 0x83, 0x83, 0x34, 0xf3, 0x9b, 0xe3, 0xcd, 0xc6, 0xe8, 0x23, 0x72, 0xaf,
	0x7e, 0xff, 0x17, 0x00, 0x00, 0xff, 0xff, 0x98, 0x5e, 0xdb, 0x3b, 0x07, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ActivityClient is the client API for Activity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ActivityClient interface {
	ListAll(ctx context.Context, in *ListAllRequest, opts ...grpc.CallOption) (*ListAllResponse, error)
	ListByType(ctx context.Context, in *ListByTypeRequest, opts ...grpc.CallOption) (*ListByTypeResponse, error)
}

type activityClient struct {
	cc *grpc.ClientConn
}

func NewActivityClient(cc *grpc.ClientConn) ActivityClient {
	return &activityClient{cc}
}

func (c *activityClient) ListAll(ctx context.Context, in *ListAllRequest, opts ...grpc.CallOption) (*ListAllResponse, error) {
	out := new(ListAllResponse)
	err := c.cc.Invoke(ctx, "/github.com/chabby-hao/acitivity-rpc.Activity/ListAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) ListByType(ctx context.Context, in *ListByTypeRequest, opts ...grpc.CallOption) (*ListByTypeResponse, error) {
	out := new(ListByTypeResponse)
	err := c.cc.Invoke(ctx, "/github.com/chabby-hao/acitivity-rpc.Activity/ListByType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActivityServer is the server API for Activity service.
type ActivityServer interface {
	ListAll(context.Context, *ListAllRequest) (*ListAllResponse, error)
	ListByType(context.Context, *ListByTypeRequest) (*ListByTypeResponse, error)
}

// UnimplementedActivityServer can be embedded to have forward compatible implementations.
type UnimplementedActivityServer struct {
}

func (*UnimplementedActivityServer) ListAll(ctx context.Context, req *ListAllRequest) (*ListAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAll not implemented")
}
func (*UnimplementedActivityServer) ListByType(ctx context.Context, req *ListByTypeRequest) (*ListByTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListByType not implemented")
}

func RegisterActivityServer(s *grpc.Server, srv ActivityServer) {
	s.RegisterService(&_Activity_serviceDesc, srv)
}

func _Activity_ListAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).ListAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com/chabby-hao/acitivity-rpc.Activity/ListAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).ListAll(ctx, req.(*ListAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_ListByType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListByTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).ListByType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com/chabby-hao/acitivity-rpc.Activity/ListByType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).ListByType(ctx, req.(*ListByTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Activity_serviceDesc = grpc.ServiceDesc{
	ServiceName: "github.com/chabby-hao/acitivity-rpc.Activity",
	HandlerType: (*ActivityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAll",
			Handler:    _Activity_ListAll_Handler,
		},
		{
			MethodName: "ListByType",
			Handler:    _Activity_ListByType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/chabby-hao/acitivity-rpc.proto",
}
