// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.0
// source: auth_service.proto

package user

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

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	AuthCreateAccommodation(ctx context.Context, in *AuthCreateAccommodationRequest, opts ...grpc.CallOption) (*AuthCreateAccommodationResponse, error)
	AuthUpdateUser(ctx context.Context, in *AuthUpdateUserRequest, opts ...grpc.CallOption) (*AuthUpdateUserResponse, error)
	AuthDeleteUser(ctx context.Context, in *AuthDeleteUserRequest, opts ...grpc.CallOption) (*AuthDeleteUserResponse, error)
	AuthGuestReserveAccommodation(ctx context.Context, in *AuthGuestReserveAccommodationRequest, opts ...grpc.CallOption) (*AuthGuestReserveAccommodationResponse, error)
	AuthBookingAccept(ctx context.Context, in *AuthBookingAcceptRequest, opts ...grpc.CallOption) (*AuthBookingAcceptResponse, error)
	AuthBookingDeny(ctx context.Context, in *AuthBookingDenyRequest, opts ...grpc.CallOption) (*AuthBookingDenyResponse, error)
	AuthReservationCanceling(ctx context.Context, in *AuthReservationCancelingRequest, opts ...grpc.CallOption) (*AuthReservationCancelingResponse, error)
	AuthDefineCustomPrice(ctx context.Context, in *AuthDefineCustomPriceRequest, opts ...grpc.CallOption) (*AuthDefineCustomPriceResponse, error)
	AuthGetAccommodationByHostId(ctx context.Context, in *AuthGetAccommodationsByHostIdRequest, opts ...grpc.CallOption) (*AuthGetAccommodationsByHostIdResponse, error)
	AuthGetBookingsByAccommodationId(ctx context.Context, in *AuthGetBookingsByAccommodationIdRequest, opts ...grpc.CallOption) (*AuthGetBookingsByAccommodationIdResponse, error)
	AuthCreateRating(ctx context.Context, in *AuthCreateRatingRequest, opts ...grpc.CallOption) (*AuthCreateRatingResponse, error)
	AuthDeleteRating(ctx context.Context, in *AuthDeleteRatingRequest, opts ...grpc.CallOption) (*AuthDeleteRatingResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/user.AuthService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AuthCreateAccommodation(ctx context.Context, in *AuthCreateAccommodationRequest, opts ...grpc.CallOption) (*AuthCreateAccommodationResponse, error) {
	out := new(AuthCreateAccommodationResponse)
	err := c.cc.Invoke(ctx, "/user.AuthService/AuthCreateAccommodation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AuthUpdateUser(ctx context.Context, in *AuthUpdateUserRequest, opts ...grpc.CallOption) (*AuthUpdateUserResponse, error) {
	out := new(AuthUpdateUserResponse)
	err := c.cc.Invoke(ctx, "/user.AuthService/AuthUpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AuthDeleteUser(ctx context.Context, in *AuthDeleteUserRequest, opts ...grpc.CallOption) (*AuthDeleteUserResponse, error) {
	out := new(AuthDeleteUserResponse)
	err := c.cc.Invoke(ctx, "/user.AuthService/AuthDeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AuthGuestReserveAccommodation(ctx context.Context, in *AuthGuestReserveAccommodationRequest, opts ...grpc.CallOption) (*AuthGuestReserveAccommodationResponse, error) {
	out := new(AuthGuestReserveAccommodationResponse)
	err := c.cc.Invoke(ctx, "/user.AuthService/AuthGuestReserveAccommodation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AuthBookingAccept(ctx context.Context, in *AuthBookingAcceptRequest, opts ...grpc.CallOption) (*AuthBookingAcceptResponse, error) {
	out := new(AuthBookingAcceptResponse)
	err := c.cc.Invoke(ctx, "/user.AuthService/AuthBookingAccept", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AuthBookingDeny(ctx context.Context, in *AuthBookingDenyRequest, opts ...grpc.CallOption) (*AuthBookingDenyResponse, error) {
	out := new(AuthBookingDenyResponse)
	err := c.cc.Invoke(ctx, "/user.AuthService/AuthBookingDeny", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AuthReservationCanceling(ctx context.Context, in *AuthReservationCancelingRequest, opts ...grpc.CallOption) (*AuthReservationCancelingResponse, error) {
	out := new(AuthReservationCancelingResponse)
	err := c.cc.Invoke(ctx, "/user.AuthService/AuthReservationCanceling", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AuthDefineCustomPrice(ctx context.Context, in *AuthDefineCustomPriceRequest, opts ...grpc.CallOption) (*AuthDefineCustomPriceResponse, error) {
	out := new(AuthDefineCustomPriceResponse)
	err := c.cc.Invoke(ctx, "/user.AuthService/AuthDefineCustomPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AuthGetAccommodationByHostId(ctx context.Context, in *AuthGetAccommodationsByHostIdRequest, opts ...grpc.CallOption) (*AuthGetAccommodationsByHostIdResponse, error) {
	out := new(AuthGetAccommodationsByHostIdResponse)
	err := c.cc.Invoke(ctx, "/user.AuthService/AuthGetAccommodationByHostId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AuthGetBookingsByAccommodationId(ctx context.Context, in *AuthGetBookingsByAccommodationIdRequest, opts ...grpc.CallOption) (*AuthGetBookingsByAccommodationIdResponse, error) {
	out := new(AuthGetBookingsByAccommodationIdResponse)
	err := c.cc.Invoke(ctx, "/user.AuthService/AuthGetBookingsByAccommodationId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AuthCreateRating(ctx context.Context, in *AuthCreateRatingRequest, opts ...grpc.CallOption) (*AuthCreateRatingResponse, error) {
	out := new(AuthCreateRatingResponse)
	err := c.cc.Invoke(ctx, "/user.AuthService/AuthCreateRating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AuthDeleteRating(ctx context.Context, in *AuthDeleteRatingRequest, opts ...grpc.CallOption) (*AuthDeleteRatingResponse, error) {
	out := new(AuthDeleteRatingResponse)
	err := c.cc.Invoke(ctx, "/user.AuthService/AuthDeleteRating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	AuthCreateAccommodation(context.Context, *AuthCreateAccommodationRequest) (*AuthCreateAccommodationResponse, error)
	AuthUpdateUser(context.Context, *AuthUpdateUserRequest) (*AuthUpdateUserResponse, error)
	AuthDeleteUser(context.Context, *AuthDeleteUserRequest) (*AuthDeleteUserResponse, error)
	AuthGuestReserveAccommodation(context.Context, *AuthGuestReserveAccommodationRequest) (*AuthGuestReserveAccommodationResponse, error)
	AuthBookingAccept(context.Context, *AuthBookingAcceptRequest) (*AuthBookingAcceptResponse, error)
	AuthBookingDeny(context.Context, *AuthBookingDenyRequest) (*AuthBookingDenyResponse, error)
	AuthReservationCanceling(context.Context, *AuthReservationCancelingRequest) (*AuthReservationCancelingResponse, error)
	AuthDefineCustomPrice(context.Context, *AuthDefineCustomPriceRequest) (*AuthDefineCustomPriceResponse, error)
	AuthGetAccommodationByHostId(context.Context, *AuthGetAccommodationsByHostIdRequest) (*AuthGetAccommodationsByHostIdResponse, error)
	AuthGetBookingsByAccommodationId(context.Context, *AuthGetBookingsByAccommodationIdRequest) (*AuthGetBookingsByAccommodationIdResponse, error)
	AuthCreateRating(context.Context, *AuthCreateRatingRequest) (*AuthCreateRatingResponse, error)
	AuthDeleteRating(context.Context, *AuthDeleteRatingRequest) (*AuthDeleteRatingResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServiceServer) AuthCreateAccommodation(context.Context, *AuthCreateAccommodationRequest) (*AuthCreateAccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthCreateAccommodation not implemented")
}
func (UnimplementedAuthServiceServer) AuthUpdateUser(context.Context, *AuthUpdateUserRequest) (*AuthUpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthUpdateUser not implemented")
}
func (UnimplementedAuthServiceServer) AuthDeleteUser(context.Context, *AuthDeleteUserRequest) (*AuthDeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthDeleteUser not implemented")
}
func (UnimplementedAuthServiceServer) AuthGuestReserveAccommodation(context.Context, *AuthGuestReserveAccommodationRequest) (*AuthGuestReserveAccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthGuestReserveAccommodation not implemented")
}
func (UnimplementedAuthServiceServer) AuthBookingAccept(context.Context, *AuthBookingAcceptRequest) (*AuthBookingAcceptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthBookingAccept not implemented")
}
func (UnimplementedAuthServiceServer) AuthBookingDeny(context.Context, *AuthBookingDenyRequest) (*AuthBookingDenyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthBookingDeny not implemented")
}
func (UnimplementedAuthServiceServer) AuthReservationCanceling(context.Context, *AuthReservationCancelingRequest) (*AuthReservationCancelingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthReservationCanceling not implemented")
}
func (UnimplementedAuthServiceServer) AuthDefineCustomPrice(context.Context, *AuthDefineCustomPriceRequest) (*AuthDefineCustomPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthDefineCustomPrice not implemented")
}
func (UnimplementedAuthServiceServer) AuthGetAccommodationByHostId(context.Context, *AuthGetAccommodationsByHostIdRequest) (*AuthGetAccommodationsByHostIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthGetAccommodationByHostId not implemented")
}
func (UnimplementedAuthServiceServer) AuthGetBookingsByAccommodationId(context.Context, *AuthGetBookingsByAccommodationIdRequest) (*AuthGetBookingsByAccommodationIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthGetBookingsByAccommodationId not implemented")
}
func (UnimplementedAuthServiceServer) AuthCreateRating(context.Context, *AuthCreateRatingRequest) (*AuthCreateRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthCreateRating not implemented")
}
func (UnimplementedAuthServiceServer) AuthDeleteRating(context.Context, *AuthDeleteRatingRequest) (*AuthDeleteRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthDeleteRating not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.AuthService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AuthCreateAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthCreateAccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AuthCreateAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.AuthService/AuthCreateAccommodation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AuthCreateAccommodation(ctx, req.(*AuthCreateAccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AuthUpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthUpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AuthUpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.AuthService/AuthUpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AuthUpdateUser(ctx, req.(*AuthUpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AuthDeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthDeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AuthDeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.AuthService/AuthDeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AuthDeleteUser(ctx, req.(*AuthDeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AuthGuestReserveAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthGuestReserveAccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AuthGuestReserveAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.AuthService/AuthGuestReserveAccommodation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AuthGuestReserveAccommodation(ctx, req.(*AuthGuestReserveAccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AuthBookingAccept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthBookingAcceptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AuthBookingAccept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.AuthService/AuthBookingAccept",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AuthBookingAccept(ctx, req.(*AuthBookingAcceptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AuthBookingDeny_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthBookingDenyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AuthBookingDeny(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.AuthService/AuthBookingDeny",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AuthBookingDeny(ctx, req.(*AuthBookingDenyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AuthReservationCanceling_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthReservationCancelingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AuthReservationCanceling(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.AuthService/AuthReservationCanceling",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AuthReservationCanceling(ctx, req.(*AuthReservationCancelingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AuthDefineCustomPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthDefineCustomPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AuthDefineCustomPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.AuthService/AuthDefineCustomPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AuthDefineCustomPrice(ctx, req.(*AuthDefineCustomPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AuthGetAccommodationByHostId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthGetAccommodationsByHostIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AuthGetAccommodationByHostId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.AuthService/AuthGetAccommodationByHostId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AuthGetAccommodationByHostId(ctx, req.(*AuthGetAccommodationsByHostIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AuthGetBookingsByAccommodationId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthGetBookingsByAccommodationIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AuthGetBookingsByAccommodationId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.AuthService/AuthGetBookingsByAccommodationId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AuthGetBookingsByAccommodationId(ctx, req.(*AuthGetBookingsByAccommodationIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AuthCreateRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthCreateRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AuthCreateRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.AuthService/AuthCreateRating",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AuthCreateRating(ctx, req.(*AuthCreateRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AuthDeleteRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthDeleteRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AuthDeleteRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.AuthService/AuthDeleteRating",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AuthDeleteRating(ctx, req.(*AuthDeleteRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "AuthCreateAccommodation",
			Handler:    _AuthService_AuthCreateAccommodation_Handler,
		},
		{
			MethodName: "AuthUpdateUser",
			Handler:    _AuthService_AuthUpdateUser_Handler,
		},
		{
			MethodName: "AuthDeleteUser",
			Handler:    _AuthService_AuthDeleteUser_Handler,
		},
		{
			MethodName: "AuthGuestReserveAccommodation",
			Handler:    _AuthService_AuthGuestReserveAccommodation_Handler,
		},
		{
			MethodName: "AuthBookingAccept",
			Handler:    _AuthService_AuthBookingAccept_Handler,
		},
		{
			MethodName: "AuthBookingDeny",
			Handler:    _AuthService_AuthBookingDeny_Handler,
		},
		{
			MethodName: "AuthReservationCanceling",
			Handler:    _AuthService_AuthReservationCanceling_Handler,
		},
		{
			MethodName: "AuthDefineCustomPrice",
			Handler:    _AuthService_AuthDefineCustomPrice_Handler,
		},
		{
			MethodName: "AuthGetAccommodationByHostId",
			Handler:    _AuthService_AuthGetAccommodationByHostId_Handler,
		},
		{
			MethodName: "AuthGetBookingsByAccommodationId",
			Handler:    _AuthService_AuthGetBookingsByAccommodationId_Handler,
		},
		{
			MethodName: "AuthCreateRating",
			Handler:    _AuthService_AuthCreateRating_Handler,
		},
		{
			MethodName: "AuthDeleteRating",
			Handler:    _AuthService_AuthDeleteRating_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth_service.proto",
}
