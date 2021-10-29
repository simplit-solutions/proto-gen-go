// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.14.0
// source: simpli-echo-me/simpliEchoMe.proto

package simpli_echo_me

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SimpliEchoMeServiceClient is the client API for SimpliEchoMeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SimpliEchoMeServiceClient interface {
	// Say Hi.
	SayHi(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SayHiResponse, error)
	// Say Hi echoing my name.
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
	// Remember my name.
	EchoRemember(ctx context.Context, in *EchoRememberRequest, opts ...grpc.CallOption) (*EchoRememberResponse, error)
	// Say Hi given my remembered id.
	SayHiRemember(ctx context.Context, in *SayHiRememberRequest, opts ...grpc.CallOption) (*SayHiRememberResponse, error)
}

type simpliEchoMeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSimpliEchoMeServiceClient(cc grpc.ClientConnInterface) SimpliEchoMeServiceClient {
	return &simpliEchoMeServiceClient{cc}
}

func (c *simpliEchoMeServiceClient) SayHi(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SayHiResponse, error) {
	out := new(SayHiResponse)
	err := c.cc.Invoke(ctx, "/simpli.simpli_echo_me.SimpliEchoMeService/SayHi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpliEchoMeServiceClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/simpli.simpli_echo_me.SimpliEchoMeService/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpliEchoMeServiceClient) EchoRemember(ctx context.Context, in *EchoRememberRequest, opts ...grpc.CallOption) (*EchoRememberResponse, error) {
	out := new(EchoRememberResponse)
	err := c.cc.Invoke(ctx, "/simpli.simpli_echo_me.SimpliEchoMeService/EchoRemember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpliEchoMeServiceClient) SayHiRemember(ctx context.Context, in *SayHiRememberRequest, opts ...grpc.CallOption) (*SayHiRememberResponse, error) {
	out := new(SayHiRememberResponse)
	err := c.cc.Invoke(ctx, "/simpli.simpli_echo_me.SimpliEchoMeService/SayHiRemember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimpliEchoMeServiceServer is the server API for SimpliEchoMeService service.
// All implementations should embed UnimplementedSimpliEchoMeServiceServer
// for forward compatibility
type SimpliEchoMeServiceServer interface {
	// Say Hi.
	SayHi(context.Context, *emptypb.Empty) (*SayHiResponse, error)
	// Say Hi echoing my name.
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
	// Remember my name.
	EchoRemember(context.Context, *EchoRememberRequest) (*EchoRememberResponse, error)
	// Say Hi given my remembered id.
	SayHiRemember(context.Context, *SayHiRememberRequest) (*SayHiRememberResponse, error)
}

// UnimplementedSimpliEchoMeServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSimpliEchoMeServiceServer struct {
}

func (UnimplementedSimpliEchoMeServiceServer) SayHi(context.Context, *emptypb.Empty) (*SayHiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHi not implemented")
}
func (UnimplementedSimpliEchoMeServiceServer) Echo(context.Context, *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedSimpliEchoMeServiceServer) EchoRemember(context.Context, *EchoRememberRequest) (*EchoRememberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EchoRemember not implemented")
}
func (UnimplementedSimpliEchoMeServiceServer) SayHiRemember(context.Context, *SayHiRememberRequest) (*SayHiRememberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHiRemember not implemented")
}

// UnsafeSimpliEchoMeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SimpliEchoMeServiceServer will
// result in compilation errors.
type UnsafeSimpliEchoMeServiceServer interface {
	mustEmbedUnimplementedSimpliEchoMeServiceServer()
}

func RegisterSimpliEchoMeServiceServer(s grpc.ServiceRegistrar, srv SimpliEchoMeServiceServer) {
	s.RegisterService(&SimpliEchoMeService_ServiceDesc, srv)
}

func _SimpliEchoMeService_SayHi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpliEchoMeServiceServer).SayHi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simpli.simpli_echo_me.SimpliEchoMeService/SayHi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpliEchoMeServiceServer).SayHi(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimpliEchoMeService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpliEchoMeServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simpli.simpli_echo_me.SimpliEchoMeService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpliEchoMeServiceServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimpliEchoMeService_EchoRemember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRememberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpliEchoMeServiceServer).EchoRemember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simpli.simpli_echo_me.SimpliEchoMeService/EchoRemember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpliEchoMeServiceServer).EchoRemember(ctx, req.(*EchoRememberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimpliEchoMeService_SayHiRemember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayHiRememberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpliEchoMeServiceServer).SayHiRemember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simpli.simpli_echo_me.SimpliEchoMeService/SayHiRemember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpliEchoMeServiceServer).SayHiRemember(ctx, req.(*SayHiRememberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SimpliEchoMeService_ServiceDesc is the grpc.ServiceDesc for SimpliEchoMeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SimpliEchoMeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "simpli.simpli_echo_me.SimpliEchoMeService",
	HandlerType: (*SimpliEchoMeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHi",
			Handler:    _SimpliEchoMeService_SayHi_Handler,
		},
		{
			MethodName: "Echo",
			Handler:    _SimpliEchoMeService_Echo_Handler,
		},
		{
			MethodName: "EchoRemember",
			Handler:    _SimpliEchoMeService_EchoRemember_Handler,
		},
		{
			MethodName: "SayHiRemember",
			Handler:    _SimpliEchoMeService_SayHiRemember_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "simpli-echo-me/simpliEchoMe.proto",
}
