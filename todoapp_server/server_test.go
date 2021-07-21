package main

import (
	"context"
	"testing"

	"github.com/akshay196/todoapp/todoapppb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestTodoapp(t *testing.T) {
	server := New()
	ctx := context.Background()

	// Test empty list
	res, _ := server.ListTodo(ctx, &emptypb.Empty{})
	if len(res.GetTasks()) != 0 {
		t.Errorf("TodoListResponse returned wrong response, Got: %v, Want: %v.\n", len(res.GetTasks()), 0)
	}

	// Add element to the list
	server.AddTodoItem(ctx, &todoapppb.TodoAddItemRequest{
		Task: "Buy vegetable",
	})

	res, _ = server.ListTodo(ctx, &emptypb.Empty{})
	if len(res.GetTasks()) != 1 {
		t.Errorf("TodoListResponse returned wrong response, Got: %v, Want: %v.\n", len(res.GetTasks()), 1)
	}

	// Add same element to the list
	_, err := server.AddTodoItem(ctx, &todoapppb.TodoAddItemRequest{
		Task: "Buy vegetable",
	})
	statusErr, _ := status.FromError(err)
	if statusErr.Code() != codes.InvalidArgument {
		t.Errorf("InvalidArgument error is expected, Got: %v, wants: %v", statusErr.Code(), codes.InvalidArgument)
	}

	// Delete element to the list
	server.DeleteTodoItem(ctx, &todoapppb.TodoDeleteItemRequest{
		Id: 1,
	})

	res, _ = server.ListTodo(ctx, &emptypb.Empty{})
	if len(res.GetTasks()) != 0 {
		t.Errorf("TodoListResponse returned wrong response, Got: %v, Want: %v.\n", len(res.GetTasks()), 0)
	}

	// Delete element not present
	_, err = server.DeleteTodoItem(ctx, &todoapppb.TodoDeleteItemRequest{
		Id: 1,
	})
	statusErr, _ = status.FromError(err)
	if statusErr.Code() != codes.InvalidArgument {
		t.Errorf("InvalidArgument error is expected, Got: %v, wants: %v", statusErr.Code(), codes.InvalidArgument)
	}
}
