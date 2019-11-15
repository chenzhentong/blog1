package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	//fmt.Println("get")
	flag := c.Ctx.Input.Param(":flag")
	switch flag {
	case "exit":
		{
			c.exit()
			break;
		}

	}

	c.Redirect("/", 301)
}
func (c *LoginController) exit() {
	username := c.GetString("username")
	password := c.GetString("password")
	//fmt.Println(user.Username)
	c.Ctx.SetCookie("username", username, -1)
	c.Ctx.SetCookie("password", password, -1)
	/*user.Password=c.Ctx.GetCookie("password")
	fmt.Println(user.Password)*/
}

func (c *LoginController) Post() {
	//fmt.Println("post")
	username := c.GetString("Username")
	password := c.GetString("Password")
	//fmt.Println("user:",username,password)
	result := models.CheckUser(username, password)
	//fmt.Println(user.Id)
	if !result {
		c.Ctx.WriteString("<html><script>alert('用户名或密码错误');window.history.back();</script><html>")
	} else {
		autoLogin := c.GetString("AutoLogin")
		//fmt.Println("autoLogin:", autoLogin)
		if autoLogin == "on" {

			c.Ctx.SetCookie("username", username, 1000)
			c.Ctx.SetCookie("password", password, 1000)

		} else {

			c.Ctx.SetCookie("username", username, 10)
			c.Ctx.SetCookie("password", password, 10)
		}
		/*c.SetSession("password",user.Password)
		password:=c.GetSession("password")
		fmt.Println("session:",password)*/
		/*username:=c.Ctx.GetCookie("username")
		fmt.Println(username,reflect.TypeOf(username))*/
		c.Redirect("/", 301)
		return
	}

}
