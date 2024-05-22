package routers

import (
	"gvb_server/api"
)

func (router RouterGroup) SettingsRouter() {
	settingsAPI := api.ApiGroupAPP.SettingsApi
	router.GET("", settingsAPI.SettingsInfoView)
}
