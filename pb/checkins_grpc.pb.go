// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// CheckinsClient is the client API for Checkins service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CheckinsClient interface {
	Create(ctx context.Context, in *CreateCheckinRequest, opts ...grpc.CallOption) (*CreateCheckinResponse, error)
	Get(ctx context.Context, in *GetCheckinRequest, opts ...grpc.CallOption) (*GetCheckinResponse, error)
}

type checkinsClient struct {
	cc grpc.ClientConnInterface
}

func NewCheckinsClient(cc grpc.ClientConnInterface) CheckinsClient {
	return &checkinsClient{cc}
}

func (c *checkinsClient) Create(ctx context.Context, in *CreateCheckinRequest, opts ...grpc.CallOption) (*CreateCheckinResponse, error) {
	out := new(CreateCheckinResponse)
	err := c.cc.Invoke(ctx, "/shibafu528.utissue.Checkins/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkinsClient) Get(ctx context.Context, in *GetCheckinRequest, opts ...grpc.CallOption) (*GetCheckinResponse, error) {
	out := new(GetCheckinResponse)
	err := c.cc.Invoke(ctx, "/shibafu528.utissue.Checkins/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckinsServer is the server API for Checkins service.
// All implementations must embed UnimplementedCheckinsServer
// for forward compatibility
type CheckinsServer interface {
	Create(context.Context, *CreateCheckinRequest) (*CreateCheckinResponse, error)
	Get(context.Context, *GetCheckinRequest) (*GetCheckinResponse, error)
	mustEmbedUnimplementedCheckinsServer()
}

// UnimplementedCheckinsServer must be embedded to have forward compatible implementations.
type UnimplementedCheckinsServer struct {
}

func (UnimplementedCheckinsServer) Create(context.Context, *CreateCheckinRequest) (*CreateCheckinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCheckinsServer) Get(context.Context, *GetCheckinRequest) (*GetCheckinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCheckinsServer) mustEmbedUnimplementedCheckinsServer() {}

// UnsafeCheckinsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CheckinsServer will
// result in compilation errors.
type UnsafeCheckinsServer interface {
	mustEmbedUnimplementedCheckinsServer()
}

func RegisterCheckinsServer(s grpc.ServiceRegistrar, srv CheckinsServer) {
	s.RegisterService(&Checkins_ServiceDesc, srv)
}

func _Checkins_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCheckinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckinsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shibafu528.utissue.Checkins/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckinsServer).Create(ctx, req.(*CreateCheckinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checkins_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCheckinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckinsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shibafu528.utissue.Checkins/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckinsServer).Get(ctx, req.(*GetCheckinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Checkins_ServiceDesc is the grpc.ServiceDesc for Checkins service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Checkins_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shibafu528.utissue.Checkins",
	HandlerType: (*CheckinsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Checkins_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Checkins_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shibafu528/utissue/checkins.proto",
}
