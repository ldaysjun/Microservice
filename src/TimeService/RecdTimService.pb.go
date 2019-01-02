// Code generated by protoc-gen-go. DO NOT EDIT.
// source: RecdTimService.proto

// 定义你的包名

package echo

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type TimedTaskRequest struct {
	Start                int32    `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	End                  int32    `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	Restall              int32    `protobuf:"varint,3,opt,name=restall,proto3" json:"restall,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimedTaskRequest) Reset()         { *m = TimedTaskRequest{} }
func (m *TimedTaskRequest) String() string { return proto.CompactTextString(m) }
func (*TimedTaskRequest) ProtoMessage()    {}
func (*TimedTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_295c14269c9257d7, []int{0}
}

func (m *TimedTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimedTaskRequest.Unmarshal(m, b)
}
func (m *TimedTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimedTaskRequest.Marshal(b, m, deterministic)
}
func (m *TimedTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimedTaskRequest.Merge(m, src)
}
func (m *TimedTaskRequest) XXX_Size() int {
	return xxx_messageInfo_TimedTaskRequest.Size(m)
}
func (m *TimedTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TimedTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TimedTaskRequest proto.InternalMessageInfo

func (m *TimedTaskRequest) GetStart() int32 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *TimedTaskRequest) GetEnd() int32 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *TimedTaskRequest) GetRestall() int32 {
	if m != nil {
		return m.Restall
	}
	return 0
}

type TimedTaskResponse struct {
	ErrCode              int32    `protobuf:"varint,1,opt,name=err_code,json=errCode,proto3" json:"err_code,omitempty"`
	EreMsg               string   `protobuf:"bytes,2,opt,name=ere_msg,json=ereMsg,proto3" json:"ere_msg,omitempty"`
	ExpData              string   `protobuf:"bytes,3,opt,name=exp_data,json=expData,proto3" json:"exp_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimedTaskResponse) Reset()         { *m = TimedTaskResponse{} }
func (m *TimedTaskResponse) String() string { return proto.CompactTextString(m) }
func (*TimedTaskResponse) ProtoMessage()    {}
func (*TimedTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_295c14269c9257d7, []int{1}
}

func (m *TimedTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimedTaskResponse.Unmarshal(m, b)
}
func (m *TimedTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimedTaskResponse.Marshal(b, m, deterministic)
}
func (m *TimedTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimedTaskResponse.Merge(m, src)
}
func (m *TimedTaskResponse) XXX_Size() int {
	return xxx_messageInfo_TimedTaskResponse.Size(m)
}
func (m *TimedTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TimedTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TimedTaskResponse proto.InternalMessageInfo

func (m *TimedTaskResponse) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *TimedTaskResponse) GetEreMsg() string {
	if m != nil {
		return m.EreMsg
	}
	return ""
}

func (m *TimedTaskResponse) GetExpData() string {
	if m != nil {
		return m.ExpData
	}
	return ""
}

func init() {
	proto.RegisterType((*TimedTaskRequest)(nil), "echo.TimedTaskRequest")
	proto.RegisterType((*TimedTaskResponse)(nil), "echo.TimedTaskResponse")
}

func init() { proto.RegisterFile("RecdTimService.proto", fileDescriptor_295c14269c9257d7) }

var fileDescriptor_295c14269c9257d7 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x50, 0x41, 0x4e, 0xc3, 0x40,
	0x0c, 0x54, 0x5b, 0xda, 0x10, 0x8b, 0x43, 0x59, 0x55, 0x34, 0x54, 0x1c, 0x50, 0x4e, 0x9c, 0x12,
	0x09, 0x9e, 0x00, 0x57, 0x24, 0x14, 0x72, 0x26, 0xda, 0x66, 0x47, 0x21, 0x22, 0xc9, 0x06, 0xaf,
	0x41, 0x3d, 0xf3, 0x05, 0x9e, 0xc6, 0x17, 0x78, 0x08, 0xea, 0x86, 0x02, 0xa2, 0x37, 0xcf, 0xd8,
	0x9e, 0xf1, 0x98, 0x16, 0x19, 0x4a, 0x93, 0xd7, 0xed, 0x3d, 0xf8, 0xb5, 0x2e, 0x91, 0xf4, 0x6c,
	0xc5, 0xaa, 0x03, 0x94, 0x8f, 0x76, 0x75, 0x56, 0x59, 0x5b, 0x35, 0x48, 0x75, 0x5f, 0xa7, 0xba,
	0xeb, 0xac, 0x68, 0xa9, 0x6d, 0xe7, 0x86, 0x99, 0x38, 0xa7, 0x79, 0x5e, 0xb7, 0x30, 0xb9, 0x76,
	0x4f, 0x19, 0x9e, 0x5f, 0xe0, 0x44, 0x2d, 0x68, 0xea, 0x44, 0xb3, 0x44, 0xa3, 0xf3, 0xd1, 0xc5,
	0x34, 0x1b, 0x80, 0x9a, 0xd3, 0x04, 0x9d, 0x89, 0xc6, 0x9e, 0xdb, 0x96, 0x2a, 0xa2, 0x80, 0xe1,
	0x44, 0x37, 0x4d, 0x34, 0xf1, 0xec, 0x0e, 0xc6, 0x6b, 0x3a, 0xfe, 0xa3, 0xea, 0x7a, 0xdb, 0x39,
	0xa8, 0x53, 0x3a, 0x04, 0x73, 0x51, 0x5a, 0x83, 0x6f, 0xe5, 0x00, 0xcc, 0xd7, 0xd6, 0x40, 0x2d,
	0x29, 0x00, 0xa3, 0x68, 0x5d, 0xe5, 0xf5, 0xc3, 0x6c, 0x06, 0xc6, 0xad, 0xab, 0xfc, 0xce, 0xa6,
	0x2f, 0x8c, 0x16, 0xed, 0x3d, 0xc2, 0x2c, 0xc0, 0xa6, 0xbf, 0xd1, 0xa2, 0x2f, 0x1f, 0x88, 0x7e,
	0x13, 0xab, 0x3b, 0x0a, 0x7f, 0x1c, 0xd5, 0x49, 0xb2, 0x4d, 0x9e, 0xfc, 0x0f, 0xb6, 0x5a, 0xee,
	0xf1, 0xc3, 0x69, 0xb1, 0x7a, 0xfb, 0xf8, 0x7c, 0x1f, 0x1f, 0x29, 0x4a, 0x65, 0xd7, 0x5b, 0xcf,
	0xfc, 0x83, 0xae, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x34, 0x60, 0x8a, 0x5c, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TimServiceClient is the client API for TimService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TimServiceClient interface {
	// 定义一个方法Echo,输入 TimedTaskRequest ,输出 TimedTaskResponse
	TimedTask(ctx context.Context, in *TimedTaskRequest, opts ...grpc.CallOption) (*TimedTaskResponse, error)
}

type timServiceClient struct {
	cc *grpc.ClientConn
}

func NewTimServiceClient(cc *grpc.ClientConn) TimServiceClient {
	return &timServiceClient{cc}
}

func (c *timServiceClient) TimedTask(ctx context.Context, in *TimedTaskRequest, opts ...grpc.CallOption) (*TimedTaskResponse, error) {
	out := new(TimedTaskResponse)
	err := c.cc.Invoke(ctx, "/echo.TimService/TimedTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimServiceServer is the server API for TimService service.
type TimServiceServer interface {
	// 定义一个方法Echo,输入 TimedTaskRequest ,输出 TimedTaskResponse
	TimedTask(context.Context, *TimedTaskRequest) (*TimedTaskResponse, error)
}

func RegisterTimServiceServer(s *grpc.Server, srv TimServiceServer) {
	s.RegisterService(&_TimService_serviceDesc, srv)
}

func _TimService_TimedTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimedTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimServiceServer).TimedTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.TimService/TimedTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimServiceServer).TimedTask(ctx, req.(*TimedTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TimService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "echo.TimService",
	HandlerType: (*TimServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TimedTask",
			Handler:    _TimService_TimedTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "RecdTimService.proto",
}
