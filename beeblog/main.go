package main

import (
	"os"
	_ "beeblog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"beeblog/models"
	"beeblog/controllers"
)

func init(){
	//初始化数据库
	models.RegisterDB()	
}

func main() {
	//建表
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	
	//创建附件目录
	os.Mkdir("attachment", os.ModePerm)
	//作为静态文件处理(第一种方法 简单)
	beego.SetStaticPath("/attachment", "attachment")

	//作为一个单独的控制器处理
	beego.Router("/attachment/:all",  &controllers.AttachController{})

	beego.Run()
}

