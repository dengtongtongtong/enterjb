package main

import (
	"enterbj/globals"
	_ "enterbj/routers"
	"fmt"

	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego/config"

	_ "enterbj/models"
	"enterbj/tasks"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/test002?charset=utf8")
	// orm.RunSyncdb("default", true, true)
}

func main() {
	// toolbox.StartTask()
	t := tasks.GetToken("310F4BBB8DC340A3BA1E02A1A9B43A6C", "1516456816000")
	fmt.Printf("toke = %v\n", t)

	if beego.BConfig.RunMode == "dev" {
		globals.EnterBJConfig, _ = config.NewConfig("ini", "conf/enterbj_alpha.conf")
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
