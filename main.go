package main

import (
	"enterbj/globals"
	_ "enterbj/routers"

	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego/config"

	_ "enterbj/models"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/test002?charset=utf8")
	// orm.RunSyncdb("default", true, true)
}

func main() {
	// iniconf, err := config.NewConfig()

	if beego.BConfig.RunMode == "dev" {
		globals.EnterBJConfig, _ = config.NewConfig("ini", "conf/enterbj_alpha.conf")
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
