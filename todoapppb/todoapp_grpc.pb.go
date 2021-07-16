// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package todoapppb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoServiceClient interface {
	// Returns all Todo list items.
	ListTodo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TodoListResponse, error)
	// Add an item to the Todo list. Returns error of type
	// INVALID_ARGUMENT when adding an item which is already present in
	// the Todo list.
	AddTodoItem(ctx context.Context, in *TodoAddItemRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Delete an item from the Todo list. Returns error of type
	// INVALID_ARGUMENT when deleting an item which is not present in
	// the Todo list.
	DeleteTodoItem(ctx context.Context, in *TodoDeleteItemRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type todoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoServiceClient(cc grpc.ClientConnInterface) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) ListTodo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TodoListResponse, error) {
	out := new(TodoListResponse)
	err := c.cc.Invoke(ctx, "/todoapppb.TodoService/ListTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) AddTodoItem(ctx context.Context, in *TodoAddItemRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/todoapppb.TodoService/AddTodoItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) DeleteTodoItem(ctx context.Context, in *TodoDeleteItemRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/todoapppb.TodoService/DeleteTodoItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
// All implementations must embed UnimplementedTodoServiceServer
// for forward compatibility
type TodoServiceServer interface {
	// Returns all Todo list items.
	ListTodo(context.Context, *emptypb.Empty) (*TodoListResponse, error)
	// Add an item to the Todo list. Returns error of type
	// INVALID_ARGUMENT when adding an item which is already present in
	// the Todo list.
	AddTodoItem(context.Context, *TodoAddItemRequest) (*emptypb.Empty, error)
	// Delete an item from the Todo list. Returns error of type
	// INVALID_ARGUMENT when deleting an item which is not present in
	// the Todo list.
	DeleteTodoItem(context.Context, *TodoDeleteItemRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedTodoServiceServer()
}

// UnimplementedTodoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTodoServiceServer struct {
}

func (UnimplementedTodoServiceServer) ListTodo(context.Context, *emptypb.Empty) (*TodoListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTodo not implemented")
}
func (UnimplementedTodoServiceServer) AddTodoItem(context.Context, *TodoAddItemRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTodoItem not implemented")
}
func (UnimplementedTodoServiceServer) DeleteTodoItem(context.Context, *TodoDeleteItemRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTodoItem not implemented")
}
func (UnimplementedTodoServiceServer) mustEmbedUnimplementedTodoServiceServer() {}

// UnsafeTodoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoServiceServer will
// result in compilation errors.
type UnsafeTodoServiceServer interface {
	mustEmbedUnimplementedTodoServiceServer()
}

func RegisterTodoServiceServer(s grpc.ServiceRegistrar, srv TodoServiceServer) {
	s.RegisterService(&TodoService_ServiceDesc, srv)
}

func _TodoService_ListTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).ListTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todoapppb.TodoService/ListTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).ListTodo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_AddTodoItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoAddItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).AddTodoItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todoapppb.TodoService/AddTodoItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).AddTodoItem(ctx, req.(*TodoAddItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_DeleteTodoItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoDeleteItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).DeleteTodoItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todoapppb.TodoService/DeleteTodoItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).DeleteTodoItem(ctx, req.(*TodoDeleteItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TodoService_ServiceDesc is the grpc.ServiceDesc for TodoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todoapppb.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTodo",
			Handler:    _TodoService_ListTodo_Handler,
		},
		{
			MethodName: "AddTodoItem",
			Handler:    _TodoService_AddTodoItem_Handler,
		},
		{
			MethodName: "DeleteTodoItem",
			Handler:    _TodoService_DeleteTodoItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todoapppb/todoapp.proto",
}
