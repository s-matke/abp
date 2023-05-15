// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: reservation_service/reservation_service.proto

package reservation

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

const (
	ReservationService_GetAll_FullMethodName                       = "/reservation.ReservationService/GetAll"
	ReservationService_GetByAccommodation_FullMethodName           = "/reservation.ReservationService/GetByAccommodation"
	ReservationService_CreateReservation_FullMethodName            = "/reservation.ReservationService/CreateReservation"
	ReservationService_GetAllPendingByAccommodation_FullMethodName = "/reservation.ReservationService/GetAllPendingByAccommodation"
	ReservationService_ConfirmReservation_FullMethodName           = "/reservation.ReservationService/ConfirmReservation"
	ReservationService_GetByGuest_FullMethodName                   = "/reservation.ReservationService/GetByGuest"
)

// ReservationServiceClient is the client API for ReservationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReservationServiceClient interface {
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	GetByAccommodation(ctx context.Context, in *GetByAccommodationRequest, opts ...grpc.CallOption) (*GetByAccommodationResponse, error)
	CreateReservation(ctx context.Context, in *CreateReservationRequest, opts ...grpc.CallOption) (*CreateReservationResponse, error)
	GetAllPendingByAccommodation(ctx context.Context, in *GetAllPendingByAccommodationRequest, opts ...grpc.CallOption) (*GetAllPendingByAccommodationResponse, error)
	ConfirmReservation(ctx context.Context, in *ConfirmReservationRequest, opts ...grpc.CallOption) (*ConfirmReservationResponse, error)
	GetByGuest(ctx context.Context, in *GetByGuestRequest, opts ...grpc.CallOption) (*GetByGuestResponse, error)
}

type reservationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReservationServiceClient(cc grpc.ClientConnInterface) ReservationServiceClient {
	return &reservationServiceClient{cc}
}

func (c *reservationServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, ReservationService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetByAccommodation(ctx context.Context, in *GetByAccommodationRequest, opts ...grpc.CallOption) (*GetByAccommodationResponse, error) {
	out := new(GetByAccommodationResponse)
	err := c.cc.Invoke(ctx, ReservationService_GetByAccommodation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) CreateReservation(ctx context.Context, in *CreateReservationRequest, opts ...grpc.CallOption) (*CreateReservationResponse, error) {
	out := new(CreateReservationResponse)
	err := c.cc.Invoke(ctx, ReservationService_CreateReservation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetAllPendingByAccommodation(ctx context.Context, in *GetAllPendingByAccommodationRequest, opts ...grpc.CallOption) (*GetAllPendingByAccommodationResponse, error) {
	out := new(GetAllPendingByAccommodationResponse)
	err := c.cc.Invoke(ctx, ReservationService_GetAllPendingByAccommodation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) ConfirmReservation(ctx context.Context, in *ConfirmReservationRequest, opts ...grpc.CallOption) (*ConfirmReservationResponse, error) {
	out := new(ConfirmReservationResponse)
	err := c.cc.Invoke(ctx, ReservationService_ConfirmReservation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetByGuest(ctx context.Context, in *GetByGuestRequest, opts ...grpc.CallOption) (*GetByGuestResponse, error) {
	out := new(GetByGuestResponse)
	err := c.cc.Invoke(ctx, ReservationService_GetByGuest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReservationServiceServer is the server API for ReservationService service.
// All implementations must embed UnimplementedReservationServiceServer
// for forward compatibility
type ReservationServiceServer interface {
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	GetByAccommodation(context.Context, *GetByAccommodationRequest) (*GetByAccommodationResponse, error)
	CreateReservation(context.Context, *CreateReservationRequest) (*CreateReservationResponse, error)
	GetAllPendingByAccommodation(context.Context, *GetAllPendingByAccommodationRequest) (*GetAllPendingByAccommodationResponse, error)
	ConfirmReservation(context.Context, *ConfirmReservationRequest) (*ConfirmReservationResponse, error)
	GetByGuest(context.Context, *GetByGuestRequest) (*GetByGuestResponse, error)
	mustEmbedUnimplementedReservationServiceServer()
}

// UnimplementedReservationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReservationServiceServer struct {
}

func (UnimplementedReservationServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedReservationServiceServer) GetByAccommodation(context.Context, *GetByAccommodationRequest) (*GetByAccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByAccommodation not implemented")
}
func (UnimplementedReservationServiceServer) CreateReservation(context.Context, *CreateReservationRequest) (*CreateReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReservation not implemented")
}
func (UnimplementedReservationServiceServer) GetAllPendingByAccommodation(context.Context, *GetAllPendingByAccommodationRequest) (*GetAllPendingByAccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPendingByAccommodation not implemented")
}
func (UnimplementedReservationServiceServer) ConfirmReservation(context.Context, *ConfirmReservationRequest) (*ConfirmReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmReservation not implemented")
}
func (UnimplementedReservationServiceServer) GetByGuest(context.Context, *GetByGuestRequest) (*GetByGuestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByGuest not implemented")
}
func (UnimplementedReservationServiceServer) mustEmbedUnimplementedReservationServiceServer() {}

// UnsafeReservationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReservationServiceServer will
// result in compilation errors.
type UnsafeReservationServiceServer interface {
	mustEmbedUnimplementedReservationServiceServer()
}

func RegisterReservationServiceServer(s grpc.ServiceRegistrar, srv ReservationServiceServer) {
	s.RegisterService(&ReservationService_ServiceDesc, srv)
}

func _ReservationService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetByAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByAccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetByAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_GetByAccommodation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetByAccommodation(ctx, req.(*GetByAccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_CreateReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).CreateReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_CreateReservation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).CreateReservation(ctx, req.(*CreateReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetAllPendingByAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllPendingByAccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetAllPendingByAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_GetAllPendingByAccommodation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetAllPendingByAccommodation(ctx, req.(*GetAllPendingByAccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_ConfirmReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).ConfirmReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_ConfirmReservation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).ConfirmReservation(ctx, req.(*ConfirmReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetByGuest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByGuestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetByGuest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_GetByGuest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetByGuest(ctx, req.(*GetByGuestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReservationService_ServiceDesc is the grpc.ServiceDesc for ReservationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReservationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reservation.ReservationService",
	HandlerType: (*ReservationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _ReservationService_GetAll_Handler,
		},
		{
			MethodName: "GetByAccommodation",
			Handler:    _ReservationService_GetByAccommodation_Handler,
		},
		{
			MethodName: "CreateReservation",
			Handler:    _ReservationService_CreateReservation_Handler,
		},
		{
			MethodName: "GetAllPendingByAccommodation",
			Handler:    _ReservationService_GetAllPendingByAccommodation_Handler,
		},
		{
			MethodName: "ConfirmReservation",
			Handler:    _ReservationService_ConfirmReservation_Handler,
		},
		{
			MethodName: "GetByGuest",
			Handler:    _ReservationService_GetByGuest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reservation_service/reservation_service.proto",
}
