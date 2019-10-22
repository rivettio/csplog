package config

import (
	"fmt"
	"github.com/astaxie/beego"
	"rivettio/github-csplog/csplog/admin/model"
	"time"
)

type LogConfigController struct {
	beego.Controller
}

var LogConfigModel *model.LogConfig = model.NewLogConfigModel()

func (lcc *LogConfigController) Error(err interface{}) {
	url := lcc.Ctx.Request.Referer()
	lcc.Data["error"] = fmt.Sprintln(err)
	fmt.Println(err)
	lcc.Redirect(url, 302)
	return
}

func (lcc *LogConfigController) Success(url string, msg interface{}) {
	lcc.Data["message"] = msg
	fmt.Println(msg)
	lcc.Redirect(url, 302)
	return
}

func (lcc *LogConfigController) Index() {
	LogConfigList, err := LogConfigModel.GetLogConfigList(0, 20)
	if err != nil {
		lcc.Error(err)
		return
	}
	lcc.Data["LogConfigList"] = LogConfigList
	lcc.TplName = "config/index.html"
	return
}

func (lcc *LogConfigController) AddLogConfig() {
	if lcc.Ctx.Input.IsGet() {
		lcc.TplName = "config/add.html"
		return
	}
	if lcc.Ctx.Input.IsPost() {
		LogPath := lcc.GetString("LogPath")
		if len(LogPath) == 0 {
			lcc.Error("please enter log path")
			return
		}
		Topic := lcc.GetString("Topic")
		if len(Topic) == 0 {
			lcc.Error("please enter the Topic")
			return
		}
		Service := lcc.GetString("Service")
		if len(Service) == 0 {
			lcc.Error("please enter service name")
			return
		}
		SendRate, err := lcc.GetInt("SendRate")
		if err != nil {
			err = fmt.Errorf("get send kafka rate err : %v", err)
			lcc.Error(err)
			return
		}
		Status, err := lcc.GetInt("Status")
		if err != nil {
			err = fmt.Errorf("get is collection err : %v", err)
			lcc.Error(err)
			return
		}
		LogConfig := &model.LogConfig{
			LogPath:  LogPath,
			Topic:    Topic,
			Service:  Service,
			SendRate: SendRate,
			Status:   Status,
			AddTime:  time.Now().Local(),
		}
		_, err = LogConfigModel.InsertLogConfig(LogConfig)
		if err != nil {
			lcc.Error(err)
			return
		}
		lcc.Success("index", "add success")
		return
	}
	lcc.Ctx.WriteString("data")
	return
}

func (lcc *LogConfigController) UpdateLogConfig() {
	lcc.Ctx.WriteString("print string")
	return
}
