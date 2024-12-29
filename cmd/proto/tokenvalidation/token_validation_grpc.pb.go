// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: token_validation.proto

package tokenValidation

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TokenValidation_ValidateToken_FullMethodName = "/tokenvalidation.TokenValidation/ValidateToken"
)

// TokenValidationClient is the client API for TokenValidation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Define the service
type TokenValidationClient interface {
	// The RPC method to validate the token
	ValidateToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error)
}

type tokenValidationClient struct {
	cc grpc.ClientConnInterface
}

func NewTokenValidationClient(cc grpc.ClientConnInterface) TokenValidationClient {
	return &tokenValidationClient{cc}
}

func (c *tokenValidationClient) ValidateToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TokenResponse)
	err := c.cc.Invoke(ctx, TokenValidation_ValidateToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenValidationServer is the server API for TokenValidation service.
// All implementations must embed UnimplementedTokenValidationServer
// for forward compatibility.
//
// Define the service
type TokenValidationServer interface {
	// The RPC method to validate the token
	ValidateToken(context.Context, *TokenRequest) (*TokenResponse, error)
	mustEmbedUnimplementedTokenValidationServer()
}

// UnimplementedTokenValidationServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTokenValidationServer struct{}

func (UnimplementedTokenValidationServer) ValidateToken(context.Context, *TokenRequest) (*TokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateToken not implemented")
}
func (UnimplementedTokenValidationServer) mustEmbedUnimplementedTokenValidationServer() {}
func (UnimplementedTokenValidationServer) testEmbeddedByValue()                         {}

// UnsafeTokenValidationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TokenValidationServer will
// result in compilation errors.
type UnsafeTokenValidationServer interface {
	mustEmbedUnimplementedTokenValidationServer()
}

func RegisterTokenValidationServer(s grpc.ServiceRegistrar, srv TokenValidationServer) {
	// If the following call pancis, it indicates UnimplementedTokenValidationServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TokenValidation_ServiceDesc, srv)
}

func _TokenValidation_ValidateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenValidationServer).ValidateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TokenValidation_ValidateToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenValidationServer).ValidateToken(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TokenValidation_ServiceDesc is the grpc.ServiceDesc for TokenValidation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TokenValidation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tokenvalidation.TokenValidation",
	HandlerType: (*TokenValidationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateToken",
			Handler:    _TokenValidation_ValidateToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "token_validation.proto",
}
