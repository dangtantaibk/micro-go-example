package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["tng/cron-job-service/controllers:ScheduleController"] = append(beego.GlobalControllerRouter["tng/cron-job-service/controllers:ScheduleController"],
        beego.ControllerComments{
            Method: "WarmUp",
            Router: `/warm-up`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

}
