package routers

import (
	"DataCertPlatform/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.RegisterController{})
    //用户登录的接口
    beego.Router("/login",&controllers.LoginController{})
    beego.Router("/login.html",&controllers.LoginController{})
}
