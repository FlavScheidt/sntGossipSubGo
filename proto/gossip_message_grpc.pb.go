// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// GossipMessageClient is the client API for GossipMessage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GossipMessageClient interface {
	ToLibP2P(ctx context.Context, in *Gossip, opts ...grpc.CallOption) (*Control, error)
	ToRippled(ctx context.Context, in *Gossip, opts ...grpc.CallOption) (*Control, error)
}

type gossipMessageClient struct {
	cc grpc.ClientConnInterface
}

func NewGossipMessageClient(cc grpc.ClientConnInterface) GossipMessageClient {
	return &gossipMessageClient{cc}
}

func (c *gossipMessageClient) ToLibP2P(ctx context.Context, in *Gossip, opts ...grpc.CallOption) (*Control, error) {
	out := new(Control)
	err := c.cc.Invoke(ctx, "/gossipmessage.GossipMessage/toLibP2P", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gossipMessageClient) ToRippled(ctx context.Context, in *Gossip, opts ...grpc.CallOption) (*Control, error) {
	out := new(Control)
	err := c.cc.Invoke(ctx, "/gossipmessage.GossipMessage/toRippled", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GossipMessageServer is the server API for GossipMessage service.
// All implementations must embed UnimplementedGossipMessageServer
// for forward compatibility
type GossipMessageServer interface {
	ToLibP2P(context.Context, *Gossip) (*Control, error)
	ToRippled(context.Context, *Gossip) (*Control, error)
	mustEmbedUnimplementedGossipMessageServer()
}

// UnimplementedGossipMessageServer must be embedded to have forward compatible implementations.
type UnimplementedGossipMessageServer struct {
}

func (UnimplementedGossipMessageServer) ToLibP2P(context.Context, *Gossip) (*Control, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToLibP2P not implemented")
}
func (UnimplementedGossipMessageServer) ToRippled(context.Context, *Gossip) (*Control, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToRippled not implemented")
}
func (UnimplementedGossipMessageServer) mustEmbedUnimplementedGossipMessageServer() {}

// UnsafeGossipMessageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GossipMessageServer will
// result in compilation errors.
type UnsafeGossipMessageServer interface {
	mustEmbedUnimplementedGossipMessageServer()
}

func RegisterGossipMessageServer(s grpc.ServiceRegistrar, srv GossipMessageServer) {
	s.RegisterService(&GossipMessage_ServiceDesc, srv)
}

func _GossipMessage_ToLibP2P_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Gossip)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GossipMessageServer).ToLibP2P(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gossipmessage.GossipMessage/toLibP2P",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GossipMessageServer).ToLibP2P(ctx, req.(*Gossip))
	}
	return interceptor(ctx, in, info, handler)
}

func _GossipMessage_ToRippled_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Gossip)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GossipMessageServer).ToRippled(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gossipmessage.GossipMessage/toRippled",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GossipMessageServer).ToRippled(ctx, req.(*Gossip))
	}
	return interceptor(ctx, in, info, handler)
}

// GossipMessage_ServiceDesc is the grpc.ServiceDesc for GossipMessage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GossipMessage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gossipmessage.GossipMessage",
	HandlerType: (*GossipMessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "toLibP2P",
			Handler:    _GossipMessage_ToLibP2P_Handler,
		},
		{
			MethodName: "toRippled",
			Handler:    _GossipMessage_ToRippled_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gossip_message.proto",
}