package controllers

import (
	"beeblog/models"
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"strconv"
	"time"
)

type CommentController struct{
	beego.Controller
}
func (c *CommentController)Get(){

}
func (c *CommentController)Post(){
	c.addComment()
}
func (c *CommentController)addComment(){
	topicId,err:=c.GetInt(":id")
	fmt.Println("topicId:",topicId)
	if err!=nil{
		log.Fatal(err)
	}else{
		topic:=models.FindTopicById(topicId)
		var comment models.Comment
		comment.Topic=&topic
		comment.Content=c.GetString("Content")
		comment.Name=c.GetString("Name")
		comment.CreatedTime=time.Now()
		comment.Website=c.GetString("Website")
		comment.Email=c.GetString("Email")
		result:=models.AddComment(&comment)
		if result{
			url:="/topic/detail/"+strconv.Itoa(topicId)
			fmt.Println("url:",url)
			c.Redirect(url,302)
		}else{
			c.Ctx.WriteString("<html><script>alert('添加失败!');window.history.back();</script><html>")
		}
	}
}
