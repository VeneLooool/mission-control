// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: api/v1/control/control.proto

package control

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MissionControl_SaveAnalyticResults_FullMethodName = "/mission_control.api.v1.control.MissionControl/SaveAnalyticResults"
)

// MissionControlClient is the client API for MissionControl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MissionControlClient interface {
	SaveAnalyticResults(ctx context.Context, in *SaveAnalyticResults_Request, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type missionControlClient struct {
	cc grpc.ClientConnInterface
}

func NewMissionControlClient(cc grpc.ClientConnInterface) MissionControlClient {
	return &missionControlClient{cc}
}

func (c *missionControlClient) SaveAnalyticResults(ctx context.Context, in *SaveAnalyticResults_Request, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MissionControl_SaveAnalyticResults_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MissionControlServer is the server API for MissionControl service.
// All implementations must embed UnimplementedMissionControlServer
// for forward compatibility.
type MissionControlServer interface {
	SaveAnalyticResults(context.Context, *SaveAnalyticResults_Request) (*emptypb.Empty, error)
	mustEmbedUnimplementedMissionControlServer()
}

// UnimplementedMissionControlServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMissionControlServer struct{}

func (UnimplementedMissionControlServer) SaveAnalyticResults(context.Context, *SaveAnalyticResults_Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveAnalyticResults not implemented")
}
func (UnimplementedMissionControlServer) mustEmbedUnimplementedMissionControlServer() {}
func (UnimplementedMissionControlServer) testEmbeddedByValue()                        {}

// UnsafeMissionControlServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MissionControlServer will
// result in compilation errors.
type UnsafeMissionControlServer interface {
	mustEmbedUnimplementedMissionControlServer()
}

func RegisterMissionControlServer(s grpc.ServiceRegistrar, srv MissionControlServer) {
	// If the following call pancis, it indicates UnimplementedMissionControlServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MissionControl_ServiceDesc, srv)
}

func _MissionControl_SaveAnalyticResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveAnalyticResults_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionControlServer).SaveAnalyticResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MissionControl_SaveAnalyticResults_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionControlServer).SaveAnalyticResults(ctx, req.(*SaveAnalyticResults_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// MissionControl_ServiceDesc is the grpc.ServiceDesc for MissionControl service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MissionControl_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mission_control.api.v1.control.MissionControl",
	HandlerType: (*MissionControlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveAnalyticResults",
			Handler:    _MissionControl_SaveAnalyticResults_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/control/control.proto",
}
