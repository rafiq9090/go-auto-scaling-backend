package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rafiq9090/go-auto-scaling-backend/internal/dto"
	"github.com/rafiq9090/go-auto-scaling-backend/internal/model"
	"github.com/rafiq9090/go-auto-scaling-backend/internal/service"
)

func GetAllTasks(c *gin.Context) {
	ctx := c.Request.Context()
	tasks, err := service.Task.GetAll(ctx)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, tasks)
}

func CreateTask(c *gin.Context) {
	var input dto.CreateTaskRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	task := model.Task{
		Name: input.Name,
	}
	if err := service.Task.Create(c, &task); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, task)
}
