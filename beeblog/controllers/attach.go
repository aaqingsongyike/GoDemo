package controllers

import(
	"net/url"
	"os"
	"io"
	"github.com/astaxie/beego"
)

type AttachController struct {
	beego.Controller
}

func (this *AttachController) Get() {
	filePath, err := url.QueryUnescape(this.Ctx.Request.RequestURI[1:])	//获取文件路径
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
	f, err := os.Open(filePath)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
	defer f.Close()

	_, err = io.Copy(this.Ctx.ResponseWriter, f)	//参数（输出流和输入流）
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
}