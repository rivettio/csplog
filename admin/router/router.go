package router

import (
	"rivettio/github-csplog/csplog/admin/controller"
	"rivettio/github-csplog/csplog/admin/controller/config"
	_ "rivettio/github-csplog/csplog/admin/model"

	"github.com/astaxie/beego"
	"github.com/beego/admin"
)

func init() {
	admin.Run()
	beego.Router("/", &controller.MainController{})

	beego.Router("/elk/LogConfig/index", &config.LogConfigController{}, "*:Index")
	beego.Router("/elk/LogConfig/add", &config.LogConfigController{}, "*:AddLogConfig")
	beego.Router("/elk/LogConfig/update", &config.LogConfigController{}, "*:UpdateLogConfig")
}
