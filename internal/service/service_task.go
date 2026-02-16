package service

import (
	"context"

	"github.com/rafiq9090/go-auto-scaling-backend/internal/model"
	"github.com/rafiq9090/go-auto-scaling-backend/internal/repositry"
)

type TaskService struct{}

var Task = TaskService{}

func (TaskService) Create(ctx context.Context, task *model.Task) error {
	return repositry.Task.Create(ctx, task)
}

func (TaskService) GetAll(ctx context.Context) ([]model.Task, error) {
	return repositry.Task.GetAll(ctx)
}

func (TaskService) GetByID(ctx context.Context, id string) (model.Task, error) {
	return repositry.Task.GetByID(ctx, id)
}

func (TaskService) Update(ctx context.Context, id string, task *model.Task) error {
	return repositry.Task.Update(ctx, id, task)
}

func (TaskService) Delete(ctx context.Context, id string) error {
	return repositry.Task.Delete(ctx, id)
}
