package main

import (
	_ "DataCertPlatform/routers"
	"github.com/astaxie/beego"
	"DataCertPlatform/db_mysql"
)

func main() {
	db_mysql.Connect()
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run( )
}

