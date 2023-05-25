// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.0
// source: booking_service.proto

package booking

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

// BookingServiceClient is the client API for BookingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookingServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	CreateBooking(ctx context.Context, in *CreateBookingRequest, opts ...grpc.CallOption) (*CreateBookingResponse, error)
	GuestReserveAccommodation(ctx context.Context, in *GuestReserveAccommodationRequest, opts ...grpc.CallOption) (*GuestReserveAccommodationResponse, error)
	BookingAccept(ctx context.Context, in *BookingAcceptRequest, opts ...grpc.CallOption) (*BookingAcceptResponse, error)
	BookingDeny(ctx context.Context, in *BookingDenyRequest, opts ...grpc.CallOption) (*BookingDenyResponse, error)
	GetByAccomodationIdandDataRange(ctx context.Context, in *GetByAccomodationIdandDataRangeRequest, opts ...grpc.CallOption) (*GetByAccomodationIdandDataRangeResponse, error)
	ReservationCanceling(ctx context.Context, in *ReservationCancelingRequest, opts ...grpc.CallOption) (*ReservationCancelingResponse, error)
	GetAllByUserAndType(ctx context.Context, in *GetAllByUserRequest, opts ...grpc.CallOption) (*GetAllByUserResponse, error)
	DeleteBooking(ctx context.Context, in *DeleteBookingRequest, opts ...grpc.CallOption) (*DeleteBookingResponse, error)
	DeleteBookingsByGuestId(ctx context.Context, in *DeleteBookingByGuestIdRequest, opts ...grpc.CallOption) (*DeleteBookingByGuestIdResponse, error)
	DeleteBookingsByAccommodationId(ctx context.Context, in *DeleteBookingByAccommodationIdRequest, opts ...grpc.CallOption) (*DeleteBookingByAccommodationIdResponse, error)
	GetByAccommodationId(ctx context.Context, in *GetByAccommodationIdRequest, opts ...grpc.CallOption) (*GetByAccommodationIdResponse, error)
	GetBookingByAccommodationAndGuestId(ctx context.Context, in *GetBookingByAccommodationAndGuestIdRequest, opts ...grpc.CallOption) (*GetBookingByAccommodationAndGuestIdResponse, error)
}

type bookingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookingServiceClient(cc grpc.ClientConnInterface) BookingServiceClient {
	return &bookingServiceClient{cc}
}

func (c *bookingServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/booking.BookingService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/booking.BookingService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) CreateBooking(ctx context.Context, in *CreateBookingRequest, opts ...grpc.CallOption) (*CreateBookingResponse, error) {
	out := new(CreateBookingResponse)
	err := c.cc.Invoke(ctx, "/booking.BookingService/CreateBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GuestReserveAccommodation(ctx context.Context, in *GuestReserveAccommodationRequest, opts ...grpc.CallOption) (*GuestReserveAccommodationResponse, error) {
	out := new(GuestReserveAccommodationResponse)
	err := c.cc.Invoke(ctx, "/booking.BookingService/GuestReserveAccommodation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) BookingAccept(ctx context.Context, in *BookingAcceptRequest, opts ...grpc.CallOption) (*BookingAcceptResponse, error) {
	out := new(BookingAcceptResponse)
	err := c.cc.Invoke(ctx, "/booking.BookingService/BookingAccept", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) BookingDeny(ctx context.Context, in *BookingDenyRequest, opts ...grpc.CallOption) (*BookingDenyResponse, error) {
	out := new(BookingDenyResponse)
	err := c.cc.Invoke(ctx, "/booking.BookingService/BookingDeny", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetByAccomodationIdandDataRange(ctx context.Context, in *GetByAccomodationIdandDataRangeRequest, opts ...grpc.CallOption) (*GetByAccomodationIdandDataRangeResponse, error) {
	out := new(GetByAccomodationIdandDataRangeResponse)
	err := c.cc.Invoke(ctx, "/booking.BookingService/GetByAccomodationIdandDataRange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) ReservationCanceling(ctx context.Context, in *ReservationCancelingRequest, opts ...grpc.CallOption) (*ReservationCancelingResponse, error) {
	out := new(ReservationCancelingResponse)
	err := c.cc.Invoke(ctx, "/booking.BookingService/ReservationCanceling", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetAllByUserAndType(ctx context.Context, in *GetAllByUserRequest, opts ...grpc.CallOption) (*GetAllByUserResponse, error) {
	out := new(GetAllByUserResponse)
	err := c.cc.Invoke(ctx, "/booking.BookingService/GetAllByUserAndType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) DeleteBooking(ctx context.Context, in *DeleteBookingRequest, opts ...grpc.CallOption) (*DeleteBookingResponse, error) {
	out := new(DeleteBookingResponse)
	err := c.cc.Invoke(ctx, "/booking.BookingService/DeleteBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) DeleteBookingsByGuestId(ctx context.Context, in *DeleteBookingByGuestIdRequest, opts ...grpc.CallOption) (*DeleteBookingByGuestIdResponse, error) {
	out := new(DeleteBookingByGuestIdResponse)
	err := c.cc.Invoke(ctx, "/booking.BookingService/DeleteBookingsByGuestId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) DeleteBookingsByAccommodationId(ctx context.Context, in *DeleteBookingByAccommodationIdRequest, opts ...grpc.CallOption) (*DeleteBookingByAccommodationIdResponse, error) {
	out := new(DeleteBookingByAccommodationIdResponse)
	err := c.cc.Invoke(ctx, "/booking.BookingService/DeleteBookingsByAccommodationId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetByAccommodationId(ctx context.Context, in *GetByAccommodationIdRequest, opts ...grpc.CallOption) (*GetByAccommodationIdResponse, error) {
	out := new(GetByAccommodationIdResponse)
	err := c.cc.Invoke(ctx, "/booking.BookingService/GetByAccommodationId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetBookingByAccommodationAndGuestId(ctx context.Context, in *GetBookingByAccommodationAndGuestIdRequest, opts ...grpc.CallOption) (*GetBookingByAccommodationAndGuestIdResponse, error) {
	out := new(GetBookingByAccommodationAndGuestIdResponse)
	err := c.cc.Invoke(ctx, "/booking.BookingService/GetBookingByAccommodationAndGuestId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookingServiceServer is the server API for BookingService service.
// All implementations must embed UnimplementedBookingServiceServer
// for forward compatibility
type BookingServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	CreateBooking(context.Context, *CreateBookingRequest) (*CreateBookingResponse, error)
	GuestReserveAccommodation(context.Context, *GuestReserveAccommodationRequest) (*GuestReserveAccommodationResponse, error)
	BookingAccept(context.Context, *BookingAcceptRequest) (*BookingAcceptResponse, error)
	BookingDeny(context.Context, *BookingDenyRequest) (*BookingDenyResponse, error)
	GetByAccomodationIdandDataRange(context.Context, *GetByAccomodationIdandDataRangeRequest) (*GetByAccomodationIdandDataRangeResponse, error)
	ReservationCanceling(context.Context, *ReservationCancelingRequest) (*ReservationCancelingResponse, error)
	GetAllByUserAndType(context.Context, *GetAllByUserRequest) (*GetAllByUserResponse, error)
	DeleteBooking(context.Context, *DeleteBookingRequest) (*DeleteBookingResponse, error)
	DeleteBookingsByGuestId(context.Context, *DeleteBookingByGuestIdRequest) (*DeleteBookingByGuestIdResponse, error)
	DeleteBookingsByAccommodationId(context.Context, *DeleteBookingByAccommodationIdRequest) (*DeleteBookingByAccommodationIdResponse, error)
	GetByAccommodationId(context.Context, *GetByAccommodationIdRequest) (*GetByAccommodationIdResponse, error)
	GetBookingByAccommodationAndGuestId(context.Context, *GetBookingByAccommodationAndGuestIdRequest) (*GetBookingByAccommodationAndGuestIdResponse, error)
	mustEmbedUnimplementedBookingServiceServer()
}

// UnimplementedBookingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookingServiceServer struct {
}

func (UnimplementedBookingServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedBookingServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedBookingServiceServer) CreateBooking(context.Context, *CreateBookingRequest) (*CreateBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBooking not implemented")
}
func (UnimplementedBookingServiceServer) GuestReserveAccommodation(context.Context, *GuestReserveAccommodationRequest) (*GuestReserveAccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GuestReserveAccommodation not implemented")
}
func (UnimplementedBookingServiceServer) BookingAccept(context.Context, *BookingAcceptRequest) (*BookingAcceptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BookingAccept not implemented")
}
func (UnimplementedBookingServiceServer) BookingDeny(context.Context, *BookingDenyRequest) (*BookingDenyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BookingDeny not implemented")
}
func (UnimplementedBookingServiceServer) GetByAccomodationIdandDataRange(context.Context, *GetByAccomodationIdandDataRangeRequest) (*GetByAccomodationIdandDataRangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByAccomodationIdandDataRange not implemented")
}
func (UnimplementedBookingServiceServer) ReservationCanceling(context.Context, *ReservationCancelingRequest) (*ReservationCancelingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReservationCanceling not implemented")
}
func (UnimplementedBookingServiceServer) GetAllByUserAndType(context.Context, *GetAllByUserRequest) (*GetAllByUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllByUserAndType not implemented")
}
func (UnimplementedBookingServiceServer) DeleteBooking(context.Context, *DeleteBookingRequest) (*DeleteBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBooking not implemented")
}
func (UnimplementedBookingServiceServer) DeleteBookingsByGuestId(context.Context, *DeleteBookingByGuestIdRequest) (*DeleteBookingByGuestIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBookingsByGuestId not implemented")
}
func (UnimplementedBookingServiceServer) DeleteBookingsByAccommodationId(context.Context, *DeleteBookingByAccommodationIdRequest) (*DeleteBookingByAccommodationIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBookingsByAccommodationId not implemented")
}
func (UnimplementedBookingServiceServer) GetByAccommodationId(context.Context, *GetByAccommodationIdRequest) (*GetByAccommodationIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByAccommodationId not implemented")
}
func (UnimplementedBookingServiceServer) GetBookingByAccommodationAndGuestId(context.Context, *GetBookingByAccommodationAndGuestIdRequest) (*GetBookingByAccommodationAndGuestIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookingByAccommodationAndGuestId not implemented")
}
func (UnimplementedBookingServiceServer) mustEmbedUnimplementedBookingServiceServer() {}

// UnsafeBookingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookingServiceServer will
// result in compilation errors.
type UnsafeBookingServiceServer interface {
	mustEmbedUnimplementedBookingServiceServer()
}

func RegisterBookingServiceServer(s grpc.ServiceRegistrar, srv BookingServiceServer) {
	s.RegisterService(&BookingService_ServiceDesc, srv)
}

func _BookingService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_CreateBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).CreateBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/CreateBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).CreateBooking(ctx, req.(*CreateBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GuestReserveAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GuestReserveAccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GuestReserveAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/GuestReserveAccommodation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GuestReserveAccommodation(ctx, req.(*GuestReserveAccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_BookingAccept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookingAcceptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).BookingAccept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/BookingAccept",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).BookingAccept(ctx, req.(*BookingAcceptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_BookingDeny_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookingDenyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).BookingDeny(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/BookingDeny",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).BookingDeny(ctx, req.(*BookingDenyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetByAccomodationIdandDataRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByAccomodationIdandDataRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetByAccomodationIdandDataRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/GetByAccomodationIdandDataRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetByAccomodationIdandDataRange(ctx, req.(*GetByAccomodationIdandDataRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_ReservationCanceling_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReservationCancelingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).ReservationCanceling(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/ReservationCanceling",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).ReservationCanceling(ctx, req.(*ReservationCancelingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetAllByUserAndType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllByUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetAllByUserAndType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/GetAllByUserAndType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetAllByUserAndType(ctx, req.(*GetAllByUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_DeleteBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).DeleteBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/DeleteBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).DeleteBooking(ctx, req.(*DeleteBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_DeleteBookingsByGuestId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookingByGuestIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).DeleteBookingsByGuestId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/DeleteBookingsByGuestId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).DeleteBookingsByGuestId(ctx, req.(*DeleteBookingByGuestIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_DeleteBookingsByAccommodationId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookingByAccommodationIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).DeleteBookingsByAccommodationId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/DeleteBookingsByAccommodationId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).DeleteBookingsByAccommodationId(ctx, req.(*DeleteBookingByAccommodationIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetByAccommodationId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByAccommodationIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetByAccommodationId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/GetByAccommodationId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetByAccommodationId(ctx, req.(*GetByAccommodationIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetBookingByAccommodationAndGuestId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookingByAccommodationAndGuestIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetBookingByAccommodationAndGuestId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.BookingService/GetBookingByAccommodationAndGuestId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetBookingByAccommodationAndGuestId(ctx, req.(*GetBookingByAccommodationAndGuestIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookingService_ServiceDesc is the grpc.ServiceDesc for BookingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "booking.BookingService",
	HandlerType: (*BookingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _BookingService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _BookingService_GetAll_Handler,
		},
		{
			MethodName: "CreateBooking",
			Handler:    _BookingService_CreateBooking_Handler,
		},
		{
			MethodName: "GuestReserveAccommodation",
			Handler:    _BookingService_GuestReserveAccommodation_Handler,
		},
		{
			MethodName: "BookingAccept",
			Handler:    _BookingService_BookingAccept_Handler,
		},
		{
			MethodName: "BookingDeny",
			Handler:    _BookingService_BookingDeny_Handler,
		},
		{
			MethodName: "GetByAccomodationIdandDataRange",
			Handler:    _BookingService_GetByAccomodationIdandDataRange_Handler,
		},
		{
			MethodName: "ReservationCanceling",
			Handler:    _BookingService_ReservationCanceling_Handler,
		},
		{
			MethodName: "GetAllByUserAndType",
			Handler:    _BookingService_GetAllByUserAndType_Handler,
		},
		{
			MethodName: "DeleteBooking",
			Handler:    _BookingService_DeleteBooking_Handler,
		},
		{
			MethodName: "DeleteBookingsByGuestId",
			Handler:    _BookingService_DeleteBookingsByGuestId_Handler,
		},
		{
			MethodName: "DeleteBookingsByAccommodationId",
			Handler:    _BookingService_DeleteBookingsByAccommodationId_Handler,
		},
		{
			MethodName: "GetByAccommodationId",
			Handler:    _BookingService_GetByAccommodationId_Handler,
		},
		{
			MethodName: "GetBookingByAccommodationAndGuestId",
			Handler:    _BookingService_GetBookingByAccommodationAndGuestId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "booking_service.proto",
}
