// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: protos/filestream.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	FileStream_UploadAndAnalyze_FullMethodName = "/filestream.FileStream/UploadAndAnalyze"
)

// FileStreamClient is the client API for FileStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileStreamClient interface {
	UploadAndAnalyze(ctx context.Context, opts ...grpc.CallOption) (FileStream_UploadAndAnalyzeClient, error)
}

type fileStreamClient struct {
	cc grpc.ClientConnInterface
}

func NewFileStreamClient(cc grpc.ClientConnInterface) FileStreamClient {
	return &fileStreamClient{cc}
}

func (c *fileStreamClient) UploadAndAnalyze(ctx context.Context, opts ...grpc.CallOption) (FileStream_UploadAndAnalyzeClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &FileStream_ServiceDesc.Streams[0], FileStream_UploadAndAnalyze_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &fileStreamUploadAndAnalyzeClient{ClientStream: stream}
	return x, nil
}

type FileStream_UploadAndAnalyzeClient interface {
	Send(*FileStreamRequest) error
	Recv() (*AnalysisResult, error)
	grpc.ClientStream
}

type fileStreamUploadAndAnalyzeClient struct {
	grpc.ClientStream
}

func (x *fileStreamUploadAndAnalyzeClient) Send(m *FileStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileStreamUploadAndAnalyzeClient) Recv() (*AnalysisResult, error) {
	m := new(AnalysisResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileStreamServer is the server API for FileStream service.
// All implementations must embed UnimplementedFileStreamServer
// for forward compatibility
type FileStreamServer interface {
	UploadAndAnalyze(FileStream_UploadAndAnalyzeServer) error
	mustEmbedUnimplementedFileStreamServer()
}

// UnimplementedFileStreamServer must be embedded to have forward compatible implementations.
type UnimplementedFileStreamServer struct {
}

func (UnimplementedFileStreamServer) UploadAndAnalyze(FileStream_UploadAndAnalyzeServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadAndAnalyze not implemented")
}
func (UnimplementedFileStreamServer) mustEmbedUnimplementedFileStreamServer() {}

// UnsafeFileStreamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileStreamServer will
// result in compilation errors.
type UnsafeFileStreamServer interface {
	mustEmbedUnimplementedFileStreamServer()
}

func RegisterFileStreamServer(s grpc.ServiceRegistrar, srv FileStreamServer) {
	s.RegisterService(&FileStream_ServiceDesc, srv)
}

func _FileStream_UploadAndAnalyze_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileStreamServer).UploadAndAnalyze(&fileStreamUploadAndAnalyzeServer{ServerStream: stream})
}

type FileStream_UploadAndAnalyzeServer interface {
	Send(*AnalysisResult) error
	Recv() (*FileStreamRequest, error)
	grpc.ServerStream
}

type fileStreamUploadAndAnalyzeServer struct {
	grpc.ServerStream
}

func (x *fileStreamUploadAndAnalyzeServer) Send(m *AnalysisResult) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileStreamUploadAndAnalyzeServer) Recv() (*FileStreamRequest, error) {
	m := new(FileStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileStream_ServiceDesc is the grpc.ServiceDesc for FileStream service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileStream_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "filestream.FileStream",
	HandlerType: (*FileStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadAndAnalyze",
			Handler:       _FileStream_UploadAndAnalyze_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protos/filestream.proto",
}
