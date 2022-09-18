// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: items.proto

package items

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

// ItemsServiceClient is the client API for ItemsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ItemsServiceClient interface {
	Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type itemsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewItemsServiceClient(cc grpc.ClientConnInterface) ItemsServiceClient {
	return &itemsServiceClient{cc}
}

func (c *itemsServiceClient) Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*FetchResponse, error) {
	out := new(FetchResponse)
	err := c.cc.Invoke(ctx, "/items.ItemsService/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemsServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/items.ItemsService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ItemsServiceServer is the server API for ItemsService service.
// All implementations should embed UnimplementedItemsServiceServer
// for forward compatibility
type ItemsServiceServer interface {
	Fetch(context.Context, *FetchRequest) (*FetchResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
}

// UnimplementedItemsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedItemsServiceServer struct {
}

func (UnimplementedItemsServiceServer) Fetch(context.Context, *FetchRequest) (*FetchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (UnimplementedItemsServiceServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

// UnsafeItemsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ItemsServiceServer will
// result in compilation errors.
type UnsafeItemsServiceServer interface {
	mustEmbedUnimplementedItemsServiceServer()
}

func RegisterItemsServiceServer(s grpc.ServiceRegistrar, srv ItemsServiceServer) {
	s.RegisterService(&ItemsService_ServiceDesc, srv)
}

func _ItemsService_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemsServiceServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/items.ItemsService/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemsServiceServer).Fetch(ctx, req.(*FetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemsService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemsServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/items.ItemsService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemsServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ItemsService_ServiceDesc is the grpc.ServiceDesc for ItemsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ItemsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "items.ItemsService",
	HandlerType: (*ItemsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fetch",
			Handler:    _ItemsService_Fetch_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ItemsService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "items.proto",
}