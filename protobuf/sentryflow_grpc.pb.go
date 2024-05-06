// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.0
// source: sentryflow.proto

package protobuf

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

const (
	SentryFlow_GetAPILog_FullMethodName       = "/protobuf.SentryFlow/GetAPILog"
	SentryFlow_GetAPIMetrics_FullMethodName   = "/protobuf.SentryFlow/GetAPIMetrics"
	SentryFlow_GetEnvoyMetrics_FullMethodName = "/protobuf.SentryFlow/GetEnvoyMetrics"
)

// SentryFlowClient is the client API for SentryFlow service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SentryFlowClient interface {
	GetAPILog(ctx context.Context, in *ClientInfo, opts ...grpc.CallOption) (SentryFlow_GetAPILogClient, error)
	GetAPIMetrics(ctx context.Context, in *ClientInfo, opts ...grpc.CallOption) (SentryFlow_GetAPIMetricsClient, error)
	GetEnvoyMetrics(ctx context.Context, in *ClientInfo, opts ...grpc.CallOption) (SentryFlow_GetEnvoyMetricsClient, error)
}

type sentryFlowClient struct {
	cc grpc.ClientConnInterface
}

func NewSentryFlowClient(cc grpc.ClientConnInterface) SentryFlowClient {
	return &sentryFlowClient{cc}
}

func (c *sentryFlowClient) GetAPILog(ctx context.Context, in *ClientInfo, opts ...grpc.CallOption) (SentryFlow_GetAPILogClient, error) {
	stream, err := c.cc.NewStream(ctx, &SentryFlow_ServiceDesc.Streams[0], SentryFlow_GetAPILog_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &sentryFlowGetAPILogClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SentryFlow_GetAPILogClient interface {
	Recv() (*APILog, error)
	grpc.ClientStream
}

type sentryFlowGetAPILogClient struct {
	grpc.ClientStream
}

func (x *sentryFlowGetAPILogClient) Recv() (*APILog, error) {
	m := new(APILog)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sentryFlowClient) GetAPIMetrics(ctx context.Context, in *ClientInfo, opts ...grpc.CallOption) (SentryFlow_GetAPIMetricsClient, error) {
	stream, err := c.cc.NewStream(ctx, &SentryFlow_ServiceDesc.Streams[1], SentryFlow_GetAPIMetrics_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &sentryFlowGetAPIMetricsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SentryFlow_GetAPIMetricsClient interface {
	Recv() (*APIMetrics, error)
	grpc.ClientStream
}

type sentryFlowGetAPIMetricsClient struct {
	grpc.ClientStream
}

func (x *sentryFlowGetAPIMetricsClient) Recv() (*APIMetrics, error) {
	m := new(APIMetrics)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sentryFlowClient) GetEnvoyMetrics(ctx context.Context, in *ClientInfo, opts ...grpc.CallOption) (SentryFlow_GetEnvoyMetricsClient, error) {
	stream, err := c.cc.NewStream(ctx, &SentryFlow_ServiceDesc.Streams[2], SentryFlow_GetEnvoyMetrics_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &sentryFlowGetEnvoyMetricsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SentryFlow_GetEnvoyMetricsClient interface {
	Recv() (*EnvoyMetrics, error)
	grpc.ClientStream
}

type sentryFlowGetEnvoyMetricsClient struct {
	grpc.ClientStream
}

func (x *sentryFlowGetEnvoyMetricsClient) Recv() (*EnvoyMetrics, error) {
	m := new(EnvoyMetrics)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SentryFlowServer is the server API for SentryFlow service.
// All implementations should embed UnimplementedSentryFlowServer
// for forward compatibility
type SentryFlowServer interface {
	GetAPILog(*ClientInfo, SentryFlow_GetAPILogServer) error
	GetAPIMetrics(*ClientInfo, SentryFlow_GetAPIMetricsServer) error
	GetEnvoyMetrics(*ClientInfo, SentryFlow_GetEnvoyMetricsServer) error
}

// UnimplementedSentryFlowServer should be embedded to have forward compatible implementations.
type UnimplementedSentryFlowServer struct {
}

func (UnimplementedSentryFlowServer) GetAPILog(*ClientInfo, SentryFlow_GetAPILogServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAPILog not implemented")
}
func (UnimplementedSentryFlowServer) GetAPIMetrics(*ClientInfo, SentryFlow_GetAPIMetricsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAPIMetrics not implemented")
}
func (UnimplementedSentryFlowServer) GetEnvoyMetrics(*ClientInfo, SentryFlow_GetEnvoyMetricsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetEnvoyMetrics not implemented")
}

// UnsafeSentryFlowServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SentryFlowServer will
// result in compilation errors.
type UnsafeSentryFlowServer interface {
	mustEmbedUnimplementedSentryFlowServer()
}

func RegisterSentryFlowServer(s grpc.ServiceRegistrar, srv SentryFlowServer) {
	s.RegisterService(&SentryFlow_ServiceDesc, srv)
}

func _SentryFlow_GetAPILog_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ClientInfo)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SentryFlowServer).GetAPILog(m, &sentryFlowGetAPILogServer{stream})
}

type SentryFlow_GetAPILogServer interface {
	Send(*APILog) error
	grpc.ServerStream
}

type sentryFlowGetAPILogServer struct {
	grpc.ServerStream
}

func (x *sentryFlowGetAPILogServer) Send(m *APILog) error {
	return x.ServerStream.SendMsg(m)
}

func _SentryFlow_GetAPIMetrics_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ClientInfo)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SentryFlowServer).GetAPIMetrics(m, &sentryFlowGetAPIMetricsServer{stream})
}

type SentryFlow_GetAPIMetricsServer interface {
	Send(*APIMetrics) error
	grpc.ServerStream
}

type sentryFlowGetAPIMetricsServer struct {
	grpc.ServerStream
}

func (x *sentryFlowGetAPIMetricsServer) Send(m *APIMetrics) error {
	return x.ServerStream.SendMsg(m)
}

func _SentryFlow_GetEnvoyMetrics_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ClientInfo)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SentryFlowServer).GetEnvoyMetrics(m, &sentryFlowGetEnvoyMetricsServer{stream})
}

type SentryFlow_GetEnvoyMetricsServer interface {
	Send(*EnvoyMetrics) error
	grpc.ServerStream
}

type sentryFlowGetEnvoyMetricsServer struct {
	grpc.ServerStream
}

func (x *sentryFlowGetEnvoyMetricsServer) Send(m *EnvoyMetrics) error {
	return x.ServerStream.SendMsg(m)
}

// SentryFlow_ServiceDesc is the grpc.ServiceDesc for SentryFlow service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SentryFlow_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.SentryFlow",
	HandlerType: (*SentryFlowServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAPILog",
			Handler:       _SentryFlow_GetAPILog_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetAPIMetrics",
			Handler:       _SentryFlow_GetAPIMetrics_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetEnvoyMetrics",
			Handler:       _SentryFlow_GetEnvoyMetrics_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "sentryflow.proto",
}
