// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: pb/order.proto

package order

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

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderClient interface {
	// 查询订单
	GetOrder(ctx context.Context, in *GetOrderReq, opts ...grpc.CallOption) (*GetOrderReply, error)
	// 创建订单
	CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*CreateOrderReply, error)
	// 关闭订单
	CloseOrder(ctx context.Context, in *CloseOrderReq, opts ...grpc.CallOption) (*CloseOrderReply, error)
	// 取消订单
	CancelOrder(ctx context.Context, in *CancelOrderReq, opts ...grpc.CallOption) (*CancelOrderReply, error)
	// 删除订单
	DeleteOrder(ctx context.Context, in *DeleteOrderReq, opts ...grpc.CallOption) (*DeleteOrderReply, error)
	// 订单支付成功TccTry
	PaymentSuccessOrderTccTry(ctx context.Context, in *PaymentSuccessTccReq, opts ...grpc.CallOption) (*PaymentSuccessTccReply, error)
	// 订单支付成功TccConfirm
	PaymentSuccessOrderTccConfirm(ctx context.Context, in *PaymentSuccessTccReq, opts ...grpc.CallOption) (*PaymentSuccessTccReply, error)
	// 订单支付成功TccCancel
	PaymentSuccessOrderTccCancel(ctx context.Context, in *PaymentSuccessTccReq, opts ...grpc.CallOption) (*PaymentSuccessTccReply, error)
}

type orderClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderClient(cc grpc.ClientConnInterface) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) GetOrder(ctx context.Context, in *GetOrderReq, opts ...grpc.CallOption) (*GetOrderReply, error) {
	out := new(GetOrderReply)
	err := c.cc.Invoke(ctx, "/order.order/getOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*CreateOrderReply, error) {
	out := new(CreateOrderReply)
	err := c.cc.Invoke(ctx, "/order.order/createOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) CloseOrder(ctx context.Context, in *CloseOrderReq, opts ...grpc.CallOption) (*CloseOrderReply, error) {
	out := new(CloseOrderReply)
	err := c.cc.Invoke(ctx, "/order.order/closeOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) CancelOrder(ctx context.Context, in *CancelOrderReq, opts ...grpc.CallOption) (*CancelOrderReply, error) {
	out := new(CancelOrderReply)
	err := c.cc.Invoke(ctx, "/order.order/cancelOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) DeleteOrder(ctx context.Context, in *DeleteOrderReq, opts ...grpc.CallOption) (*DeleteOrderReply, error) {
	out := new(DeleteOrderReply)
	err := c.cc.Invoke(ctx, "/order.order/deleteOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) PaymentSuccessOrderTccTry(ctx context.Context, in *PaymentSuccessTccReq, opts ...grpc.CallOption) (*PaymentSuccessTccReply, error) {
	out := new(PaymentSuccessTccReply)
	err := c.cc.Invoke(ctx, "/order.order/paymentSuccessOrderTccTry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) PaymentSuccessOrderTccConfirm(ctx context.Context, in *PaymentSuccessTccReq, opts ...grpc.CallOption) (*PaymentSuccessTccReply, error) {
	out := new(PaymentSuccessTccReply)
	err := c.cc.Invoke(ctx, "/order.order/paymentSuccessOrderTccConfirm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) PaymentSuccessOrderTccCancel(ctx context.Context, in *PaymentSuccessTccReq, opts ...grpc.CallOption) (*PaymentSuccessTccReply, error) {
	out := new(PaymentSuccessTccReply)
	err := c.cc.Invoke(ctx, "/order.order/paymentSuccessOrderTccCancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
// All implementations must embed UnimplementedOrderServer
// for forward compatibility
type OrderServer interface {
	// 查询订单
	GetOrder(context.Context, *GetOrderReq) (*GetOrderReply, error)
	// 创建订单
	CreateOrder(context.Context, *CreateOrderReq) (*CreateOrderReply, error)
	// 关闭订单
	CloseOrder(context.Context, *CloseOrderReq) (*CloseOrderReply, error)
	// 取消订单
	CancelOrder(context.Context, *CancelOrderReq) (*CancelOrderReply, error)
	// 删除订单
	DeleteOrder(context.Context, *DeleteOrderReq) (*DeleteOrderReply, error)
	// 订单支付成功TccTry
	PaymentSuccessOrderTccTry(context.Context, *PaymentSuccessTccReq) (*PaymentSuccessTccReply, error)
	// 订单支付成功TccConfirm
	PaymentSuccessOrderTccConfirm(context.Context, *PaymentSuccessTccReq) (*PaymentSuccessTccReply, error)
	// 订单支付成功TccCancel
	PaymentSuccessOrderTccCancel(context.Context, *PaymentSuccessTccReq) (*PaymentSuccessTccReply, error)
	mustEmbedUnimplementedOrderServer()
}

// UnimplementedOrderServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServer struct {
}

func (UnimplementedOrderServer) GetOrder(context.Context, *GetOrderReq) (*GetOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedOrderServer) CreateOrder(context.Context, *CreateOrderReq) (*CreateOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderServer) CloseOrder(context.Context, *CloseOrderReq) (*CloseOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseOrder not implemented")
}
func (UnimplementedOrderServer) CancelOrder(context.Context, *CancelOrderReq) (*CancelOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (UnimplementedOrderServer) DeleteOrder(context.Context, *DeleteOrderReq) (*DeleteOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrder not implemented")
}
func (UnimplementedOrderServer) PaymentSuccessOrderTccTry(context.Context, *PaymentSuccessTccReq) (*PaymentSuccessTccReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PaymentSuccessOrderTccTry not implemented")
}
func (UnimplementedOrderServer) PaymentSuccessOrderTccConfirm(context.Context, *PaymentSuccessTccReq) (*PaymentSuccessTccReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PaymentSuccessOrderTccConfirm not implemented")
}
func (UnimplementedOrderServer) PaymentSuccessOrderTccCancel(context.Context, *PaymentSuccessTccReq) (*PaymentSuccessTccReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PaymentSuccessOrderTccCancel not implemented")
}
func (UnimplementedOrderServer) mustEmbedUnimplementedOrderServer() {}

// UnsafeOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServer will
// result in compilation errors.
type UnsafeOrderServer interface {
	mustEmbedUnimplementedOrderServer()
}

func RegisterOrderServer(s grpc.ServiceRegistrar, srv OrderServer) {
	s.RegisterService(&Order_ServiceDesc, srv)
}

func _Order_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.order/getOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).GetOrder(ctx, req.(*GetOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.order/createOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CreateOrder(ctx, req.(*CreateOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_CloseOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CloseOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.order/closeOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CloseOrder(ctx, req.(*CloseOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.order/cancelOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CancelOrder(ctx, req.(*CancelOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_DeleteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).DeleteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.order/deleteOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).DeleteOrder(ctx, req.(*DeleteOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_PaymentSuccessOrderTccTry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentSuccessTccReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).PaymentSuccessOrderTccTry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.order/paymentSuccessOrderTccTry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).PaymentSuccessOrderTccTry(ctx, req.(*PaymentSuccessTccReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_PaymentSuccessOrderTccConfirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentSuccessTccReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).PaymentSuccessOrderTccConfirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.order/paymentSuccessOrderTccConfirm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).PaymentSuccessOrderTccConfirm(ctx, req.(*PaymentSuccessTccReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_PaymentSuccessOrderTccCancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentSuccessTccReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).PaymentSuccessOrderTccCancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.order/paymentSuccessOrderTccCancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).PaymentSuccessOrderTccCancel(ctx, req.(*PaymentSuccessTccReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Order_ServiceDesc is the grpc.ServiceDesc for Order service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Order_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getOrder",
			Handler:    _Order_GetOrder_Handler,
		},
		{
			MethodName: "createOrder",
			Handler:    _Order_CreateOrder_Handler,
		},
		{
			MethodName: "closeOrder",
			Handler:    _Order_CloseOrder_Handler,
		},
		{
			MethodName: "cancelOrder",
			Handler:    _Order_CancelOrder_Handler,
		},
		{
			MethodName: "deleteOrder",
			Handler:    _Order_DeleteOrder_Handler,
		},
		{
			MethodName: "paymentSuccessOrderTccTry",
			Handler:    _Order_PaymentSuccessOrderTccTry_Handler,
		},
		{
			MethodName: "paymentSuccessOrderTccConfirm",
			Handler:    _Order_PaymentSuccessOrderTccConfirm_Handler,
		},
		{
			MethodName: "paymentSuccessOrderTccCancel",
			Handler:    _Order_PaymentSuccessOrderTccCancel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/order.proto",
}
