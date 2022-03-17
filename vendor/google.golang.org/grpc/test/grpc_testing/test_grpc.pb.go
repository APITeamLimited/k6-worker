// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: test/grpc_testing/test.proto

package grpc_testing

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

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestServiceClient interface ***REMOVED***
	// One empty request followed by one empty response.
	EmptyCall(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// One request followed by one response.
	// The server returns the client payload as-is.
	UnaryCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
	// One request followed by a sequence of responses (streamed download).
	// The server returns the payload with client desired type and sizes.
	StreamingOutputCall(ctx context.Context, in *StreamingOutputCallRequest, opts ...grpc.CallOption) (TestService_StreamingOutputCallClient, error)
	// A sequence of requests followed by one response (streamed upload).
	// The server returns the aggregated size of client payload as the result.
	StreamingInputCall(ctx context.Context, opts ...grpc.CallOption) (TestService_StreamingInputCallClient, error)
	// A sequence of requests with each request served by the server immediately.
	// As one request could lead to multiple responses, this interface
	// demonstrates the idea of full duplexing.
	FullDuplexCall(ctx context.Context, opts ...grpc.CallOption) (TestService_FullDuplexCallClient, error)
	// A sequence of requests followed by a sequence of responses.
	// The server buffers all the client requests and then serves them in order. A
	// stream of responses are returned to the client when the server starts with
	// first request.
	HalfDuplexCall(ctx context.Context, opts ...grpc.CallOption) (TestService_HalfDuplexCallClient, error)
***REMOVED***

type testServiceClient struct ***REMOVED***
	cc grpc.ClientConnInterface
***REMOVED***

func NewTestServiceClient(cc grpc.ClientConnInterface) TestServiceClient ***REMOVED***
	return &testServiceClient***REMOVED***cc***REMOVED***
***REMOVED***

func (c *testServiceClient) EmptyCall(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) ***REMOVED***
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/grpc.testing.TestService/EmptyCall", in, out, opts...)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return out, nil
***REMOVED***

func (c *testServiceClient) UnaryCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error) ***REMOVED***
	out := new(SimpleResponse)
	err := c.cc.Invoke(ctx, "/grpc.testing.TestService/UnaryCall", in, out, opts...)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return out, nil
***REMOVED***

func (c *testServiceClient) StreamingOutputCall(ctx context.Context, in *StreamingOutputCallRequest, opts ...grpc.CallOption) (TestService_StreamingOutputCallClient, error) ***REMOVED***
	stream, err := c.cc.NewStream(ctx, &TestService_ServiceDesc.Streams[0], "/grpc.testing.TestService/StreamingOutputCall", opts...)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	x := &testServiceStreamingOutputCallClient***REMOVED***stream***REMOVED***
	if err := x.ClientStream.SendMsg(in); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	if err := x.ClientStream.CloseSend(); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return x, nil
***REMOVED***

type TestService_StreamingOutputCallClient interface ***REMOVED***
	Recv() (*StreamingOutputCallResponse, error)
	grpc.ClientStream
***REMOVED***

type testServiceStreamingOutputCallClient struct ***REMOVED***
	grpc.ClientStream
***REMOVED***

func (x *testServiceStreamingOutputCallClient) Recv() (*StreamingOutputCallResponse, error) ***REMOVED***
	m := new(StreamingOutputCallResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return m, nil
***REMOVED***

func (c *testServiceClient) StreamingInputCall(ctx context.Context, opts ...grpc.CallOption) (TestService_StreamingInputCallClient, error) ***REMOVED***
	stream, err := c.cc.NewStream(ctx, &TestService_ServiceDesc.Streams[1], "/grpc.testing.TestService/StreamingInputCall", opts...)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	x := &testServiceStreamingInputCallClient***REMOVED***stream***REMOVED***
	return x, nil
***REMOVED***

type TestService_StreamingInputCallClient interface ***REMOVED***
	Send(*StreamingInputCallRequest) error
	CloseAndRecv() (*StreamingInputCallResponse, error)
	grpc.ClientStream
***REMOVED***

type testServiceStreamingInputCallClient struct ***REMOVED***
	grpc.ClientStream
***REMOVED***

func (x *testServiceStreamingInputCallClient) Send(m *StreamingInputCallRequest) error ***REMOVED***
	return x.ClientStream.SendMsg(m)
***REMOVED***

func (x *testServiceStreamingInputCallClient) CloseAndRecv() (*StreamingInputCallResponse, error) ***REMOVED***
	if err := x.ClientStream.CloseSend(); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	m := new(StreamingInputCallResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return m, nil
***REMOVED***

func (c *testServiceClient) FullDuplexCall(ctx context.Context, opts ...grpc.CallOption) (TestService_FullDuplexCallClient, error) ***REMOVED***
	stream, err := c.cc.NewStream(ctx, &TestService_ServiceDesc.Streams[2], "/grpc.testing.TestService/FullDuplexCall", opts...)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	x := &testServiceFullDuplexCallClient***REMOVED***stream***REMOVED***
	return x, nil
***REMOVED***

type TestService_FullDuplexCallClient interface ***REMOVED***
	Send(*StreamingOutputCallRequest) error
	Recv() (*StreamingOutputCallResponse, error)
	grpc.ClientStream
***REMOVED***

type testServiceFullDuplexCallClient struct ***REMOVED***
	grpc.ClientStream
***REMOVED***

func (x *testServiceFullDuplexCallClient) Send(m *StreamingOutputCallRequest) error ***REMOVED***
	return x.ClientStream.SendMsg(m)
***REMOVED***

func (x *testServiceFullDuplexCallClient) Recv() (*StreamingOutputCallResponse, error) ***REMOVED***
	m := new(StreamingOutputCallResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return m, nil
***REMOVED***

func (c *testServiceClient) HalfDuplexCall(ctx context.Context, opts ...grpc.CallOption) (TestService_HalfDuplexCallClient, error) ***REMOVED***
	stream, err := c.cc.NewStream(ctx, &TestService_ServiceDesc.Streams[3], "/grpc.testing.TestService/HalfDuplexCall", opts...)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	x := &testServiceHalfDuplexCallClient***REMOVED***stream***REMOVED***
	return x, nil
***REMOVED***

type TestService_HalfDuplexCallClient interface ***REMOVED***
	Send(*StreamingOutputCallRequest) error
	Recv() (*StreamingOutputCallResponse, error)
	grpc.ClientStream
***REMOVED***

type testServiceHalfDuplexCallClient struct ***REMOVED***
	grpc.ClientStream
***REMOVED***

func (x *testServiceHalfDuplexCallClient) Send(m *StreamingOutputCallRequest) error ***REMOVED***
	return x.ClientStream.SendMsg(m)
***REMOVED***

func (x *testServiceHalfDuplexCallClient) Recv() (*StreamingOutputCallResponse, error) ***REMOVED***
	m := new(StreamingOutputCallResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return m, nil
***REMOVED***

// TestServiceServer is the server API for TestService service.
// All implementations must embed UnimplementedTestServiceServer
// for forward compatibility
type TestServiceServer interface ***REMOVED***
	// One empty request followed by one empty response.
	EmptyCall(context.Context, *Empty) (*Empty, error)
	// One request followed by one response.
	// The server returns the client payload as-is.
	UnaryCall(context.Context, *SimpleRequest) (*SimpleResponse, error)
	// One request followed by a sequence of responses (streamed download).
	// The server returns the payload with client desired type and sizes.
	StreamingOutputCall(*StreamingOutputCallRequest, TestService_StreamingOutputCallServer) error
	// A sequence of requests followed by one response (streamed upload).
	// The server returns the aggregated size of client payload as the result.
	StreamingInputCall(TestService_StreamingInputCallServer) error
	// A sequence of requests with each request served by the server immediately.
	// As one request could lead to multiple responses, this interface
	// demonstrates the idea of full duplexing.
	FullDuplexCall(TestService_FullDuplexCallServer) error
	// A sequence of requests followed by a sequence of responses.
	// The server buffers all the client requests and then serves them in order. A
	// stream of responses are returned to the client when the server starts with
	// first request.
	HalfDuplexCall(TestService_HalfDuplexCallServer) error
	mustEmbedUnimplementedTestServiceServer()
***REMOVED***

// UnimplementedTestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTestServiceServer struct ***REMOVED***
***REMOVED***

func (UnimplementedTestServiceServer) EmptyCall(context.Context, *Empty) (*Empty, error) ***REMOVED***
	return nil, status.Errorf(codes.Unimplemented, "method EmptyCall not implemented")
***REMOVED***
func (UnimplementedTestServiceServer) UnaryCall(context.Context, *SimpleRequest) (*SimpleResponse, error) ***REMOVED***
	return nil, status.Errorf(codes.Unimplemented, "method UnaryCall not implemented")
***REMOVED***
func (UnimplementedTestServiceServer) StreamingOutputCall(*StreamingOutputCallRequest, TestService_StreamingOutputCallServer) error ***REMOVED***
	return status.Errorf(codes.Unimplemented, "method StreamingOutputCall not implemented")
***REMOVED***
func (UnimplementedTestServiceServer) StreamingInputCall(TestService_StreamingInputCallServer) error ***REMOVED***
	return status.Errorf(codes.Unimplemented, "method StreamingInputCall not implemented")
***REMOVED***
func (UnimplementedTestServiceServer) FullDuplexCall(TestService_FullDuplexCallServer) error ***REMOVED***
	return status.Errorf(codes.Unimplemented, "method FullDuplexCall not implemented")
***REMOVED***
func (UnimplementedTestServiceServer) HalfDuplexCall(TestService_HalfDuplexCallServer) error ***REMOVED***
	return status.Errorf(codes.Unimplemented, "method HalfDuplexCall not implemented")
***REMOVED***
func (UnimplementedTestServiceServer) mustEmbedUnimplementedTestServiceServer() ***REMOVED******REMOVED***

// UnsafeTestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestServiceServer will
// result in compilation errors.
type UnsafeTestServiceServer interface ***REMOVED***
	mustEmbedUnimplementedTestServiceServer()
***REMOVED***

func RegisterTestServiceServer(s grpc.ServiceRegistrar, srv TestServiceServer) ***REMOVED***
	s.RegisterService(&TestService_ServiceDesc, srv)
***REMOVED***

func _TestService_EmptyCall_Handler(srv interface***REMOVED******REMOVED***, ctx context.Context, dec func(interface***REMOVED******REMOVED***) error, interceptor grpc.UnaryServerInterceptor) (interface***REMOVED******REMOVED***, error) ***REMOVED***
	in := new(Empty)
	if err := dec(in); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	if interceptor == nil ***REMOVED***
		return srv.(TestServiceServer).EmptyCall(ctx, in)
	***REMOVED***
	info := &grpc.UnaryServerInfo***REMOVED***
		Server:     srv,
		FullMethod: "/grpc.testing.TestService/EmptyCall",
	***REMOVED***
	handler := func(ctx context.Context, req interface***REMOVED******REMOVED***) (interface***REMOVED******REMOVED***, error) ***REMOVED***
		return srv.(TestServiceServer).EmptyCall(ctx, req.(*Empty))
	***REMOVED***
	return interceptor(ctx, in, info, handler)
***REMOVED***

func _TestService_UnaryCall_Handler(srv interface***REMOVED******REMOVED***, ctx context.Context, dec func(interface***REMOVED******REMOVED***) error, interceptor grpc.UnaryServerInterceptor) (interface***REMOVED******REMOVED***, error) ***REMOVED***
	in := new(SimpleRequest)
	if err := dec(in); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	if interceptor == nil ***REMOVED***
		return srv.(TestServiceServer).UnaryCall(ctx, in)
	***REMOVED***
	info := &grpc.UnaryServerInfo***REMOVED***
		Server:     srv,
		FullMethod: "/grpc.testing.TestService/UnaryCall",
	***REMOVED***
	handler := func(ctx context.Context, req interface***REMOVED******REMOVED***) (interface***REMOVED******REMOVED***, error) ***REMOVED***
		return srv.(TestServiceServer).UnaryCall(ctx, req.(*SimpleRequest))
	***REMOVED***
	return interceptor(ctx, in, info, handler)
***REMOVED***

func _TestService_StreamingOutputCall_Handler(srv interface***REMOVED******REMOVED***, stream grpc.ServerStream) error ***REMOVED***
	m := new(StreamingOutputCallRequest)
	if err := stream.RecvMsg(m); err != nil ***REMOVED***
		return err
	***REMOVED***
	return srv.(TestServiceServer).StreamingOutputCall(m, &testServiceStreamingOutputCallServer***REMOVED***stream***REMOVED***)
***REMOVED***

type TestService_StreamingOutputCallServer interface ***REMOVED***
	Send(*StreamingOutputCallResponse) error
	grpc.ServerStream
***REMOVED***

type testServiceStreamingOutputCallServer struct ***REMOVED***
	grpc.ServerStream
***REMOVED***

func (x *testServiceStreamingOutputCallServer) Send(m *StreamingOutputCallResponse) error ***REMOVED***
	return x.ServerStream.SendMsg(m)
***REMOVED***

func _TestService_StreamingInputCall_Handler(srv interface***REMOVED******REMOVED***, stream grpc.ServerStream) error ***REMOVED***
	return srv.(TestServiceServer).StreamingInputCall(&testServiceStreamingInputCallServer***REMOVED***stream***REMOVED***)
***REMOVED***

type TestService_StreamingInputCallServer interface ***REMOVED***
	SendAndClose(*StreamingInputCallResponse) error
	Recv() (*StreamingInputCallRequest, error)
	grpc.ServerStream
***REMOVED***

type testServiceStreamingInputCallServer struct ***REMOVED***
	grpc.ServerStream
***REMOVED***

func (x *testServiceStreamingInputCallServer) SendAndClose(m *StreamingInputCallResponse) error ***REMOVED***
	return x.ServerStream.SendMsg(m)
***REMOVED***

func (x *testServiceStreamingInputCallServer) Recv() (*StreamingInputCallRequest, error) ***REMOVED***
	m := new(StreamingInputCallRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return m, nil
***REMOVED***

func _TestService_FullDuplexCall_Handler(srv interface***REMOVED******REMOVED***, stream grpc.ServerStream) error ***REMOVED***
	return srv.(TestServiceServer).FullDuplexCall(&testServiceFullDuplexCallServer***REMOVED***stream***REMOVED***)
***REMOVED***

type TestService_FullDuplexCallServer interface ***REMOVED***
	Send(*StreamingOutputCallResponse) error
	Recv() (*StreamingOutputCallRequest, error)
	grpc.ServerStream
***REMOVED***

type testServiceFullDuplexCallServer struct ***REMOVED***
	grpc.ServerStream
***REMOVED***

func (x *testServiceFullDuplexCallServer) Send(m *StreamingOutputCallResponse) error ***REMOVED***
	return x.ServerStream.SendMsg(m)
***REMOVED***

func (x *testServiceFullDuplexCallServer) Recv() (*StreamingOutputCallRequest, error) ***REMOVED***
	m := new(StreamingOutputCallRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return m, nil
***REMOVED***

func _TestService_HalfDuplexCall_Handler(srv interface***REMOVED******REMOVED***, stream grpc.ServerStream) error ***REMOVED***
	return srv.(TestServiceServer).HalfDuplexCall(&testServiceHalfDuplexCallServer***REMOVED***stream***REMOVED***)
***REMOVED***

type TestService_HalfDuplexCallServer interface ***REMOVED***
	Send(*StreamingOutputCallResponse) error
	Recv() (*StreamingOutputCallRequest, error)
	grpc.ServerStream
***REMOVED***

type testServiceHalfDuplexCallServer struct ***REMOVED***
	grpc.ServerStream
***REMOVED***

func (x *testServiceHalfDuplexCallServer) Send(m *StreamingOutputCallResponse) error ***REMOVED***
	return x.ServerStream.SendMsg(m)
***REMOVED***

func (x *testServiceHalfDuplexCallServer) Recv() (*StreamingOutputCallRequest, error) ***REMOVED***
	m := new(StreamingOutputCallRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return m, nil
***REMOVED***

// TestService_ServiceDesc is the grpc.ServiceDesc for TestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestService_ServiceDesc = grpc.ServiceDesc***REMOVED***
	ServiceName: "grpc.testing.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc***REMOVED***
		***REMOVED***
			MethodName: "EmptyCall",
			Handler:    _TestService_EmptyCall_Handler,
		***REMOVED***,
		***REMOVED***
			MethodName: "UnaryCall",
			Handler:    _TestService_UnaryCall_Handler,
		***REMOVED***,
	***REMOVED***,
	Streams: []grpc.StreamDesc***REMOVED***
		***REMOVED***
			StreamName:    "StreamingOutputCall",
			Handler:       _TestService_StreamingOutputCall_Handler,
			ServerStreams: true,
		***REMOVED***,
		***REMOVED***
			StreamName:    "StreamingInputCall",
			Handler:       _TestService_StreamingInputCall_Handler,
			ClientStreams: true,
		***REMOVED***,
		***REMOVED***
			StreamName:    "FullDuplexCall",
			Handler:       _TestService_FullDuplexCall_Handler,
			ServerStreams: true,
			ClientStreams: true,
		***REMOVED***,
		***REMOVED***
			StreamName:    "HalfDuplexCall",
			Handler:       _TestService_HalfDuplexCall_Handler,
			ServerStreams: true,
			ClientStreams: true,
		***REMOVED***,
	***REMOVED***,
	Metadata: "test/grpc_testing/test.proto",
***REMOVED***
