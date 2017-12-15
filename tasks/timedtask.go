package tasks

import (
	"fmt"

	"github.com/astaxie/beego/orm"

	"enterbj/models"

	"github.com/astaxie/beego/toolbox"
)

func enterCarlist() error {
	return nil
}

func applyEnterBJ() error {
	// all 依赖于qs默认limit限制 默认limit=1000
	fmt.Println("apply enterbj")
	o := orm.NewOrm()
	var userdocuments []models.Userdocument
	_, _ = o.QueryTable(new(models.Userdocument)).All(&userdocuments)
	for _, userdoc := range userdocuments {
		fmt.Println(userdoc)
	}
	// qs.Filter()
	fmt.Println(userdocuments)
	fmt.Println("finish apply enterbj")
	return nil
}

func init() {
	tkapplyenterbj := toolbox.NewTask("tkapplyenterbj", "* * * * *", applyEnterBJ)
	toolbox.AddTask("tkapplyenterbj", tkapplyenterbj)
}
