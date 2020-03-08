package routers

import (
	"github.com/astaxie/beego"
	"github.com/jadugnap/gwp/Chapter_11_Frameworks_Frameworks_Everywhere/mosaic-b/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
