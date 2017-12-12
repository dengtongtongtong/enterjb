package main

import (
	"enterbj/globals"
	_ "enterbj/routers"

	"github.com/astaxie/beego/config"

	"github.com/astaxie/beego"
)

func main() {
	// iniconf, err := config.NewConfig()

	if beego.BConfig.RunMode == "dev" {
		globals.EnterBJConfig, _ = config.NewConfig("ini", "conf/enterbj_alpha.conf")
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
