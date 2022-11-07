// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: service.proto

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

// RegulationGRPCClient is the client API for RegulationGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegulationGRPCClient interface {
	GetRegulation(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Regulation, error)
	GetChapter(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Chapter, error)
	GetAllChapters(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Chapters, error)
	GetParagraphs(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Paragraphs, error)
	Search(ctx context.Context, in *SearchRequestMessage, opts ...grpc.CallOption) (*SearchResponseMessage, error)
	SearchRegulations(ctx context.Context, in *SearchRequestMessage, opts ...grpc.CallOption) (*SearchResponseMessage, error)
	SearchChapters(ctx context.Context, in *SearchRequestMessage, opts ...grpc.CallOption) (*SearchResponseMessage, error)
	SearchPargaraphs(ctx context.Context, in *SearchRequestMessage, opts ...grpc.CallOption) (*SearchResponseMessage, error)
}

type regulationGRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewRegulationGRPCClient(cc grpc.ClientConnInterface) RegulationGRPCClient {
	return &regulationGRPCClient{cc}
}

func (c *regulationGRPCClient) GetRegulation(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Regulation, error) {
	out := new(Regulation)
	err := c.cc.Invoke(ctx, "/RegulationGRPC/GetRegulation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regulationGRPCClient) GetChapter(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Chapter, error) {
	out := new(Chapter)
	err := c.cc.Invoke(ctx, "/RegulationGRPC/GetChapter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regulationGRPCClient) GetAllChapters(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Chapters, error) {
	out := new(Chapters)
	err := c.cc.Invoke(ctx, "/RegulationGRPC/GetAllChapters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regulationGRPCClient) GetParagraphs(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Paragraphs, error) {
	out := new(Paragraphs)
	err := c.cc.Invoke(ctx, "/RegulationGRPC/GetParagraphs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regulationGRPCClient) Search(ctx context.Context, in *SearchRequestMessage, opts ...grpc.CallOption) (*SearchResponseMessage, error) {
	out := new(SearchResponseMessage)
	err := c.cc.Invoke(ctx, "/RegulationGRPC/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regulationGRPCClient) SearchRegulations(ctx context.Context, in *SearchRequestMessage, opts ...grpc.CallOption) (*SearchResponseMessage, error) {
	out := new(SearchResponseMessage)
	err := c.cc.Invoke(ctx, "/RegulationGRPC/SearchRegulations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regulationGRPCClient) SearchChapters(ctx context.Context, in *SearchRequestMessage, opts ...grpc.CallOption) (*SearchResponseMessage, error) {
	out := new(SearchResponseMessage)
	err := c.cc.Invoke(ctx, "/RegulationGRPC/SearchChapters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regulationGRPCClient) SearchPargaraphs(ctx context.Context, in *SearchRequestMessage, opts ...grpc.CallOption) (*SearchResponseMessage, error) {
	out := new(SearchResponseMessage)
	err := c.cc.Invoke(ctx, "/RegulationGRPC/SearchPargaraphs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegulationGRPCServer is the server API for RegulationGRPC service.
// All implementations must embed UnimplementedRegulationGRPCServer
// for forward compatibility
type RegulationGRPCServer interface {
	GetRegulation(context.Context, *ID) (*Regulation, error)
	GetChapter(context.Context, *ID) (*Chapter, error)
	GetAllChapters(context.Context, *ID) (*Chapters, error)
	GetParagraphs(context.Context, *ID) (*Paragraphs, error)
	Search(context.Context, *SearchRequestMessage) (*SearchResponseMessage, error)
	SearchRegulations(context.Context, *SearchRequestMessage) (*SearchResponseMessage, error)
	SearchChapters(context.Context, *SearchRequestMessage) (*SearchResponseMessage, error)
	SearchPargaraphs(context.Context, *SearchRequestMessage) (*SearchResponseMessage, error)
	mustEmbedUnimplementedRegulationGRPCServer()
}

// UnimplementedRegulationGRPCServer must be embedded to have forward compatible implementations.
type UnimplementedRegulationGRPCServer struct {
}

func (UnimplementedRegulationGRPCServer) GetRegulation(context.Context, *ID) (*Regulation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegulation not implemented")
}
func (UnimplementedRegulationGRPCServer) GetChapter(context.Context, *ID) (*Chapter, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChapter not implemented")
}
func (UnimplementedRegulationGRPCServer) GetAllChapters(context.Context, *ID) (*Chapters, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllChapters not implemented")
}
func (UnimplementedRegulationGRPCServer) GetParagraphs(context.Context, *ID) (*Paragraphs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParagraphs not implemented")
}
func (UnimplementedRegulationGRPCServer) Search(context.Context, *SearchRequestMessage) (*SearchResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedRegulationGRPCServer) SearchRegulations(context.Context, *SearchRequestMessage) (*SearchResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchRegulations not implemented")
}
func (UnimplementedRegulationGRPCServer) SearchChapters(context.Context, *SearchRequestMessage) (*SearchResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchChapters not implemented")
}
func (UnimplementedRegulationGRPCServer) SearchPargaraphs(context.Context, *SearchRequestMessage) (*SearchResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPargaraphs not implemented")
}
func (UnimplementedRegulationGRPCServer) mustEmbedUnimplementedRegulationGRPCServer() {}

// UnsafeRegulationGRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegulationGRPCServer will
// result in compilation errors.
type UnsafeRegulationGRPCServer interface {
	mustEmbedUnimplementedRegulationGRPCServer()
}

func RegisterRegulationGRPCServer(s grpc.ServiceRegistrar, srv RegulationGRPCServer) {
	s.RegisterService(&RegulationGRPC_ServiceDesc, srv)
}

func _RegulationGRPC_GetRegulation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegulationGRPCServer).GetRegulation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RegulationGRPC/GetRegulation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegulationGRPCServer).GetRegulation(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegulationGRPC_GetChapter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegulationGRPCServer).GetChapter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RegulationGRPC/GetChapter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegulationGRPCServer).GetChapter(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegulationGRPC_GetAllChapters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegulationGRPCServer).GetAllChapters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RegulationGRPC/GetAllChapters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegulationGRPCServer).GetAllChapters(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegulationGRPC_GetParagraphs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegulationGRPCServer).GetParagraphs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RegulationGRPC/GetParagraphs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegulationGRPCServer).GetParagraphs(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegulationGRPC_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegulationGRPCServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RegulationGRPC/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegulationGRPCServer).Search(ctx, req.(*SearchRequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegulationGRPC_SearchRegulations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegulationGRPCServer).SearchRegulations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RegulationGRPC/SearchRegulations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegulationGRPCServer).SearchRegulations(ctx, req.(*SearchRequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegulationGRPC_SearchChapters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegulationGRPCServer).SearchChapters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RegulationGRPC/SearchChapters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegulationGRPCServer).SearchChapters(ctx, req.(*SearchRequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegulationGRPC_SearchPargaraphs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegulationGRPCServer).SearchPargaraphs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RegulationGRPC/SearchPargaraphs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegulationGRPCServer).SearchPargaraphs(ctx, req.(*SearchRequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// RegulationGRPC_ServiceDesc is the grpc.ServiceDesc for RegulationGRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RegulationGRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RegulationGRPC",
	HandlerType: (*RegulationGRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRegulation",
			Handler:    _RegulationGRPC_GetRegulation_Handler,
		},
		{
			MethodName: "GetChapter",
			Handler:    _RegulationGRPC_GetChapter_Handler,
		},
		{
			MethodName: "GetAllChapters",
			Handler:    _RegulationGRPC_GetAllChapters_Handler,
		},
		{
			MethodName: "GetParagraphs",
			Handler:    _RegulationGRPC_GetParagraphs_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _RegulationGRPC_Search_Handler,
		},
		{
			MethodName: "SearchRegulations",
			Handler:    _RegulationGRPC_SearchRegulations_Handler,
		},
		{
			MethodName: "SearchChapters",
			Handler:    _RegulationGRPC_SearchChapters_Handler,
		},
		{
			MethodName: "SearchPargaraphs",
			Handler:    _RegulationGRPC_SearchPargaraphs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
