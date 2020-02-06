package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["tng/promotion-service/controllers:CampaignController"] = append(beego.GlobalControllerRouter["tng/promotion-service/controllers:CampaignController"],
        beego.ControllerComments{
            Method: "Insert",
            Router: `/insert`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/promotion-service/controllers:CampaignController"] = append(beego.GlobalControllerRouter["tng/promotion-service/controllers:CampaignController"],
        beego.ControllerComments{
            Method: "List",
            Router: `/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/promotion-service/controllers:PromotionController"] = append(beego.GlobalControllerRouter["tng/promotion-service/controllers:PromotionController"],
        beego.ControllerComments{
            Method: "Insert",
            Router: `/insert`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

}
