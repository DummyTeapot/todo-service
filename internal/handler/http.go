package handler

import (
	"todo-service/internal/repository"

	"github.com/rs/zerolog"
)

type TaskHandler struct {
	repo   *repository.TaskRepository
	logger zerolog.Logger
}

func NewTaskHandler(repo *repository.TaskRepository, logger zerolog.Logger) *TaskHandler {
	return &TaskHandler{repo: repo, logger: logger}
}
