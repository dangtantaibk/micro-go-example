// Package routers routes all url of system.
// @APIVersion 1.0.0
// @Title Pos Service
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
// @URL /tng
package routers

import (
	_stdCtx "context"
	"fmt"
	"os"
	"strings"
	"time"
	"tng/common/models"
	"tng/common/models/pos"
	"tng/common/utils/merchantutil"
	dtos2 "tng/pos-service/dtos"

	"tng/common/httpcode"
	"tng/common/logger"
	"tng/common/utils/cfgutil"
	"tng/common/utils/msgutil"
	"tng/common/utils/redisutil"
	"tng/pos-service/controllers"
	"tng/pos-service/services"

	"github.com/astaxie/beego"
	_beeCtx "github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
)

// Main constants definition.
const (
	ApplicationLoadFail = 1
)

func init() {
	var (
		errMsgFilePath   = cfgutil.Load("ERROR_MESSAGE_FILE_PATH")
		merchantFilePath = cfgutil.Load("MERCHANT_FILE_PATH")
	)
	if err := msgutil.InitialErrorMessageResource(errMsgFilePath); err != nil {
		logs.Error("Initializing error message resource: %v", err)
		os.Exit(ApplicationLoadFail)
	}
	if err := merchantutil.InitialResource(merchantFilePath); err != nil {
		logs.Error("Initializing merchant message resource: %v", err)
		os.Exit(ApplicationLoadFail)
	}
	// Initial DB, JWTHelper, ServiceProvider
	initialDB()
	_services := services.InitialServices()

	var (
		_redisStore redisutil.Cache
	)
	_ = _services.Invoke(func(store redisutil.Cache) {
		_redisStore = store
	})

	// Add contexts into each request
	beego.BConfig.RecoverFunc = recoverPanic
	beego.InsertFilter("*", beego.BeforeRouter, beforeRouterHandler)
	beego.InsertFilter("*", beego.AfterExec, afterExecHandler, false)

	// CORS
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     strings.Split(cfgutil.Load("ALLOW_DOMAIN_FE"), ","),
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PATCH", "PUT"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type", "timezoneoffset"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           time.Second * 86400}))
	// Initial routers
	ns := beego.NewNamespace("/vpos/api/v1",
		beego.NSNamespace("/invoice",
			beego.NSInclude(&controllers.InvoiceController{}),
		),
	)

	beego.AddNamespace(ns)
}

func beforeRouterHandler(c *_beeCtx.Context) {
	rqCtx := c.Request.Context()
	rqCtx = _stdCtx.WithValue(rqCtx, logger.RqIDCtxKey, logger.NewRequestID())
	rqCtx = _stdCtx.WithValue(rqCtx, logger.RqExecTimeCtxKey, time.Now())   // add exec time
	rqCtx = _stdCtx.WithValue(rqCtx, logger.RqClientIPCtxKey, c.Input.IP()) // add client IP
	rqCtx = _stdCtx.WithValue(rqCtx, logger.RqURICtxKey, c.Input.URI())
	c.Request = c.Request.WithContext(rqCtx)
}

func afterExecHandler(c *_beeCtx.Context) {
	logger.CtxLog(c.Request.Context(), c.ResponseWriter.Status, beego.BConfig.AppName)
}

// initialDB Initialize DB connection
func initialDB() {
	err := orm.RegisterDriver(models.DBDriver, orm.DRMySQL)
	if err != nil {
		logs.Error("Registering DB driver: %v", err)
		os.Exit(ApplicationLoadFail)
	}

	mysqlUser := cfgutil.Load("DB_USERNAME")
	mysqlPassword := cfgutil.Load("DB_PASSWORD")
	mysqlHost := cfgutil.Load("DB_HOST")
	mysqlPort := cfgutil.Load("DB_PORT")
	maxIdle, _ := cfgutil.LoadInt("DB_MAX_IDLE")
	maxConn, _ := cfgutil.LoadInt("DB_MAX_CONNECTION")
	mysqlDatabase := cfgutil.Load("DB_DATABASE")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		mysqlUser,
		mysqlPassword,
		mysqlHost,
		mysqlPort,
		mysqlDatabase)

	err = orm.RegisterDataBase(models.DBDefaultAlias, models.DBDriver, dataSource, maxIdle, maxConn)
	if err != nil {
		logs.Error("Connecting DB: %v", err)
		os.Exit(ApplicationLoadFail)
	}
	for _, mc := range merchantutil.GetKeys() {
		mysqlDatabase := fmt.Sprintf("vpos_%v", mc)
		dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
			mysqlUser,
			mysqlPassword,
			mysqlHost,
			mysqlPort,
			mysqlDatabase)

		err = orm.RegisterDataBase(mysqlDatabase, models.DBDriver, dataSource, maxIdle, maxConn)
		if err != nil {
			logs.Error("Connecting DB: %v", err)
			os.Exit(ApplicationLoadFail)
		}
	}

	orm.DefaultTimeLoc = time.UTC
	orm.RegisterModel(new(pos.Item))
	orm.RegisterModel(new(pos.Invoice))
	orm.RegisterModel(new(pos.InvoiceDetail))
}

func recoverPanic(c *_beeCtx.Context) {
	if err := recover(); err != nil {
		if err == beego.ErrAbort {
			return
		}
		logger.Errorf(c.Request.Context(), "Panicking request: %v", err)
		httpStatusCode := httpcode.ServerErrorCode
		c.Output.SetStatus(httpStatusCode)
		_ = c.Output.JSON(dtos2.NewAppError(httpStatusCode), false, false)
		logger.CtxLog(c.Request.Context(), c.ResponseWriter.Status, beego.BConfig.AppName)
	}
}
