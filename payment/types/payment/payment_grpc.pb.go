// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: pb/payment.proto

package payment

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

// PaymentClient is the client API for Payment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentClient interface {
	// 创建支付单
	CreatePayment(ctx context.Context, in *CreatePaymentReq, opts ...grpc.CallOption) (*CreatePaymentReply, error)
	// 支付结果通知
	OrderPaymentNoticeTccTry(ctx context.Context, in *PaymentNoticeReq, opts ...grpc.CallOption) (*PaymentNoticePayReply, error)
	OrderPaymentNoticeTccConfirm(ctx context.Context, in *PaymentNoticeReq, opts ...grpc.CallOption) (*PaymentNoticePayReply, error)
	OrderPaymentNoticeTccCancel(ctx context.Context, in *PaymentNoticeReq, opts ...grpc.CallOption) (*PaymentNoticePayReply, error)
}

type paymentClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentClient(cc grpc.ClientConnInterface) PaymentClient {
	return &paymentClient{cc}
}

func (c *paymentClient) CreatePayment(ctx context.Context, in *CreatePaymentReq, opts ...grpc.CallOption) (*CreatePaymentReply, error) {
	out := new(CreatePaymentReply)
	err := c.cc.Invoke(ctx, "/payment.payment/createPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentClient) OrderPaymentNoticeTccTry(ctx context.Context, in *PaymentNoticeReq, opts ...grpc.CallOption) (*PaymentNoticePayReply, error) {
	out := new(PaymentNoticePayReply)
	err := c.cc.Invoke(ctx, "/payment.payment/orderPaymentNoticeTccTry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentClient) OrderPaymentNoticeTccConfirm(ctx context.Context, in *PaymentNoticeReq, opts ...grpc.CallOption) (*PaymentNoticePayReply, error) {
	out := new(PaymentNoticePayReply)
	err := c.cc.Invoke(ctx, "/payment.payment/orderPaymentNoticeTccConfirm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentClient) OrderPaymentNoticeTccCancel(ctx context.Context, in *PaymentNoticeReq, opts ...grpc.CallOption) (*PaymentNoticePayReply, error) {
	out := new(PaymentNoticePayReply)
	err := c.cc.Invoke(ctx, "/payment.payment/orderPaymentNoticeTccCancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServer is the server API for Payment service.
// All implementations must embed UnimplementedPaymentServer
// for forward compatibility
type PaymentServer interface {
	// 创建支付单
	CreatePayment(context.Context, *CreatePaymentReq) (*CreatePaymentReply, error)
	// 支付结果通知
	OrderPaymentNoticeTccTry(context.Context, *PaymentNoticeReq) (*PaymentNoticePayReply, error)
	OrderPaymentNoticeTccConfirm(context.Context, *PaymentNoticeReq) (*PaymentNoticePayReply, error)
	OrderPaymentNoticeTccCancel(context.Context, *PaymentNoticeReq) (*PaymentNoticePayReply, error)
	mustEmbedUnimplementedPaymentServer()
}

// UnimplementedPaymentServer must be embedded to have forward compatible implementations.
type UnimplementedPaymentServer struct {
}

func (UnimplementedPaymentServer) CreatePayment(context.Context, *CreatePaymentReq) (*CreatePaymentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePayment not implemented")
}
func (UnimplementedPaymentServer) OrderPaymentNoticeTccTry(context.Context, *PaymentNoticeReq) (*PaymentNoticePayReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderPaymentNoticeTccTry not implemented")
}
func (UnimplementedPaymentServer) OrderPaymentNoticeTccConfirm(context.Context, *PaymentNoticeReq) (*PaymentNoticePayReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderPaymentNoticeTccConfirm not implemented")
}
func (UnimplementedPaymentServer) OrderPaymentNoticeTccCancel(context.Context, *PaymentNoticeReq) (*PaymentNoticePayReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderPaymentNoticeTccCancel not implemented")
}
func (UnimplementedPaymentServer) mustEmbedUnimplementedPaymentServer() {}

// UnsafePaymentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentServer will
// result in compilation errors.
type UnsafePaymentServer interface {
	mustEmbedUnimplementedPaymentServer()
}

func RegisterPaymentServer(s grpc.ServiceRegistrar, srv PaymentServer) {
	s.RegisterService(&Payment_ServiceDesc, srv)
}

func _Payment_CreatePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePaymentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).CreatePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.payment/createPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).CreatePayment(ctx, req.(*CreatePaymentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_OrderPaymentNoticeTccTry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentNoticeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).OrderPaymentNoticeTccTry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.payment/orderPaymentNoticeTccTry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).OrderPaymentNoticeTccTry(ctx, req.(*PaymentNoticeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_OrderPaymentNoticeTccConfirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentNoticeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).OrderPaymentNoticeTccConfirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.payment/orderPaymentNoticeTccConfirm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).OrderPaymentNoticeTccConfirm(ctx, req.(*PaymentNoticeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_OrderPaymentNoticeTccCancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentNoticeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).OrderPaymentNoticeTccCancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.payment/orderPaymentNoticeTccCancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).OrderPaymentNoticeTccCancel(ctx, req.(*PaymentNoticeReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Payment_ServiceDesc is the grpc.ServiceDesc for Payment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Payment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "payment.payment",
	HandlerType: (*PaymentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createPayment",
			Handler:    _Payment_CreatePayment_Handler,
		},
		{
			MethodName: "orderPaymentNoticeTccTry",
			Handler:    _Payment_OrderPaymentNoticeTccTry_Handler,
		},
		{
			MethodName: "orderPaymentNoticeTccConfirm",
			Handler:    _Payment_OrderPaymentNoticeTccConfirm_Handler,
		},
		{
			MethodName: "orderPaymentNoticeTccCancel",
			Handler:    _Payment_OrderPaymentNoticeTccCancel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/payment.proto",
}
