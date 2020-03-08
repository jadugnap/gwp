package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/jadugnap/gwp/Chapter_11_Frameworks_Frameworks_Everywhere/ws-b/controllers:PostController"] = append(beego.GlobalControllerRouter["github.com/jadugnap/gwp/Chapter_11_Frameworks_Frameworks_Everywhere/ws-b/controllers:PostController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/jadugnap/gwp/Chapter_11_Frameworks_Frameworks_Everywhere/ws-b/controllers:PostController"] = append(beego.GlobalControllerRouter["github.com/jadugnap/gwp/Chapter_11_Frameworks_Frameworks_Everywhere/ws-b/controllers:PostController"],
		beego.ControllerComments{
			"Get",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/jadugnap/gwp/Chapter_11_Frameworks_Frameworks_Everywhere/ws-b/controllers:PostController"] = append(beego.GlobalControllerRouter["github.com/jadugnap/gwp/Chapter_11_Frameworks_Frameworks_Everywhere/ws-b/controllers:PostController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/jadugnap/gwp/Chapter_11_Frameworks_Frameworks_Everywhere/ws-b/controllers:PostController"] = append(beego.GlobalControllerRouter["github.com/jadugnap/gwp/Chapter_11_Frameworks_Frameworks_Everywhere/ws-b/controllers:PostController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
