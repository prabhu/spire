// Code generated by protoc-gen-go. DO NOT EDIT.
// source: plugin.proto

/*
Package sriplugin is a generated protocol buffer package.

It is generated from these files:
	plugin.proto

It has these top-level messages:
	ConfigureRequest
	ConfigureResponse
	GetPluginInfoRequest
	GetPluginInfoResponse
	PluginInfoRequest
	PluginInfoReply
	StopRequest
	StopReply
*/
package sriplugin

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// * Represents the plugin-specific configuration string.
type ConfigureRequest struct {
	Configuration string `protobuf:"bytes,1,opt,name=configuration" json:"configuration,omitempty"`
}

func (m *ConfigureRequest) Reset()                    { *m = ConfigureRequest{} }
func (m *ConfigureRequest) String() string            { return proto.CompactTextString(m) }
func (*ConfigureRequest) ProtoMessage()               {}
func (*ConfigureRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ConfigureRequest) GetConfiguration() string {
	if m != nil {
		return m.Configuration
	}
	return ""
}

// * Represents a list of configuration problems found in the configuration string.
type ConfigureResponse struct {
	ErrorList []string `protobuf:"bytes,1,rep,name=errorList" json:"errorList,omitempty"`
}

func (m *ConfigureResponse) Reset()                    { *m = ConfigureResponse{} }
func (m *ConfigureResponse) String() string            { return proto.CompactTextString(m) }
func (*ConfigureResponse) ProtoMessage()               {}
func (*ConfigureResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ConfigureResponse) GetErrorList() []string {
	if m != nil {
		return m.ErrorList
	}
	return nil
}

// * Represents an empty request.
type GetPluginInfoRequest struct {
}

func (m *GetPluginInfoRequest) Reset()                    { *m = GetPluginInfoRequest{} }
func (m *GetPluginInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*GetPluginInfoRequest) ProtoMessage()               {}
func (*GetPluginInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// * Represents the plugin metadata.
type GetPluginInfoResponse struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Category    string `protobuf:"bytes,2,opt,name=category" json:"category,omitempty"`
	Type        string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	DateCreated string `protobuf:"bytes,5,opt,name=dateCreated" json:"dateCreated,omitempty"`
	Location    string `protobuf:"bytes,6,opt,name=location" json:"location,omitempty"`
	Version     string `protobuf:"bytes,7,opt,name=version" json:"version,omitempty"`
	Author      string `protobuf:"bytes,8,opt,name=author" json:"author,omitempty"`
	Company     string `protobuf:"bytes,9,opt,name=company" json:"company,omitempty"`
}

func (m *GetPluginInfoResponse) Reset()                    { *m = GetPluginInfoResponse{} }
func (m *GetPluginInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*GetPluginInfoResponse) ProtoMessage()               {}
func (*GetPluginInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetPluginInfoResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetPluginInfoResponse) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *GetPluginInfoResponse) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GetPluginInfoResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *GetPluginInfoResponse) GetDateCreated() string {
	if m != nil {
		return m.DateCreated
	}
	return ""
}

func (m *GetPluginInfoResponse) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *GetPluginInfoResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *GetPluginInfoResponse) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *GetPluginInfoResponse) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

type PluginInfoRequest struct {
}

func (m *PluginInfoRequest) Reset()                    { *m = PluginInfoRequest{} }
func (m *PluginInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*PluginInfoRequest) ProtoMessage()               {}
func (*PluginInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type PluginInfoReply struct {
	PluginInfo []*GetPluginInfoResponse `protobuf:"bytes,1,rep,name=pluginInfo" json:"pluginInfo,omitempty"`
}

func (m *PluginInfoReply) Reset()                    { *m = PluginInfoReply{} }
func (m *PluginInfoReply) String() string            { return proto.CompactTextString(m) }
func (*PluginInfoReply) ProtoMessage()               {}
func (*PluginInfoReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PluginInfoReply) GetPluginInfo() []*GetPluginInfoResponse {
	if m != nil {
		return m.PluginInfo
	}
	return nil
}

type StopRequest struct {
}

func (m *StopRequest) Reset()                    { *m = StopRequest{} }
func (m *StopRequest) String() string            { return proto.CompactTextString(m) }
func (*StopRequest) ProtoMessage()               {}
func (*StopRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type StopReply struct {
}

func (m *StopReply) Reset()                    { *m = StopReply{} }
func (m *StopReply) String() string            { return proto.CompactTextString(m) }
func (*StopReply) ProtoMessage()               {}
func (*StopReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func init() {
	proto.RegisterType((*ConfigureRequest)(nil), "spire.common.plugin.ConfigureRequest")
	proto.RegisterType((*ConfigureResponse)(nil), "spire.common.plugin.ConfigureResponse")
	proto.RegisterType((*GetPluginInfoRequest)(nil), "spire.common.plugin.GetPluginInfoRequest")
	proto.RegisterType((*GetPluginInfoResponse)(nil), "spire.common.plugin.GetPluginInfoResponse")
	proto.RegisterType((*PluginInfoRequest)(nil), "spire.common.plugin.PluginInfoRequest")
	proto.RegisterType((*PluginInfoReply)(nil), "spire.common.plugin.PluginInfoReply")
	proto.RegisterType((*StopRequest)(nil), "spire.common.plugin.StopRequest")
	proto.RegisterType((*StopReply)(nil), "spire.common.plugin.StopReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Server service

type ServerClient interface {
	Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopReply, error)
	PluginInfo(ctx context.Context, in *PluginInfoRequest, opts ...grpc.CallOption) (*PluginInfoReply, error)
}

type serverClient struct {
	cc *grpc.ClientConn
}

func NewServerClient(cc *grpc.ClientConn) ServerClient {
	return &serverClient{cc}
}

func (c *serverClient) Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopReply, error) {
	out := new(StopReply)
	err := grpc.Invoke(ctx, "/spire.common.plugin.Server/Stop", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) PluginInfo(ctx context.Context, in *PluginInfoRequest, opts ...grpc.CallOption) (*PluginInfoReply, error) {
	out := new(PluginInfoReply)
	err := grpc.Invoke(ctx, "/spire.common.plugin.Server/PluginInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Server service

type ServerServer interface {
	Stop(context.Context, *StopRequest) (*StopReply, error)
	PluginInfo(context.Context, *PluginInfoRequest) (*PluginInfoReply, error)
}

func RegisterServerServer(s *grpc.Server, srv ServerServer) {
	s.RegisterService(&_Server_serviceDesc, srv)
}

func _Server_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.common.plugin.Server/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).Stop(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_PluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).PluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.common.plugin.Server/PluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).PluginInfo(ctx, req.(*PluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Server_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spire.common.plugin.Server",
	HandlerType: (*ServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Stop",
			Handler:    _Server_Stop_Handler,
		},
		{
			MethodName: "PluginInfo",
			Handler:    _Server_PluginInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plugin.proto",
}

func init() { proto.RegisterFile("plugin.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x6f, 0xda, 0x30,
	0x18, 0x56, 0x80, 0x05, 0xf2, 0x66, 0x68, 0xc3, 0x6c, 0xc8, 0x42, 0xd3, 0x14, 0x45, 0x68, 0x42,
	0x3b, 0x44, 0x1a, 0xbb, 0xec, 0x3c, 0x0e, 0xfd, 0x50, 0x0f, 0x55, 0xb8, 0x21, 0xf5, 0x90, 0x86,
	0x17, 0x1a, 0x29, 0xb1, 0x5d, 0xdb, 0x41, 0xca, 0x2f, 0xe8, 0xbf, 0xe9, 0x6f, 0xac, 0xe2, 0x04,
	0x48, 0x4b, 0xd4, 0xde, 0xfc, 0x7c, 0xbc, 0x7e, 0xde, 0x3c, 0x31, 0x7c, 0x16, 0x69, 0xbe, 0x4b,
	0x58, 0x20, 0x24, 0xd7, 0x9c, 0x8c, 0x95, 0x48, 0x24, 0x06, 0x31, 0xcf, 0x32, 0xce, 0x82, 0x4a,
	0xf2, 0xff, 0xc1, 0xd7, 0x25, 0x67, 0xdb, 0x64, 0x97, 0x4b, 0x0c, 0xf1, 0x31, 0x47, 0xa5, 0xc9,
	0x0c, 0x86, 0x71, 0xcd, 0x45, 0x3a, 0xe1, 0x8c, 0x5a, 0x9e, 0x35, 0x77, 0xc2, 0xd7, 0xa4, 0xff,
	0x07, 0x46, 0x8d, 0x49, 0x25, 0x38, 0x53, 0x48, 0x7e, 0x80, 0x83, 0x52, 0x72, 0x79, 0x93, 0x28,
	0x4d, 0x2d, 0xaf, 0x3b, 0x77, 0xc2, 0x13, 0xe1, 0x4f, 0xe0, 0xdb, 0x05, 0xea, 0x5b, 0x93, 0x7c,
	0xc5, 0xb6, 0xbc, 0x0e, 0xf4, 0x9f, 0x3a, 0xf0, 0xfd, 0x8d, 0x50, 0xdf, 0x47, 0xa0, 0xc7, 0xa2,
	0x0c, 0xeb, 0x0d, 0xcc, 0x99, 0x4c, 0x61, 0x10, 0x47, 0x1a, 0x77, 0x5c, 0x16, 0xb4, 0x63, 0xf8,
	0x23, 0x2e, 0xfd, 0xba, 0x10, 0x48, 0xbb, 0x95, 0xbf, 0x3c, 0x13, 0x0f, 0xdc, 0x0d, 0xaa, 0x58,
	0x26, 0xc2, 0x7c, 0x4c, 0xcf, 0x48, 0x4d, 0xca, 0x38, 0x22, 0x8d, 0x4b, 0x89, 0x91, 0xc6, 0x0d,
	0xfd, 0x54, 0x3b, 0x4e, 0x54, 0x99, 0x99, 0xf2, 0xb8, 0x6a, 0xc3, 0xae, 0x32, 0x0f, 0x98, 0x50,
	0xe8, 0xef, 0x51, 0xaa, 0x52, 0xea, 0x1b, 0xe9, 0x00, 0xc9, 0x04, 0xec, 0x28, 0xd7, 0x0f, 0x5c,
	0xd2, 0x81, 0x11, 0x6a, 0x54, 0x4e, 0xc4, 0x3c, 0x13, 0x11, 0x2b, 0xa8, 0x53, 0x4d, 0xd4, 0xd0,
	0x1f, 0xc3, 0xe8, 0xbc, 0x9e, 0x3b, 0xf8, 0xd2, 0x24, 0x45, 0x5a, 0x90, 0x6b, 0x00, 0x71, 0xa4,
	0x4c, 0xd1, 0xee, 0xe2, 0x77, 0xd0, 0xf2, 0x83, 0x83, 0xd6, 0x5e, 0xc3, 0xc6, 0xb4, 0x3f, 0x04,
	0x77, 0xa5, 0xb9, 0x38, 0xa4, 0xb9, 0xe0, 0x54, 0x50, 0xa4, 0xc5, 0xe2, 0xd9, 0x02, 0x7b, 0x85,
	0x72, 0x8f, 0x92, 0x5c, 0x42, 0xaf, 0xe4, 0x89, 0xd7, 0x1a, 0xd3, 0xb8, 0x61, 0xfa, 0xf3, 0x1d,
	0x47, 0xb9, 0xfc, 0x1a, 0xe0, 0xb4, 0x12, 0xf9, 0xd5, 0xea, 0x3e, 0x6b, 0x61, 0x3a, 0xfb, 0xd0,
	0x27, 0xd2, 0xe2, 0xbf, 0xbb, 0x76, 0x94, 0x4c, 0x2a, 0xf1, 0xde, 0x36, 0x0f, 0xff, 0xef, 0x4b,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x91, 0x27, 0xa0, 0x08, 0x03, 0x00, 0x00,
}
