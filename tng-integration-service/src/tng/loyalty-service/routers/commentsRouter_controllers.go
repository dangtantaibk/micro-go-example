package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["tng/loyalty-service/controllers:ClassTrackingController"] = append(beego.GlobalControllerRouter["tng/loyalty-service/controllers:ClassTrackingController"],
        beego.ControllerComments{
            Method: "DeleteClassTracking",
            Router: `/delete`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/loyalty-service/controllers:ClassTrackingController"] = append(beego.GlobalControllerRouter["tng/loyalty-service/controllers:ClassTrackingController"],
        beego.ControllerComments{
            Method: "GetByID",
            Router: `/getbyid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/loyalty-service/controllers:ClassTrackingController"] = append(beego.GlobalControllerRouter["tng/loyalty-service/controllers:ClassTrackingController"],
        beego.ControllerComments{
            Method: "InsertOrUpdate",
            Router: `/insertorupdate`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/loyalty-service/controllers:ClassTrackingController"] = append(beego.GlobalControllerRouter["tng/loyalty-service/controllers:ClassTrackingController"],
        beego.ControllerComments{
            Method: "ListClassTracking",
            Router: `/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/loyalty-service/controllers:PointClassController"] = append(beego.GlobalControllerRouter["tng/loyalty-service/controllers:PointClassController"],
        beego.ControllerComments{
            Method: "DeletePointClass",
            Router: `/delete`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/loyalty-service/controllers:PointClassController"] = append(beego.GlobalControllerRouter["tng/loyalty-service/controllers:PointClassController"],
        beego.ControllerComments{
            Method: "GetByID",
            Router: `/getbyid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/loyalty-service/controllers:PointClassController"] = append(beego.GlobalControllerRouter["tng/loyalty-service/controllers:PointClassController"],
        beego.ControllerComments{
            Method: "InsertOrUpdate",
            Router: `/insertorupdate`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/loyalty-service/controllers:PointClassController"] = append(beego.GlobalControllerRouter["tng/loyalty-service/controllers:PointClassController"],
        beego.ControllerComments{
            Method: "ListPointClass",
            Router: `/list`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/loyalty-service/controllers:PointController"] = append(beego.GlobalControllerRouter["tng/loyalty-service/controllers:PointController"],
        beego.ControllerComments{
            Method: "AddPoint",
            Router: `/addpoint`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/loyalty-service/controllers:PointController"] = append(beego.GlobalControllerRouter["tng/loyalty-service/controllers:PointController"],
        beego.ControllerComments{
            Method: "CheckOldPoint",
            Router: `/checkoldpoint`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/loyalty-service/controllers:PointController"] = append(beego.GlobalControllerRouter["tng/loyalty-service/controllers:PointController"],
        beego.ControllerComments{
            Method: "CheckPoint",
            Router: `/checkpoint`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/loyalty-service/controllers:PointController"] = append(beego.GlobalControllerRouter["tng/loyalty-service/controllers:PointController"],
        beego.ControllerComments{
            Method: "DeletePoint",
            Router: `/delete`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/loyalty-service/controllers:PointController"] = append(beego.GlobalControllerRouter["tng/loyalty-service/controllers:PointController"],
        beego.ControllerComments{
            Method: "GetByID",
            Router: `/getbyid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/loyalty-service/controllers:PointController"] = append(beego.GlobalControllerRouter["tng/loyalty-service/controllers:PointController"],
        beego.ControllerComments{
            Method: "InsertOrUpdate",
            Router: `/insertorupdate`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/loyalty-service/controllers:PointController"] = append(beego.GlobalControllerRouter["tng/loyalty-service/controllers:PointController"],
        beego.ControllerComments{
            Method: "ListPoint",
            Router: `/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/loyalty-service/controllers:PointController"] = append(beego.GlobalControllerRouter["tng/loyalty-service/controllers:PointController"],
        beego.ControllerComments{
            Method: "PointHistory",
            Router: `/pointhistory`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/loyalty-service/controllers:PointController"] = append(beego.GlobalControllerRouter["tng/loyalty-service/controllers:PointController"],
        beego.ControllerComments{
            Method: "SearchPoint",
            Router: `/search`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/loyalty-service/controllers:PointTypeController"] = append(beego.GlobalControllerRouter["tng/loyalty-service/controllers:PointTypeController"],
        beego.ControllerComments{
            Method: "DeletePointType",
            Router: `/delete`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/loyalty-service/controllers:PointTypeController"] = append(beego.GlobalControllerRouter["tng/loyalty-service/controllers:PointTypeController"],
        beego.ControllerComments{
            Method: "GetByID",
            Router: `/getbyid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/loyalty-service/controllers:PointTypeController"] = append(beego.GlobalControllerRouter["tng/loyalty-service/controllers:PointTypeController"],
        beego.ControllerComments{
            Method: "InsertOrUpdate",
            Router: `/insertorupdate`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/loyalty-service/controllers:PointTypeController"] = append(beego.GlobalControllerRouter["tng/loyalty-service/controllers:PointTypeController"],
        beego.ControllerComments{
            Method: "ListPointType",
            Router: `/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/loyalty-service/controllers:SettingController"] = append(beego.GlobalControllerRouter["tng/loyalty-service/controllers:SettingController"],
        beego.ControllerComments{
            Method: "DeleteSetting",
            Router: `/delete`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/loyalty-service/controllers:SettingController"] = append(beego.GlobalControllerRouter["tng/loyalty-service/controllers:SettingController"],
        beego.ControllerComments{
            Method: "GetByID",
            Router: `/getbyid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/loyalty-service/controllers:SettingController"] = append(beego.GlobalControllerRouter["tng/loyalty-service/controllers:SettingController"],
        beego.ControllerComments{
            Method: "InsertOrUpdate",
            Router: `/insertorupdate`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/loyalty-service/controllers:SettingController"] = append(beego.GlobalControllerRouter["tng/loyalty-service/controllers:SettingController"],
        beego.ControllerComments{
            Method: "ListSetting",
            Router: `/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/loyalty-service/controllers:WalletController"] = append(beego.GlobalControllerRouter["tng/loyalty-service/controllers:WalletController"],
        beego.ControllerComments{
            Method: "DeleteWallet",
            Router: `/delete`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/loyalty-service/controllers:WalletController"] = append(beego.GlobalControllerRouter["tng/loyalty-service/controllers:WalletController"],
        beego.ControllerComments{
            Method: "GetByID",
            Router: `/getbyid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/loyalty-service/controllers:WalletController"] = append(beego.GlobalControllerRouter["tng/loyalty-service/controllers:WalletController"],
        beego.ControllerComments{
            Method: "InsertOrUpdate",
            Router: `/insertorupdate`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/loyalty-service/controllers:WalletController"] = append(beego.GlobalControllerRouter["tng/loyalty-service/controllers:WalletController"],
        beego.ControllerComments{
            Method: "ListWallet",
            Router: `/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
