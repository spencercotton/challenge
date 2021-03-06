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

// CouponServiceClient is the client API for CouponService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CouponServiceClient interface {
	CreateCoupon(ctx context.Context, in *CreateCouponRequest, opts ...grpc.CallOption) (*Coupon, error)
	GetCoupon(ctx context.Context, in *GetCouponRequest, opts ...grpc.CallOption) (*Coupon, error)
	GetCoupons(ctx context.Context, in *GetCouponsRequest, opts ...grpc.CallOption) (*GetCouponsResponse, error)
	DeleteCoupon(ctx context.Context, in *DeleteCouponRequest, opts ...grpc.CallOption) (*DeleteCouponResponse, error)
}

type couponServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCouponServiceClient(cc grpc.ClientConnInterface) CouponServiceClient {
	return &couponServiceClient{cc}
}

func (c *couponServiceClient) CreateCoupon(ctx context.Context, in *CreateCouponRequest, opts ...grpc.CallOption) (*Coupon, error) {
	out := new(Coupon)
	err := c.cc.Invoke(ctx, "/coupon.CouponService/CreateCoupon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *couponServiceClient) GetCoupon(ctx context.Context, in *GetCouponRequest, opts ...grpc.CallOption) (*Coupon, error) {
	out := new(Coupon)
	err := c.cc.Invoke(ctx, "/coupon.CouponService/GetCoupon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *couponServiceClient) GetCoupons(ctx context.Context, in *GetCouponsRequest, opts ...grpc.CallOption) (*GetCouponsResponse, error) {
	out := new(GetCouponsResponse)
	err := c.cc.Invoke(ctx, "/coupon.CouponService/GetCoupons", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *couponServiceClient) DeleteCoupon(ctx context.Context, in *DeleteCouponRequest, opts ...grpc.CallOption) (*DeleteCouponResponse, error) {
	out := new(DeleteCouponResponse)
	err := c.cc.Invoke(ctx, "/coupon.CouponService/DeleteCoupon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CouponServiceServer is the server API for CouponService service.
// All implementations must embed UnimplementedCouponServiceServer
// for forward compatibility
type CouponServiceServer interface {
	CreateCoupon(context.Context, *CreateCouponRequest) (*Coupon, error)
	GetCoupon(context.Context, *GetCouponRequest) (*Coupon, error)
	GetCoupons(context.Context, *GetCouponsRequest) (*GetCouponsResponse, error)
	DeleteCoupon(context.Context, *DeleteCouponRequest) (*DeleteCouponResponse, error)
	mustEmbedUnimplementedCouponServiceServer()
}

// UnimplementedCouponServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCouponServiceServer struct {
}

func (UnimplementedCouponServiceServer) CreateCoupon(context.Context, *CreateCouponRequest) (*Coupon, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCoupon not implemented")
}
func (UnimplementedCouponServiceServer) GetCoupon(context.Context, *GetCouponRequest) (*Coupon, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoupon not implemented")
}
func (UnimplementedCouponServiceServer) GetCoupons(context.Context, *GetCouponsRequest) (*GetCouponsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoupons not implemented")
}
func (UnimplementedCouponServiceServer) DeleteCoupon(context.Context, *DeleteCouponRequest) (*DeleteCouponResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCoupon not implemented")
}
func (UnimplementedCouponServiceServer) mustEmbedUnimplementedCouponServiceServer() {}

// UnsafeCouponServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CouponServiceServer will
// result in compilation errors.
type UnsafeCouponServiceServer interface {
	mustEmbedUnimplementedCouponServiceServer()
}

func RegisterCouponServiceServer(s grpc.ServiceRegistrar, srv CouponServiceServer) {
	s.RegisterService(&CouponService_ServiceDesc, srv)
}

func _CouponService_CreateCoupon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCouponRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CouponServiceServer).CreateCoupon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coupon.CouponService/CreateCoupon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CouponServiceServer).CreateCoupon(ctx, req.(*CreateCouponRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CouponService_GetCoupon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCouponRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CouponServiceServer).GetCoupon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coupon.CouponService/GetCoupon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CouponServiceServer).GetCoupon(ctx, req.(*GetCouponRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CouponService_GetCoupons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCouponsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CouponServiceServer).GetCoupons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coupon.CouponService/GetCoupons",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CouponServiceServer).GetCoupons(ctx, req.(*GetCouponsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CouponService_DeleteCoupon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCouponRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CouponServiceServer).DeleteCoupon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coupon.CouponService/DeleteCoupon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CouponServiceServer).DeleteCoupon(ctx, req.(*DeleteCouponRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CouponService_ServiceDesc is the grpc.ServiceDesc for CouponService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CouponService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "coupon.CouponService",
	HandlerType: (*CouponServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCoupon",
			Handler:    _CouponService_CreateCoupon_Handler,
		},
		{
			MethodName: "GetCoupon",
			Handler:    _CouponService_GetCoupon_Handler,
		},
		{
			MethodName: "GetCoupons",
			Handler:    _CouponService_GetCoupons_Handler,
		},
		{
			MethodName: "DeleteCoupon",
			Handler:    _CouponService_DeleteCoupon_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
