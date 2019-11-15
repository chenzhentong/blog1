package main

import (
	"beeblog/models"
	_ "beeblog/routers"
	"beeblog/views"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"os"
)
func init(){
	models.RegisterDB()
}

func main() {
	beego.BConfig.WebConfig.Session.SessionOn=true
	orm.Debug=true
	//orm.RunSyncdb("default",false,true)
    beego.AddFuncMap("Add",views.Add)
	beego.AddFuncMap("IsArrayNull",views.IsArrayNull)
	beego.AddFuncMap("IsStringNull",views.IsStringNull)

	beego.AddFuncMap("TimeParseString",views.TimeParseString)

	beego.AddFuncMap("JudgeTopicType",views.JudgeTopicType)

	beego.AddFuncMap("JudgeCategory",views.JudgeCategory)

	beego.AddFuncMap("JudgeMonth",views.JudgeMonth)
	os.Mkdir("attachment",os.ModePerm)
	beego.Run()

}

