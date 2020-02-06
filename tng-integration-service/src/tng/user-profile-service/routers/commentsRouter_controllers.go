package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["tng/user-profile-service/controllers:AuthenticationController"] = append(beego.GlobalControllerRouter["tng/user-profile-service/controllers:AuthenticationController"],
        beego.ControllerComments{
            Method: "CheckLogin",
            Router: `/check-login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/user-profile-service/controllers:AuthenticationController"] = append(beego.GlobalControllerRouter["tng/user-profile-service/controllers:AuthenticationController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/user-profile-service/controllers:UserProfileController"] = append(beego.GlobalControllerRouter["tng/user-profile-service/controllers:UserProfileController"],
        beego.ControllerComments{
            Method: "Create",
            Router: `/create`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("body", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/user-profile-service/controllers:UserProfileController"] = append(beego.GlobalControllerRouter["tng/user-profile-service/controllers:UserProfileController"],
        beego.ControllerComments{
            Method: "GetByID",
            Router: `/get-by-id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["tng/user-profile-service/controllers:UserProfileController"] = append(beego.GlobalControllerRouter["tng/user-profile-service/controllers:UserProfileController"],
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
