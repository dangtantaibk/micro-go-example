package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["tng/shipper-service/controllers:InvoiceController"] = append(beego.GlobalControllerRouter["tng/shipper-service/controllers:InvoiceController"],
        beego.ControllerComments{
            Method: "AllList",
            Router: `/all-list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/shipper-service/controllers:InvoiceController"] = append(beego.GlobalControllerRouter["tng/shipper-service/controllers:InvoiceController"],
        beego.ControllerComments{
            Method: "GetInvoiceDetail",
            Router: `/invoice-detail`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/shipper-service/controllers:InvoiceController"] = append(beego.GlobalControllerRouter["tng/shipper-service/controllers:InvoiceController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/shipper-service/controllers:InvoiceController"] = append(beego.GlobalControllerRouter["tng/shipper-service/controllers:InvoiceController"],
        beego.ControllerComments{
            Method: "ScanQRCode",
            Router: `/scan-qr-code`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/shipper-service/controllers:InvoiceController"] = append(beego.GlobalControllerRouter["tng/shipper-service/controllers:InvoiceController"],
        beego.ControllerComments{
            Method: "UpdateStatus",
            Router: `/update-status`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/shipper-service/controllers:ShipperController"] = append(beego.GlobalControllerRouter["tng/shipper-service/controllers:ShipperController"],
        beego.ControllerComments{
            Method: "DeleteShipper",
            Router: `/delete`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/shipper-service/controllers:ShipperController"] = append(beego.GlobalControllerRouter["tng/shipper-service/controllers:ShipperController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/shipper-service/controllers:ShipperController"] = append(beego.GlobalControllerRouter["tng/shipper-service/controllers:ShipperController"],
        beego.ControllerComments{
            Method: "Signup",
            Router: `/signup`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/shipper-service/controllers:ShipperController"] = append(beego.GlobalControllerRouter["tng/shipper-service/controllers:ShipperController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/update`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/shipper-service/controllers:UserController"] = append(beego.GlobalControllerRouter["tng/shipper-service/controllers:UserController"],
        beego.ControllerComments{
            Method: "LoginWithPassword",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/shipper-service/controllers:UserController"] = append(beego.GlobalControllerRouter["tng/shipper-service/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/loginzalo`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/shipper-service/controllers:UserController"] = append(beego.GlobalControllerRouter["tng/shipper-service/controllers:UserController"],
        beego.ControllerComments{
            Method: "RefreshToken",
            Router: `/refresh`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/shipper-service/controllers:UserController"] = append(beego.GlobalControllerRouter["tng/shipper-service/controllers:UserController"],
        beego.ControllerComments{
            Method: "Signup",
            Router: `/signup`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/shipper-service/controllers:UserController"] = append(beego.GlobalControllerRouter["tng/shipper-service/controllers:UserController"],
        beego.ControllerComments{
            Method: "VerifyPhoneNumber",
            Router: `/verify`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

}
