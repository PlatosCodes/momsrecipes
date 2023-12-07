// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: service_moms_recipes.proto

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

const (
	MomsRecipes_CreateUser_FullMethodName    = "/pb.MomsRecipes/CreateUser"
	MomsRecipes_LoginUser_FullMethodName     = "/pb.MomsRecipes/LoginUser"
	MomsRecipes_UpdateUser_FullMethodName    = "/pb.MomsRecipes/UpdateUser"
	MomsRecipes_CreateRecipe_FullMethodName  = "/pb.MomsRecipes/CreateRecipe"
	MomsRecipes_GetRecipeByID_FullMethodName = "/pb.MomsRecipes/GetRecipeByID"
)

// MomsRecipesClient is the client API for MomsRecipes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MomsRecipesClient interface {
	// Creates a new user
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	// Logs in a user
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
	// Updates a user's password
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	// Creates a new recipe
	CreateRecipe(ctx context.Context, in *CreateRecipeRequest, opts ...grpc.CallOption) (*CreateRecipeResponse, error)
	// Retrieves a recipe by ID
	GetRecipeByID(ctx context.Context, in *GetRecipeByIDRequest, opts ...grpc.CallOption) (*GetRecipeByIDResponse, error)
}

type momsRecipesClient struct {
	cc grpc.ClientConnInterface
}

func NewMomsRecipesClient(cc grpc.ClientConnInterface) MomsRecipesClient {
	return &momsRecipesClient{cc}
}

func (c *momsRecipesClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, MomsRecipes_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momsRecipesClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, MomsRecipes_LoginUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momsRecipesClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, MomsRecipes_UpdateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momsRecipesClient) CreateRecipe(ctx context.Context, in *CreateRecipeRequest, opts ...grpc.CallOption) (*CreateRecipeResponse, error) {
	out := new(CreateRecipeResponse)
	err := c.cc.Invoke(ctx, MomsRecipes_CreateRecipe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momsRecipesClient) GetRecipeByID(ctx context.Context, in *GetRecipeByIDRequest, opts ...grpc.CallOption) (*GetRecipeByIDResponse, error) {
	out := new(GetRecipeByIDResponse)
	err := c.cc.Invoke(ctx, MomsRecipes_GetRecipeByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MomsRecipesServer is the server API for MomsRecipes service.
// All implementations must embed UnimplementedMomsRecipesServer
// for forward compatibility
type MomsRecipesServer interface {
	// Creates a new user
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	// Logs in a user
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	// Updates a user's password
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	// Creates a new recipe
	CreateRecipe(context.Context, *CreateRecipeRequest) (*CreateRecipeResponse, error)
	// Retrieves a recipe by ID
	GetRecipeByID(context.Context, *GetRecipeByIDRequest) (*GetRecipeByIDResponse, error)
	mustEmbedUnimplementedMomsRecipesServer()
}

// UnimplementedMomsRecipesServer must be embedded to have forward compatible implementations.
type UnimplementedMomsRecipesServer struct {
}

func (UnimplementedMomsRecipesServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedMomsRecipesServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedMomsRecipesServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedMomsRecipesServer) CreateRecipe(context.Context, *CreateRecipeRequest) (*CreateRecipeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRecipe not implemented")
}
func (UnimplementedMomsRecipesServer) GetRecipeByID(context.Context, *GetRecipeByIDRequest) (*GetRecipeByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecipeByID not implemented")
}
func (UnimplementedMomsRecipesServer) mustEmbedUnimplementedMomsRecipesServer() {}

// UnsafeMomsRecipesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MomsRecipesServer will
// result in compilation errors.
type UnsafeMomsRecipesServer interface {
	mustEmbedUnimplementedMomsRecipesServer()
}

func RegisterMomsRecipesServer(s grpc.ServiceRegistrar, srv MomsRecipesServer) {
	s.RegisterService(&MomsRecipes_ServiceDesc, srv)
}

func _MomsRecipes_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomsRecipesServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MomsRecipes_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomsRecipesServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MomsRecipes_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomsRecipesServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MomsRecipes_LoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomsRecipesServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MomsRecipes_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomsRecipesServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MomsRecipes_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomsRecipesServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MomsRecipes_CreateRecipe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRecipeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomsRecipesServer).CreateRecipe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MomsRecipes_CreateRecipe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomsRecipesServer).CreateRecipe(ctx, req.(*CreateRecipeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MomsRecipes_GetRecipeByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecipeByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomsRecipesServer).GetRecipeByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MomsRecipes_GetRecipeByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomsRecipesServer).GetRecipeByID(ctx, req.(*GetRecipeByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MomsRecipes_ServiceDesc is the grpc.ServiceDesc for MomsRecipes service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MomsRecipes_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MomsRecipes",
	HandlerType: (*MomsRecipesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _MomsRecipes_CreateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _MomsRecipes_LoginUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _MomsRecipes_UpdateUser_Handler,
		},
		{
			MethodName: "CreateRecipe",
			Handler:    _MomsRecipes_CreateRecipe_Handler,
		},
		{
			MethodName: "GetRecipeByID",
			Handler:    _MomsRecipes_GetRecipeByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_moms_recipes.proto",
}
