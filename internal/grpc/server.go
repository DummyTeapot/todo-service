package grpc

import (
	"context"
	"net"
	"todo-service/internal/repository"
	"todo-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type TaskGRPCServer struct {
	proto.UnimplementedTaskServiceServer
	repo *repository.TaskRepository
}

func NewTaskGRPCServer(repo *repository.TaskRepository) *TaskGRPCServer {
	return &TaskGRPCServer{repo: repo}
}

func (s *TaskGRPCServer) GetTasks(ctx context.Context, req *proto.GetTasksRequest) (*proto.GetTasksResponse, error) {
	tasks, err := s.repo.GetAllTasks(ctx)
	if err != nil {
		return nil, err
	}
	var protoTasks []*proto.Task
	for _, task := range tasks {
		protoTasks = append(protoTasks, &proto.Task{
			Id:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			CreatedAt:   task.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			Completed:   task.Completed,
		})
	}
	return &proto.GetTasksResponse{Tasks: protoTasks}, nil
}

func (s *TaskGRPCServer) UpdateTaskStatus(ctx context.Context, req *proto.UpdateTaskStatusRequest) (*proto.UpdateTaskStatusResponse, error) {
	task, err := s.repo.GetTaskByID(ctx, req.Id)
	if err != nil {
		return &proto.UpdateTaskStatusResponse{Success: false}, err
	}
	task.Completed = req.Completed
	if err := s.repo.UpdateTask(ctx, task); err != nil {
		return &proto.UpdateTaskStatusResponse{Success: false}, err
	}
	return &proto.UpdateTaskStatusResponse{Success: true}, nil
}

func StartGRPCServer(grpcAddr string, repo *repository.TaskRepository) error {
	server := grpc.NewServer()
	taskServer := NewTaskGRPCServer(repo)
	proto.RegisterTaskServiceServer(server, taskServer)
	reflection.Register(server)
	return server.Serve(lis(grpcAddr))
}

func lis(addr string) net.Listener {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	return listener
}
