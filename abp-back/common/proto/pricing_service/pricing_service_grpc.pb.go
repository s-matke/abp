// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: pricing_service/pricing_service.proto

package pricing

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
	PricingService_Get_FullMethodName                = "/pricing.PricingService/Get"
	PricingService_GetByAccommodation_FullMethodName = "/pricing.PricingService/GetByAccommodation"
	PricingService_GetAll_FullMethodName             = "/pricing.PricingService/GetAll"
	PricingService_CreatePricing_FullMethodName      = "/pricing.PricingService/CreatePricing"
)

// PricingServiceClient is the client API for PricingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PricingServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetByAccommodation(ctx context.Context, in *GetByAccommodationRequest, opts ...grpc.CallOption) (*GetByAccommodationResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	CreatePricing(ctx context.Context, in *CreatePricingRequest, opts ...grpc.CallOption) (*CreatePricingResponse, error)
}

type pricingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPricingServiceClient(cc grpc.ClientConnInterface) PricingServiceClient {
	return &pricingServiceClient{cc}
}

func (c *pricingServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, PricingService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pricingServiceClient) GetByAccommodation(ctx context.Context, in *GetByAccommodationRequest, opts ...grpc.CallOption) (*GetByAccommodationResponse, error) {
	out := new(GetByAccommodationResponse)
	err := c.cc.Invoke(ctx, PricingService_GetByAccommodation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pricingServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, PricingService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pricingServiceClient) CreatePricing(ctx context.Context, in *CreatePricingRequest, opts ...grpc.CallOption) (*CreatePricingResponse, error) {
	out := new(CreatePricingResponse)
	err := c.cc.Invoke(ctx, PricingService_CreatePricing_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PricingServiceServer is the server API for PricingService service.
// All implementations must embed UnimplementedPricingServiceServer
// for forward compatibility
type PricingServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetByAccommodation(context.Context, *GetByAccommodationRequest) (*GetByAccommodationResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	CreatePricing(context.Context, *CreatePricingRequest) (*CreatePricingResponse, error)
	mustEmbedUnimplementedPricingServiceServer()
}

// UnimplementedPricingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPricingServiceServer struct {
}

func (UnimplementedPricingServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPricingServiceServer) GetByAccommodation(context.Context, *GetByAccommodationRequest) (*GetByAccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByAccommodation not implemented")
}
func (UnimplementedPricingServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedPricingServiceServer) CreatePricing(context.Context, *CreatePricingRequest) (*CreatePricingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePricing not implemented")
}
func (UnimplementedPricingServiceServer) mustEmbedUnimplementedPricingServiceServer() {}

// UnsafePricingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PricingServiceServer will
// result in compilation errors.
type UnsafePricingServiceServer interface {
	mustEmbedUnimplementedPricingServiceServer()
}

func RegisterPricingServiceServer(s grpc.ServiceRegistrar, srv PricingServiceServer) {
	s.RegisterService(&PricingService_ServiceDesc, srv)
}

func _PricingService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricingServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PricingService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricingServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PricingService_GetByAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByAccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricingServiceServer).GetByAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PricingService_GetByAccommodation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricingServiceServer).GetByAccommodation(ctx, req.(*GetByAccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PricingService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricingServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PricingService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricingServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PricingService_CreatePricing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePricingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PricingServiceServer).CreatePricing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PricingService_CreatePricing_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PricingServiceServer).CreatePricing(ctx, req.(*CreatePricingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PricingService_ServiceDesc is the grpc.ServiceDesc for PricingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PricingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pricing.PricingService",
	HandlerType: (*PricingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _PricingService_Get_Handler,
		},
		{
			MethodName: "GetByAccommodation",
			Handler:    _PricingService_GetByAccommodation_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _PricingService_GetAll_Handler,
		},
		{
			MethodName: "CreatePricing",
			Handler:    _PricingService_CreatePricing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pricing_service/pricing_service.proto",
}
