// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rentalsvc.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type CheckAvailabilityResponse_Availability int32

const (
	CheckAvailabilityResponse_AVAILABLE   CheckAvailabilityResponse_Availability = 0
	CheckAvailabilityResponse_UNAVAILABLE CheckAvailabilityResponse_Availability = 1
)

var CheckAvailabilityResponse_Availability_name = map[int32]string{
	0: "AVAILABLE",
	1: "UNAVAILABLE",
}

var CheckAvailabilityResponse_Availability_value = map[string]int32{
	"AVAILABLE":   0,
	"UNAVAILABLE": 1,
}

func (x CheckAvailabilityResponse_Availability) String() string {
	return proto.EnumName(CheckAvailabilityResponse_Availability_name, int32(x))
}

func (CheckAvailabilityResponse_Availability) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_94881bd0765b2a96, []int{1, 0}
}

type CheckAvailabilityRequest struct {
	BookId               int32    `protobuf:"varint,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckAvailabilityRequest) Reset()         { *m = CheckAvailabilityRequest{} }
func (m *CheckAvailabilityRequest) String() string { return proto.CompactTextString(m) }
func (*CheckAvailabilityRequest) ProtoMessage()    {}
func (*CheckAvailabilityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_94881bd0765b2a96, []int{0}
}

func (m *CheckAvailabilityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckAvailabilityRequest.Unmarshal(m, b)
}
func (m *CheckAvailabilityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckAvailabilityRequest.Marshal(b, m, deterministic)
}
func (m *CheckAvailabilityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckAvailabilityRequest.Merge(m, src)
}
func (m *CheckAvailabilityRequest) XXX_Size() int {
	return xxx_messageInfo_CheckAvailabilityRequest.Size(m)
}
func (m *CheckAvailabilityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckAvailabilityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckAvailabilityRequest proto.InternalMessageInfo

func (m *CheckAvailabilityRequest) GetBookId() int32 {
	if m != nil {
		return m.BookId
	}
	return 0
}

type CheckAvailabilityResponse struct {
	Availability         CheckAvailabilityResponse_Availability `protobuf:"varint,1,opt,name=availability,proto3,enum=proto.CheckAvailabilityResponse_Availability" json:"availability,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *CheckAvailabilityResponse) Reset()         { *m = CheckAvailabilityResponse{} }
func (m *CheckAvailabilityResponse) String() string { return proto.CompactTextString(m) }
func (*CheckAvailabilityResponse) ProtoMessage()    {}
func (*CheckAvailabilityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_94881bd0765b2a96, []int{1}
}

func (m *CheckAvailabilityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckAvailabilityResponse.Unmarshal(m, b)
}
func (m *CheckAvailabilityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckAvailabilityResponse.Marshal(b, m, deterministic)
}
func (m *CheckAvailabilityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckAvailabilityResponse.Merge(m, src)
}
func (m *CheckAvailabilityResponse) XXX_Size() int {
	return xxx_messageInfo_CheckAvailabilityResponse.Size(m)
}
func (m *CheckAvailabilityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckAvailabilityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckAvailabilityResponse proto.InternalMessageInfo

func (m *CheckAvailabilityResponse) GetAvailability() CheckAvailabilityResponse_Availability {
	if m != nil {
		return m.Availability
	}
	return CheckAvailabilityResponse_AVAILABLE
}

type RentBookRequest struct {
	BookId               int32    `protobuf:"varint,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	UserId               int32    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RentBookRequest) Reset()         { *m = RentBookRequest{} }
func (m *RentBookRequest) String() string { return proto.CompactTextString(m) }
func (*RentBookRequest) ProtoMessage()    {}
func (*RentBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_94881bd0765b2a96, []int{2}
}

func (m *RentBookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RentBookRequest.Unmarshal(m, b)
}
func (m *RentBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RentBookRequest.Marshal(b, m, deterministic)
}
func (m *RentBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RentBookRequest.Merge(m, src)
}
func (m *RentBookRequest) XXX_Size() int {
	return xxx_messageInfo_RentBookRequest.Size(m)
}
func (m *RentBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RentBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RentBookRequest proto.InternalMessageInfo

func (m *RentBookRequest) GetBookId() int32 {
	if m != nil {
		return m.BookId
	}
	return 0
}

func (m *RentBookRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type RentBookResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RentBookResponse) Reset()         { *m = RentBookResponse{} }
func (m *RentBookResponse) String() string { return proto.CompactTextString(m) }
func (*RentBookResponse) ProtoMessage()    {}
func (*RentBookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_94881bd0765b2a96, []int{3}
}

func (m *RentBookResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RentBookResponse.Unmarshal(m, b)
}
func (m *RentBookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RentBookResponse.Marshal(b, m, deterministic)
}
func (m *RentBookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RentBookResponse.Merge(m, src)
}
func (m *RentBookResponse) XXX_Size() int {
	return xxx_messageInfo_RentBookResponse.Size(m)
}
func (m *RentBookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RentBookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RentBookResponse proto.InternalMessageInfo

func (m *RentBookResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterEnum("proto.CheckAvailabilityResponse_Availability", CheckAvailabilityResponse_Availability_name, CheckAvailabilityResponse_Availability_value)
	proto.RegisterType((*CheckAvailabilityRequest)(nil), "proto.CheckAvailabilityRequest")
	proto.RegisterType((*CheckAvailabilityResponse)(nil), "proto.CheckAvailabilityResponse")
	proto.RegisterType((*RentBookRequest)(nil), "proto.RentBookRequest")
	proto.RegisterType((*RentBookResponse)(nil), "proto.RentBookResponse")
}

func init() { proto.RegisterFile("rentalsvc.proto", fileDescriptor_94881bd0765b2a96) }

var fileDescriptor_94881bd0765b2a96 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x4a, 0xcd, 0x2b,
	0x49, 0xcc, 0x29, 0x2e, 0x4b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a,
	0xc6, 0x5c, 0x12, 0xce, 0x19, 0xa9, 0xc9, 0xd9, 0x8e, 0x65, 0x89, 0x99, 0x39, 0x89, 0x49, 0x99,
	0x39, 0x99, 0x25, 0x95, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0xe2, 0x5c, 0xec, 0x49,
	0xf9, 0xf9, 0xd9, 0xf1, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x6c, 0x20, 0xae,
	0x67, 0x8a, 0xd2, 0x3c, 0x46, 0x2e, 0x49, 0x2c, 0xba, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85,
	0x02, 0xb9, 0x78, 0x12, 0x91, 0xc4, 0xc1, 0x7a, 0xf9, 0x8c, 0x74, 0x21, 0xf6, 0xea, 0xe1, 0xd4,
	0xa7, 0x87, 0x22, 0x88, 0x62, 0x84, 0x92, 0x1e, 0x17, 0x0f, 0xb2, 0xac, 0x10, 0x2f, 0x17, 0xa7,
	0x63, 0x98, 0xa3, 0xa7, 0x8f, 0xa3, 0x93, 0x8f, 0xab, 0x00, 0x83, 0x10, 0x3f, 0x17, 0x77, 0xa8,
	0x1f, 0x42, 0x80, 0x51, 0xc9, 0x99, 0x8b, 0x3f, 0x28, 0x35, 0xaf, 0xc4, 0x29, 0x3f, 0x3f, 0x9b,
	0x90, 0x67, 0x40, 0x12, 0xa5, 0xc5, 0xa9, 0x45, 0x20, 0x09, 0x26, 0x88, 0x04, 0x88, 0xeb, 0x99,
	0xa2, 0xa4, 0xc5, 0x25, 0x80, 0x30, 0x04, 0xea, 0x37, 0x31, 0x2e, 0xb6, 0xe2, 0x92, 0xc4, 0x92,
	0xd2, 0x62, 0xb0, 0x21, 0x9c, 0x41, 0x50, 0x9e, 0xd1, 0x0a, 0x46, 0x2e, 0xde, 0x20, 0x70, 0x08,
	0x07, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0x45, 0x70, 0x09, 0x62, 0x78, 0x55, 0x48, 0x1e,
	0x77, 0x20, 0x80, 0x5d, 0x29, 0xa5, 0x40, 0x28, 0x94, 0x94, 0x18, 0x84, 0x6c, 0xb9, 0x38, 0x60,
	0xee, 0x12, 0x12, 0x83, 0xaa, 0x47, 0xf3, 0xad, 0x94, 0x38, 0x86, 0x38, 0x4c, 0x7b, 0x12, 0x1b,
	0x58, 0xc6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xe6, 0xe5, 0x40, 0x11, 0x12, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RentalServiceClient is the client API for RentalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RentalServiceClient interface {
	CheckAvailability(ctx context.Context, in *CheckAvailabilityRequest, opts ...grpc.CallOption) (*CheckAvailabilityResponse, error)
	RentBook(ctx context.Context, in *RentBookRequest, opts ...grpc.CallOption) (*RentBookResponse, error)
}

type rentalServiceClient struct {
	cc *grpc.ClientConn
}

func NewRentalServiceClient(cc *grpc.ClientConn) RentalServiceClient {
	return &rentalServiceClient{cc}
}

func (c *rentalServiceClient) CheckAvailability(ctx context.Context, in *CheckAvailabilityRequest, opts ...grpc.CallOption) (*CheckAvailabilityResponse, error) {
	out := new(CheckAvailabilityResponse)
	err := c.cc.Invoke(ctx, "/proto.RentalService/CheckAvailability", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rentalServiceClient) RentBook(ctx context.Context, in *RentBookRequest, opts ...grpc.CallOption) (*RentBookResponse, error) {
	out := new(RentBookResponse)
	err := c.cc.Invoke(ctx, "/proto.RentalService/RentBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RentalServiceServer is the server API for RentalService service.
type RentalServiceServer interface {
	CheckAvailability(context.Context, *CheckAvailabilityRequest) (*CheckAvailabilityResponse, error)
	RentBook(context.Context, *RentBookRequest) (*RentBookResponse, error)
}

func RegisterRentalServiceServer(s *grpc.Server, srv RentalServiceServer) {
	s.RegisterService(&_RentalService_serviceDesc, srv)
}

func _RentalService_CheckAvailability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAvailabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RentalServiceServer).CheckAvailability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RentalService/CheckAvailability",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RentalServiceServer).CheckAvailability(ctx, req.(*CheckAvailabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RentalService_RentBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RentBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RentalServiceServer).RentBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RentalService/RentBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RentalServiceServer).RentBook(ctx, req.(*RentBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RentalService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.RentalService",
	HandlerType: (*RentalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckAvailability",
			Handler:    _RentalService_CheckAvailability_Handler,
		},
		{
			MethodName: "RentBook",
			Handler:    _RentalService_RentBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rentalsvc.proto",
}