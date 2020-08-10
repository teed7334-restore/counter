package route

import (
	"os"

	"github.com/teed7334-restore/counter/controllers"

	"github.com/gin-gonic/gin"
)

//API Restful路由
func API() *gin.Engine {
	env := os.Getenv("env")
	gin.SetMode(env)
	route := gin.Default()
	route.POST("/Mail/SendMail", controllers.SendMail)
	return route
}
