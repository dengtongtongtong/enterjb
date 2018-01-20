package tasks

import (
	"crypto/md5"
	"fmt"

	"github.com/astaxie/beego/orm"

	"enterbj/models"

	"gocommon/datastructutils"

	"github.com/astaxie/beego/toolbox"
)

func CryptoToken(unordered map[string]interface{}, secret string) (signature string) {
	ordered, _ := datastructutils.SortMapByStringKey(unordered)
	var str string
	for _, x := range ordered {
		key := x.(map[string]interface{})["key"]
		str += key.(string)
		value := x.(map[string]interface{})["value"]
		str += value.(string)
	}
	fmt.Printf("secret %v\n", secret)
	fmt.Printf("str %v\n", str)
	strwithsecret := secret + str + secret
	bytesignature := md5.Sum([]byte(strwithsecret))
	signature = fmt.Sprintf("%X", bytesignature)
	return signature
}

func CryptoSign(imageid string) string {
	return ""
}

func GetToken(userid, timestamp string) string {
	var unorder = map[string]interface{}{"userid": userid, "appkey": "kkk", "deviceid": "ddd", "timestamp": timestamp}
	secret := timestamp
	return CryptoToken(unorder, secret)
}

func GetSign(userid, timestamp string) string {
	imageid := userid + timestamp
	return CryptoSign(imageid)
}

func EnterCarlist() error {
	return nil
}

func ApplyEnterBJ() error {
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
	tkapplyenterbj := toolbox.NewTask("tkapplyenterbj", "* * * * *", ApplyEnterBJ)
	toolbox.AddTask("tkapplyenterbj", tkapplyenterbj)
}
