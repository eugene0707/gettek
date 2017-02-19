// @APIVersion 1.0.0
// @Title GettEK API
// @Description Home assignment for golang from Gett
// @Contact eugene0707@gmail.com
// @TermsOfServiceUrl http://gettek.herokuapp.com/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/eugene0707/gettek/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{}, "*:Index")

	ns := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/drivers",
			beego.NSInclude(
				&controllers.DriversController{},
			),
		),
		beego.NSNamespace("/metrics",
			beego.NSInclude(
				&controllers.MetricsController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
