package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	_ "tng/cron-job-service/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"os"
	hdl "tng/cron-job-service/cronjob"
)

func main() {
	orm.RunCommand()
	runMode := os.Getenv("RUNMODE")
	if runMode == "" {
		runMode = beego.BConfig.RunMode
	}
	if runMode == beego.DEV {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/cron-job-service/swagger"] = "swagger"
	}
	go func() {
		h := hdl.NewCronHandler(context.Background())
		h.InitCronJob(context.Background())
		h.Run()
	}()
	beego.Run()
}
