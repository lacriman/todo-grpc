package handlers

import (
	"context"

	pb "github.com/lacriman/todo-grpc/pb/todo"
)

type Server struct {
	pb.UnimplementedToDoServiceServer
	todos map[int64]*pb.ToDo
}

func NewServer() *Server {
	return &Server{
		todos: make(map[int64]*pb.ToDo),
	}
}

func (s *Server) CreateToDo(ctx context.Context, req *pb.CreateToDoRequest) (*pb.CreateToDoResponse, error) {
	id := int64(len(s.todos) + 1)

	todo := &pb.ToDo{
		Id:          id,
		Title:       req.Title,
		Description: req.Description,
		Done:        false,
	}

	s.todos[id] = todo
	return &pb.CreateToDoResponse{Todo: todo}, nil
}

func (s *Server) GetToDo(ctx context.Context, req *pb.GetToDoRequest) (*pb.GetToDoResponse, error) {
	return &pb.GetToDoResponse{}, nil
}

func (s *Server) ListToDo(ctx context.Context, req *pb.ListToDoRequest) (*pb.ListToDoResponse, error) {
	return &pb.ListToDoResponse{}, nil
}

func (s *Server) UpdateToDo(ctx context.Context, req *pb.UpdateToDoRequest) (*pb.UpdateToDoResponse, error) {
	return &pb.UpdateToDoResponse{}, nil
}

func (s *Server) DeleteToDo(ctx context.Context, req *pb.DeleteToDoRequest) (*pb.DeleteToDoResponse, error) {
	return &pb.DeleteToDoResponse{}, nil
}
