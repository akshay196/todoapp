package main

import (
	"context"
	"log"
	"net"
	"sync"

	pb "github.com/akshay196/todoapp/todoapppb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// database stores task with its unique id. id is not re-used.
var database = make(map[int32]string)

// Last id used in database. Increase id by 1 and use as id, before
// adding element to the database.
var lastId int32

var mutex = &sync.Mutex{}

type todoappService struct {
	pb.UnimplementedTodoServiceServer
}

func (*todoappService) ListTodo(ctx context.Context, req *emptypb.Empty) (*pb.TodoListResponse, error) {
	log.Printf("Received request to list: %v\n", req)
	var list []*pb.TodoItem
	for id, value := range database {
		list = append(list, &pb.TodoItem{
			Id:    id,
			Value: value,
		})
	}
	res := &pb.TodoListResponse{
		Tasks: list,
	}
	return res, nil
}

func (*todoappService) AddTodoItem(ctx context.Context, req *pb.TodoAddItemRequest) (*emptypb.Empty, error) {
	log.Printf("Received request to add: %v\n", req)
	newVal := req.GetTask()
	for _, val := range database {
		if newVal == val {
			err := status.Errorf(codes.InvalidArgument, "Task already present: %v", newVal)
			return &emptypb.Empty{}, err
		}
	}
	mutex.Lock()
	lastId += 1
	database[lastId] = newVal
	mutex.Unlock()

	return &emptypb.Empty{}, nil
}

func (*todoappService) DeleteTodoItem(ctx context.Context, req *pb.TodoDeleteItemRequest) (*emptypb.Empty, error) {
	log.Printf("Received request to delete: %v\n", req)
	id := req.GetId()
	_, ok := database[id]
	if !ok {
		err := status.Errorf(codes.InvalidArgument, "Id is not present: %v", id)
		return &emptypb.Empty{}, err
	}
	mutex.Lock()
	delete(database, id)
	mutex.Unlock()
	return &emptypb.Empty{}, nil
}

func main() {
	log.Println("Listening on :50051")
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Creating listener failed: %v\n", err)
	}

	s := grpc.NewServer()
	pb.RegisterTodoServiceServer(s, &todoappService{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("GRPC server failed: %v\n", err)
	}
}
