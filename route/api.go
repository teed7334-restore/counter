package route

import (
	"github.com/teed7334-restore/counter/controllers"

	"github.com/teed7334-restore/counter/env"

	"github.com/gin-gonic/gin"
)

var cfg = env.GetEnv()

//API Restful路由
func API() *gin.Engine {
	gin.SetMode(cfg.Env)
	route := gin.Default()
	route.POST("/Mail/SendMail", controllers.SendMail)
	route.POST("/PunchClock/UploadDailyPunchclockData", controllers.UploadDailyPunchclockData)
	return route
}
