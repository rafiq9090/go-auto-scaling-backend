package route

import (
	"github.com/gin-gonic/gin"
	"github.com/rafiq9090/go-auto-scaling-backend/internal/handler"
)

func SetupTaskRoute(route *gin.RouterGroup) {
	tasks := route.Group("/tasks")
	tasks.GET("/", handler.GetAllTasks)
	tasks.POST("/", handler.CreateTask)
	tasks.GET("/:id", handler.GetTaskByID)
	tasks.PUT("/:id", handler.UpdateTask)
	tasks.DELETE("/:id", handler.DeleteTask)
}
