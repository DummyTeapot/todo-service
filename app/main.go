package main

import (
	"database/sql"
	"fmt"
	"todo-service/config"
	"todo-service/internal/handler"
	"todo-service/internal/repository"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func main() {
	dsn := "postgres://postgres:postgres@localhost:5432/todo?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())
	// fmt.Println(db)

	taskRepo := repository.NewTaskRepository(db)

	logger := config.InitLogger()
	taskHandler := handler.NewTaskHandler(taskRepo, logger)

	fmt.Println(taskHandler)
}
