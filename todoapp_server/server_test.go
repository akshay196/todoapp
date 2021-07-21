package main

import (
	"context"
	"testing"

	"github.com/akshay196/todoapp/todoapppb"
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

	// Delete element to the list
	server.DeleteTodoItem(ctx, &todoapppb.TodoDeleteItemRequest{
		Id: 1,
	})

	res, _ = server.ListTodo(ctx, &emptypb.Empty{})
	if len(res.GetTasks()) != 0 {
		t.Errorf("TodoListResponse returned wrong response, Got: %v, Want: %v.\n", len(res.GetTasks()), 0)
	}
}
