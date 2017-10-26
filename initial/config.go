package initial

import (
	"flag"
	"github.com/astaxie/beego"
	"fmt"
)

const usage = `
Usage: runmode [options]
    default runmode is dev
    options can be any of [dev test prod]
	`
const adapterConf = "ini"

var runmode string

func flagBind() {
	flag.StringVar(&runmode, "r", "test", usage)
	flag.Parse()
}

func checkRunMode() (b bool) {
	rms := []string{"dev", "test", "prod"}
	for _, r := range rms {
		if r == runmode {
			b = true
		}
	}
	return b
}

func flagCheck() {
	if checkRunMode() == false {
		panic("flag runmode must be one of [dev, test, prod]")
	}
}

func initConfig() {
	flagBind()
	flagCheck()
	fmt.Println(fmt.Sprintf("conf/%s.conf", runmode))
	beego.LoadAppConfig(adapterConf, fmt.Sprintf("conf/%s.conf", runmode))
}
