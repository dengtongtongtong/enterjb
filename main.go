package main

import (
	"enterbj/globals"
	_ "enterbj/routers"
	"fmt"

	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego/config"

	"enterbj/models"
	"enterbj/tasks"

	"gocommon/stringutils"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func initMockData() {
	o := orm.NewOrm()
	uid := stringutils.Uuid4()
	userInfo := models.UserInfo{
		Uid:             uid,
		Phone:           "18610486211",
		UidByBJJJ:       "310F4BBB8DC340A3BA1E02A1A9B43A6C",
		Engineno:        "7059735",
		Cartypecode:     "02",
		Driverlicenseno: "50010619831028081X",
		Carid:           "18484313",
	}
	userStatus := models.UserStatus{
		Uid:        uid,
		Status:     0,
		ExpireTime: 1616453292,
	}
	applyStaus := models.ApplyStatus{
		Uid: uid,
	}
	o.Insert(&userInfo)
	o.Insert(&userStatus)
	o.Insert(&applyStaus)
	fmt.Println("auto gened uid", userInfo.Id)
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/test003?charset=utf8")
	orm.RunSyncdb("default", true, true)
	initMockData()
}

func main() {
	// toolbox.StartTask()
	// t := tasks.GetToken("310F4BBB8DC340A3BA1E02A1A9B43A6C", "1516461724000")
	// fmt.Printf("toke = %v\n", t)
	imageid := tasks.GetMyImageId()
	fmt.Printf("imageid=%v\n", imageid)

	if beego.BConfig.RunMode == "dev" {
		globals.EnterBJConfig, _ = config.NewConfig("ini", "conf/enterbj_alpha.conf")
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
