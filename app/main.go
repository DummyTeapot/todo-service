package main

import (
	"database/sql"
	"todo-service/config"
	"todo-service/internal/handler"
	"todo-service/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func main() {
	dsn := "postgres://postgres:postgres@localhost:5432/todo?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())

	taskRepo := repository.NewTaskRepository(db)

	router := gin.Default()

	logger := config.InitLogger()
	taskHandler := handler.NewTaskHandler(taskRepo, logger)

	router.POST("/tasks", taskHandler.CreateTask)
	router.GET("/tasks/:id", taskHandler.GetTask)
	router.GET("/tasks", taskHandler.GetAllTasks)
	router.PUT("/tasks/:id", taskHandler.UpdateTask)
	router.DELETE("/tasks/:id", taskHandler.DeleteTask)

	logger.Info().Msgf("Start HTTP 8080")
	if err := router.Run(":8080"); err != nil {
		logger.Fatal().Err(err).Msg("HTTP error")
	}
}
