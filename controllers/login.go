package controllers

import (
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	isExit := c.Input().Get("exit") == "true"
	if isExit {
		c.Ctx.SetCookie("uname", "", -1, "/") //设置为-1，cookie立即过期
		c.Ctx.SetCookie("password", "", -1, "/")
		c.Redirect("/", 301)
		return
	}
	c.TplNames = "login.html"
}

func (c *LoginController) Post() {

	/*
		//在浏览器中打印接收到的收据
		c.Ctx.WriteString(fmt.Sprint(c.Input()))
		return*/
	uname := c.Input().Get("uname")
	password := c.Input().Get("password")
	autoLogin := c.Input().Get("autoLogin") == "on"

	if beego.AppConfig.String("uname") == uname && beego.AppConfig.String("password") == password {
		maxAge := 0 //cookie存活时间
		if autoLogin {
			maxAge = 1<<31 - 1
		}
		c.Ctx.SetCookie("uname", uname, maxAge, "/")
		c.Ctx.SetCookie("password", password, maxAge, "/")
	}

	//重定向
	c.Redirect("/", 301)
	return //加上return ，防止重复渲染
}

//检查cookie,判断帐号是否登录
func checkAccount(ctx *context.Context) bool {

	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}
	uname := ck.Value
	ck, err = ctx.Request.Cookie("password")

	if err != nil {
		return false
	}
	password := ck.Value
	return beego.AppConfig.String("uname") == uname && beego.AppConfig.String("password") == password

}
