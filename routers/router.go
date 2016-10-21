package routers

import (
	"git.chunyu.me/feiwang/appserver/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/api/mp/:app_id/", &controllers.MainController{}, "get:MobileProvision4Key")
	beego.Router("/api/icon/:app_id/", &controllers.MainController{}, "get:AppIcon")
	beego.Router("/api/plist/:app_id/", &controllers.MainController{}, "get:PlistFile")
	beego.Router("/api/ipa/:app_id/", &controllers.MainController{}, "get:AppIpa")
	beego.Router("/api/apk/:app_id/", &controllers.MainController{}, "get:AndroidApk")
}
