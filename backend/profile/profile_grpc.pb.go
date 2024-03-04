// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: profile.proto

package profile

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

// ProfileRetrievalServiceClient is the client API for ProfileRetrievalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfileRetrievalServiceClient interface {
	GetProfileById(ctx context.Context, in *ProfileByIdRequest, opts ...grpc.CallOption) (*Profile, error)
	AddProfile(ctx context.Context, in *Profile, opts ...grpc.CallOption) (*ErrorResponse, error)
}

type profileRetrievalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileRetrievalServiceClient(cc grpc.ClientConnInterface) ProfileRetrievalServiceClient {
	return &profileRetrievalServiceClient{cc}
}

func (c *profileRetrievalServiceClient) GetProfileById(ctx context.Context, in *ProfileByIdRequest, opts ...grpc.CallOption) (*Profile, error) {
	out := new(Profile)
	err := c.cc.Invoke(ctx, "/profiles.ProfileRetrievalService/GetProfileById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileRetrievalServiceClient) AddProfile(ctx context.Context, in *Profile, opts ...grpc.CallOption) (*ErrorResponse, error) {
	out := new(ErrorResponse)
	err := c.cc.Invoke(ctx, "/profiles.ProfileRetrievalService/AddProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileRetrievalServiceServer is the server API for ProfileRetrievalService service.
// All implementations must embed UnimplementedProfileRetrievalServiceServer
// for forward compatibility
type ProfileRetrievalServiceServer interface {
	GetProfileById(context.Context, *ProfileByIdRequest) (*Profile, error)
	AddProfile(context.Context, *Profile) (*ErrorResponse, error)
	mustEmbedUnimplementedProfileRetrievalServiceServer()
}

// UnimplementedProfileRetrievalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProfileRetrievalServiceServer struct {
}

func (UnimplementedProfileRetrievalServiceServer) GetProfileById(context.Context, *ProfileByIdRequest) (*Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileById not implemented")
}
func (UnimplementedProfileRetrievalServiceServer) AddProfile(context.Context, *Profile) (*ErrorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProfile not implemented")
}
func (UnimplementedProfileRetrievalServiceServer) mustEmbedUnimplementedProfileRetrievalServiceServer() {
}

// UnsafeProfileRetrievalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfileRetrievalServiceServer will
// result in compilation errors.
type UnsafeProfileRetrievalServiceServer interface {
	mustEmbedUnimplementedProfileRetrievalServiceServer()
}

func RegisterProfileRetrievalServiceServer(s grpc.ServiceRegistrar, srv ProfileRetrievalServiceServer) {
	s.RegisterService(&ProfileRetrievalService_ServiceDesc, srv)
}

func _ProfileRetrievalService_GetProfileById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileRetrievalServiceServer).GetProfileById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profiles.ProfileRetrievalService/GetProfileById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileRetrievalServiceServer).GetProfileById(ctx, req.(*ProfileByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileRetrievalService_AddProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Profile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileRetrievalServiceServer).AddProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profiles.ProfileRetrievalService/AddProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileRetrievalServiceServer).AddProfile(ctx, req.(*Profile))
	}
	return interceptor(ctx, in, info, handler)
}

// ProfileRetrievalService_ServiceDesc is the grpc.ServiceDesc for ProfileRetrievalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProfileRetrievalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "profiles.ProfileRetrievalService",
	HandlerType: (*ProfileRetrievalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProfileById",
			Handler:    _ProfileRetrievalService_GetProfileById_Handler,
		},
		{
			MethodName: "AddProfile",
			Handler:    _ProfileRetrievalService_AddProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profile.proto",
}