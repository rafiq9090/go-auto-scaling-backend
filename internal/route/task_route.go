package route

import (
	"github.com/gin-gonic/gin"
	"github.com/rafiq9090/go-auto-scaling-backend/internal/handler"
)

func SetupTaskRoute(route *gin.RouterGroup) {
	tasks := route.Group("/tasks")
	tasks.GET("/", handler.GetAllTasks)
	tasks.POST("/", handler.CreateTask)
}
