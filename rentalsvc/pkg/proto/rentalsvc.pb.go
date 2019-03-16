// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rentalsvc.proto

package stella_rental_v1

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

type ResourceStatus_Availability int32

const (
	ResourceStatus_AVAILABLE   ResourceStatus_Availability = 0
	ResourceStatus_UNAVAILABLE ResourceStatus_Availability = 1
)

var ResourceStatus_Availability_name = map[int32]string{
	0: "AVAILABLE",
	1: "UNAVAILABLE",
}

var ResourceStatus_Availability_value = map[string]int32{
	"AVAILABLE":   0,
	"UNAVAILABLE": 1,
}

func (x ResourceStatus_Availability) String() string {
	return proto.EnumName(ResourceStatus_Availability_name, int32(x))
}

func (ResourceStatus_Availability) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_94881bd0765b2a96, []int{0, 0}
}

type ResourceStatus struct {
	EntityId             int64                       `protobuf:"varint,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	Availability         ResourceStatus_Availability `protobuf:"varint,2,opt,name=availability,proto3,enum=stella.rental.v1.ResourceStatus_Availability" json:"availability,omitempty"`
	ReservedUserIds      []int64                     `protobuf:"varint,3,rep,packed,name=reserved_user_ids,json=reservedUserIds,proto3" json:"reserved_user_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ResourceStatus) Reset()         { *m = ResourceStatus{} }
func (m *ResourceStatus) String() string { return proto.CompactTextString(m) }
func (*ResourceStatus) ProtoMessage()    {}
func (*ResourceStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_94881bd0765b2a96, []int{0}
}

func (m *ResourceStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceStatus.Unmarshal(m, b)
}
func (m *ResourceStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceStatus.Marshal(b, m, deterministic)
}
func (m *ResourceStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceStatus.Merge(m, src)
}
func (m *ResourceStatus) XXX_Size() int {
	return xxx_messageInfo_ResourceStatus.Size(m)
}
func (m *ResourceStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceStatus proto.InternalMessageInfo

func (m *ResourceStatus) GetEntityId() int64 {
	if m != nil {
		return m.EntityId
	}
	return 0
}

func (m *ResourceStatus) GetAvailability() ResourceStatus_Availability {
	if m != nil {
		return m.Availability
	}
	return ResourceStatus_AVAILABLE
}

func (m *ResourceStatus) GetReservedUserIds() []int64 {
	if m != nil {
		return m.ReservedUserIds
	}
	return nil
}

type GetResourceStatusRequest struct {
	EntityId             int64    `protobuf:"varint,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResourceStatusRequest) Reset()         { *m = GetResourceStatusRequest{} }
func (m *GetResourceStatusRequest) String() string { return proto.CompactTextString(m) }
func (*GetResourceStatusRequest) ProtoMessage()    {}
func (*GetResourceStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_94881bd0765b2a96, []int{1}
}

func (m *GetResourceStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResourceStatusRequest.Unmarshal(m, b)
}
func (m *GetResourceStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResourceStatusRequest.Marshal(b, m, deterministic)
}
func (m *GetResourceStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResourceStatusRequest.Merge(m, src)
}
func (m *GetResourceStatusRequest) XXX_Size() int {
	return xxx_messageInfo_GetResourceStatusRequest.Size(m)
}
func (m *GetResourceStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResourceStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetResourceStatusRequest proto.InternalMessageInfo

func (m *GetResourceStatusRequest) GetEntityId() int64 {
	if m != nil {
		return m.EntityId
	}
	return 0
}

type RentResourceRequest struct {
	EntityId             int64    `protobuf:"varint,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RentResourceRequest) Reset()         { *m = RentResourceRequest{} }
func (m *RentResourceRequest) String() string { return proto.CompactTextString(m) }
func (*RentResourceRequest) ProtoMessage()    {}
func (*RentResourceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_94881bd0765b2a96, []int{2}
}

func (m *RentResourceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RentResourceRequest.Unmarshal(m, b)
}
func (m *RentResourceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RentResourceRequest.Marshal(b, m, deterministic)
}
func (m *RentResourceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RentResourceRequest.Merge(m, src)
}
func (m *RentResourceRequest) XXX_Size() int {
	return xxx_messageInfo_RentResourceRequest.Size(m)
}
func (m *RentResourceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RentResourceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RentResourceRequest proto.InternalMessageInfo

func (m *RentResourceRequest) GetEntityId() int64 {
	if m != nil {
		return m.EntityId
	}
	return 0
}

func (m *RentResourceRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type ReturnResourceRequest struct {
	EntityId             int64    `protobuf:"varint,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReturnResourceRequest) Reset()         { *m = ReturnResourceRequest{} }
func (m *ReturnResourceRequest) String() string { return proto.CompactTextString(m) }
func (*ReturnResourceRequest) ProtoMessage()    {}
func (*ReturnResourceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_94881bd0765b2a96, []int{3}
}

func (m *ReturnResourceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReturnResourceRequest.Unmarshal(m, b)
}
func (m *ReturnResourceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReturnResourceRequest.Marshal(b, m, deterministic)
}
func (m *ReturnResourceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReturnResourceRequest.Merge(m, src)
}
func (m *ReturnResourceRequest) XXX_Size() int {
	return xxx_messageInfo_ReturnResourceRequest.Size(m)
}
func (m *ReturnResourceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReturnResourceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReturnResourceRequest proto.InternalMessageInfo

func (m *ReturnResourceRequest) GetEntityId() int64 {
	if m != nil {
		return m.EntityId
	}
	return 0
}

func (m *ReturnResourceRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type ReserveResourceRequest struct {
	EntityId             int64    `protobuf:"varint,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReserveResourceRequest) Reset()         { *m = ReserveResourceRequest{} }
func (m *ReserveResourceRequest) String() string { return proto.CompactTextString(m) }
func (*ReserveResourceRequest) ProtoMessage()    {}
func (*ReserveResourceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_94881bd0765b2a96, []int{4}
}

func (m *ReserveResourceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReserveResourceRequest.Unmarshal(m, b)
}
func (m *ReserveResourceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReserveResourceRequest.Marshal(b, m, deterministic)
}
func (m *ReserveResourceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReserveResourceRequest.Merge(m, src)
}
func (m *ReserveResourceRequest) XXX_Size() int {
	return xxx_messageInfo_ReserveResourceRequest.Size(m)
}
func (m *ReserveResourceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReserveResourceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReserveResourceRequest proto.InternalMessageInfo

func (m *ReserveResourceRequest) GetEntityId() int64 {
	if m != nil {
		return m.EntityId
	}
	return 0
}

func (m *ReserveResourceRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type CancelResourceRequest struct {
	EntityId             int64    `protobuf:"varint,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CancelResourceRequest) Reset()         { *m = CancelResourceRequest{} }
func (m *CancelResourceRequest) String() string { return proto.CompactTextString(m) }
func (*CancelResourceRequest) ProtoMessage()    {}
func (*CancelResourceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_94881bd0765b2a96, []int{5}
}

func (m *CancelResourceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelResourceRequest.Unmarshal(m, b)
}
func (m *CancelResourceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelResourceRequest.Marshal(b, m, deterministic)
}
func (m *CancelResourceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelResourceRequest.Merge(m, src)
}
func (m *CancelResourceRequest) XXX_Size() int {
	return xxx_messageInfo_CancelResourceRequest.Size(m)
}
func (m *CancelResourceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelResourceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CancelResourceRequest proto.InternalMessageInfo

func (m *CancelResourceRequest) GetEntityId() int64 {
	if m != nil {
		return m.EntityId
	}
	return 0
}

func (m *CancelResourceRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func init() {
	proto.RegisterEnum("stella.rental.v1.ResourceStatus_Availability", ResourceStatus_Availability_name, ResourceStatus_Availability_value)
	proto.RegisterType((*ResourceStatus)(nil), "stella.rental.v1.ResourceStatus")
	proto.RegisterType((*GetResourceStatusRequest)(nil), "stella.rental.v1.GetResourceStatusRequest")
	proto.RegisterType((*RentResourceRequest)(nil), "stella.rental.v1.RentResourceRequest")
	proto.RegisterType((*ReturnResourceRequest)(nil), "stella.rental.v1.ReturnResourceRequest")
	proto.RegisterType((*ReserveResourceRequest)(nil), "stella.rental.v1.ReserveResourceRequest")
	proto.RegisterType((*CancelResourceRequest)(nil), "stella.rental.v1.CancelResourceRequest")
}

func init() { proto.RegisterFile("rentalsvc.proto", fileDescriptor_94881bd0765b2a96) }

var fileDescriptor_94881bd0765b2a96 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcf, 0x4a, 0xeb, 0x40,
	0x14, 0xc6, 0x9b, 0x1b, 0xe8, 0xbd, 0x3d, 0xb7, 0x6d, 0xda, 0xb9, 0x5c, 0x0d, 0xba, 0x09, 0x01,
	0x31, 0x14, 0x0c, 0x58, 0x17, 0xae, 0xa3, 0x88, 0x14, 0x6b, 0xc1, 0x29, 0x15, 0x5c, 0x48, 0x98,
	0x26, 0x67, 0x31, 0x10, 0x52, 0x9d, 0x99, 0x04, 0xfa, 0xae, 0x3e, 0x83, 0xcf, 0x20, 0xa6, 0xfe,
	0x49, 0xda, 0x51, 0xb3, 0xe8, 0xf6, 0x84, 0xef, 0xc7, 0x97, 0xf3, 0x3b, 0x09, 0x58, 0x02, 0x53,
	0xc5, 0x12, 0x99, 0x47, 0xfe, 0x83, 0x58, 0xa8, 0x05, 0xe9, 0x49, 0x85, 0x49, 0xc2, 0xfc, 0xd5,
	0xdc, 0xcf, 0x8f, 0xdd, 0x27, 0x03, 0xba, 0x14, 0xe5, 0x22, 0x13, 0x11, 0x4e, 0x15, 0x53, 0x99,
	0x24, 0xfb, 0xd0, 0xc2, 0x54, 0x71, 0xb5, 0x0c, 0x79, 0x6c, 0x1b, 0x8e, 0xe1, 0x99, 0xf4, 0xcf,
	0x6a, 0x30, 0x8a, 0xc9, 0x0d, 0xb4, 0x59, 0xce, 0x78, 0xc2, 0xe6, 0x3c, 0xe1, 0x6a, 0x69, 0xff,
	0x72, 0x0c, 0xaf, 0x3b, 0x3c, 0xf2, 0xd7, 0xc1, 0x7e, 0x15, 0xea, 0x07, 0xa5, 0x10, 0xad, 0x20,
	0xc8, 0x00, 0xfa, 0x02, 0x25, 0x8a, 0x1c, 0xe3, 0x30, 0x93, 0x28, 0x42, 0x1e, 0x4b, 0xdb, 0x74,
	0x4c, 0xcf, 0xa4, 0xd6, 0xfb, 0x83, 0x99, 0x44, 0x31, 0x8a, 0xa5, 0xeb, 0x43, 0xbb, 0x4c, 0x22,
	0x1d, 0x68, 0x05, 0xb7, 0xc1, 0x68, 0x1c, 0x9c, 0x8d, 0x2f, 0x7a, 0x0d, 0x62, 0xc1, 0xdf, 0xd9,
	0xe4, 0x73, 0x60, 0xb8, 0xa7, 0x60, 0x5f, 0xa2, 0xaa, 0x76, 0xa1, 0xf8, 0x98, 0xa1, 0x54, 0xdf,
	0xbe, 0xa7, 0x7b, 0x05, 0xff, 0x28, 0xa6, 0x1f, 0xc9, 0x3a, 0x19, 0xb2, 0x0b, 0xbf, 0xdf, 0xfa,
	0x17, 0x6b, 0x31, 0x69, 0x33, 0x2b, 0x6a, 0xbb, 0xd7, 0xf0, 0x9f, 0xa2, 0xca, 0x44, 0xba, 0x1d,
	0xdc, 0x04, 0x76, 0xe8, 0x6a, 0x2f, 0x5b, 0xab, 0x77, 0xce, 0xd2, 0x08, 0x93, 0xad, 0xe0, 0x86,
	0xcf, 0x26, 0x74, 0x68, 0x71, 0x07, 0x53, 0x14, 0x39, 0x8f, 0x90, 0x44, 0xd0, 0xdf, 0xb0, 0x40,
	0x06, 0x9b, 0x37, 0xf3, 0x95, 0xaa, 0x3d, 0xe7, 0xa7, 0xfb, 0x72, 0x1b, 0xe4, 0x0e, 0xda, 0x65,
	0x63, 0xe4, 0x40, 0x97, 0xd9, 0x30, 0x5a, 0x0b, 0x7d, 0xff, 0xfa, 0x8d, 0x94, 0xfd, 0x91, 0x43,
	0x5d, 0x4a, 0x63, 0xb8, 0x16, 0x3e, 0x04, 0x6b, 0xcd, 0x27, 0xf1, 0xb4, 0x31, 0x8d, 0xf2, 0xba,
	0xfd, 0xab, 0x82, 0x75, 0xfd, 0xb5, 0x27, 0x50, 0x07, 0x3f, 0x6f, 0x16, 0x3f, 0x97, 0x93, 0x97,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x57, 0xd9, 0x6d, 0xca, 0x6f, 0x04, 0x00, 0x00,
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
	GetResourceStatus(ctx context.Context, in *GetResourceStatusRequest, opts ...grpc.CallOption) (*ResourceStatus, error)
	RentResource(ctx context.Context, in *RentResourceRequest, opts ...grpc.CallOption) (*ResourceStatus, error)
	ReturnResource(ctx context.Context, in *ReturnResourceRequest, opts ...grpc.CallOption) (*ResourceStatus, error)
	ReserveResource(ctx context.Context, in *ReserveResourceRequest, opts ...grpc.CallOption) (*ResourceStatus, error)
	CancelResource(ctx context.Context, in *CancelResourceRequest, opts ...grpc.CallOption) (*ResourceStatus, error)
}

type rentalServiceClient struct {
	cc *grpc.ClientConn
}

func NewRentalServiceClient(cc *grpc.ClientConn) RentalServiceClient {
	return &rentalServiceClient{cc}
}

func (c *rentalServiceClient) GetResourceStatus(ctx context.Context, in *GetResourceStatusRequest, opts ...grpc.CallOption) (*ResourceStatus, error) {
	out := new(ResourceStatus)
	err := c.cc.Invoke(ctx, "/stella.rental.v1.RentalService/GetResourceStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rentalServiceClient) RentResource(ctx context.Context, in *RentResourceRequest, opts ...grpc.CallOption) (*ResourceStatus, error) {
	out := new(ResourceStatus)
	err := c.cc.Invoke(ctx, "/stella.rental.v1.RentalService/RentResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rentalServiceClient) ReturnResource(ctx context.Context, in *ReturnResourceRequest, opts ...grpc.CallOption) (*ResourceStatus, error) {
	out := new(ResourceStatus)
	err := c.cc.Invoke(ctx, "/stella.rental.v1.RentalService/ReturnResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rentalServiceClient) ReserveResource(ctx context.Context, in *ReserveResourceRequest, opts ...grpc.CallOption) (*ResourceStatus, error) {
	out := new(ResourceStatus)
	err := c.cc.Invoke(ctx, "/stella.rental.v1.RentalService/ReserveResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rentalServiceClient) CancelResource(ctx context.Context, in *CancelResourceRequest, opts ...grpc.CallOption) (*ResourceStatus, error) {
	out := new(ResourceStatus)
	err := c.cc.Invoke(ctx, "/stella.rental.v1.RentalService/CancelResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RentalServiceServer is the server API for RentalService service.
type RentalServiceServer interface {
	GetResourceStatus(context.Context, *GetResourceStatusRequest) (*ResourceStatus, error)
	RentResource(context.Context, *RentResourceRequest) (*ResourceStatus, error)
	ReturnResource(context.Context, *ReturnResourceRequest) (*ResourceStatus, error)
	ReserveResource(context.Context, *ReserveResourceRequest) (*ResourceStatus, error)
	CancelResource(context.Context, *CancelResourceRequest) (*ResourceStatus, error)
}

func RegisterRentalServiceServer(s *grpc.Server, srv RentalServiceServer) {
	s.RegisterService(&_RentalService_serviceDesc, srv)
}

func _RentalService_GetResourceStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResourceStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RentalServiceServer).GetResourceStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stella.rental.v1.RentalService/GetResourceStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RentalServiceServer).GetResourceStatus(ctx, req.(*GetResourceStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RentalService_RentResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RentResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RentalServiceServer).RentResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stella.rental.v1.RentalService/RentResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RentalServiceServer).RentResource(ctx, req.(*RentResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RentalService_ReturnResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReturnResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RentalServiceServer).ReturnResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stella.rental.v1.RentalService/ReturnResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RentalServiceServer).ReturnResource(ctx, req.(*ReturnResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RentalService_ReserveResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReserveResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RentalServiceServer).ReserveResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stella.rental.v1.RentalService/ReserveResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RentalServiceServer).ReserveResource(ctx, req.(*ReserveResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RentalService_CancelResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RentalServiceServer).CancelResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stella.rental.v1.RentalService/CancelResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RentalServiceServer).CancelResource(ctx, req.(*CancelResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RentalService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stella.rental.v1.RentalService",
	HandlerType: (*RentalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetResourceStatus",
			Handler:    _RentalService_GetResourceStatus_Handler,
		},
		{
			MethodName: "RentResource",
			Handler:    _RentalService_RentResource_Handler,
		},
		{
			MethodName: "ReturnResource",
			Handler:    _RentalService_ReturnResource_Handler,
		},
		{
			MethodName: "ReserveResource",
			Handler:    _RentalService_ReserveResource_Handler,
		},
		{
			MethodName: "CancelResource",
			Handler:    _RentalService_CancelResource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rentalsvc.proto",
}
