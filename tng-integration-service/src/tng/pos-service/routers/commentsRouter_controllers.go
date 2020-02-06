package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["tng/pos-service/controllers:InvoiceController"] = append(beego.GlobalControllerRouter["tng/pos-service/controllers:InvoiceController"],
        beego.ControllerComments{
            Method: "Cancel",
            Router: `/cancel`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/pos-service/controllers:InvoiceController"] = append(beego.GlobalControllerRouter["tng/pos-service/controllers:InvoiceController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/create`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/pos-service/controllers:InvoiceController"] = append(beego.GlobalControllerRouter["tng/pos-service/controllers:InvoiceController"],
        beego.ControllerComments{
            Method: "Refund",
            Router: `/refund`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

}
