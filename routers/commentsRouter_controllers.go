package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/eugene0707/gettek/controllers:DriversController"] = append(beego.GlobalControllerRouter["github.com/eugene0707/gettek/controllers:DriversController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/eugene0707/gettek/controllers:DriversController"] = append(beego.GlobalControllerRouter["github.com/eugene0707/gettek/controllers:DriversController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/eugene0707/gettek/controllers:DriversController"] = append(beego.GlobalControllerRouter["github.com/eugene0707/gettek/controllers:DriversController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/eugene0707/gettek/controllers:DriversController"] = append(beego.GlobalControllerRouter["github.com/eugene0707/gettek/controllers:DriversController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/eugene0707/gettek/controllers:DriversController"] = append(beego.GlobalControllerRouter["github.com/eugene0707/gettek/controllers:DriversController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/eugene0707/gettek/controllers:MetricsController"] = append(beego.GlobalControllerRouter["github.com/eugene0707/gettek/controllers:MetricsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/eugene0707/gettek/controllers:MetricsController"] = append(beego.GlobalControllerRouter["github.com/eugene0707/gettek/controllers:MetricsController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/eugene0707/gettek/controllers:MetricsController"] = append(beego.GlobalControllerRouter["github.com/eugene0707/gettek/controllers:MetricsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/eugene0707/gettek/controllers:MetricsController"] = append(beego.GlobalControllerRouter["github.com/eugene0707/gettek/controllers:MetricsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/eugene0707/gettek/controllers:MetricsController"] = append(beego.GlobalControllerRouter["github.com/eugene0707/gettek/controllers:MetricsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
