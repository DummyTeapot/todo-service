package main

import (
	"database/sql"
	"todo-service/config"
	"todo-service/internal/grpc"
	"todo-service/internal/handler"
	"todo-service/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func main() {
	logger := config.InitLogger()

	dsn := "postgres://postgres:postgres@localhost:5432/todo?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())
	defer func() {
		if err := db.Close(); err != nil {
			logger.Error().Err(err).Msg("Не удалось закрыть соединение с БД")
		}
	}()

	if err := db.Ping(); err != nil {
		logger.Fatal().Err(err).Msg("Не удалось подключиться к БД")
	}
	logger.Info().Msg("Успешное подключение к БД")

	taskRepo := repository.NewTaskRepository(db)

	router := gin.Default()

	taskHandler := handler.NewTaskHandler(taskRepo, logger)

	router.POST("/tasks", taskHandler.CreateTask)
	router.GET("/tasks/:id", taskHandler.GetTask)
	router.GET("/tasks", taskHandler.GetAllTasks)
	router.PUT("/tasks/:id", taskHandler.UpdateTask)
	router.DELETE("/tasks/:id", taskHandler.DeleteTask)

	go func() {
		logger.Info().Msgf("Старт gRPC сервера на 50051")
		if err := grpc.StartGRPCServer(":50051", taskRepo); err != nil {
			logger.Fatal().Err(err).Msg("Ошибка gRPC сервера")
		}
	}()

	logger.Info().Msgf("Старт HTTP сервера на 8080")
	if err := router.Run(":8080"); err != nil {
		logger.Fatal().Err(err).Msg("Ошибка HTTP сервера")
	}
}
