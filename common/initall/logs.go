package initall

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func InitLogs() (err error) {
	logPath := LogConfAll.LogConf.LogPath
	config := fmt.Sprintf(`{"filename":"%s"}`, logPath)

	beego.SetLogger("file", config)

	beego.SetLevel(LogConfAll.LogConf.LogLevel)

	beego.SetLogFuncCall(true)

	logs.Error("init logs success")
	return
}
