package handlers

import (
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

