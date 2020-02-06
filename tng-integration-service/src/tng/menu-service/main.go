package main

import (
	_ "tng/menu-service/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"os"
)

func main() {
	orm.RunCommand()
	runMode := os.Getenv("RUNMODE")
	if runMode == "" {
		runMode = beego.BConfig.RunMode
	}
	if runMode == beego.DEV {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/menu-service/swagger"] = "swagger"
	}
	beego.Run()
}
