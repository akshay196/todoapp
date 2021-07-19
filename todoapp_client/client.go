package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/akshay196/todoapp/todoapppb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to create client connection: %v", err)
	}
	defer cc.Close()

	c := todoapppb.NewTodoServiceClient(cc)

	addTask(c, "Buy vegetables")
	doList(c)
	delTask(c, 1)
	doList(c)
}

func doList(c todoapppb.TodoServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	res, err := c.ListTodo(ctx, &emptypb.Empty{})
	if err != nil {
		resErr, ok := status.FromError(err)
		if ok {
			if resErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timed out!")
				return
			}
		}
		fmt.Printf("Failed to fetch todo list: %v\n", err)
		return
	}
	list := res.GetTasks()
	if len(list) == 0 {
		fmt.Println("No tasks in the list.")
		return
	}
	fmt.Println("Todo tasks:")
	for _, l := range list {
		fmt.Printf("%d: %s\n", l.GetId(), l.GetValue())
	}
}

func addTask(c todoapppb.TodoServiceClient, newTask string) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req := &todoapppb.TodoAddItemRequest{
		Task: newTask,
	}
	_, err := c.AddTodoItem(ctx, req)
	if err != nil {
		resErr, ok := status.FromError(err)
		if ok {
			if resErr.Code() == codes.InvalidArgument {
				fmt.Println(resErr.Message())
				return
			}
			if resErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timed out!")
				return
			}
		}
		fmt.Printf("Error adding task: %v", err)
		return
	}
	// Falling here means no errors occurred
	fmt.Println("Task is added successfully")
}

func delTask(c todoapppb.TodoServiceClient, id int32) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req := &todoapppb.TodoDeleteItemRequest{
		Id: id,
	}
	_, err := c.DeleteTodoItem(ctx, req)
	if err != nil {
		resErr, ok := status.FromError(err)
		if ok {
			if resErr.Code() == codes.InvalidArgument {
				fmt.Println(resErr.Message())
				return
			}
			if resErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timed out!")
				return
			}
		}
		fmt.Printf("Error deleting task: %v", err)
		return
	}
	// Falling here means no errors occurred
	fmt.Println("Task is deleted successfully")
}
