package server

import (
	"context"
	pb "git.local/jmercado/todo/api"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc"
)

type apiServer struct {
	pb.UnimplementedTodoServiceServer
}

func NewApiServer(ctx context.Context) (pb.TodoServiceServer, error) {

	var options []grpc.ServerOption

	server := &apiServer{}

	s := grpc.NewServer(options...)
	pb.RegisterTodoServiceServer(s, &apiServer{})

	return server, nil
}

func (a *apiServer) CreatTask(ctx context.Context, in *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	return &pb.CreateTaskResponse{Task: in.Task}, nil
}
