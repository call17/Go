// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Event struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Text                 string               `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	Owner                string               `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	StartTime            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,6,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Event) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Event) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Event) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Event) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

type CreateEventRequest struct {
	Title                string               `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Text                 string               `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	StartTime            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CreateEventRequest) Reset()         { *m = CreateEventRequest{} }
func (m *CreateEventRequest) String() string { return proto.CompactTextString(m) }
func (*CreateEventRequest) ProtoMessage()    {}
func (*CreateEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{1}
}

func (m *CreateEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEventRequest.Unmarshal(m, b)
}
func (m *CreateEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEventRequest.Marshal(b, m, deterministic)
}
func (m *CreateEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEventRequest.Merge(m, src)
}
func (m *CreateEventRequest) XXX_Size() int {
	return xxx_messageInfo_CreateEventRequest.Size(m)
}
func (m *CreateEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEventRequest proto.InternalMessageInfo

func (m *CreateEventRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CreateEventRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *CreateEventRequest) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *CreateEventRequest) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

type CreateEventResponse struct {
	// Types that are valid to be assigned to Result:
	//	*CreateEventResponse_Event
	//	*CreateEventResponse_Error
	Result               isCreateEventResponse_Result `protobuf_oneof:"result"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *CreateEventResponse) Reset()         { *m = CreateEventResponse{} }
func (m *CreateEventResponse) String() string { return proto.CompactTextString(m) }
func (*CreateEventResponse) ProtoMessage()    {}
func (*CreateEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{2}
}

func (m *CreateEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEventResponse.Unmarshal(m, b)
}
func (m *CreateEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEventResponse.Marshal(b, m, deterministic)
}
func (m *CreateEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEventResponse.Merge(m, src)
}
func (m *CreateEventResponse) XXX_Size() int {
	return xxx_messageInfo_CreateEventResponse.Size(m)
}
func (m *CreateEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEventResponse proto.InternalMessageInfo

type isCreateEventResponse_Result interface {
	isCreateEventResponse_Result()
}

type CreateEventResponse_Event struct {
	Event *Event `protobuf:"bytes,1,opt,name=event,proto3,oneof"`
}

type CreateEventResponse_Error struct {
	Error string `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*CreateEventResponse_Event) isCreateEventResponse_Result() {}

func (*CreateEventResponse_Error) isCreateEventResponse_Result() {}

func (m *CreateEventResponse) GetResult() isCreateEventResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *CreateEventResponse) GetEvent() *Event {
	if x, ok := m.GetResult().(*CreateEventResponse_Event); ok {
		return x.Event
	}
	return nil
}

func (m *CreateEventResponse) GetError() string {
	if x, ok := m.GetResult().(*CreateEventResponse_Error); ok {
		return x.Error
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CreateEventResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CreateEventResponse_Event)(nil),
		(*CreateEventResponse_Error)(nil),
	}
}

func init() {
	proto.RegisterType((*Event)(nil), "Event")
	proto.RegisterType((*CreateEventRequest)(nil), "CreateEventRequest")
	proto.RegisterType((*CreateEventResponse)(nil), "CreateEventResponse")
}

func init() { proto.RegisterFile("api/api.proto", fileDescriptor_1b40cafcd4234784) }

var fileDescriptor_1b40cafcd4234784 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x9b, 0x6d, 0x77, 0x6d, 0xa7, 0xa8, 0x90, 0x16, 0x59, 0x7a, 0xd0, 0xd2, 0x53, 0x4f,
	0x29, 0x54, 0x3c, 0xe8, 0xd1, 0x22, 0x78, 0xf1, 0xb2, 0x0a, 0x1e, 0x25, 0x75, 0xc7, 0x12, 0xd8,
	0x26, 0x6b, 0x32, 0xad, 0xfe, 0x2c, 0xff, 0x86, 0xff, 0x4a, 0x92, 0x50, 0x6c, 0xb1, 0xa0, 0x78,
	0xcb, 0x7b, 0x79, 0xc3, 0x7c, 0x6f, 0xe0, 0x50, 0xd6, 0x6a, 0x22, 0x6b, 0x25, 0x6a, 0x6b, 0xc8,
	0x0c, 0xce, 0x16, 0xc6, 0x2c, 0x2a, 0x9c, 0x04, 0x35, 0x5f, 0xbd, 0x4c, 0x48, 0x2d, 0xd1, 0x91,
	0x5c, 0xd6, 0x31, 0x30, 0xfa, 0x64, 0x90, 0xde, 0xac, 0x51, 0x13, 0x3f, 0x82, 0x44, 0x95, 0x39,
	0x1b, 0xb2, 0x71, 0xa7, 0x48, 0x54, 0xc9, 0xfb, 0x90, 0x92, 0xa2, 0x0a, 0xf3, 0x24, 0x58, 0x51,
	0x70, 0x0e, 0x2d, 0xc2, 0x77, 0xca, 0x9b, 0xc1, 0x0c, 0x6f, 0x9f, 0x34, 0x6f, 0x1a, 0x6d, 0xde,
	0x8a, 0xc9, 0x20, 0xf8, 0x25, 0x80, 0x23, 0x69, 0xe9, 0xc9, 0xaf, 0xcc, 0xd3, 0x21, 0x1b, 0x77,
	0xa7, 0x03, 0x11, 0x79, 0xc4, 0x86, 0x47, 0x3c, 0x6c, 0x78, 0x8a, 0x4e, 0x48, 0x7b, 0xcd, 0x2f,
	0xa0, 0x8d, 0xba, 0x8c, 0x83, 0xd9, 0xaf, 0x83, 0x07, 0xa8, 0x4b, 0xaf, 0x46, 0x1f, 0x0c, 0xf8,
	0xcc, 0xa2, 0x24, 0x0c, 0x8d, 0x0a, 0x7c, 0x5d, 0xa1, 0xa3, 0xef, 0x22, 0x6c, 0x5f, 0x91, 0x64,
	0xab, 0xc8, 0x2e, 0x72, 0xf3, 0xbf, 0xc8, 0xad, 0xbf, 0x23, 0x3f, 0x42, 0x6f, 0x87, 0xd8, 0xd5,
	0x46, 0x3b, 0xe4, 0xa7, 0x90, 0xa2, 0x37, 0x02, 0x72, 0x77, 0x9a, 0x89, 0xf0, 0x7d, 0xdb, 0x28,
	0xa2, 0xcd, 0x4f, 0x20, 0x45, 0x6b, 0x8d, 0x8d, 0xf4, 0xc1, 0xf7, 0xf2, 0xba, 0x0d, 0x99, 0x45,
	0xb7, 0xaa, 0x68, 0x7a, 0x07, 0xc7, 0x33, 0x59, 0xa1, 0x2e, 0xa5, 0xbd, 0x47, 0xbb, 0x56, 0xcf,
	0xc8, 0xaf, 0xa0, 0xbb, 0xb5, 0x8b, 0xf7, 0xc4, 0xcf, 0x5b, 0x0d, 0xfa, 0x62, 0x0f, 0xce, 0xa8,
	0x31, 0xcf, 0x42, 0x89, 0xf3, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x14, 0xcf, 0xed, 0xde, 0x5f,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalendarServiceClient is the client API for CalendarService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalendarServiceClient interface {
	CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*CreateEventResponse, error)
}

type calendarServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalendarServiceClient(cc *grpc.ClientConn) CalendarServiceClient {
	return &calendarServiceClient{cc}
}

func (c *calendarServiceClient) CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*CreateEventResponse, error) {
	out := new(CreateEventResponse)
	err := c.cc.Invoke(ctx, "/CalendarService/CreateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalendarServiceServer is the server API for CalendarService service.
type CalendarServiceServer interface {
	CreateEvent(context.Context, *CreateEventRequest) (*CreateEventResponse, error)
}

// UnimplementedCalendarServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalendarServiceServer struct {
}

func (*UnimplementedCalendarServiceServer) CreateEvent(ctx context.Context, req *CreateEventRequest) (*CreateEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}

func RegisterCalendarServiceServer(s *grpc.Server, srv CalendarServiceServer) {
	s.RegisterService(&_CalendarService_serviceDesc, srv)
}

func _CalendarService_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CalendarService/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).CreateEvent(ctx, req.(*CreateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalendarService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CalendarService",
	HandlerType: (*CalendarServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEvent",
			Handler:    _CalendarService_CreateEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}