package initial

import (
	//"github.com/astaxie/beego"
	"sc-git.com/beeApi/utils/logs"
)

func initLog() {
	//logs.InitLog(beego.AppConfig.String("emailName"), beego.AppConfig.String("emailPwd"))
	logs.InitLog("", "")
}
