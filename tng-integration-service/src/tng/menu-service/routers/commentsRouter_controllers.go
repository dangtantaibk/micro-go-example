package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["tng/menu-service/controllers:AreaController"] = append(beego.GlobalControllerRouter["tng/menu-service/controllers:AreaController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/menu-service/controllers:CategoryController"] = append(beego.GlobalControllerRouter["tng/menu-service/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/create`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/menu-service/controllers:CategoryController"] = append(beego.GlobalControllerRouter["tng/menu-service/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "DeleteCategory",
            Router: `/delete`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/menu-service/controllers:CategoryController"] = append(beego.GlobalControllerRouter["tng/menu-service/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/menu-service/controllers:CategoryController"] = append(beego.GlobalControllerRouter["tng/menu-service/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/update`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/menu-service/controllers:CategoryController"] = append(beego.GlobalControllerRouter["tng/menu-service/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "UpdateStatus",
            Router: `/update-status`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/menu-service/controllers:ItemTypeController"] = append(beego.GlobalControllerRouter["tng/menu-service/controllers:ItemTypeController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/create`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/menu-service/controllers:ItemTypeController"] = append(beego.GlobalControllerRouter["tng/menu-service/controllers:ItemTypeController"],
        beego.ControllerComments{
            Method: "DeleteItemType",
            Router: `/delete`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/menu-service/controllers:ItemTypeController"] = append(beego.GlobalControllerRouter["tng/menu-service/controllers:ItemTypeController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/menu-service/controllers:ItemTypeController"] = append(beego.GlobalControllerRouter["tng/menu-service/controllers:ItemTypeController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/update`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

}
