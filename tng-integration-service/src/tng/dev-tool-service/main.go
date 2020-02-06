package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "tng/dev-tool-service/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"os"
)

func main() {
	orm.RunCommand()
	runMode := os.Getenv("RunMode")
	if runMode == "" {
		runMode = beego.BConfig.RunMode
	}
	if runMode == beego.DEV {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/dev-tool-service/swagger"] = "swagger"
	}
	beego.Run()
}
