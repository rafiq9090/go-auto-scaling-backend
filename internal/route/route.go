package route

import "github.com/gin-gonic/gin"

func SetupRoute(api *gin.RouterGroup) {
	SetupTaskRoute(api)
}
