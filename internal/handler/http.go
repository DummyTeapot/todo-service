package handler

import (
	"net/http"
	"strconv"
	"todo-service/internal/model"
	"todo-service/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type TaskHandler struct {
	repo   *repository.TaskRepository
	logger zerolog.Logger
}

func NewTaskHandler(repo *repository.TaskRepository, logger zerolog.Logger) *TaskHandler {
	return &TaskHandler{repo: repo, logger: logger}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		h.logger.Error().Err(err).Msg("Неверный ввод")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.repo.CreateTask(c.Request.Context(), &task); err != nil {
		h.logger.Error().Err(err).Msg("Не удалось создать таск")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать таск"})
		return
	}
	c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) GetTask(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		h.logger.Error().Err(err).Msg("Неверный ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}
	task, err := h.repo.GetTaskByID(c.Request.Context(), id)
	if err != nil {
		h.logger.Error().Err(err).Msg("Таск не найден")
		c.JSON(http.StatusNotFound, gin.H{"error": "Таск не найден"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	tasks, err := h.repo.GetAllTasks(c.Request.Context())
	if err != nil {
		h.logger.Error().Err(err).Msg("Не удалось получить таск")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить таск"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		h.logger.Error().Err(err).Msg("Неверный ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		h.logger.Error().Err(err).Msg("Неверный ввод")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task.ID = id
	if err := h.repo.UpdateTask(c.Request.Context(), &task); err != nil {
		h.logger.Error().Err(err).Msg("Не удалось обновить таск")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить таск"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		h.logger.Error().Err(err).Msg("Неверный ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}
	if err := h.repo.DeleteTask(c.Request.Context(), id); err != nil {
		h.logger.Error().Err(err).Msg("Не удалось удалить таск")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить таск"})
		return
	}
	c.Status(http.StatusNoContent)
}
