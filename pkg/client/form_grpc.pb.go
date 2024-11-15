// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.1
// source: form.proto

package client

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Form_FindByID_FullMethodName              = "/api.Form/FindByID"
	Form_FindByTeacherUsername_FullMethodName = "/api.Form/FindByTeacherUsername"
)

// FormClient is the client API for Form service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FormClient interface {
	FindByID(ctx context.Context, in *FormFindByIDRequest, opts ...grpc.CallOption) (*FormFindByIDResponse, error)
	FindByTeacherUsername(ctx context.Context, in *FormFindByTeacherUsernameRequest, opts ...grpc.CallOption) (*FormFindByTeacherUsernameResponse, error)
}

type formClient struct {
	cc grpc.ClientConnInterface
}

func NewFormClient(cc grpc.ClientConnInterface) FormClient {
	return &formClient{cc}
}

func (c *formClient) FindByID(ctx context.Context, in *FormFindByIDRequest, opts ...grpc.CallOption) (*FormFindByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FormFindByIDResponse)
	err := c.cc.Invoke(ctx, Form_FindByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *formClient) FindByTeacherUsername(ctx context.Context, in *FormFindByTeacherUsernameRequest, opts ...grpc.CallOption) (*FormFindByTeacherUsernameResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FormFindByTeacherUsernameResponse)
	err := c.cc.Invoke(ctx, Form_FindByTeacherUsername_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FormServer is the server API for Form service.
// All implementations must embed UnimplementedFormServer
// for forward compatibility.
type FormServer interface {
	FindByID(context.Context, *FormFindByIDRequest) (*FormFindByIDResponse, error)
	FindByTeacherUsername(context.Context, *FormFindByTeacherUsernameRequest) (*FormFindByTeacherUsernameResponse, error)
	mustEmbedUnimplementedFormServer()
}

// UnimplementedFormServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFormServer struct{}

func (UnimplementedFormServer) FindByID(context.Context, *FormFindByIDRequest) (*FormFindByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByID not implemented")
}
func (UnimplementedFormServer) FindByTeacherUsername(context.Context, *FormFindByTeacherUsernameRequest) (*FormFindByTeacherUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByTeacherUsername not implemented")
}
func (UnimplementedFormServer) mustEmbedUnimplementedFormServer() {}
func (UnimplementedFormServer) testEmbeddedByValue()              {}

// UnsafeFormServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FormServer will
// result in compilation errors.
type UnsafeFormServer interface {
	mustEmbedUnimplementedFormServer()
}

func RegisterFormServer(s grpc.ServiceRegistrar, srv FormServer) {
	// If the following call pancis, it indicates UnimplementedFormServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Form_ServiceDesc, srv)
}

func _Form_FindByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FormFindByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FormServer).FindByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Form_FindByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FormServer).FindByID(ctx, req.(*FormFindByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Form_FindByTeacherUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FormFindByTeacherUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FormServer).FindByTeacherUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Form_FindByTeacherUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FormServer).FindByTeacherUsername(ctx, req.(*FormFindByTeacherUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Form_ServiceDesc is the grpc.ServiceDesc for Form service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Form_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Form",
	HandlerType: (*FormServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindByID",
			Handler:    _Form_FindByID_Handler,
		},
		{
			MethodName: "FindByTeacherUsername",
			Handler:    _Form_FindByTeacherUsername_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "form.proto",
}
