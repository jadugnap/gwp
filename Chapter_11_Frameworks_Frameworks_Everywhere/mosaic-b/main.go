package main

import (
	"github.com/astaxie/beego"
	"github.com/jadugnap/gwp/Chapter_11_Frameworks_Frameworks_Everywhere/mosaic-b/mosaic"
	_ "github.com/jadugnap/gwp/Chapter_11_Frameworks_Frameworks_Everywhere/mosaic-b/routers"
)

func main() {
	go mosaic.TilesDB()
	beego.Run()
}
