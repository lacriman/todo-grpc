package handlers

import (
	"context"

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

func (s *Server) ListToDo(req *pb.ListToDoRequest, stream pb.ToDoService_ListToDoServer) error {
	for _, todo := range s.todos {
		res := &pb.ListToDoResponse{Todos: todo}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
	return nil
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
