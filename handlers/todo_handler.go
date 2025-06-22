package handlers

import (
	"context"
	"maps"
	"slices"

	pb "github.com/lacriman/todo-grpc/pb/todo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	todo, exists := s.todos[req.Id]
	if !exists {
		return nil, status.Error(codes.NotFound, "todo not found")
	}
	return &pb.GetToDoResponse{Todo: todo}, nil
}

func (s *Server) ListToDo(ctx context.Context, req *pb.ListToDoRequest) (*pb.ListToDoResponse, error) {
	vals := slices.Collect(maps.Values(s.todos))
	return &pb.ListToDoResponse{Todos: vals}, nil
}

func (s *Server) UpdateToDo(ctx context.Context, req *pb.UpdateToDoRequest) (*pb.UpdateToDoResponse, error) {
	todo := s.todos[req.Id]
	todo.Title = req.Title
	todo.Description = req.Description
	todo.Done = req.Done
	return &pb.UpdateToDoResponse{Todo: todo}, nil
}

func (s *Server) DeleteToDo(ctx context.Context, req *pb.DeleteToDoRequest) (*pb.DeleteToDoResponse, error) {
	if _, exists := s.todos[req.Id]; exists {
		delete(s.todos, req.Id)
		return &pb.DeleteToDoResponse{}, nil
	} else {
		return &pb.DeleteToDoResponse{}, nil
	}
}
