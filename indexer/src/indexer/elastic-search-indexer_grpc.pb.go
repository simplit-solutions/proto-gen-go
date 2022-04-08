// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: indexer/elastic-search-indexer.proto

package indexer

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// IndexerServiceClient is the client API for IndexerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IndexerServiceClient interface {
	// Locations querying
	GetIndexById(ctx context.Context, in *GetIndexByIdRequest, opts ...grpc.CallOption) (*GetIndexByIdResponse, error)
}

type indexerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIndexerServiceClient(cc grpc.ClientConnInterface) IndexerServiceClient {
	return &indexerServiceClient{cc}
}

func (c *indexerServiceClient) GetIndexById(ctx context.Context, in *GetIndexByIdRequest, opts ...grpc.CallOption) (*GetIndexByIdResponse, error) {
	out := new(GetIndexByIdResponse)
	err := c.cc.Invoke(ctx, "/simpliroute.indexer.IndexerService/GetIndexById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IndexerServiceServer is the server API for IndexerService service.
// All implementations should embed UnimplementedIndexerServiceServer
// for forward compatibility
type IndexerServiceServer interface {
	// Locations querying
	GetIndexById(context.Context, *GetIndexByIdRequest) (*GetIndexByIdResponse, error)
}

// UnimplementedIndexerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedIndexerServiceServer struct {
}

func (UnimplementedIndexerServiceServer) GetIndexById(context.Context, *GetIndexByIdRequest) (*GetIndexByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIndexById not implemented")
}

// UnsafeIndexerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IndexerServiceServer will
// result in compilation errors.
type UnsafeIndexerServiceServer interface {
	mustEmbedUnimplementedIndexerServiceServer()
}

func RegisterIndexerServiceServer(s grpc.ServiceRegistrar, srv IndexerServiceServer) {
	s.RegisterService(&IndexerService_ServiceDesc, srv)
}

func _IndexerService_GetIndexById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIndexByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IndexerServiceServer).GetIndexById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simpliroute.indexer.IndexerService/GetIndexById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IndexerServiceServer).GetIndexById(ctx, req.(*GetIndexByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IndexerService_ServiceDesc is the grpc.ServiceDesc for IndexerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IndexerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "simpliroute.indexer.IndexerService",
	HandlerType: (*IndexerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetIndexById",
			Handler:    _IndexerService_GetIndexById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "indexer/elastic-search-indexer.proto",
}
