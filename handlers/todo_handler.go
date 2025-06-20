package handlers

import (
	"context"

	pb "github.com/lacriman/todo-grpc/pb/todo"
)

type Server struct {
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
