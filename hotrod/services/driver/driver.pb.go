package driver

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type DriverLocationRequest struct {
	Location             string   `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DriverLocationRequest) Reset()         { *m = DriverLocationRequest{} }
func (m *DriverLocationRequest) String() string { return proto.CompactTextString(m) }
func (*DriverLocationRequest) ProtoMessage()    {}
func (*DriverLocationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdcd28b7ebdcd54f, []int{0}
}
func (m *DriverLocationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DriverLocationRequest.Unmarshal(m, b)
}
func (m *DriverLocationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DriverLocationRequest.Marshal(b, m, deterministic)
}
func (m *DriverLocationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DriverLocationRequest.Merge(m, src)
}
func (m *DriverLocationRequest) XXX_Size() int {
	return xxx_messageInfo_DriverLocationRequest.Size(m)
}
func (m *DriverLocationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DriverLocationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DriverLocationRequest proto.InternalMessageInfo

func (m *DriverLocationRequest) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

type DriverLocation struct {
	DriverID             string   `protobuf:"bytes,1,opt,name=driverID,proto3" json:"driverID,omitempty"`
	Location             string   `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DriverLocation) Reset()         { *m = DriverLocation{} }
func (m *DriverLocation) String() string { return proto.CompactTextString(m) }
func (*DriverLocation) ProtoMessage()    {}
func (*DriverLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdcd28b7ebdcd54f, []int{1}
}
func (m *DriverLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DriverLocation.Unmarshal(m, b)
}
func (m *DriverLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DriverLocation.Marshal(b, m, deterministic)
}
func (m *DriverLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DriverLocation.Merge(m, src)
}
func (m *DriverLocation) XXX_Size() int {
	return xxx_messageInfo_DriverLocation.Size(m)
}
func (m *DriverLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_DriverLocation.DiscardUnknown(m)
}

var xxx_messageInfo_DriverLocation proto.InternalMessageInfo

func (m *DriverLocation) GetDriverID() string {
	if m != nil {
		return m.DriverID
	}
	return ""
}

func (m *DriverLocation) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

type DriverLocationResponse struct {
	Locations            []*DriverLocation `protobuf:"bytes,1,rep,name=locations,proto3" json:"locations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DriverLocationResponse) Reset()         { *m = DriverLocationResponse{} }
func (m *DriverLocationResponse) String() string { return proto.CompactTextString(m) }
func (*DriverLocationResponse) ProtoMessage()    {}
func (*DriverLocationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdcd28b7ebdcd54f, []int{2}
}
func (m *DriverLocationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DriverLocationResponse.Unmarshal(m, b)
}
func (m *DriverLocationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DriverLocationResponse.Marshal(b, m, deterministic)
}
func (m *DriverLocationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DriverLocationResponse.Merge(m, src)
}
func (m *DriverLocationResponse) XXX_Size() int {
	return xxx_messageInfo_DriverLocationResponse.Size(m)
}
func (m *DriverLocationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DriverLocationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DriverLocationResponse proto.InternalMessageInfo

func (m *DriverLocationResponse) GetLocations() []*DriverLocation {
	if m != nil {
		return m.Locations
	}
	return nil
}

func init() {
	proto.RegisterType((*DriverLocationRequest)(nil), "driver.DriverLocationRequest")
	proto.RegisterType((*DriverLocation)(nil), "driver.DriverLocation")
	proto.RegisterType((*DriverLocationResponse)(nil), "driver.DriverLocationResponse")
}

func init() {
	proto.RegisterFile("examples/hotrod/services/driver/driver.proto", fileDescriptor_cdcd28b7ebdcd54f)
}

var fileDescriptor_cdcd28b7ebdcd54f = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x49, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0xcf, 0xc8, 0x2f, 0x29, 0xca, 0x4f, 0xd1, 0x2f, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0x4f, 0x29, 0xca, 0x2c, 0x4b, 0x2d, 0x82, 0x52, 0x7a, 0x05, 0x45,
	0xf9, 0x25, 0xf9, 0x42, 0x6c, 0x10, 0x9e, 0x92, 0x31, 0x97, 0xa8, 0x0b, 0x98, 0xe5, 0x93, 0x9f,
	0x9c, 0x58, 0x92, 0x99, 0x9f, 0x17, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x24, 0xc5, 0xc5,
	0x91, 0x03, 0x15, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf3, 0x95, 0x3c, 0xb8, 0xf8,
	0x50, 0x35, 0x81, 0x54, 0x43, 0x0c, 0xf4, 0x74, 0x81, 0xa9, 0x86, 0xf1, 0x51, 0x4c, 0x62, 0x42,
	0x33, 0xc9, 0x8f, 0x4b, 0x0c, 0xdd, 0xfa, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x13, 0x2e,
	0x4e, 0x98, 0xaa, 0x62, 0x09, 0x46, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x31, 0x3d, 0xa8, 0x17, 0xd0,
	0xb4, 0x20, 0x14, 0x1a, 0xc5, 0x72, 0xf1, 0x42, 0x24, 0x83, 0x21, 0x9e, 0x17, 0xf2, 0xe1, 0xe2,
	0x76, 0xcb, 0xcc, 0x4b, 0xf1, 0x4b, 0x4d, 0x2c, 0x02, 0xf9, 0x4a, 0x16, 0x87, 0x11, 0x10, 0x4f,
	0x4b, 0xc9, 0xe1, 0x92, 0x86, 0x38, 0xca, 0x89, 0x23, 0x0a, 0x1a, 0x6e, 0x49, 0x6c, 0xe0, 0x60,
	0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x90, 0x0b, 0x66, 0x76, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DriverServiceClient is the client API for DriverService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DriverServiceClient interface {
	FindNearest(ctx context.Context, in *DriverLocationRequest, opts ...grpc.CallOption) (*DriverLocationResponse, error)
}

type driverServiceClient struct {
	cc *grpc.ClientConn
}

func NewDriverServiceClient(cc *grpc.ClientConn) DriverServiceClient {
	return &driverServiceClient{cc}
}

func (c *driverServiceClient) FindNearest(ctx context.Context, in *DriverLocationRequest, opts ...grpc.CallOption) (*DriverLocationResponse, error) {
	out := new(DriverLocationResponse)
	err := c.cc.Invoke(ctx, "/driver.DriverService/FindNearest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DriverServiceServer is the server API for DriverService service.
type DriverServiceServer interface {
	FindNearest(context.Context, *DriverLocationRequest) (*DriverLocationResponse, error)
}

// UnimplementedDriverServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDriverServiceServer struct {
}

func (*UnimplementedDriverServiceServer) FindNearest(ctx context.Context, req *DriverLocationRequest) (*DriverLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindNearest not implemented")
}

func RegisterDriverServiceServer(s *grpc.Server, srv DriverServiceServer) {
	s.RegisterService(&_DriverService_serviceDesc, srv)
}

func _DriverService_FindNearest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DriverLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServiceServer).FindNearest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/driver.DriverService/FindNearest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServiceServer).FindNearest(ctx, req.(*DriverLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DriverService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "driver.DriverService",
	HandlerType: (*DriverServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindNearest",
			Handler:    _DriverService_FindNearest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "examples/hotrod/services/driver/driver.proto",
}
