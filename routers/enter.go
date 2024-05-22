package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	//gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	SettingsRouter(router)
	return router

}
