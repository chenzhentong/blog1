package routers

import (
	"beeblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/:page", &controllers.MainController{})
	beego.Router("/:page/:flag/:id", &controllers.MainController{})

	beego.Router("/index", &controllers.MainController{})
    beego.Router("/test",&controllers.TestController{})
    beego.Router("/login/:flag",&controllers.LoginController{},"get:Get;post:Post")
	beego.Router("/category",&controllers.CategoryController{},"get:Get;post:Post")
    beego.Router("/category/:flag/:id",&controllers.CategoryController{},"get:Get;post:Post")


	beego.Router("/topic/:flag",&controllers.TopicController{},"get:Get;post:Post")
	beego.Router("/topic/:flag/:id",&controllers.TopicController{},"get:Get;post:Post")
	beego.Router("/comment/:id",&controllers.CommentController{},"get:Get;post:Post")
    beego.Router("/attachment/:all",&controllers.AttachmentController{},"get:Get;post:Post")
	beego.Router("/attachment",&controllers.AttachmentController{},"get:Get;post:Post")
    beego.Router("/base",&controllers.BaseController{})

}
