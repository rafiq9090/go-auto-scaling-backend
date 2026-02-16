package repositry

import (
	"context"

	"github.com/rafiq9090/go-auto-scaling-backend/internal/db"
	"github.com/rafiq9090/go-auto-scaling-backend/internal/model"
)

type TaskRepository struct{}

var Task = TaskRepository{}

func (TaskRepository) Create(ctx context.Context, task *model.Task) error {
	return db.DB.WithContext(ctx).Create(task).Error
}

func (TaskRepository) GetAll(ctx context.Context) ([]model.Task, error) {
	var tasks []model.Task
	err := db.DB.WithContext(ctx).Find(&tasks).Error
	return tasks, err
}

func (TaskRepository) GetByID(ctx context.Context, id string) (model.Task, error) {
	var task model.Task
	err := db.DB.WithContext(ctx).First(&task, id).Error
	return task, err
}

func (TaskRepository) Update(ctx context.Context, id string, task *model.Task) error {
	return db.DB.WithContext(ctx).Where("id = ?", id).Updates(task).Error
}

func (TaskRepository) Delete(ctx context.Context, id string) error {
	return db.DB.WithContext(ctx).Delete(&model.Task{}, id).Error
}
