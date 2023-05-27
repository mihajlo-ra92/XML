// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.0
// source: rating_service.proto

package rating

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

// RatingServiceClient is the client API for RatingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RatingServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	CreateRating(ctx context.Context, in *CreateRatingRequest, opts ...grpc.CallOption) (*CreateRatingResponse, error)
	DeleteRating(ctx context.Context, in *DeleteRatingRequest, opts ...grpc.CallOption) (*DeleteRatingResponse, error)
	GetUserRatingByAccommodationId(ctx context.Context, in *GetUserRatingByAccommodationIdRequest, opts ...grpc.CallOption) (*GetUserRatingByAccommodationIdResponse, error)
	GetUserRatingByHostId(ctx context.Context, in *GetUserRatingByHostIdRequest, opts ...grpc.CallOption) (*GetUserRatingByHostIdResponse, error)
}

type ratingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRatingServiceClient(cc grpc.ClientConnInterface) RatingServiceClient {
	return &ratingServiceClient{cc}
}

func (c *ratingServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/rating.RatingService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/rating.RatingService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) CreateRating(ctx context.Context, in *CreateRatingRequest, opts ...grpc.CallOption) (*CreateRatingResponse, error) {
	out := new(CreateRatingResponse)
	err := c.cc.Invoke(ctx, "/rating.RatingService/CreateRating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) DeleteRating(ctx context.Context, in *DeleteRatingRequest, opts ...grpc.CallOption) (*DeleteRatingResponse, error) {
	out := new(DeleteRatingResponse)
	err := c.cc.Invoke(ctx, "/rating.RatingService/DeleteRating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetUserRatingByAccommodationId(ctx context.Context, in *GetUserRatingByAccommodationIdRequest, opts ...grpc.CallOption) (*GetUserRatingByAccommodationIdResponse, error) {
	out := new(GetUserRatingByAccommodationIdResponse)
	err := c.cc.Invoke(ctx, "/rating.RatingService/GetUserRatingByAccommodationId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetUserRatingByHostId(ctx context.Context, in *GetUserRatingByHostIdRequest, opts ...grpc.CallOption) (*GetUserRatingByHostIdResponse, error) {
	out := new(GetUserRatingByHostIdResponse)
	err := c.cc.Invoke(ctx, "/rating.RatingService/GetUserRatingByHostId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RatingServiceServer is the server API for RatingService service.
// All implementations must embed UnimplementedRatingServiceServer
// for forward compatibility
type RatingServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	CreateRating(context.Context, *CreateRatingRequest) (*CreateRatingResponse, error)
	DeleteRating(context.Context, *DeleteRatingRequest) (*DeleteRatingResponse, error)
	GetUserRatingByAccommodationId(context.Context, *GetUserRatingByAccommodationIdRequest) (*GetUserRatingByAccommodationIdResponse, error)
	GetUserRatingByHostId(context.Context, *GetUserRatingByHostIdRequest) (*GetUserRatingByHostIdResponse, error)
	mustEmbedUnimplementedRatingServiceServer()
}

// UnimplementedRatingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRatingServiceServer struct {
}

func (UnimplementedRatingServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedRatingServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedRatingServiceServer) CreateRating(context.Context, *CreateRatingRequest) (*CreateRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRating not implemented")
}
func (UnimplementedRatingServiceServer) DeleteRating(context.Context, *DeleteRatingRequest) (*DeleteRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRating not implemented")
}
func (UnimplementedRatingServiceServer) GetUserRatingByAccommodationId(context.Context, *GetUserRatingByAccommodationIdRequest) (*GetUserRatingByAccommodationIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserRatingByAccommodationId not implemented")
}
func (UnimplementedRatingServiceServer) GetUserRatingByHostId(context.Context, *GetUserRatingByHostIdRequest) (*GetUserRatingByHostIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserRatingByHostId not implemented")
}
func (UnimplementedRatingServiceServer) mustEmbedUnimplementedRatingServiceServer() {}

// UnsafeRatingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RatingServiceServer will
// result in compilation errors.
type UnsafeRatingServiceServer interface {
	mustEmbedUnimplementedRatingServiceServer()
}

func RegisterRatingServiceServer(s grpc.ServiceRegistrar, srv RatingServiceServer) {
	s.RegisterService(&RatingService_ServiceDesc, srv)
}

func _RatingService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_CreateRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).CreateRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/CreateRating",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).CreateRating(ctx, req.(*CreateRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_DeleteRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).DeleteRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/DeleteRating",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).DeleteRating(ctx, req.(*DeleteRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetUserRatingByAccommodationId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRatingByAccommodationIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetUserRatingByAccommodationId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/GetUserRatingByAccommodationId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetUserRatingByAccommodationId(ctx, req.(*GetUserRatingByAccommodationIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetUserRatingByHostId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRatingByHostIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetUserRatingByHostId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/GetUserRatingByHostId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetUserRatingByHostId(ctx, req.(*GetUserRatingByHostIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RatingService_ServiceDesc is the grpc.ServiceDesc for RatingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RatingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rating.RatingService",
	HandlerType: (*RatingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _RatingService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _RatingService_GetAll_Handler,
		},
		{
			MethodName: "CreateRating",
			Handler:    _RatingService_CreateRating_Handler,
		},
		{
			MethodName: "DeleteRating",
			Handler:    _RatingService_DeleteRating_Handler,
		},
		{
			MethodName: "GetUserRatingByAccommodationId",
			Handler:    _RatingService_GetUserRatingByAccommodationId_Handler,
		},
		{
			MethodName: "GetUserRatingByHostId",
			Handler:    _RatingService_GetUserRatingByHostId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rating_service.proto",
}
