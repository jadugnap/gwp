package main

import (
	_ "github.com/jadugnap/gwp/Chapter_11_Frameworks_Frameworks_Everywhere/ws-b/docs"
	_ "github.com/jadugnap/gwp/Chapter_11_Frameworks_Frameworks_Everywhere/ws-b/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
