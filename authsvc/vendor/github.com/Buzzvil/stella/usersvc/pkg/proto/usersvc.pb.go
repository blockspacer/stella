// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usersvc.proto

package stella_user_v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type User struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	SlackUserId          string   `protobuf:"bytes,3,opt,name=slack_user_id,json=slackUserId,proto3" json:"slack_user_id,omitempty"`
	SlackTeamId          string   `protobuf:"bytes,4,opt,name=slack_team_id,json=slackTeamId,proto3" json:"slack_team_id,omitempty"`
	Image                string   `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_c11b3307f9b338e4, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetSlackUserId() string {
	if m != nil {
		return m.SlackUserId
	}
	return ""
}

func (m *User) GetSlackTeamId() string {
	if m != nil {
		return m.SlackTeamId
	}
	return ""
}

func (m *User) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

type ListUsersRequest struct {
	Ids                  []int64  `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersRequest) Reset()         { *m = ListUsersRequest{} }
func (m *ListUsersRequest) String() string { return proto.CompactTextString(m) }
func (*ListUsersRequest) ProtoMessage()    {}
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c11b3307f9b338e4, []int{1}
}

func (m *ListUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersRequest.Unmarshal(m, b)
}
func (m *ListUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersRequest.Marshal(b, m, deterministic)
}
func (m *ListUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersRequest.Merge(m, src)
}
func (m *ListUsersRequest) XXX_Size() int {
	return xxx_messageInfo_ListUsersRequest.Size(m)
}
func (m *ListUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersRequest proto.InternalMessageInfo

func (m *ListUsersRequest) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

type ListUsersResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersResponse) Reset()         { *m = ListUsersResponse{} }
func (m *ListUsersResponse) String() string { return proto.CompactTextString(m) }
func (*ListUsersResponse) ProtoMessage()    {}
func (*ListUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c11b3307f9b338e4, []int{2}
}

func (m *ListUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersResponse.Unmarshal(m, b)
}
func (m *ListUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersResponse.Marshal(b, m, deterministic)
}
func (m *ListUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersResponse.Merge(m, src)
}
func (m *ListUsersResponse) XXX_Size() int {
	return xxx_messageInfo_ListUsersResponse.Size(m)
}
func (m *ListUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersResponse proto.InternalMessageInfo

func (m *ListUsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type GetUserRequest struct {
	// Types that are valid to be assigned to Identifier:
	//	*GetUserRequest_Id
	//	*GetUserRequest_SlackUserId
	Identifier           isGetUserRequest_Identifier `protobuf_oneof:"identifier"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c11b3307f9b338e4, []int{3}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

type isGetUserRequest_Identifier interface {
	isGetUserRequest_Identifier()
}

type GetUserRequest_Id struct {
	Id int64 `protobuf:"varint,1,opt,name=id,proto3,oneof"`
}

type GetUserRequest_SlackUserId struct {
	SlackUserId string `protobuf:"bytes,2,opt,name=slack_user_id,json=slackUserId,proto3,oneof"`
}

func (*GetUserRequest_Id) isGetUserRequest_Identifier() {}

func (*GetUserRequest_SlackUserId) isGetUserRequest_Identifier() {}

func (m *GetUserRequest) GetIdentifier() isGetUserRequest_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *GetUserRequest) GetId() int64 {
	if x, ok := m.GetIdentifier().(*GetUserRequest_Id); ok {
		return x.Id
	}
	return 0
}

func (m *GetUserRequest) GetSlackUserId() string {
	if x, ok := m.GetIdentifier().(*GetUserRequest_SlackUserId); ok {
		return x.SlackUserId
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GetUserRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GetUserRequest_Id)(nil),
		(*GetUserRequest_SlackUserId)(nil),
	}
}

type CreateUserRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SlackUserId          string   `protobuf:"bytes,2,opt,name=slack_user_id,json=slackUserId,proto3" json:"slack_user_id,omitempty"`
	SlackTeamId          string   `protobuf:"bytes,3,opt,name=slack_team_id,json=slackTeamId,proto3" json:"slack_team_id,omitempty"`
	Image                string   `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c11b3307f9b338e4, []int{4}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateUserRequest) GetSlackUserId() string {
	if m != nil {
		return m.SlackUserId
	}
	return ""
}

func (m *CreateUserRequest) GetSlackTeamId() string {
	if m != nil {
		return m.SlackTeamId
	}
	return ""
}

func (m *CreateUserRequest) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "stella.user.v1.User")
	proto.RegisterType((*ListUsersRequest)(nil), "stella.user.v1.ListUsersRequest")
	proto.RegisterType((*ListUsersResponse)(nil), "stella.user.v1.ListUsersResponse")
	proto.RegisterType((*GetUserRequest)(nil), "stella.user.v1.GetUserRequest")
	proto.RegisterType((*CreateUserRequest)(nil), "stella.user.v1.CreateUserRequest")
}

func init() { proto.RegisterFile("usersvc.proto", fileDescriptor_c11b3307f9b338e4) }

var fileDescriptor_c11b3307f9b338e4 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xdb, 0xaa, 0x1a, 0x31,
	0x14, 0x86, 0xe7, 0x64, 0x8b, 0xcb, 0x2a, 0x1a, 0xa4, 0x0c, 0x16, 0xca, 0x34, 0x78, 0x21, 0xbd,
	0x88, 0xd4, 0x3e, 0x40, 0x0f, 0x52, 0x54, 0xe8, 0xd5, 0xf4, 0x70, 0x2b, 0xa3, 0xb3, 0x94, 0xd0,
	0x39, 0xd8, 0x24, 0x23, 0xf4, 0x01, 0x0a, 0xfb, 0x6d, 0xf6, 0x2b, 0x6e, 0x92, 0xd9, 0xea, 0x38,
	0x6e, 0x65, 0xdf, 0xc5, 0xe4, 0xcb, 0xef, 0xfc, 0xdf, 0x0a, 0xb4, 0x0b, 0x89, 0x42, 0xee, 0xd7,
	0x6c, 0x27, 0x72, 0x95, 0x93, 0x8e, 0x54, 0x98, 0x24, 0x11, 0xd3, 0xbb, 0x6c, 0xff, 0x61, 0xf0,
	0x66, 0x9b, 0xe7, 0xdb, 0x04, 0xc7, 0xe6, 0x74, 0x55, 0x6c, 0xc6, 0x98, 0xee, 0xd4, 0xbf, 0x12,
	0xa6, 0x77, 0x36, 0x78, 0xbf, 0x24, 0x0a, 0xd2, 0x01, 0x87, 0xc7, 0xbe, 0x1d, 0xd8, 0x23, 0x37,
	0x74, 0x78, 0x4c, 0x08, 0x78, 0x59, 0x94, 0xa2, 0xef, 0x04, 0xf6, 0xa8, 0x19, 0x9a, 0x35, 0xa1,
	0xd0, 0x96, 0x49, 0xb4, 0xfe, 0xb3, 0xd4, 0xd1, 0x4b, 0x1e, 0xfb, 0xae, 0x39, 0x6c, 0x99, 0x4d,
	0x9d, 0xb2, 0x88, 0x4f, 0x8c, 0xc2, 0x28, 0xd5, 0x8c, 0x57, 0x61, 0x7e, 0x62, 0x94, 0x2e, 0x62,
	0xd2, 0x87, 0x06, 0x4f, 0xa3, 0x2d, 0xfa, 0x0d, 0x73, 0x56, 0xfe, 0xa0, 0x43, 0xe8, 0x7e, 0xe7,
	0x52, 0xe9, 0x1c, 0x19, 0xe2, 0xdf, 0x02, 0xa5, 0x22, 0x5d, 0x70, 0x79, 0x2c, 0x7d, 0x3b, 0x70,
	0x47, 0x6e, 0xa8, 0x97, 0xf4, 0x13, 0xf4, 0x2a, 0x94, 0xdc, 0xe5, 0x99, 0x44, 0xf2, 0x1e, 0x1a,
	0xc6, 0x81, 0x01, 0x5b, 0x93, 0x3e, 0x3b, 0x57, 0xc0, 0x34, 0x1d, 0x96, 0x08, 0xfd, 0x0d, 0x9d,
	0x19, 0x9a, 0xfb, 0xa7, 0x3f, 0x39, 0x56, 0x9f, 0x5b, 0xa6, 0xfc, 0xb0, 0x5e, 0xd4, 0x58, 0x98,
	0x5b, 0x67, 0x55, 0xbf, 0xbe, 0x02, 0xe0, 0x31, 0x66, 0x8a, 0x6f, 0x38, 0x0a, 0xfa, 0xdf, 0x86,
	0xde, 0x54, 0x60, 0xa4, 0xb0, 0x9a, 0x7d, 0xd0, 0x68, 0xdf, 0xd2, 0xe8, 0x3c, 0x43, 0xa3, 0x7b,
	0x43, 0xa3, 0x57, 0xd1, 0x38, 0xb9, 0x77, 0xa0, 0xa5, 0x43, 0x7e, 0xa0, 0xd8, 0xf3, 0x35, 0x92,
	0xcf, 0xa6, 0xef, 0xb4, 0x10, 0x02, 0x33, 0x53, 0x9b, 0xbc, 0x66, 0xe5, 0x8b, 0x60, 0x87, 0x17,
	0xc1, 0xbe, 0xe9, 0x17, 0x31, 0x78, 0x52, 0x1b, 0xb5, 0xc8, 0x17, 0x78, 0xf9, 0x68, 0x8c, 0xbc,
	0xad, 0x23, 0xe7, 0x2a, 0xaf, 0x46, 0x84, 0xd0, 0x3c, 0x4e, 0x8d, 0x04, 0x75, 0xa8, 0x3e, 0xf6,
	0xc1, 0xbb, 0x1b, 0x44, 0x39, 0x72, 0x6a, 0x91, 0x19, 0xc0, 0xc9, 0x37, 0xb9, 0xb8, 0x72, 0x31,
	0x8b, 0x6b, 0x1f, 0xb7, 0x7a, 0x61, 0x3c, 0x7c, 0x7c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x77, 0x6c,
	0xaa, 0xdc, 0x48, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	GetCurrentUser(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*User, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetCurrentUser(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/stella.user.v1.UserService/GetCurrentUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/stella.user.v1.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, "/stella.user.v1.UserService/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/stella.user.v1.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	GetCurrentUser(context.Context, *empty.Empty) (*User, error)
	GetUser(context.Context, *GetUserRequest) (*User, error)
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) GetCurrentUser(ctx context.Context, req *empty.Empty) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentUser not implemented")
}
func (*UnimplementedUserServiceServer) GetUser(ctx context.Context, req *GetUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedUserServiceServer) ListUsers(ctx context.Context, req *ListUsersRequest) (*ListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (*UnimplementedUserServiceServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_GetCurrentUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetCurrentUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stella.user.v1.UserService/GetCurrentUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetCurrentUser(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stella.user.v1.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stella.user.v1.UserService/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListUsers(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stella.user.v1.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stella.user.v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCurrentUser",
			Handler:    _UserService_GetCurrentUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _UserService_ListUsers_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usersvc.proto",
}
