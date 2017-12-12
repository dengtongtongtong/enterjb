package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["enterbj/controllers:UserdocumentController"] = append(beego.GlobalControllerRouter["enterbj/controllers:UserdocumentController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["enterbj/controllers:UserdocumentController"] = append(beego.GlobalControllerRouter["enterbj/controllers:UserdocumentController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["enterbj/controllers:UserdocumentController"] = append(beego.GlobalControllerRouter["enterbj/controllers:UserdocumentController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["enterbj/controllers:UserdocumentController"] = append(beego.GlobalControllerRouter["enterbj/controllers:UserdocumentController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["enterbj/controllers:UserdocumentController"] = append(beego.GlobalControllerRouter["enterbj/controllers:UserdocumentController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
