package repository

import (
	"context"
	"todo-service/internal/model"

	"github.com/uptrace/bun"
)

type TaskRepository struct {
	db *bun.DB
}

func NewTaskRepository(db *bun.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) CreateTask(ctx context.Context, task *model.Task) error {
	_, err := r.db.NewInsert().Model(task).Exec(ctx)
	return err
}

func (r *TaskRepository) GetTaskByID(ctx context.Context, id int64) (*model.Task, error) {
	task := new(model.Task)
	err := r.db.NewSelect().Model(task).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (r *TaskRepository) GetAllTasks(ctx context.Context) ([]*model.Task, error) {
	var tasks []*model.Task
	err := r.db.NewSelect().Model(&tasks).Order("created_at DESC").Scan(ctx)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepository) UpdateTask(ctx context.Context, task *model.Task) error {
	_, err := r.db.NewUpdate().Model(task).WherePK().Exec(ctx)
	return err
}

func (r *TaskRepository) DeleteTask(ctx context.Context, id int64) error {
	_, err := r.db.NewDelete().Model(&model.Task{}).Where("id = ?", id).Exec(ctx)
	return err
}
