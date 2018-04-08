package routers

import (
	"talkGoNotify/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/notify",
			beego.NSInclude(
				&controllers.NotifyController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
