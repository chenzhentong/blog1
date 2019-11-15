package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
	"log"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	username := c.Ctx.GetCookie("username")
	//fmt.Println("id:", id)
	if username != "" {
		c.Data["isLogin"] = true
		//fmt.Println("cookie有效")
	} else {
		//fmt.Println(false)
		c.Data["isLogin"] = false
		//fmt.Println("session无效，cookie无效")
	}

	pageSize:=3
	pageNow,err:=c.GetInt(":page")
	if err!=nil{
		pageNow=1
	}
	var topics []models.Topic
	var pageCount int
	flag:=c.GetString(":flag")
	if flag=="category"{
		categroyId,err:=c.GetInt(":id")
		if err==nil{
			topics,err=models.ReadCategoryTopicByPage(categroyId,pageNow,pageSize)
			if err!=nil{
				log.Fatal(err)
			}
			pageCount=models.GetCategoryTopicPageCount(categroyId,pageSize)
			c.Data["flag"]="category"
			c.Data["id"]=categroyId
		}
	}else if flag=="month"{
		month:=c.GetString(":id")
		if month!=""{
			topics,err=models.ReadCategoryTopicByMonth(month,pageNow,pageSize)
			if err!=nil{
				log.Fatal(err)
			}
			pageCount=models.GetMonthTopicPageCount(month,pageSize)
			c.Data["flag"]="month"
			c.Data["id"]=month
		}

	}else{
		topics,err=models.ReadTopicByPage(pageNow,pageSize)
		if err!=nil{
			log.Fatal(err)
		}
		pageCount=models.GetTopicPageCount(pageSize)
		c.Data["flag"]=""
	}
	topicRecords:=models.ReadAllTopicsByTime()
	categories:=models.ReadAllTopicByCategory()
	c.Data["categories"]=categories
	c.Data["pageCount"]=pageCount
	c.Data["pageNow"]=pageNow
	c.Data["topics"]=topics
	c.TplName = "home.html"
	c.Data["isHome"] = true
	c.Data["topicRecords"]=topicRecords
}
