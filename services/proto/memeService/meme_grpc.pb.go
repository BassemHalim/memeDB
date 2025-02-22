// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: meme.proto

package memeService

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
	MemeService_UploadMeme_FullMethodName        = "/meme.MemeService/UploadMeme"
	MemeService_GetMeme_FullMethodName           = "/meme.MemeService/GetMeme"
	MemeService_DeleteMeme_FullMethodName        = "/meme.MemeService/DeleteMeme"
	MemeService_GetTimelineMemes_FullMethodName  = "/meme.MemeService/GetTimelineMemes"
	MemeService_FilterMemesByTags_FullMethodName = "/meme.MemeService/FilterMemesByTags"
	MemeService_SearchMemes_FullMethodName       = "/meme.MemeService/SearchMemes"
	MemeService_SearchTags_FullMethodName        = "/meme.MemeService/SearchTags"
)

// MemeServiceClient is the client API for MemeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MemeServiceClient interface {
	UploadMeme(ctx context.Context, in *UploadMemeRequest, opts ...grpc.CallOption) (*MemeResponse, error)
	GetMeme(ctx context.Context, in *GetMemeRequest, opts ...grpc.CallOption) (*MemeResponse, error)
	DeleteMeme(ctx context.Context, in *DeleteMemeRequest, opts ...grpc.CallOption) (*DeleteMemeResponse, error)
	GetTimelineMemes(ctx context.Context, in *GetTimelineRequest, opts ...grpc.CallOption) (*MemesResponse, error)
	FilterMemesByTags(ctx context.Context, in *FilterMemesByTagsRequest, opts ...grpc.CallOption) (*MemesResponse, error)
	SearchMemes(ctx context.Context, in *SearchMemesRequest, opts ...grpc.CallOption) (*MemesResponse, error)
	SearchTags(ctx context.Context, in *SearchTagsRequest, opts ...grpc.CallOption) (*TagsResponse, error)
}

type memeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMemeServiceClient(cc grpc.ClientConnInterface) MemeServiceClient {
	return &memeServiceClient{cc}
}

func (c *memeServiceClient) UploadMeme(ctx context.Context, in *UploadMemeRequest, opts ...grpc.CallOption) (*MemeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MemeResponse)
	err := c.cc.Invoke(ctx, MemeService_UploadMeme_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memeServiceClient) GetMeme(ctx context.Context, in *GetMemeRequest, opts ...grpc.CallOption) (*MemeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MemeResponse)
	err := c.cc.Invoke(ctx, MemeService_GetMeme_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memeServiceClient) DeleteMeme(ctx context.Context, in *DeleteMemeRequest, opts ...grpc.CallOption) (*DeleteMemeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteMemeResponse)
	err := c.cc.Invoke(ctx, MemeService_DeleteMeme_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memeServiceClient) GetTimelineMemes(ctx context.Context, in *GetTimelineRequest, opts ...grpc.CallOption) (*MemesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MemesResponse)
	err := c.cc.Invoke(ctx, MemeService_GetTimelineMemes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memeServiceClient) FilterMemesByTags(ctx context.Context, in *FilterMemesByTagsRequest, opts ...grpc.CallOption) (*MemesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MemesResponse)
	err := c.cc.Invoke(ctx, MemeService_FilterMemesByTags_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memeServiceClient) SearchMemes(ctx context.Context, in *SearchMemesRequest, opts ...grpc.CallOption) (*MemesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MemesResponse)
	err := c.cc.Invoke(ctx, MemeService_SearchMemes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memeServiceClient) SearchTags(ctx context.Context, in *SearchTagsRequest, opts ...grpc.CallOption) (*TagsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TagsResponse)
	err := c.cc.Invoke(ctx, MemeService_SearchTags_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MemeServiceServer is the server API for MemeService service.
// All implementations must embed UnimplementedMemeServiceServer
// for forward compatibility.
type MemeServiceServer interface {
	UploadMeme(context.Context, *UploadMemeRequest) (*MemeResponse, error)
	GetMeme(context.Context, *GetMemeRequest) (*MemeResponse, error)
	DeleteMeme(context.Context, *DeleteMemeRequest) (*DeleteMemeResponse, error)
	GetTimelineMemes(context.Context, *GetTimelineRequest) (*MemesResponse, error)
	FilterMemesByTags(context.Context, *FilterMemesByTagsRequest) (*MemesResponse, error)
	SearchMemes(context.Context, *SearchMemesRequest) (*MemesResponse, error)
	SearchTags(context.Context, *SearchTagsRequest) (*TagsResponse, error)
	mustEmbedUnimplementedMemeServiceServer()
}

// UnimplementedMemeServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMemeServiceServer struct{}

func (UnimplementedMemeServiceServer) UploadMeme(context.Context, *UploadMemeRequest) (*MemeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadMeme not implemented")
}
func (UnimplementedMemeServiceServer) GetMeme(context.Context, *GetMemeRequest) (*MemeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeme not implemented")
}
func (UnimplementedMemeServiceServer) DeleteMeme(context.Context, *DeleteMemeRequest) (*DeleteMemeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMeme not implemented")
}
func (UnimplementedMemeServiceServer) GetTimelineMemes(context.Context, *GetTimelineRequest) (*MemesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimelineMemes not implemented")
}
func (UnimplementedMemeServiceServer) FilterMemesByTags(context.Context, *FilterMemesByTagsRequest) (*MemesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FilterMemesByTags not implemented")
}
func (UnimplementedMemeServiceServer) SearchMemes(context.Context, *SearchMemesRequest) (*MemesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchMemes not implemented")
}
func (UnimplementedMemeServiceServer) SearchTags(context.Context, *SearchTagsRequest) (*TagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchTags not implemented")
}
func (UnimplementedMemeServiceServer) mustEmbedUnimplementedMemeServiceServer() {}
func (UnimplementedMemeServiceServer) testEmbeddedByValue()                     {}

// UnsafeMemeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MemeServiceServer will
// result in compilation errors.
type UnsafeMemeServiceServer interface {
	mustEmbedUnimplementedMemeServiceServer()
}

func RegisterMemeServiceServer(s grpc.ServiceRegistrar, srv MemeServiceServer) {
	// If the following call pancis, it indicates UnimplementedMemeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MemeService_ServiceDesc, srv)
}

func _MemeService_UploadMeme_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadMemeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemeServiceServer).UploadMeme(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemeService_UploadMeme_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemeServiceServer).UploadMeme(ctx, req.(*UploadMemeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemeService_GetMeme_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMemeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemeServiceServer).GetMeme(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemeService_GetMeme_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemeServiceServer).GetMeme(ctx, req.(*GetMemeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemeService_DeleteMeme_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMemeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemeServiceServer).DeleteMeme(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemeService_DeleteMeme_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemeServiceServer).DeleteMeme(ctx, req.(*DeleteMemeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemeService_GetTimelineMemes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTimelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemeServiceServer).GetTimelineMemes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemeService_GetTimelineMemes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemeServiceServer).GetTimelineMemes(ctx, req.(*GetTimelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemeService_FilterMemesByTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterMemesByTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemeServiceServer).FilterMemesByTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemeService_FilterMemesByTags_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemeServiceServer).FilterMemesByTags(ctx, req.(*FilterMemesByTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemeService_SearchMemes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchMemesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemeServiceServer).SearchMemes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemeService_SearchMemes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemeServiceServer).SearchMemes(ctx, req.(*SearchMemesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemeService_SearchTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemeServiceServer).SearchTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MemeService_SearchTags_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemeServiceServer).SearchTags(ctx, req.(*SearchTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MemeService_ServiceDesc is the grpc.ServiceDesc for MemeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MemeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "meme.MemeService",
	HandlerType: (*MemeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadMeme",
			Handler:    _MemeService_UploadMeme_Handler,
		},
		{
			MethodName: "GetMeme",
			Handler:    _MemeService_GetMeme_Handler,
		},
		{
			MethodName: "DeleteMeme",
			Handler:    _MemeService_DeleteMeme_Handler,
		},
		{
			MethodName: "GetTimelineMemes",
			Handler:    _MemeService_GetTimelineMemes_Handler,
		},
		{
			MethodName: "FilterMemesByTags",
			Handler:    _MemeService_FilterMemesByTags_Handler,
		},
		{
			MethodName: "SearchMemes",
			Handler:    _MemeService_SearchMemes_Handler,
		},
		{
			MethodName: "SearchTags",
			Handler:    _MemeService_SearchTags_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "meme.proto",
}
