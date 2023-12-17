// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: answerspb/api.proto

package answerspb

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

// AnswersServiceClient is the client API for AnswersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnswersServiceClient interface {
	CreateAnswer(ctx context.Context, in *CreateAnswerRequest, opts ...grpc.CallOption) (*CreateAnswerResponse, error)
	GetAnswers(ctx context.Context, in *GetAnswersRequest, opts ...grpc.CallOption) (*GetAnswersResponse, error)
}

type answersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAnswersServiceClient(cc grpc.ClientConnInterface) AnswersServiceClient {
	return &answersServiceClient{cc}
}

func (c *answersServiceClient) CreateAnswer(ctx context.Context, in *CreateAnswerRequest, opts ...grpc.CallOption) (*CreateAnswerResponse, error) {
	out := new(CreateAnswerResponse)
	err := c.cc.Invoke(ctx, "/answerspb.AnswersService/CreateAnswer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *answersServiceClient) GetAnswers(ctx context.Context, in *GetAnswersRequest, opts ...grpc.CallOption) (*GetAnswersResponse, error) {
	out := new(GetAnswersResponse)
	err := c.cc.Invoke(ctx, "/answerspb.AnswersService/GetAnswers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnswersServiceServer is the server API for AnswersService service.
// All implementations must embed UnimplementedAnswersServiceServer
// for forward compatibility
type AnswersServiceServer interface {
	CreateAnswer(context.Context, *CreateAnswerRequest) (*CreateAnswerResponse, error)
	GetAnswers(context.Context, *GetAnswersRequest) (*GetAnswersResponse, error)
	mustEmbedUnimplementedAnswersServiceServer()
}

// UnimplementedAnswersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAnswersServiceServer struct {
}

func (UnimplementedAnswersServiceServer) CreateAnswer(context.Context, *CreateAnswerRequest) (*CreateAnswerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAnswer not implemented")
}
func (UnimplementedAnswersServiceServer) GetAnswers(context.Context, *GetAnswersRequest) (*GetAnswersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnswers not implemented")
}
func (UnimplementedAnswersServiceServer) mustEmbedUnimplementedAnswersServiceServer() {}

// UnsafeAnswersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnswersServiceServer will
// result in compilation errors.
type UnsafeAnswersServiceServer interface {
	mustEmbedUnimplementedAnswersServiceServer()
}

func RegisterAnswersServiceServer(s grpc.ServiceRegistrar, srv AnswersServiceServer) {
	s.RegisterService(&AnswersService_ServiceDesc, srv)
}

func _AnswersService_CreateAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAnswerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnswersServiceServer).CreateAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/answerspb.AnswersService/CreateAnswer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnswersServiceServer).CreateAnswer(ctx, req.(*CreateAnswerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnswersService_GetAnswers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAnswersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnswersServiceServer).GetAnswers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/answerspb.AnswersService/GetAnswers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnswersServiceServer).GetAnswers(ctx, req.(*GetAnswersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AnswersService_ServiceDesc is the grpc.ServiceDesc for AnswersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AnswersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "answerspb.AnswersService",
	HandlerType: (*AnswersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAnswer",
			Handler:    _AnswersService_CreateAnswer_Handler,
		},
		{
			MethodName: "GetAnswers",
			Handler:    _AnswersService_GetAnswers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "answerspb/api.proto",
}