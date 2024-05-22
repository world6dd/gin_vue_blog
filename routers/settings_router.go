package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

func SettingsRouter(router *gin.Engine) {
	settingsAPI := api.ApiGroupAPP.SettingsApi
	router.GET("", settingsAPI.SettingsInfoView)
}
