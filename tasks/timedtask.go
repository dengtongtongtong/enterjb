package tasks

import (
	"crypto/md5"
	"fmt"

	"gocommon/datastructutils"

	"github.com/astaxie/beego/toolbox"
)

type FinalSubmitData struct {
	Appsource         string
	Hiddentime        string
	Inbjentrancecode1 string
	Inbjentrancecode  string
	Inbjduration      string
	Inbjtime          string
	Deviceid          string
	Timestamp         string
	Userid            string
	Licenseno         string
	Engineno          string
	Cartypecode       string
	Vehicletype       string
	Drivername        string
	Driverlicenseno   string
	Gpslon            string
	Gpslat            string
	Imei              string
	Carid             string
	Carmodel          string
	Carregtime        string
	EnvGrade          string
	ImageId           string
	Sign              string
	Platform          string

	Appkey       string
	Token        string
	Drivingphoto string
	Carphoto     string
	Driverphoto  string
	Personphoto  string
	Phoneno      string
	Imsi         string
	Code         string
}

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

func GetFinalSubmitImageId(data FinalSubmitData) (imageId string) {
	imageId = data.Inbjentrancecode + data.Inbjduration + data.Inbjtime + data.Userid + data.Engineno + data.Cartypecode + data.Driverlicenseno + data.Carid + data.Timestamp
	return imageId
}

func GetFinalSubmitImageIdWithSecret(data FinalSubmitData) (imageId string) {
	appkey := "0791682354"
	imageId = appkey + GetFinalSubmitImageId(data) + appkey
	return imageId
}

func GetMyImageId() (imageId string) {
	var data = FinalSubmitData{}
	data.Inbjduration = "2"
	data.Inbjtime = "2018-1-21"
	data.Userid = "310F4BBB8DC340A3BA1E02A1A9B43A6C"
	data.Engineno = "7059735"
	data.Cartypecode = "02"
	data.Driverlicenseno = "50010619831028081X"
	data.Carid = "18484313"
	data.Timestamp = "2018-01-20 22:26:50"
	imageId = GetFinalSubmitImageIdWithSecret(data)
	// imageId = GetFinalSubmitImageId(data)
	return imageId
}

func ApplyEnterBJ() error {
	// all 依赖于qs默认limit限制 默认limit=1000
	// fmt.Println("apply enterbj")
	// o := orm.NewOrm()
	// var userdocuments []models.Userdocument
	// _, _ = o.QueryTable(new(models.Userdocument)).All(&userdocuments)
	// for _, userdoc := range userdocuments {
	// 	fmt.Println(userdoc)
	// }
	// // qs.Filter()
	// fmt.Println(userdocuments)
	// fmt.Println("finish apply enterbj")
	return nil
}

func init() {
	tkapplyenterbj := toolbox.NewTask("tkapplyenterbj", "* * * * *", ApplyEnterBJ)
	toolbox.AddTask("tkapplyenterbj", tkapplyenterbj)
}
