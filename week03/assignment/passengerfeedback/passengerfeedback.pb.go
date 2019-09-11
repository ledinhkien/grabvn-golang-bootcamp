// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passengerfeedback/passengerfeedback.proto

package passengerfeedback

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

type PassengerFeedback struct {
	BookingCode          string   `protobuf:"bytes,1,opt,name=bookingCode,proto3" json:"bookingCode,omitempty"`
	PassengerID          int32    `protobuf:"varint,2,opt,name=passengerID,proto3" json:"passengerID,omitempty"`
	Feedback             string   `protobuf:"bytes,3,opt,name=feedback,proto3" json:"feedback,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PassengerFeedback) Reset()         { *m = PassengerFeedback{} }
func (m *PassengerFeedback) String() string { return proto.CompactTextString(m) }
func (*PassengerFeedback) ProtoMessage()    {}
func (*PassengerFeedback) Descriptor() ([]byte, []int) {
	return fileDescriptor_02fdbf3897eefc3d, []int{0}
}

func (m *PassengerFeedback) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PassengerFeedback.Unmarshal(m, b)
}
func (m *PassengerFeedback) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PassengerFeedback.Marshal(b, m, deterministic)
}
func (m *PassengerFeedback) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PassengerFeedback.Merge(m, src)
}
func (m *PassengerFeedback) XXX_Size() int {
	return xxx_messageInfo_PassengerFeedback.Size(m)
}
func (m *PassengerFeedback) XXX_DiscardUnknown() {
	xxx_messageInfo_PassengerFeedback.DiscardUnknown(m)
}

var xxx_messageInfo_PassengerFeedback proto.InternalMessageInfo

func (m *PassengerFeedback) GetBookingCode() string {
	if m != nil {
		return m.BookingCode
	}
	return ""
}

func (m *PassengerFeedback) GetPassengerID() int32 {
	if m != nil {
		return m.PassengerID
	}
	return 0
}

func (m *PassengerFeedback) GetFeedback() string {
	if m != nil {
		return m.Feedback
	}
	return ""
}

type FeedbackList struct {
	Feedbacks            []*PassengerFeedback `protobuf:"bytes,1,rep,name=feedbacks,proto3" json:"feedbacks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FeedbackList) Reset()         { *m = FeedbackList{} }
func (m *FeedbackList) String() string { return proto.CompactTextString(m) }
func (*FeedbackList) ProtoMessage()    {}
func (*FeedbackList) Descriptor() ([]byte, []int) {
	return fileDescriptor_02fdbf3897eefc3d, []int{1}
}

func (m *FeedbackList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedbackList.Unmarshal(m, b)
}
func (m *FeedbackList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedbackList.Marshal(b, m, deterministic)
}
func (m *FeedbackList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedbackList.Merge(m, src)
}
func (m *FeedbackList) XXX_Size() int {
	return xxx_messageInfo_FeedbackList.Size(m)
}
func (m *FeedbackList) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedbackList.DiscardUnknown(m)
}

var xxx_messageInfo_FeedbackList proto.InternalMessageInfo

func (m *FeedbackList) GetFeedbacks() []*PassengerFeedback {
	if m != nil {
		return m.Feedbacks
	}
	return nil
}

type FeedbackReq struct {
	Feedback             *PassengerFeedback `protobuf:"bytes,1,opt,name=Feedback,proto3" json:"Feedback,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *FeedbackReq) Reset()         { *m = FeedbackReq{} }
func (m *FeedbackReq) String() string { return proto.CompactTextString(m) }
func (*FeedbackReq) ProtoMessage()    {}
func (*FeedbackReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_02fdbf3897eefc3d, []int{2}
}

func (m *FeedbackReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedbackReq.Unmarshal(m, b)
}
func (m *FeedbackReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedbackReq.Marshal(b, m, deterministic)
}
func (m *FeedbackReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedbackReq.Merge(m, src)
}
func (m *FeedbackReq) XXX_Size() int {
	return xxx_messageInfo_FeedbackReq.Size(m)
}
func (m *FeedbackReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedbackReq.DiscardUnknown(m)
}

var xxx_messageInfo_FeedbackReq proto.InternalMessageInfo

func (m *FeedbackReq) GetFeedback() *PassengerFeedback {
	if m != nil {
		return m.Feedback
	}
	return nil
}

type FeedbackRes struct {
	Feedback             *PassengerFeedback `protobuf:"bytes,1,opt,name=Feedback,proto3" json:"Feedback,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *FeedbackRes) Reset()         { *m = FeedbackRes{} }
func (m *FeedbackRes) String() string { return proto.CompactTextString(m) }
func (*FeedbackRes) ProtoMessage()    {}
func (*FeedbackRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_02fdbf3897eefc3d, []int{3}
}

func (m *FeedbackRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedbackRes.Unmarshal(m, b)
}
func (m *FeedbackRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedbackRes.Marshal(b, m, deterministic)
}
func (m *FeedbackRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedbackRes.Merge(m, src)
}
func (m *FeedbackRes) XXX_Size() int {
	return xxx_messageInfo_FeedbackRes.Size(m)
}
func (m *FeedbackRes) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedbackRes.DiscardUnknown(m)
}

var xxx_messageInfo_FeedbackRes proto.InternalMessageInfo

func (m *FeedbackRes) GetFeedback() *PassengerFeedback {
	if m != nil {
		return m.Feedback
	}
	return nil
}

type PassengerIDReq struct {
	PassengerID          int32    `protobuf:"varint,1,opt,name=passengerID,proto3" json:"passengerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PassengerIDReq) Reset()         { *m = PassengerIDReq{} }
func (m *PassengerIDReq) String() string { return proto.CompactTextString(m) }
func (*PassengerIDReq) ProtoMessage()    {}
func (*PassengerIDReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_02fdbf3897eefc3d, []int{4}
}

func (m *PassengerIDReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PassengerIDReq.Unmarshal(m, b)
}
func (m *PassengerIDReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PassengerIDReq.Marshal(b, m, deterministic)
}
func (m *PassengerIDReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PassengerIDReq.Merge(m, src)
}
func (m *PassengerIDReq) XXX_Size() int {
	return xxx_messageInfo_PassengerIDReq.Size(m)
}
func (m *PassengerIDReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PassengerIDReq.DiscardUnknown(m)
}

var xxx_messageInfo_PassengerIDReq proto.InternalMessageInfo

func (m *PassengerIDReq) GetPassengerID() int32 {
	if m != nil {
		return m.PassengerID
	}
	return 0
}

type BookingCodeReq struct {
	BookingCode          string   `protobuf:"bytes,1,opt,name=bookingCode,proto3" json:"bookingCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookingCodeReq) Reset()         { *m = BookingCodeReq{} }
func (m *BookingCodeReq) String() string { return proto.CompactTextString(m) }
func (*BookingCodeReq) ProtoMessage()    {}
func (*BookingCodeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_02fdbf3897eefc3d, []int{5}
}

func (m *BookingCodeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookingCodeReq.Unmarshal(m, b)
}
func (m *BookingCodeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookingCodeReq.Marshal(b, m, deterministic)
}
func (m *BookingCodeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookingCodeReq.Merge(m, src)
}
func (m *BookingCodeReq) XXX_Size() int {
	return xxx_messageInfo_BookingCodeReq.Size(m)
}
func (m *BookingCodeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BookingCodeReq.DiscardUnknown(m)
}

var xxx_messageInfo_BookingCodeReq proto.InternalMessageInfo

func (m *BookingCodeReq) GetBookingCode() string {
	if m != nil {
		return m.BookingCode
	}
	return ""
}

type DeleteFeedbackByPassengerIDRes struct {
	Deleted              int32    `protobuf:"varint,1,opt,name=deleted,proto3" json:"deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteFeedbackByPassengerIDRes) Reset()         { *m = DeleteFeedbackByPassengerIDRes{} }
func (m *DeleteFeedbackByPassengerIDRes) String() string { return proto.CompactTextString(m) }
func (*DeleteFeedbackByPassengerIDRes) ProtoMessage()    {}
func (*DeleteFeedbackByPassengerIDRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_02fdbf3897eefc3d, []int{6}
}

func (m *DeleteFeedbackByPassengerIDRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteFeedbackByPassengerIDRes.Unmarshal(m, b)
}
func (m *DeleteFeedbackByPassengerIDRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteFeedbackByPassengerIDRes.Marshal(b, m, deterministic)
}
func (m *DeleteFeedbackByPassengerIDRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteFeedbackByPassengerIDRes.Merge(m, src)
}
func (m *DeleteFeedbackByPassengerIDRes) XXX_Size() int {
	return xxx_messageInfo_DeleteFeedbackByPassengerIDRes.Size(m)
}
func (m *DeleteFeedbackByPassengerIDRes) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteFeedbackByPassengerIDRes.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteFeedbackByPassengerIDRes proto.InternalMessageInfo

func (m *DeleteFeedbackByPassengerIDRes) GetDeleted() int32 {
	if m != nil {
		return m.Deleted
	}
	return 0
}

func init() {
	proto.RegisterType((*PassengerFeedback)(nil), "passengerfeedback.PassengerFeedback")
	proto.RegisterType((*FeedbackList)(nil), "passengerfeedback.FeedbackList")
	proto.RegisterType((*FeedbackReq)(nil), "passengerfeedback.FeedbackReq")
	proto.RegisterType((*FeedbackRes)(nil), "passengerfeedback.FeedbackRes")
	proto.RegisterType((*PassengerIDReq)(nil), "passengerfeedback.PassengerIDReq")
	proto.RegisterType((*BookingCodeReq)(nil), "passengerfeedback.BookingCodeReq")
	proto.RegisterType((*DeleteFeedbackByPassengerIDRes)(nil), "passengerfeedback.DeleteFeedbackByPassengerIDRes")
}

func init() {
	proto.RegisterFile("passengerfeedback/passengerfeedback.proto", fileDescriptor_02fdbf3897eefc3d)
}

var fileDescriptor_02fdbf3897eefc3d = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x4d, 0x4f, 0x83, 0x40,
	0x10, 0x2d, 0x12, 0xb5, 0x1d, 0x4c, 0x93, 0xee, 0x69, 0x43, 0x93, 0x8a, 0x1b, 0x0f, 0x78, 0xa9,
	0x11, 0x6f, 0x9e, 0x0c, 0x6d, 0x34, 0x26, 0x1a, 0x1b, 0x7e, 0x81, 0x20, 0x23, 0x69, 0xaa, 0x50,
	0x59, 0x3c, 0xf8, 0x9f, 0xfc, 0x91, 0x06, 0xc2, 0xf2, 0xe1, 0x6e, 0x36, 0x1c, 0x3c, 0xce, 0x9b,
	0x99, 0xf7, 0x66, 0xde, 0xec, 0xc2, 0xc5, 0x3e, 0xe4, 0x1c, 0xd3, 0x04, 0xf3, 0x37, 0xc4, 0x38,
	0x0a, 0x5f, 0x77, 0x97, 0x12, 0xb2, 0xdc, 0xe7, 0x59, 0x91, 0x91, 0x99, 0x94, 0x60, 0x1c, 0x66,
	0x1b, 0x01, 0xde, 0xd5, 0x20, 0x71, 0xc0, 0x8a, 0xb2, 0x6c, 0xb7, 0x4d, 0x93, 0x55, 0x16, 0x23,
	0x35, 0x1c, 0xc3, 0x9d, 0x04, 0x5d, 0xa8, 0xac, 0x68, 0xb8, 0x1e, 0xd6, 0xf4, 0xc0, 0x31, 0xdc,
	0xc3, 0xa0, 0x0b, 0x11, 0x1b, 0xc6, 0x42, 0x84, 0x9a, 0x15, 0x41, 0x13, 0xb3, 0x00, 0x4e, 0x84,
	0xd6, 0xe3, 0x96, 0x17, 0xc4, 0x87, 0x89, 0xc8, 0x71, 0x6a, 0x38, 0xa6, 0x6b, 0x79, 0xe7, 0x4b,
	0x79, 0x09, 0x69, 0xd0, 0xa0, 0x6d, 0x63, 0xcf, 0x60, 0x35, 0x30, 0x7e, 0x92, 0x5b, 0x18, 0x8b,
	0xb0, 0x9a, 0x7f, 0x28, 0x63, 0xd3, 0xd5, 0x27, 0xe4, 0xff, 0x40, 0xe8, 0xc1, 0x74, 0xd3, 0x1a,
	0x54, 0x0e, 0xf9, 0xc7, 0x45, 0x43, 0x72, 0xb1, 0xec, 0xf1, 0x5b, 0xdb, 0xeb, 0x1e, 0xfd, 0x6d,
	0xd8, 0x0d, 0x2c, 0xd6, 0xf8, 0x8e, 0x05, 0x0a, 0x65, 0xff, 0xbb, 0xa7, 0xcb, 0x09, 0x85, 0xe3,
	0xb8, 0xaa, 0x88, 0x6b, 0x4d, 0x11, 0x7a, 0x3f, 0x26, 0xcc, 0xa5, 0x1d, 0x9e, 0xc2, 0x34, 0x4c,
	0xf0, 0x03, 0xd3, 0x82, 0x04, 0x30, 0x5d, 0xe5, 0x18, 0xb6, 0xdc, 0x64, 0xa1, 0x70, 0xa1, 0x73,
	0x08, 0x5b, 0x9f, 0xe7, 0x6c, 0x44, 0x5e, 0x80, 0xde, 0x63, 0xa1, 0x1c, 0x96, 0x9c, 0xe9, 0x3c,
	0xae, 0x4c, 0xb4, 0x4f, 0x35, 0x02, 0xe5, 0xeb, 0x52, 0x28, 0x74, 0x2c, 0x55, 0x2a, 0xf4, 0x2d,
	0x1f, 0xa2, 0xf0, 0x05, 0x73, 0x8d, 0xe7, 0x43, 0xd6, 0xb8, 0x52, 0x94, 0xe8, 0xcf, 0xc8, 0x46,
	0xd1, 0x51, 0xf5, 0xaf, 0xaf, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x75, 0x36, 0x7d, 0x04,
	0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PassengerFeedbackManagementClient is the client API for PassengerFeedbackManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PassengerFeedbackManagementClient interface {
	CreateFeedback(ctx context.Context, in *FeedbackReq, opts ...grpc.CallOption) (*FeedbackRes, error)
	GetFeedbackByPassengerID(ctx context.Context, in *PassengerIDReq, opts ...grpc.CallOption) (*FeedbackList, error)
	GetFeedbackByBookingCode(ctx context.Context, in *BookingCodeReq, opts ...grpc.CallOption) (*FeedbackList, error)
	DeleteFeedbackByPassengerID(ctx context.Context, in *PassengerIDReq, opts ...grpc.CallOption) (*DeleteFeedbackByPassengerIDRes, error)
}

type passengerFeedbackManagementClient struct {
	cc *grpc.ClientConn
}

func NewPassengerFeedbackManagementClient(cc *grpc.ClientConn) PassengerFeedbackManagementClient {
	return &passengerFeedbackManagementClient{cc}
}

func (c *passengerFeedbackManagementClient) CreateFeedback(ctx context.Context, in *FeedbackReq, opts ...grpc.CallOption) (*FeedbackRes, error) {
	out := new(FeedbackRes)
	err := c.cc.Invoke(ctx, "/passengerfeedback.PassengerFeedbackManagement/CreateFeedback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passengerFeedbackManagementClient) GetFeedbackByPassengerID(ctx context.Context, in *PassengerIDReq, opts ...grpc.CallOption) (*FeedbackList, error) {
	out := new(FeedbackList)
	err := c.cc.Invoke(ctx, "/passengerfeedback.PassengerFeedbackManagement/GetFeedbackByPassengerID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passengerFeedbackManagementClient) GetFeedbackByBookingCode(ctx context.Context, in *BookingCodeReq, opts ...grpc.CallOption) (*FeedbackList, error) {
	out := new(FeedbackList)
	err := c.cc.Invoke(ctx, "/passengerfeedback.PassengerFeedbackManagement/GetFeedbackByBookingCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passengerFeedbackManagementClient) DeleteFeedbackByPassengerID(ctx context.Context, in *PassengerIDReq, opts ...grpc.CallOption) (*DeleteFeedbackByPassengerIDRes, error) {
	out := new(DeleteFeedbackByPassengerIDRes)
	err := c.cc.Invoke(ctx, "/passengerfeedback.PassengerFeedbackManagement/DeleteFeedbackByPassengerID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PassengerFeedbackManagementServer is the server API for PassengerFeedbackManagement service.
type PassengerFeedbackManagementServer interface {
	CreateFeedback(context.Context, *FeedbackReq) (*FeedbackRes, error)
	GetFeedbackByPassengerID(context.Context, *PassengerIDReq) (*FeedbackList, error)
	GetFeedbackByBookingCode(context.Context, *BookingCodeReq) (*FeedbackList, error)
	DeleteFeedbackByPassengerID(context.Context, *PassengerIDReq) (*DeleteFeedbackByPassengerIDRes, error)
}

// UnimplementedPassengerFeedbackManagementServer can be embedded to have forward compatible implementations.
type UnimplementedPassengerFeedbackManagementServer struct {
}

func (*UnimplementedPassengerFeedbackManagementServer) CreateFeedback(ctx context.Context, req *FeedbackReq) (*FeedbackRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFeedback not implemented")
}
func (*UnimplementedPassengerFeedbackManagementServer) GetFeedbackByPassengerID(ctx context.Context, req *PassengerIDReq) (*FeedbackList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeedbackByPassengerID not implemented")
}
func (*UnimplementedPassengerFeedbackManagementServer) GetFeedbackByBookingCode(ctx context.Context, req *BookingCodeReq) (*FeedbackList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeedbackByBookingCode not implemented")
}
func (*UnimplementedPassengerFeedbackManagementServer) DeleteFeedbackByPassengerID(ctx context.Context, req *PassengerIDReq) (*DeleteFeedbackByPassengerIDRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFeedbackByPassengerID not implemented")
}

func RegisterPassengerFeedbackManagementServer(s *grpc.Server, srv PassengerFeedbackManagementServer) {
	s.RegisterService(&_PassengerFeedbackManagement_serviceDesc, srv)
}

func _PassengerFeedbackManagement_CreateFeedback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedbackReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassengerFeedbackManagementServer).CreateFeedback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passengerfeedback.PassengerFeedbackManagement/CreateFeedback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassengerFeedbackManagementServer).CreateFeedback(ctx, req.(*FeedbackReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PassengerFeedbackManagement_GetFeedbackByPassengerID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PassengerIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassengerFeedbackManagementServer).GetFeedbackByPassengerID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passengerfeedback.PassengerFeedbackManagement/GetFeedbackByPassengerID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassengerFeedbackManagementServer).GetFeedbackByPassengerID(ctx, req.(*PassengerIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PassengerFeedbackManagement_GetFeedbackByBookingCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookingCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassengerFeedbackManagementServer).GetFeedbackByBookingCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passengerfeedback.PassengerFeedbackManagement/GetFeedbackByBookingCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassengerFeedbackManagementServer).GetFeedbackByBookingCode(ctx, req.(*BookingCodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PassengerFeedbackManagement_DeleteFeedbackByPassengerID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PassengerIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassengerFeedbackManagementServer).DeleteFeedbackByPassengerID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/passengerfeedback.PassengerFeedbackManagement/DeleteFeedbackByPassengerID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassengerFeedbackManagementServer).DeleteFeedbackByPassengerID(ctx, req.(*PassengerIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _PassengerFeedbackManagement_serviceDesc = grpc.ServiceDesc{
	ServiceName: "passengerfeedback.PassengerFeedbackManagement",
	HandlerType: (*PassengerFeedbackManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFeedback",
			Handler:    _PassengerFeedbackManagement_CreateFeedback_Handler,
		},
		{
			MethodName: "GetFeedbackByPassengerID",
			Handler:    _PassengerFeedbackManagement_GetFeedbackByPassengerID_Handler,
		},
		{
			MethodName: "GetFeedbackByBookingCode",
			Handler:    _PassengerFeedbackManagement_GetFeedbackByBookingCode_Handler,
		},
		{
			MethodName: "DeleteFeedbackByPassengerID",
			Handler:    _PassengerFeedbackManagement_DeleteFeedbackByPassengerID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "passengerfeedback/passengerfeedback.proto",
}