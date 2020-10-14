package controllers

import (
	"DataCertPlatform/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	l.TplName = "login.html"
}
/**
*post方法处理用户的登录请求
 */
func (l *LoginController) post() {
	//1.解析客户端用户提交的登陆数据
	var user models.User
	err := l.ParseForm(&user)
	if err != nil {
		fmt.Print(err.Error())
		l.Ctx.WriteString("抱歉，用户登录信息解析失败，请重试")
		return
	}
	//2.根据解析到的数据，执行数据库查询操作
    u, err := user.QueryUser()
	//3.判断数据库查询结果
	if err != nil{
		l.Ctx.WriteString("抱歉，用户登陆失败，请重试")
		return
	}
	//4.根据查询结果返回客户端相应的信息或者页面跳转
	l.Data["Phone"] = u.Phone
	l.TplName = "home.html"
}