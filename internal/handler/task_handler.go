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

func GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	task, err := service.Task.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(200, task)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var input dto.UpdateTaskRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var task model.Task
	if input.Name != "" {
		task.Name = input.Name
	}
	if input.Completed != nil {
		task.Completed = *input.Completed
	}

	if err := service.Task.Update(c.Request.Context(), id, &task); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	// Fetch updated task to return
	updatedTask, err := service.Task.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(200, task) // Fallback if fetch fails, though unlikely
		return
	}
	c.JSON(200, updatedTask)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := service.Task.Delete(c.Request.Context(), id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Task deleted successfully"})
}
