// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: exposer/service.proto

package exposer

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

// ExposerClient is the client API for Exposer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExposerClient interface {
	// GetNodeID will returns current node id
	GetNodeID(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ResponseNodeID, error)
	// GetPubKey returns the public key of the current node
	GetPubKey(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ResponsePubKey, error)
	// GetGenesisFile returns the genesis file of the node
	GetGenesisFile(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ResponseFileData, error)
	// GetKeysFile returns the keys of the node
	GetKeysFile(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ResponseFileData, error)
	// GetPrivKeysFile returns the keys of the node
	GetPrivKeysFile(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ResponseFileData, error)
}

type exposerClient struct {
	cc grpc.ClientConnInterface
}

func NewExposerClient(cc grpc.ClientConnInterface) ExposerClient {
	return &exposerClient{cc}
}

func (c *exposerClient) GetNodeID(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ResponseNodeID, error) {
	out := new(ResponseNodeID)
	err := c.cc.Invoke(ctx, "/exposer.Exposer/GetNodeID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exposerClient) GetPubKey(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ResponsePubKey, error) {
	out := new(ResponsePubKey)
	err := c.cc.Invoke(ctx, "/exposer.Exposer/GetPubKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exposerClient) GetGenesisFile(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ResponseFileData, error) {
	out := new(ResponseFileData)
	err := c.cc.Invoke(ctx, "/exposer.Exposer/GetGenesisFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exposerClient) GetKeysFile(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ResponseFileData, error) {
	out := new(ResponseFileData)
	err := c.cc.Invoke(ctx, "/exposer.Exposer/GetKeysFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exposerClient) GetPrivKeysFile(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ResponseFileData, error) {
	out := new(ResponseFileData)
	err := c.cc.Invoke(ctx, "/exposer.Exposer/GetPrivKeysFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExposerServer is the server API for Exposer service.
// All implementations must embed UnimplementedExposerServer
// for forward compatibility
type ExposerServer interface {
	// GetNodeID will returns current node id
	GetNodeID(context.Context, *emptypb.Empty) (*ResponseNodeID, error)
	// GetPubKey returns the public key of the current node
	GetPubKey(context.Context, *emptypb.Empty) (*ResponsePubKey, error)
	// GetGenesisFile returns the genesis file of the node
	GetGenesisFile(context.Context, *emptypb.Empty) (*ResponseFileData, error)
	// GetKeysFile returns the keys of the node
	GetKeysFile(context.Context, *emptypb.Empty) (*ResponseFileData, error)
	// GetPrivKeysFile returns the keys of the node
	GetPrivKeysFile(context.Context, *emptypb.Empty) (*ResponseFileData, error)
	mustEmbedUnimplementedExposerServer()
}

// UnimplementedExposerServer must be embedded to have forward compatible implementations.
type UnimplementedExposerServer struct {
}

func (UnimplementedExposerServer) GetNodeID(context.Context, *emptypb.Empty) (*ResponseNodeID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeID not implemented")
}
func (UnimplementedExposerServer) GetPubKey(context.Context, *emptypb.Empty) (*ResponsePubKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPubKey not implemented")
}
func (UnimplementedExposerServer) GetGenesisFile(context.Context, *emptypb.Empty) (*ResponseFileData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGenesisFile not implemented")
}
func (UnimplementedExposerServer) GetKeysFile(context.Context, *emptypb.Empty) (*ResponseFileData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKeysFile not implemented")
}
func (UnimplementedExposerServer) GetPrivKeysFile(context.Context, *emptypb.Empty) (*ResponseFileData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrivKeysFile not implemented")
}
func (UnimplementedExposerServer) mustEmbedUnimplementedExposerServer() {}

// UnsafeExposerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExposerServer will
// result in compilation errors.
type UnsafeExposerServer interface {
	mustEmbedUnimplementedExposerServer()
}

func RegisterExposerServer(s grpc.ServiceRegistrar, srv ExposerServer) {
	s.RegisterService(&Exposer_ServiceDesc, srv)
}

func _Exposer_GetNodeID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExposerServer).GetNodeID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exposer.Exposer/GetNodeID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExposerServer).GetNodeID(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exposer_GetPubKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExposerServer).GetPubKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exposer.Exposer/GetPubKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExposerServer).GetPubKey(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exposer_GetGenesisFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExposerServer).GetGenesisFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exposer.Exposer/GetGenesisFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExposerServer).GetGenesisFile(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exposer_GetKeysFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExposerServer).GetKeysFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exposer.Exposer/GetKeysFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExposerServer).GetKeysFile(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exposer_GetPrivKeysFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExposerServer).GetPrivKeysFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exposer.Exposer/GetPrivKeysFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExposerServer).GetPrivKeysFile(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Exposer_ServiceDesc is the grpc.ServiceDesc for Exposer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Exposer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "exposer.Exposer",
	HandlerType: (*ExposerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNodeID",
			Handler:    _Exposer_GetNodeID_Handler,
		},
		{
			MethodName: "GetPubKey",
			Handler:    _Exposer_GetPubKey_Handler,
		},
		{
			MethodName: "GetGenesisFile",
			Handler:    _Exposer_GetGenesisFile_Handler,
		},
		{
			MethodName: "GetKeysFile",
			Handler:    _Exposer_GetKeysFile_Handler,
		},
		{
			MethodName: "GetPrivKeysFile",
			Handler:    _Exposer_GetPrivKeysFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "exposer/service.proto",
}
