package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["tng/dev-tool-service/controllers:Pi3UpdateInfoController"] = append(beego.GlobalControllerRouter["tng/dev-tool-service/controllers:Pi3UpdateInfoController"],
        beego.ControllerComments{
            Method: "GetPi3UpdateInfo",
            Router: `/get-pi3-update-info`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/dev-tool-service/controllers:Pi3UpdateInfoController"] = append(beego.GlobalControllerRouter["tng/dev-tool-service/controllers:Pi3UpdateInfoController"],
        beego.ControllerComments{
            Method: "RegisterPi3",
            Router: `/register-pi3`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

}
