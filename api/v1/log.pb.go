// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/v1/log.proto

package log_v1

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/proto"
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

type Record struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Offset               uint64   `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Term                 uint64   `protobuf:"varint,3,opt,name=term,proto3" json:"term,omitempty"`
	Type                 uint32   `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_19a5c3fde3f7ae80, []int{0}
}

func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Record) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Record) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *Record) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type Server struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RpcAddr              string   `protobuf:"bytes,2,opt,name=rpc_addr,json=rpcAddr,proto3" json:"rpc_addr,omitempty"`
	IsLeader             bool     `protobuf:"varint,3,opt,name=is_leader,json=isLeader,proto3" json:"is_leader,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Server) Reset()         { *m = Server{} }
func (m *Server) String() string { return proto.CompactTextString(m) }
func (*Server) ProtoMessage()    {}
func (*Server) Descriptor() ([]byte, []int) {
	return fileDescriptor_19a5c3fde3f7ae80, []int{1}
}

func (m *Server) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Server.Unmarshal(m, b)
}
func (m *Server) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Server.Marshal(b, m, deterministic)
}
func (m *Server) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Server.Merge(m, src)
}
func (m *Server) XXX_Size() int {
	return xxx_messageInfo_Server.Size(m)
}
func (m *Server) XXX_DiscardUnknown() {
	xxx_messageInfo_Server.DiscardUnknown(m)
}

var xxx_messageInfo_Server proto.InternalMessageInfo

func (m *Server) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Server) GetRpcAddr() string {
	if m != nil {
		return m.RpcAddr
	}
	return ""
}

func (m *Server) GetIsLeader() bool {
	if m != nil {
		return m.IsLeader
	}
	return false
}

type ProduceRequest struct {
	Record               *Record  `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProduceRequest) Reset()         { *m = ProduceRequest{} }
func (m *ProduceRequest) String() string { return proto.CompactTextString(m) }
func (*ProduceRequest) ProtoMessage()    {}
func (*ProduceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_19a5c3fde3f7ae80, []int{2}
}

func (m *ProduceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProduceRequest.Unmarshal(m, b)
}
func (m *ProduceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProduceRequest.Marshal(b, m, deterministic)
}
func (m *ProduceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProduceRequest.Merge(m, src)
}
func (m *ProduceRequest) XXX_Size() int {
	return xxx_messageInfo_ProduceRequest.Size(m)
}
func (m *ProduceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProduceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProduceRequest proto.InternalMessageInfo

func (m *ProduceRequest) GetRecord() *Record {
	if m != nil {
		return m.Record
	}
	return nil
}

type ProduceResponse struct {
	Offset               uint64   `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProduceResponse) Reset()         { *m = ProduceResponse{} }
func (m *ProduceResponse) String() string { return proto.CompactTextString(m) }
func (*ProduceResponse) ProtoMessage()    {}
func (*ProduceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_19a5c3fde3f7ae80, []int{3}
}

func (m *ProduceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProduceResponse.Unmarshal(m, b)
}
func (m *ProduceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProduceResponse.Marshal(b, m, deterministic)
}
func (m *ProduceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProduceResponse.Merge(m, src)
}
func (m *ProduceResponse) XXX_Size() int {
	return xxx_messageInfo_ProduceResponse.Size(m)
}
func (m *ProduceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProduceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProduceResponse proto.InternalMessageInfo

func (m *ProduceResponse) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ConsumeRequest struct {
	Offset               uint64   `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsumeRequest) Reset()         { *m = ConsumeRequest{} }
func (m *ConsumeRequest) String() string { return proto.CompactTextString(m) }
func (*ConsumeRequest) ProtoMessage()    {}
func (*ConsumeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_19a5c3fde3f7ae80, []int{4}
}

func (m *ConsumeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsumeRequest.Unmarshal(m, b)
}
func (m *ConsumeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsumeRequest.Marshal(b, m, deterministic)
}
func (m *ConsumeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsumeRequest.Merge(m, src)
}
func (m *ConsumeRequest) XXX_Size() int {
	return xxx_messageInfo_ConsumeRequest.Size(m)
}
func (m *ConsumeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsumeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConsumeRequest proto.InternalMessageInfo

func (m *ConsumeRequest) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ConsumeResponse struct {
	Record               *Record  `protobuf:"bytes,2,opt,name=record,proto3" json:"record,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsumeResponse) Reset()         { *m = ConsumeResponse{} }
func (m *ConsumeResponse) String() string { return proto.CompactTextString(m) }
func (*ConsumeResponse) ProtoMessage()    {}
func (*ConsumeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_19a5c3fde3f7ae80, []int{5}
}

func (m *ConsumeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsumeResponse.Unmarshal(m, b)
}
func (m *ConsumeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsumeResponse.Marshal(b, m, deterministic)
}
func (m *ConsumeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsumeResponse.Merge(m, src)
}
func (m *ConsumeResponse) XXX_Size() int {
	return xxx_messageInfo_ConsumeResponse.Size(m)
}
func (m *ConsumeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsumeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConsumeResponse proto.InternalMessageInfo

func (m *ConsumeResponse) GetRecord() *Record {
	if m != nil {
		return m.Record
	}
	return nil
}

type GetServersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServersRequest) Reset()         { *m = GetServersRequest{} }
func (m *GetServersRequest) String() string { return proto.CompactTextString(m) }
func (*GetServersRequest) ProtoMessage()    {}
func (*GetServersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_19a5c3fde3f7ae80, []int{6}
}

func (m *GetServersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServersRequest.Unmarshal(m, b)
}
func (m *GetServersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServersRequest.Marshal(b, m, deterministic)
}
func (m *GetServersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServersRequest.Merge(m, src)
}
func (m *GetServersRequest) XXX_Size() int {
	return xxx_messageInfo_GetServersRequest.Size(m)
}
func (m *GetServersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetServersRequest proto.InternalMessageInfo

type GetServersResponse struct {
	Servers              []*Server `protobuf:"bytes,1,rep,name=servers,proto3" json:"servers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetServersResponse) Reset()         { *m = GetServersResponse{} }
func (m *GetServersResponse) String() string { return proto.CompactTextString(m) }
func (*GetServersResponse) ProtoMessage()    {}
func (*GetServersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_19a5c3fde3f7ae80, []int{7}
}

func (m *GetServersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServersResponse.Unmarshal(m, b)
}
func (m *GetServersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServersResponse.Marshal(b, m, deterministic)
}
func (m *GetServersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServersResponse.Merge(m, src)
}
func (m *GetServersResponse) XXX_Size() int {
	return xxx_messageInfo_GetServersResponse.Size(m)
}
func (m *GetServersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetServersResponse proto.InternalMessageInfo

func (m *GetServersResponse) GetServers() []*Server {
	if m != nil {
		return m.Servers
	}
	return nil
}

func init() {
	proto.RegisterType((*Record)(nil), "log.v1.Record")
	proto.RegisterType((*Server)(nil), "log.v1.Server")
	proto.RegisterType((*ProduceRequest)(nil), "log.v1.ProduceRequest")
	proto.RegisterType((*ProduceResponse)(nil), "log.v1.ProduceResponse")
	proto.RegisterType((*ConsumeRequest)(nil), "log.v1.ConsumeRequest")
	proto.RegisterType((*ConsumeResponse)(nil), "log.v1.ConsumeResponse")
	proto.RegisterType((*GetServersRequest)(nil), "log.v1.GetServersRequest")
	proto.RegisterType((*GetServersResponse)(nil), "log.v1.GetServersResponse")
}

func init() {
	proto.RegisterFile("api/v1/log.proto", fileDescriptor_19a5c3fde3f7ae80)
}

var fileDescriptor_19a5c3fde3f7ae80 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xed, 0x3a, 0xc1, 0x49, 0x86, 0xc6, 0x85, 0xa1, 0x2a, 0xae, 0x91, 0x22, 0x6b, 0x0f, 0xc8,
	0x5c, 0x92, 0xb6, 0x5c, 0x40, 0x42, 0x48, 0x7c, 0x5f, 0x7a, 0xa8, 0xb6, 0x77, 0x2a, 0xe3, 0xdd,
	0x5a, 0x96, 0x92, 0xae, 0xd9, 0xb5, 0x2d, 0xf1, 0x0f, 0x39, 0x72, 0xe2, 0x8c, 0xf2, 0x4b, 0x50,
	0x76, 0xd7, 0xb1, 0xdb, 0x02, 0x42, 0xbd, 0xcd, 0xbe, 0x79, 0xf3, 0x66, 0xde, 0x93, 0x16, 0x1e,
	0xa4, 0x65, 0xb1, 0x68, 0x8e, 0x17, 0x4b, 0x99, 0xcf, 0x4b, 0x25, 0x2b, 0x89, 0xfe, 0xa6, 0x6c,
	0x8e, 0xa3, 0xfd, 0x5c, 0xe6, 0xd2, 0x40, 0x8b, 0x4d, 0x65, 0xbb, 0xf4, 0x33, 0xf8, 0x4c, 0x64,
	0x52, 0x71, 0xdc, 0x87, 0x7b, 0x4d, 0xba, 0xac, 0x45, 0x48, 0x62, 0x92, 0xec, 0x32, 0xfb, 0xc0,
	0x03, 0xf0, 0xe5, 0xe5, 0xa5, 0x16, 0x55, 0xe8, 0xc5, 0x24, 0x19, 0x32, 0xf7, 0x42, 0x84, 0x61,
	0x25, 0xd4, 0x2a, 0x1c, 0x18, 0xd4, 0xd4, 0x06, 0xfb, 0x56, 0x8a, 0x70, 0x18, 0x93, 0x64, 0xca,
	0x4c, 0x4d, 0xcf, 0xc0, 0x3f, 0x17, 0xaa, 0x11, 0x0a, 0x03, 0xf0, 0x0a, 0x6e, 0xc4, 0x27, 0xcc,
	0x2b, 0x38, 0x1e, 0xc2, 0x58, 0x95, 0xd9, 0x45, 0xca, 0xb9, 0x32, 0xda, 0x13, 0x36, 0x52, 0x65,
	0xf6, 0x86, 0x73, 0x85, 0x4f, 0x60, 0x52, 0xe8, 0x8b, 0xa5, 0x48, 0xb9, 0x50, 0x66, 0xc3, 0x98,
	0x8d, 0x0b, 0x7d, 0x6a, 0xde, 0xf4, 0x05, 0x04, 0x67, 0x4a, 0xf2, 0x3a, 0x13, 0x4c, 0x7c, 0xad,
	0x85, 0xae, 0xf0, 0x29, 0xf8, 0xca, 0x78, 0x30, 0xea, 0xf7, 0x4f, 0x82, 0xb9, 0xb5, 0x3c, 0xb7,
	0xce, 0x98, 0xeb, 0xd2, 0x67, 0xb0, 0xb7, 0x9d, 0xd4, 0xa5, 0xbc, 0xd2, 0x7d, 0x7b, 0xa4, 0x6f,
	0x8f, 0x26, 0x10, 0xbc, 0x93, 0x57, 0xba, 0x5e, 0x6d, 0x97, 0xfc, 0x8d, 0xf9, 0x12, 0xf6, 0xb6,
	0x4c, 0x27, 0xda, 0xdd, 0xe3, 0xfd, 0xf3, 0x9e, 0x47, 0xf0, 0xf0, 0x93, 0xa8, 0x6c, 0x3c, 0xda,
	0xed, 0xa1, 0xaf, 0x01, 0xfb, 0xa0, 0x93, 0x4c, 0x60, 0xa4, 0x2d, 0x14, 0x92, 0x78, 0xd0, 0xd7,
	0xb4, 0x4c, 0xd6, 0xb6, 0x4f, 0x7e, 0x7a, 0x30, 0x38, 0x95, 0x39, 0xbe, 0x82, 0x91, 0x33, 0x8b,
	0x07, 0x2d, 0xf7, 0x7a, 0x6e, 0xd1, 0xe3, 0x5b, 0xb8, 0xdd, 0x46, 0x77, 0x36, 0xd3, 0xce, 0x55,
	0x37, 0x7d, 0x3d, 0x90, 0x6e, 0xfa, 0x86, 0x7d, 0xba, 0x83, 0xef, 0x61, 0xea, 0xc0, 0xf3, 0x4a,
	0x89, 0x74, 0x75, 0x07, 0x8d, 0x23, 0x82, 0x1f, 0x61, 0xea, 0x0e, 0xbb, 0xa9, 0xf2, 0xdf, 0x3e,
	0x12, 0x72, 0x44, 0xf0, 0x03, 0x40, 0x97, 0x28, 0x1e, 0xb6, 0xe4, 0x5b, 0xd1, 0x47, 0xd1, 0x9f,
	0x5a, 0xad, 0xd4, 0xdb, 0xdd, 0xef, 0xeb, 0x19, 0xf9, 0xb1, 0x9e, 0x91, 0x5f, 0xeb, 0x19, 0xf9,
	0xe2, 0x9b, 0xef, 0xf3, 0xfc, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xa9, 0x8c, 0xf4, 0x70,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LogClient is the client API for Log service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogClient interface {
	Produce(ctx context.Context, in *ProduceRequest, opts ...grpc.CallOption) (*ProduceResponse, error)
	Consume(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (*ConsumeResponse, error)
	ConsumeStream(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (Log_ConsumeStreamClient, error)
	ProduceStream(ctx context.Context, opts ...grpc.CallOption) (Log_ProduceStreamClient, error)
	GetServers(ctx context.Context, in *GetServersRequest, opts ...grpc.CallOption) (*GetServersResponse, error)
}

type logClient struct {
	cc grpc.ClientConnInterface
}

func NewLogClient(cc grpc.ClientConnInterface) LogClient {
	return &logClient{cc}
}

func (c *logClient) Produce(ctx context.Context, in *ProduceRequest, opts ...grpc.CallOption) (*ProduceResponse, error) {
	out := new(ProduceResponse)
	err := c.cc.Invoke(ctx, "/log.v1.Log/Produce", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logClient) Consume(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (*ConsumeResponse, error) {
	out := new(ConsumeResponse)
	err := c.cc.Invoke(ctx, "/log.v1.Log/Consume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logClient) ConsumeStream(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (Log_ConsumeStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Log_serviceDesc.Streams[0], "/log.v1.Log/ConsumeStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &logConsumeStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Log_ConsumeStreamClient interface {
	Recv() (*ConsumeResponse, error)
	grpc.ClientStream
}

type logConsumeStreamClient struct {
	grpc.ClientStream
}

func (x *logConsumeStreamClient) Recv() (*ConsumeResponse, error) {
	m := new(ConsumeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *logClient) ProduceStream(ctx context.Context, opts ...grpc.CallOption) (Log_ProduceStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Log_serviceDesc.Streams[1], "/log.v1.Log/ProduceStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &logProduceStreamClient{stream}
	return x, nil
}

type Log_ProduceStreamClient interface {
	Send(*ProduceRequest) error
	Recv() (*ProduceResponse, error)
	grpc.ClientStream
}

type logProduceStreamClient struct {
	grpc.ClientStream
}

func (x *logProduceStreamClient) Send(m *ProduceRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *logProduceStreamClient) Recv() (*ProduceResponse, error) {
	m := new(ProduceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *logClient) GetServers(ctx context.Context, in *GetServersRequest, opts ...grpc.CallOption) (*GetServersResponse, error) {
	out := new(GetServersResponse)
	err := c.cc.Invoke(ctx, "/log.v1.Log/GetServers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogServer is the server API for Log service.
type LogServer interface {
	Produce(context.Context, *ProduceRequest) (*ProduceResponse, error)
	Consume(context.Context, *ConsumeRequest) (*ConsumeResponse, error)
	ConsumeStream(*ConsumeRequest, Log_ConsumeStreamServer) error
	ProduceStream(Log_ProduceStreamServer) error
	GetServers(context.Context, *GetServersRequest) (*GetServersResponse, error)
}

// UnimplementedLogServer can be embedded to have forward compatible implementations.
type UnimplementedLogServer struct {
}

func (*UnimplementedLogServer) Produce(ctx context.Context, req *ProduceRequest) (*ProduceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Produce not implemented")
}
func (*UnimplementedLogServer) Consume(ctx context.Context, req *ConsumeRequest) (*ConsumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Consume not implemented")
}
func (*UnimplementedLogServer) ConsumeStream(req *ConsumeRequest, srv Log_ConsumeStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ConsumeStream not implemented")
}
func (*UnimplementedLogServer) ProduceStream(srv Log_ProduceStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ProduceStream not implemented")
}
func (*UnimplementedLogServer) GetServers(ctx context.Context, req *GetServersRequest) (*GetServersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServers not implemented")
}

func RegisterLogServer(s *grpc.Server, srv LogServer) {
	s.RegisterService(&_Log_serviceDesc, srv)
}

func _Log_Produce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProduceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServer).Produce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/log.v1.Log/Produce",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServer).Produce(ctx, req.(*ProduceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Log_Consume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServer).Consume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/log.v1.Log/Consume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServer).Consume(ctx, req.(*ConsumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Log_ConsumeStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConsumeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogServer).ConsumeStream(m, &logConsumeStreamServer{stream})
}

type Log_ConsumeStreamServer interface {
	Send(*ConsumeResponse) error
	grpc.ServerStream
}

type logConsumeStreamServer struct {
	grpc.ServerStream
}

func (x *logConsumeStreamServer) Send(m *ConsumeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Log_ProduceStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LogServer).ProduceStream(&logProduceStreamServer{stream})
}

type Log_ProduceStreamServer interface {
	Send(*ProduceResponse) error
	Recv() (*ProduceRequest, error)
	grpc.ServerStream
}

type logProduceStreamServer struct {
	grpc.ServerStream
}

func (x *logProduceStreamServer) Send(m *ProduceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *logProduceStreamServer) Recv() (*ProduceRequest, error) {
	m := new(ProduceRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Log_GetServers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServer).GetServers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/log.v1.Log/GetServers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServer).GetServers(ctx, req.(*GetServersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Log_serviceDesc = grpc.ServiceDesc{
	ServiceName: "log.v1.Log",
	HandlerType: (*LogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Produce",
			Handler:    _Log_Produce_Handler,
		},
		{
			MethodName: "Consume",
			Handler:    _Log_Consume_Handler,
		},
		{
			MethodName: "GetServers",
			Handler:    _Log_GetServers_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConsumeStream",
			Handler:       _Log_ConsumeStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ProduceStream",
			Handler:       _Log_ProduceStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/v1/log.proto",
}
