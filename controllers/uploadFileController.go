package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type UploadFileController struct {
	beego.Controller
}

func (u *UploadFileController) post() {
	//1.解析用户上传的数据及文件内容
	//用户上传的自定义的标题
	title := u.Ctx.Request.PostFormValue("upload_title")//用户输入的标题
	//用户上传的文件
	file, header, err := u.GetFile("wangyuanjing")
	if err != nil{//解析客户端提交的文件出现的错误
		u.Ctx.WriteString("抱歉，文件解析失败，请重试")
		return
	}
	defer file.Close()
	fmt.Println("自定义的标题：",title)
	//获得到了上传的文件
	fmt.Println("上传的文件名称：",header.Filename)
	fileNameSlice := strings.Split(header.Filename,".")
	fileType := fileNameSlice[1]
	fmt.Println(fileNameSlice)
	fmt.Println(":",fileType)
	isJpg := strings.HasSuffix(header.Filename,".")
	isPng := strings.HasSuffix(header.Filename,".")
	if !isJpg && !isPng{//文件不支持
		u.Ctx.WriteString("抱歉，文件类型不符合，请上传符合格式的文件")
		return
	}
	config := beego.AppConfig
	fileSize,err := config.Int64("file_size")
	if header.Size / 1024 > fileSize{
		u.Ctx.WriteString("抱歉，文件的大小超出范围，请上传符合要求的文件")
		return
	}
	fmt.Println("上传的文件的大小：",header.Size)//字节大小

	//perm:permission 权限
	//权限的组成：a+b+c
	  //a: 文件所有者对文件的操作权限,读   写   执行
	  //b:文件所有者所在组的用户的操作权限，读   写   执行
	  //c: 其他用户的操作权限，读   写   执行
	//os.Mkdir()
	//fromFile:文"static/upload"件
	//toFile:要保存的文件
	//u.SaveToFile()





	fmt.Println("上传的文件：",file)
	u.Ctx.WriteString("已获取到上传的文件")
}