package controllers

import (
	"beeblog/models"
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"path"
	"time"
)

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get() {
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
	flag := c.Ctx.Input.Param(":flag")
	switch flag {
	case "add":{
		c.goAddTopic()
		break
	}
	case "edit":{
		c.goEditTopic()
		break
	}
	case "delete":{
		c.deleteTopic()
		break
	}
	case "detail":{
		c.viewDetailTopic()
		break
	}
	case "list":{
		c.viewTopicList()
		break
	}

	}

	c.Data["isTopic"] = true
}
func (c *TopicController) goEditTopic(){
	username := c.Ctx.GetCookie("username")
	if username == "" {
		c.Ctx.WriteString("<html><script>alert('请登录后再操作!');window.history.back();</script><html>")
		return
	}
	id, err := c.GetInt(":id")
	if err != nil {
		log.Fatal(err)
	}
	var categories []models.Category
	models.ReadAllCategories(&categories)
	topic := models.FindTopicById(id)
	c.Data["topic"] = topic
	c.Data["categories"] = categories
	fmt.Println("categories", categories)
	c.Data["isManage"] = true
	fmt.Println("isManage:", c.Data["isManage"])
	c.TplName = "topic_view.html"
}
func (c *TopicController) goAddTopic(){
	username := c.Ctx.GetCookie("username")
	if username == "" {
		c.Ctx.WriteString("<html><script>alert('请登录后再操作!');window.history.back();</script><html>")
		return
	}
	var categories []models.Category
	models.ReadAllCategories(&categories)
	c.Data["categories"] = categories
	c.TplName = "topic_add.html"
	c.Data["isAdd"] = true
	fmt.Println("isAdd:", c.Data["isAdd"])

}
func (c *TopicController)deleteTopic(){
	username := c.Ctx.GetCookie("username")
	//fmt.Println("id:", id)
	if username == "" {
		c.Ctx.WriteString("<html><script>alert('请登录后再操作!');window.history.back();</script><html>")
		return
	}
	id, err := c.GetInt(":id")
	if err != nil {
		log.Fatal(err)
	}
	topic := models.FindTopicById(id)
	result := models.DeleteTopic(topic)
	if !result {
		c.Ctx.WriteString("<html><script>alert('删除失败!');window.history.back();</script><html>")
		//panic(err)
	} else {
		//c.Ctx.WriteString("<html><script>alert('添加成功');</script><html>")
		c.Data["isManage"] = true
		c.Redirect("/topic/list", 302)
	}
}
func (c *TopicController) viewDetailTopic(){
	id, err := c.GetInt(":id")
	if err != nil {
		log.Fatal(err)
	}
	//	fmt.Println("topicId:",id)
	comments, err := models.ReadTopicComments(id)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("comments:",comments)
	/*for _,comment:=range comments{
		fmt.Println(comment.Id,comment.Name,comment.CreatedTime,comment.Website,comment.Email)
	}*/
	topic := models.FindTopicById(id)
	c.TplName = "topic_detail.html"
	c.Data["topic"] = topic
	c.Data["comments"] = comments
}
func (c *TopicController) viewTopicList(){
	var topics []models.Topic
	topics, err := models.ReadAlltopic(false, topics)
	if err != nil {
		log.Fatal(err)
	}
	c.Data["isManage"] = true
	fmt.Println("isManage:", c.Data["isManage"])
	c.Data["topics"] = topics
	c.TplName = "topic.html"
}
func (c *TopicController) Post() {
	flag := c.Ctx.Input.Param(":flag")
	fmt.Println("flag:", flag)
	switch flag {
	case "add":
		{
			c.addTopic()
			break

		}
	case "update":
		{
			c.updateTopic()
			break
		}

	}

}
func (c *TopicController) addTopic() {
	var topic models.Topic
	topic.Title = c.GetString("Title")
	topic.Content = c.GetString("Content")

	categoryId, err := c.GetInt("Category")

	if err != nil {
		log.Fatal(err)
	}
	category := models.FindCategoryById(categoryId)
	loc, _ := time.LoadLocation("Asia/Chongqing")
	topic.Updated = time.Now().In(loc)
	topic.Created = time.Now().In(loc)
	topic.ReplyTime = time.Now().In(loc)
	topic.Category = &category
	var attachment string
	_, fileHander, err := c.GetFile("attachment")
	if fileHander != nil {
		//保存附件
		attachment = fileHander.Filename
		fmt.Println(attachment)
		c.SaveToFile("attachment", path.Join("attachment", attachment))
		if err != nil {
			fmt.Println(err)
		}
	}
	topic.Attachment = attachment
	if err != nil {
		fmt.Println(err)
	}
	result := models.AddTopic(&topic)
	if !result {
		c.Ctx.WriteString("<html><script>alert('添加失败!');window.history.back();</script><html>")
		//panic(err)
	} else {
		//c.Ctx.WriteString("<html><script>alert('添加成功');</script><html>")
		c.Redirect("/topic/list", 302)
	}
}
func (c *TopicController) updateTopic() {
	//fmt.Println("flag:", flag)

	username := c.Ctx.GetCookie("username")
	if username == "" {
		c.Ctx.WriteString("<html><script>alert('请登录后再操作!');window.history.back();</script><html>")
		return
	}
	id, err := c.GetInt(":id")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("id:", id)
	topic := models.FindTopicById(id)
	topic.Title = c.GetString("Title")
	topic.Content = c.GetString("Content")
	loc, _ := time.LoadLocation("Asia/Chongqing")

	topic.Updated = time.Now().In(loc)
	categoryId, err := c.GetInt("Category")
	//fmt.Println("categoryId:",categoryId)
	if err != nil {
		log.Fatal(err)
	}
	category := models.FindCategoryById(categoryId)

	topic.Category = &category
	topic.Views++
	result := models.UpdateTopic(&topic)
	if !result {

		c.Ctx.WriteString("<html><script>alert('更新失败!');window.history.back();</script><html>")
		//panic(err)
	} else {
		//c.Ctx.WriteString("<html><script>alert('添加成功');</script><html>")
		c.Redirect("/topic/list", 302)
	}
}
