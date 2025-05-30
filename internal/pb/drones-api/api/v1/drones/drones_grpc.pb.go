// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: drones-api/api/v1/drones/drones.proto

package drones

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
	Drones_CreateDrone_FullMethodName       = "/drones_api.api.v1.drones.Drones/CreateDrone"
	Drones_UpdateDrone_FullMethodName       = "/drones_api.api.v1.drones.Drones/UpdateDrone"
	Drones_GetDroneByID_FullMethodName      = "/drones_api.api.v1.drones.Drones/GetDroneByID"
	Drones_GetDronesByAuthor_FullMethodName = "/drones_api.api.v1.drones.Drones/GetDronesByAuthor"
	Drones_DeleteDrone_FullMethodName       = "/drones_api.api.v1.drones.Drones/DeleteDrone"
	Drones_StartDroneMission_FullMethodName = "/drones_api.api.v1.drones.Drones/StartDroneMission"
)

// DronesClient is the client API for Drones service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DronesClient interface {
	CreateDrone(ctx context.Context, in *CreateDrone_Request, opts ...grpc.CallOption) (*CreateDrone_Response, error)
	UpdateDrone(ctx context.Context, in *UpdateDrone_Request, opts ...grpc.CallOption) (*UpdateDrone_Response, error)
	GetDroneByID(ctx context.Context, in *GetDroneByID_Request, opts ...grpc.CallOption) (*GetDroneByID_Response, error)
	GetDronesByAuthor(ctx context.Context, in *GetDronesByAuthor_Request, opts ...grpc.CallOption) (*GetDronesByAuthor_Response, error)
	DeleteDrone(ctx context.Context, in *DeleteDrone_Request, opts ...grpc.CallOption) (*emptypb.Empty, error)
	StartDroneMission(ctx context.Context, in *StartDroneMission_Request, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type dronesClient struct {
	cc grpc.ClientConnInterface
}

func NewDronesClient(cc grpc.ClientConnInterface) DronesClient {
	return &dronesClient{cc}
}

func (c *dronesClient) CreateDrone(ctx context.Context, in *CreateDrone_Request, opts ...grpc.CallOption) (*CreateDrone_Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateDrone_Response)
	err := c.cc.Invoke(ctx, Drones_CreateDrone_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dronesClient) UpdateDrone(ctx context.Context, in *UpdateDrone_Request, opts ...grpc.CallOption) (*UpdateDrone_Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateDrone_Response)
	err := c.cc.Invoke(ctx, Drones_UpdateDrone_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dronesClient) GetDroneByID(ctx context.Context, in *GetDroneByID_Request, opts ...grpc.CallOption) (*GetDroneByID_Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDroneByID_Response)
	err := c.cc.Invoke(ctx, Drones_GetDroneByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dronesClient) GetDronesByAuthor(ctx context.Context, in *GetDronesByAuthor_Request, opts ...grpc.CallOption) (*GetDronesByAuthor_Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDronesByAuthor_Response)
	err := c.cc.Invoke(ctx, Drones_GetDronesByAuthor_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dronesClient) DeleteDrone(ctx context.Context, in *DeleteDrone_Request, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Drones_DeleteDrone_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dronesClient) StartDroneMission(ctx context.Context, in *StartDroneMission_Request, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Drones_StartDroneMission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DronesServer is the server API for Drones service.
// All implementations must embed UnimplementedDronesServer
// for forward compatibility.
type DronesServer interface {
	CreateDrone(context.Context, *CreateDrone_Request) (*CreateDrone_Response, error)
	UpdateDrone(context.Context, *UpdateDrone_Request) (*UpdateDrone_Response, error)
	GetDroneByID(context.Context, *GetDroneByID_Request) (*GetDroneByID_Response, error)
	GetDronesByAuthor(context.Context, *GetDronesByAuthor_Request) (*GetDronesByAuthor_Response, error)
	DeleteDrone(context.Context, *DeleteDrone_Request) (*emptypb.Empty, error)
	StartDroneMission(context.Context, *StartDroneMission_Request) (*emptypb.Empty, error)
	mustEmbedUnimplementedDronesServer()
}

// UnimplementedDronesServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDronesServer struct{}

func (UnimplementedDronesServer) CreateDrone(context.Context, *CreateDrone_Request) (*CreateDrone_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDrone not implemented")
}
func (UnimplementedDronesServer) UpdateDrone(context.Context, *UpdateDrone_Request) (*UpdateDrone_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDrone not implemented")
}
func (UnimplementedDronesServer) GetDroneByID(context.Context, *GetDroneByID_Request) (*GetDroneByID_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDroneByID not implemented")
}
func (UnimplementedDronesServer) GetDronesByAuthor(context.Context, *GetDronesByAuthor_Request) (*GetDronesByAuthor_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDronesByAuthor not implemented")
}
func (UnimplementedDronesServer) DeleteDrone(context.Context, *DeleteDrone_Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDrone not implemented")
}
func (UnimplementedDronesServer) StartDroneMission(context.Context, *StartDroneMission_Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartDroneMission not implemented")
}
func (UnimplementedDronesServer) mustEmbedUnimplementedDronesServer() {}
func (UnimplementedDronesServer) testEmbeddedByValue()                {}

// UnsafeDronesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DronesServer will
// result in compilation errors.
type UnsafeDronesServer interface {
	mustEmbedUnimplementedDronesServer()
}

func RegisterDronesServer(s grpc.ServiceRegistrar, srv DronesServer) {
	// If the following call pancis, it indicates UnimplementedDronesServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Drones_ServiceDesc, srv)
}

func _Drones_CreateDrone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDrone_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DronesServer).CreateDrone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Drones_CreateDrone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DronesServer).CreateDrone(ctx, req.(*CreateDrone_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Drones_UpdateDrone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDrone_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DronesServer).UpdateDrone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Drones_UpdateDrone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DronesServer).UpdateDrone(ctx, req.(*UpdateDrone_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Drones_GetDroneByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDroneByID_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DronesServer).GetDroneByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Drones_GetDroneByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DronesServer).GetDroneByID(ctx, req.(*GetDroneByID_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Drones_GetDronesByAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDronesByAuthor_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DronesServer).GetDronesByAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Drones_GetDronesByAuthor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DronesServer).GetDronesByAuthor(ctx, req.(*GetDronesByAuthor_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Drones_DeleteDrone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDrone_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DronesServer).DeleteDrone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Drones_DeleteDrone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DronesServer).DeleteDrone(ctx, req.(*DeleteDrone_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Drones_StartDroneMission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartDroneMission_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DronesServer).StartDroneMission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Drones_StartDroneMission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DronesServer).StartDroneMission(ctx, req.(*StartDroneMission_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Drones_ServiceDesc is the grpc.ServiceDesc for Drones service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Drones_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "drones_api.api.v1.drones.Drones",
	HandlerType: (*DronesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDrone",
			Handler:    _Drones_CreateDrone_Handler,
		},
		{
			MethodName: "UpdateDrone",
			Handler:    _Drones_UpdateDrone_Handler,
		},
		{
			MethodName: "GetDroneByID",
			Handler:    _Drones_GetDroneByID_Handler,
		},
		{
			MethodName: "GetDronesByAuthor",
			Handler:    _Drones_GetDronesByAuthor_Handler,
		},
		{
			MethodName: "DeleteDrone",
			Handler:    _Drones_DeleteDrone_Handler,
		},
		{
			MethodName: "StartDroneMission",
			Handler:    _Drones_StartDroneMission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "drones-api/api/v1/drones/drones.proto",
}
