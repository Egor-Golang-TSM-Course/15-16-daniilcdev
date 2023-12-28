// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: fileStorageService.proto

package protoGenerated

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

// FileStorageServiceClient is the client API for FileStorageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileStorageServiceClient interface {
	Get(ctx context.Context, in *LoadFileMessage, opts ...grpc.CallOption) (*FileData, error)
	Save(ctx context.Context, in *SaveFileMessage, opts ...grpc.CallOption) (*OperationResultMessage, error)
}

type fileStorageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileStorageServiceClient(cc grpc.ClientConnInterface) FileStorageServiceClient {
	return &fileStorageServiceClient{cc}
}

func (c *fileStorageServiceClient) Get(ctx context.Context, in *LoadFileMessage, opts ...grpc.CallOption) (*FileData, error) {
	out := new(FileData)
	err := c.cc.Invoke(ctx, "/protoGenerated.FileStorageService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileStorageServiceClient) Save(ctx context.Context, in *SaveFileMessage, opts ...grpc.CallOption) (*OperationResultMessage, error) {
	out := new(OperationResultMessage)
	err := c.cc.Invoke(ctx, "/protoGenerated.FileStorageService/Save", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileStorageServiceServer is the server API for FileStorageService service.
// All implementations must embed UnimplementedFileStorageServiceServer
// for forward compatibility
type FileStorageServiceServer interface {
	Get(context.Context, *LoadFileMessage) (*FileData, error)
	Save(context.Context, *SaveFileMessage) (*OperationResultMessage, error)
	mustEmbedUnimplementedFileStorageServiceServer()
}

// UnimplementedFileStorageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFileStorageServiceServer struct {
}

func (UnimplementedFileStorageServiceServer) Get(context.Context, *LoadFileMessage) (*FileData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedFileStorageServiceServer) Save(context.Context, *SaveFileMessage) (*OperationResultMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedFileStorageServiceServer) mustEmbedUnimplementedFileStorageServiceServer() {}

// UnsafeFileStorageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileStorageServiceServer will
// result in compilation errors.
type UnsafeFileStorageServiceServer interface {
	mustEmbedUnimplementedFileStorageServiceServer()
}

func RegisterFileStorageServiceServer(s grpc.ServiceRegistrar, srv FileStorageServiceServer) {
	s.RegisterService(&FileStorageService_ServiceDesc, srv)
}

func _FileStorageService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadFileMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileStorageServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoGenerated.FileStorageService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileStorageServiceServer).Get(ctx, req.(*LoadFileMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileStorageService_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveFileMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileStorageServiceServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoGenerated.FileStorageService/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileStorageServiceServer).Save(ctx, req.(*SaveFileMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// FileStorageService_ServiceDesc is the grpc.ServiceDesc for FileStorageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileStorageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protoGenerated.FileStorageService",
	HandlerType: (*FileStorageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _FileStorageService_Get_Handler,
		},
		{
			MethodName: "Save",
			Handler:    _FileStorageService_Save_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fileStorageService.proto",
}
