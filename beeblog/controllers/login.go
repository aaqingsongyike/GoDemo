package controllers

import(
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct{	//创建LoginController
	beego.Controller	//嵌入beego的controller结构
}

func (this *LoginController) Get() {
	//退出
	isExit := this.Input().Get("exit") == "true"
	if isExit{
		this.Ctx.SetCookie("uname", "" , -1, "/")
		this.Ctx.SetCookie("pwd", "", -1, "/")
		this.Redirect("/", 301)
		return
	}
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	uname := this.Input().Get("uname")	//获取
	pwd := this.Input().Get("pwd")
	autoLogin := this.Input().Get("autoLogin") == "on"

	//判断
	if beego.AppConfig.String("uname") == uname  && beego.AppConfig.String("pwd") == pwd {
		//Cookie操作
		maxAge := 0
		if autoLogin {
			maxAge = 100
		}
		this.Ctx.SetCookie("uname", uname, maxAge, "/")	//存Cookie
		this.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	}

	//重定向
	this.Redirect("/", 301)
	return
}

//提取Cookie再次验证
func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}
	uname := ck.Value

	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := ck.Value

	return beego.AppConfig.String("uname") == uname && beego.AppConfig.String("pwd") == pwd
	
}
