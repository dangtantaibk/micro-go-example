package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["tng/h5-integration-service/controllers:H5ZaloPayController"] = append(beego.GlobalControllerRouter["tng/h5-integration-service/controllers:H5ZaloPayController"],
        beego.ControllerComments{
            Method: "GrantedMBToken",
            Router: `/grantedmbtoken`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/h5-integration-service/controllers:H5ZaloPayController"] = append(beego.GlobalControllerRouter["tng/h5-integration-service/controllers:H5ZaloPayController"],
        beego.ControllerComments{
            Method: "GetPaymentOrderUrl",
            Router: `/paymenturl`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

}
