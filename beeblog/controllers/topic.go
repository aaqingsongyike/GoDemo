package controllers

import(
	"strings"
	"path"
	"beeblog/models"
	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsTopic"] = true
	this.TplName = "topic.html"
	
	topics, err := models.GetAllTopics("", "", false)
	if err != nil {
		beego.Error(err)
	}else {
		this.Data["Topics"] = topics
	}
	
}

func (this *TopicController) Post() {
	//身份验证
	if !checkAccount(this.Ctx){	//身份不通过
		this.Redirect("/login", 302)
		return
	}
	//表单解析
	tid := this.Input().Get("tid")
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	category := this.Input().Get("category")
	label := this.Input().Get("label")

	//获取附件
	_, fh, err := this.GetFile("attachment")
	if err != nil {
		beego.Error(err)
	}

	var attachment string	//保存filename
	if fh != nil {
		//保存附件
		attachment = fh.Filename
		beego.Info(attachment)	//打印文件名
		err = this.SaveToFile("attachment", path.Join("attachment", attachment))
		//假设 filename: tmp.go
		//存入后 attachment/tmp.go
		if err != nil {
			beego.Error(err)
		}
	}


	if len(tid) == 0{
		err = models.AddTopic(title, category, label, content, attachment)
	} else {
		err = models.ModifyTopic(tid, title, category, label, content, attachment)
	}
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic", 302)
}

func (this *TopicController) Add() { 	//添加文章
	this.TplName = "topic_add.html"
}

func (this *TopicController) View() {	//浏览文章
	this.TplName = "topic_view.html"

	tid := this.Ctx.Input.Param("0")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Topic"] = topic
	this.Data["Labels"] = strings.Split(topic.Labels, " ")

	replies, err := models.GetAllReplies(tid)
	if err != nil {
		beego.Error(err)
		return
	}

	this.Data["Replies"] = replies
	this.Data["IsLogin"] = checkAccount(this.Ctx)
}

//修改文章
func (this *TopicController) Modify() {
	this.TplName = "topic_modify.html"
	tid := this.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	this.Data["Topic"] = topic
	this.Data["Tid"] = tid	//由于是Post操作因此需再传一次tid

}

//删除操作
func (this *TopicController) Delete() {
	//身份验证
	if !checkAccount(this.Ctx){
		this.Redirect("/login", 302)
		return
	}

	err := models.DeleteModify(this.Input().Get("tid"))
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/", 302)
}