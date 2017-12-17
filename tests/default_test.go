package test

import (
	// _ "enterbj/routers"

	"testing"

	"enterbj/tasks"
	// . "github.com/smartystreets/goconvey/convey"
)

// func init() {
// 	_, file, _, _ := runtime.Caller(1)
// 	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
// 	beego.TestBeegoInit(apppath)
// }

// // TestGet is a sample to run an endpoint test
// func TestGet(t *testing.T) {
// 	r, _ := http.NewRequest("GET", "/v1/object", nil)
// 	w := httptest.NewRecorder()
// 	beego.BeeApp.Handlers.ServeHTTP(w, r)

// 	beego.Trace("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())

// 	Convey("Subject: Test Station Endpoint\n", t, func() {
// 	        Convey("Status Code Should Be 200", func() {
// 	                So(w.Code, ShouldEqual, 200)
// 	        })
// 	        Convey("The Result Should Not Be Empty", func() {
// 	                So(w.Body.Len(), ShouldBeGreaterThan, 0)
// 	        })
// 	})
// }

func TestSign(t *testing.T) {
	var unorder = map[string]interface{}{"userid": "310F4BBB8DC340A3BA1E02A1A9B43A6C", "appkey": "kkk", "deviceid": "ddd", "timestamp": "1513290532000"}
	secret := "1513290532000"
	signature := tasks.Sign(unorder, secret)
	t.Log("signature", signature)
}
