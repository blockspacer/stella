// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ratingsvc.proto

package stella_rating_v1

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

type Rating struct {
	EntityId             int64    `protobuf:"varint,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	Score                float32  `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
	UserId               int64    `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Comment              string   `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rating) Reset()         { *m = Rating{} }
func (m *Rating) String() string { return proto.CompactTextString(m) }
func (*Rating) ProtoMessage()    {}
func (*Rating) Descriptor() ([]byte, []int) {
	return fileDescriptor_61e79acbe641d6ef, []int{0}
}

func (m *Rating) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rating.Unmarshal(m, b)
}
func (m *Rating) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rating.Marshal(b, m, deterministic)
}
func (m *Rating) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rating.Merge(m, src)
}
func (m *Rating) XXX_Size() int {
	return xxx_messageInfo_Rating.Size(m)
}
func (m *Rating) XXX_DiscardUnknown() {
	xxx_messageInfo_Rating.DiscardUnknown(m)
}

var xxx_messageInfo_Rating proto.InternalMessageInfo

func (m *Rating) GetEntityId() int64 {
	if m != nil {
		return m.EntityId
	}
	return 0
}

func (m *Rating) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *Rating) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Rating) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

type GetRatingRequest struct {
	EntityId             int64    `protobuf:"varint,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRatingRequest) Reset()         { *m = GetRatingRequest{} }
func (m *GetRatingRequest) String() string { return proto.CompactTextString(m) }
func (*GetRatingRequest) ProtoMessage()    {}
func (*GetRatingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_61e79acbe641d6ef, []int{1}
}

func (m *GetRatingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRatingRequest.Unmarshal(m, b)
}
func (m *GetRatingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRatingRequest.Marshal(b, m, deterministic)
}
func (m *GetRatingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRatingRequest.Merge(m, src)
}
func (m *GetRatingRequest) XXX_Size() int {
	return xxx_messageInfo_GetRatingRequest.Size(m)
}
func (m *GetRatingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRatingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRatingRequest proto.InternalMessageInfo

func (m *GetRatingRequest) GetEntityId() int64 {
	if m != nil {
		return m.EntityId
	}
	return 0
}

type GetRatingResponse struct {
	Score                float32  `protobuf:"fixed32,1,opt,name=score,proto3" json:"score,omitempty"`
	Count                int32    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRatingResponse) Reset()         { *m = GetRatingResponse{} }
func (m *GetRatingResponse) String() string { return proto.CompactTextString(m) }
func (*GetRatingResponse) ProtoMessage()    {}
func (*GetRatingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_61e79acbe641d6ef, []int{2}
}

func (m *GetRatingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRatingResponse.Unmarshal(m, b)
}
func (m *GetRatingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRatingResponse.Marshal(b, m, deterministic)
}
func (m *GetRatingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRatingResponse.Merge(m, src)
}
func (m *GetRatingResponse) XXX_Size() int {
	return xxx_messageInfo_GetRatingResponse.Size(m)
}
func (m *GetRatingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRatingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetRatingResponse proto.InternalMessageInfo

func (m *GetRatingResponse) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *GetRatingResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type GetUserRatingRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRatingRequest) Reset()         { *m = GetUserRatingRequest{} }
func (m *GetUserRatingRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRatingRequest) ProtoMessage()    {}
func (*GetUserRatingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_61e79acbe641d6ef, []int{3}
}

func (m *GetUserRatingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRatingRequest.Unmarshal(m, b)
}
func (m *GetUserRatingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRatingRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRatingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRatingRequest.Merge(m, src)
}
func (m *GetUserRatingRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRatingRequest.Size(m)
}
func (m *GetUserRatingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRatingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRatingRequest proto.InternalMessageInfo

func (m *GetUserRatingRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type ListRatingsResponse struct {
	Ratings              []*Rating `protobuf:"bytes,1,rep,name=ratings,proto3" json:"ratings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListRatingsResponse) Reset()         { *m = ListRatingsResponse{} }
func (m *ListRatingsResponse) String() string { return proto.CompactTextString(m) }
func (*ListRatingsResponse) ProtoMessage()    {}
func (*ListRatingsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_61e79acbe641d6ef, []int{4}
}

func (m *ListRatingsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRatingsResponse.Unmarshal(m, b)
}
func (m *ListRatingsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRatingsResponse.Marshal(b, m, deterministic)
}
func (m *ListRatingsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRatingsResponse.Merge(m, src)
}
func (m *ListRatingsResponse) XXX_Size() int {
	return xxx_messageInfo_ListRatingsResponse.Size(m)
}
func (m *ListRatingsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRatingsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListRatingsResponse proto.InternalMessageInfo

func (m *ListRatingsResponse) GetRatings() []*Rating {
	if m != nil {
		return m.Ratings
	}
	return nil
}

type UpsertRatingRequest struct {
	EntityId             int64    `protobuf:"varint,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Score                int64    `protobuf:"varint,3,opt,name=score,proto3" json:"score,omitempty"`
	Comment              string   `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpsertRatingRequest) Reset()         { *m = UpsertRatingRequest{} }
func (m *UpsertRatingRequest) String() string { return proto.CompactTextString(m) }
func (*UpsertRatingRequest) ProtoMessage()    {}
func (*UpsertRatingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_61e79acbe641d6ef, []int{5}
}

func (m *UpsertRatingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpsertRatingRequest.Unmarshal(m, b)
}
func (m *UpsertRatingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpsertRatingRequest.Marshal(b, m, deterministic)
}
func (m *UpsertRatingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpsertRatingRequest.Merge(m, src)
}
func (m *UpsertRatingRequest) XXX_Size() int {
	return xxx_messageInfo_UpsertRatingRequest.Size(m)
}
func (m *UpsertRatingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpsertRatingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpsertRatingRequest proto.InternalMessageInfo

func (m *UpsertRatingRequest) GetEntityId() int64 {
	if m != nil {
		return m.EntityId
	}
	return 0
}

func (m *UpsertRatingRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UpsertRatingRequest) GetScore() int64 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *UpsertRatingRequest) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

type DeleteRequest struct {
	EntityId             int64    `protobuf:"varint,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_61e79acbe641d6ef, []int{6}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetEntityId() int64 {
	if m != nil {
		return m.EntityId
	}
	return 0
}

func (m *DeleteRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func init() {
	proto.RegisterType((*Rating)(nil), "stella.rating.v1.Rating")
	proto.RegisterType((*GetRatingRequest)(nil), "stella.rating.v1.GetRatingRequest")
	proto.RegisterType((*GetRatingResponse)(nil), "stella.rating.v1.GetRatingResponse")
	proto.RegisterType((*GetUserRatingRequest)(nil), "stella.rating.v1.GetUserRatingRequest")
	proto.RegisterType((*ListRatingsResponse)(nil), "stella.rating.v1.ListRatingsResponse")
	proto.RegisterType((*UpsertRatingRequest)(nil), "stella.rating.v1.UpsertRatingRequest")
	proto.RegisterType((*DeleteRequest)(nil), "stella.rating.v1.DeleteRequest")
}

func init() { proto.RegisterFile("ratingsvc.proto", fileDescriptor_61e79acbe641d6ef) }

var fileDescriptor_61e79acbe641d6ef = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x41, 0xcf, 0xd2, 0x40,
	0x10, 0x6d, 0xe1, 0xa3, 0xc8, 0x7c, 0x92, 0x0f, 0x17, 0xa2, 0x4d, 0x39, 0xd8, 0xac, 0xd1, 0xf4,
	0xb4, 0x8d, 0xf8, 0x03, 0x3c, 0x28, 0x21, 0x24, 0x5e, 0x2c, 0xc1, 0x83, 0x17, 0x03, 0x65, 0x24,
	0x4d, 0x4a, 0xb7, 0x76, 0xb7, 0x24, 0xc4, 0xbf, 0xe1, 0x0f, 0x36, 0xed, 0x16, 0x68, 0x69, 0x21,
	0x18, 0x8f, 0xb3, 0xfb, 0x26, 0xef, 0xcd, 0x7b, 0x93, 0x81, 0xa7, 0x64, 0x25, 0x83, 0x68, 0x2b,
	0xf6, 0x3e, 0x8b, 0x13, 0x2e, 0x39, 0x19, 0x08, 0x89, 0x61, 0xb8, 0x62, 0xea, 0x9d, 0xed, 0xdf,
	0x5b, 0xe3, 0x2d, 0xe7, 0xdb, 0x10, 0xdd, 0xfc, 0x7f, 0x9d, 0xfe, 0x74, 0x71, 0x17, 0xcb, 0x83,
	0x82, 0xd3, 0x08, 0x0c, 0x2f, 0x47, 0x92, 0x31, 0xf4, 0x30, 0x92, 0x81, 0x3c, 0xfc, 0x08, 0x36,
	0xa6, 0x6e, 0xeb, 0x4e, 0xdb, 0x7b, 0xa6, 0x1e, 0xe6, 0x1b, 0x32, 0x82, 0x8e, 0xf0, 0x79, 0x82,
	0x66, 0xcb, 0xd6, 0x9d, 0x96, 0xa7, 0x0a, 0xf2, 0x0a, 0xba, 0xa9, 0xc0, 0x24, 0x6b, 0x68, 0xe7,
	0x0d, 0x46, 0x56, 0xce, 0x37, 0xc4, 0x84, 0xae, 0xcf, 0x77, 0x3b, 0x8c, 0xa4, 0xf9, 0x60, 0xeb,
	0x4e, 0xcf, 0x3b, 0x96, 0xd4, 0x85, 0xc1, 0x0c, 0xa5, 0xa2, 0xf4, 0xf0, 0x57, 0x8a, 0x42, 0xde,
	0x64, 0xa6, 0x1f, 0xe1, 0x45, 0xa9, 0x41, 0xc4, 0x3c, 0x12, 0x78, 0x96, 0xa3, 0x97, 0xe5, 0x8c,
	0xa0, 0xe3, 0xf3, 0x34, 0x92, 0xb9, 0xc8, 0x8e, 0xa7, 0x0a, 0xea, 0xc2, 0x68, 0x86, 0x72, 0x29,
	0x30, 0xa9, 0xb2, 0x96, 0xc4, 0xeb, 0x65, 0xf1, 0x74, 0x0e, 0xc3, 0x2f, 0x81, 0x28, 0x28, 0xc5,
	0x89, 0x73, 0x02, 0xdd, 0xc2, 0x6b, 0x53, 0xb7, 0xdb, 0xce, 0xe3, 0xc4, 0x64, 0x97, 0x56, 0xb3,
	0x82, 0xe1, 0x08, 0xa4, 0xbf, 0x61, 0xb8, 0x8c, 0x05, 0x26, 0xff, 0x30, 0x70, 0x59, 0x57, 0xab,
	0x62, 0xea, 0x69, 0x68, 0xe5, 0x75, 0x31, 0xf4, 0x75, 0xab, 0xa7, 0xd0, 0xff, 0x8c, 0x21, 0x4a,
	0xfc, 0x2f, 0xda, 0xc9, 0x9f, 0x07, 0xe8, 0x2b, 0xf9, 0x0b, 0x4c, 0xf6, 0x81, 0x8f, 0xe4, 0x1b,
	0xf4, 0x4e, 0x91, 0x10, 0x5a, 0x77, 0xe1, 0x32, 0x60, 0xeb, 0xcd, 0x4d, 0x8c, 0xf2, 0x97, 0x6a,
	0x64, 0x01, 0xfd, 0x4a, 0x52, 0xe4, 0x5d, 0x63, 0x5f, 0x2d, 0x4a, 0xeb, 0x6a, 0x12, 0x54, 0x23,
	0xdf, 0xe1, 0xb1, 0x94, 0xe6, 0x5d, 0x72, 0xdf, 0xd6, 0x31, 0x0d, 0x0b, 0x41, 0x35, 0xb2, 0x86,
	0xa7, 0xec, 0xe3, 0x2c, 0x48, 0xdc, 0x2d, 0xf9, 0x6e, 0x8e, 0xaf, 0xf0, 0xbc, 0xbc, 0x42, 0xa4,
	0xa1, 0xb1, 0x61, 0xc5, 0x6e, 0x5a, 0xf2, 0x09, 0x0c, 0xb5, 0x18, 0xe4, 0x75, 0x1d, 0x55, 0x59,
	0x19, 0xeb, 0x25, 0x53, 0xc7, 0x83, 0x1d, 0x8f, 0x07, 0x9b, 0x66, 0xc7, 0x83, 0x6a, 0x6b, 0x23,
	0x7f, 0xf9, 0xf0, 0x37, 0x00, 0x00, 0xff, 0xff, 0xff, 0xb8, 0xbb, 0x49, 0x81, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RatingServiceClient is the client API for RatingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RatingServiceClient interface {
	GetRating(ctx context.Context, in *GetRatingRequest, opts ...grpc.CallOption) (*GetRatingResponse, error)
	GetUserRating(ctx context.Context, in *GetUserRatingRequest, opts ...grpc.CallOption) (*Rating, error)
	ListRatings(ctx context.Context, in *GetRatingRequest, opts ...grpc.CallOption) (*ListRatingsResponse, error)
	ListUserRatings(ctx context.Context, in *GetUserRatingRequest, opts ...grpc.CallOption) (*ListRatingsResponse, error)
	UpsertRating(ctx context.Context, in *UpsertRatingRequest, opts ...grpc.CallOption) (*Rating, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type ratingServiceClient struct {
	cc *grpc.ClientConn
}

func NewRatingServiceClient(cc *grpc.ClientConn) RatingServiceClient {
	return &ratingServiceClient{cc}
}

func (c *ratingServiceClient) GetRating(ctx context.Context, in *GetRatingRequest, opts ...grpc.CallOption) (*GetRatingResponse, error) {
	out := new(GetRatingResponse)
	err := c.cc.Invoke(ctx, "/stella.rating.v1.RatingService/GetRating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetUserRating(ctx context.Context, in *GetUserRatingRequest, opts ...grpc.CallOption) (*Rating, error) {
	out := new(Rating)
	err := c.cc.Invoke(ctx, "/stella.rating.v1.RatingService/GetUserRating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) ListRatings(ctx context.Context, in *GetRatingRequest, opts ...grpc.CallOption) (*ListRatingsResponse, error) {
	out := new(ListRatingsResponse)
	err := c.cc.Invoke(ctx, "/stella.rating.v1.RatingService/ListRatings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) ListUserRatings(ctx context.Context, in *GetUserRatingRequest, opts ...grpc.CallOption) (*ListRatingsResponse, error) {
	out := new(ListRatingsResponse)
	err := c.cc.Invoke(ctx, "/stella.rating.v1.RatingService/ListUserRatings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) UpsertRating(ctx context.Context, in *UpsertRatingRequest, opts ...grpc.CallOption) (*Rating, error) {
	out := new(Rating)
	err := c.cc.Invoke(ctx, "/stella.rating.v1.RatingService/UpsertRating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/stella.rating.v1.RatingService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RatingServiceServer is the server API for RatingService service.
type RatingServiceServer interface {
	GetRating(context.Context, *GetRatingRequest) (*GetRatingResponse, error)
	GetUserRating(context.Context, *GetUserRatingRequest) (*Rating, error)
	ListRatings(context.Context, *GetRatingRequest) (*ListRatingsResponse, error)
	ListUserRatings(context.Context, *GetUserRatingRequest) (*ListRatingsResponse, error)
	UpsertRating(context.Context, *UpsertRatingRequest) (*Rating, error)
	Delete(context.Context, *DeleteRequest) (*empty.Empty, error)
}

// UnimplementedRatingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRatingServiceServer struct {
}

func (*UnimplementedRatingServiceServer) GetRating(ctx context.Context, req *GetRatingRequest) (*GetRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRating not implemented")
}
func (*UnimplementedRatingServiceServer) GetUserRating(ctx context.Context, req *GetUserRatingRequest) (*Rating, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserRating not implemented")
}
func (*UnimplementedRatingServiceServer) ListRatings(ctx context.Context, req *GetRatingRequest) (*ListRatingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRatings not implemented")
}
func (*UnimplementedRatingServiceServer) ListUserRatings(ctx context.Context, req *GetUserRatingRequest) (*ListRatingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserRatings not implemented")
}
func (*UnimplementedRatingServiceServer) UpsertRating(ctx context.Context, req *UpsertRatingRequest) (*Rating, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertRating not implemented")
}
func (*UnimplementedRatingServiceServer) Delete(ctx context.Context, req *DeleteRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterRatingServiceServer(s *grpc.Server, srv RatingServiceServer) {
	s.RegisterService(&_RatingService_serviceDesc, srv)
}

func _RatingService_GetRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stella.rating.v1.RatingService/GetRating",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetRating(ctx, req.(*GetRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetUserRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetUserRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stella.rating.v1.RatingService/GetUserRating",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetUserRating(ctx, req.(*GetUserRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_ListRatings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).ListRatings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stella.rating.v1.RatingService/ListRatings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).ListRatings(ctx, req.(*GetRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_ListUserRatings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).ListUserRatings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stella.rating.v1.RatingService/ListUserRatings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).ListUserRatings(ctx, req.(*GetUserRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_UpsertRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).UpsertRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stella.rating.v1.RatingService/UpsertRating",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).UpsertRating(ctx, req.(*UpsertRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stella.rating.v1.RatingService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RatingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stella.rating.v1.RatingService",
	HandlerType: (*RatingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRating",
			Handler:    _RatingService_GetRating_Handler,
		},
		{
			MethodName: "GetUserRating",
			Handler:    _RatingService_GetUserRating_Handler,
		},
		{
			MethodName: "ListRatings",
			Handler:    _RatingService_ListRatings_Handler,
		},
		{
			MethodName: "ListUserRatings",
			Handler:    _RatingService_ListUserRatings_Handler,
		},
		{
			MethodName: "UpsertRating",
			Handler:    _RatingService_UpsertRating_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RatingService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ratingsvc.proto",
}
