package controllers

import (
	"encoding/json"
	"enterbj/globals"
	"enterbj/models"
	"errors"
	"gocommon/fileutils"
	"gocommon/webutils"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

//  UserdocumentController operations for Userdocument
type UserdocumentController struct {
	beego.Controller
}

// URLMapping ...
func (c *UserdocumentController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

func getUserinfo(c *UserdocumentController) (userinfo *models.Userinfo, err error) {
	uid := c.Ctx.Input.Param(":uid")
	userinfo, err = models.GetUserinfoByUId(uid)
	return userinfo, err
}

func convertParams2Model(c *UserdocumentController) (userdoc models.Userdocument) {
	userdoc.Uid = c.Ctx.Input.Param(":uid")
	userdoc.VehicleBrandType = c.GetString("vehicle_brand_type")
	userdoc.CarRegisterTime, _ = c.GetInt("car_register_time")
	userdoc.Cartype = c.GetString("car_type")
	userdoc.CarID = c.GetString("carid")
	userdoc.EngineID = c.GetString("engineid")
	userdoc.DriverName = c.GetString("driver_name")
	userdoc.DrivingLisence = c.GetString("driving_lisence")

	var err error
	var filepath string
	filedir := globals.EnterBJConfig.String("BASIC::USER_DOCUMENT_BASEDIR")
	filepath = fileutils.MakeTempFilePath(filedir)
	err = c.SaveToFile("driving_lisence_photo", filepath)
	if err == nil {
		userdoc.DrivingLisencePhotoPath = filepath
	}
	filepath = fileutils.MakeTempFilePath(filedir)
	err = c.SaveToFile("car_front_photo", filepath)
	if err == nil {
		userdoc.CarFrontPhotoPath = filepath
	}
	filepath = fileutils.MakeTempFilePath(filedir)
	if err == nil {
		userdoc.IdcardPhotoPath = filepath
	}
	return userdoc
}

// Post ...
// @Title Pos
// @Description 提交/更新档案
// @Param   uid   path   string   true   "用户ID"
// @Param   vehicle_brand_type   formdata   string   false   "汽车品牌型号"
// @Param   car_register_time    formdata   string   false   "车辆注册日期"
// @Param   car_type   formdata   string   false       "号牌类型"
// @Param   carid   formdata   string  false   "机动车号牌"
// @Param   engineid   formdata   string  false   "发动机号"
// @Param   driving_lisence_photo   formdata   sting   false   "行驶证照片"
// @Param   car_front_photo   formdata   string   false   "车辆正面照片"
// @Param   idcard_photo   formdata   string   false   "车辆正面照片"
// @Param   driver_name   formdata   string   false   "驾驶员姓名"
// @Param   driving_lisence   formdata   string   false   "驾照号"
// @Success 200 {int} models.Userdocument
// @Failure 404 body is empty
// @router /:uid [post]
func (c *UserdocumentController) Post() {
	// var err error
	// _, err = getUserinfo(c)
	// if err != nil {
	// 	c.Data["json"] = webutils.Success(nil)
	// 	return
	// }
	userdocmodel := convertParams2Model(c)
	_ = models.UpdateOrInsertUserdocument(&userdocmodel)
	c.Data["json"] = webutils.Success(nil)
	c.ServeJSON()

	//baseidr := "/Users/dengtongtong/workspace/golangworkspace/src/enterbj"
	//filepath := fileutils.MakeTempFilePath(baseidr)
	//fmt.Println("filepath", filepath)

	// imgfile, _, _ := c.GetFile("imgdata")

	//c.SaveToFile("imgdata", filepath)
	// c.GetFile
	// var v models.Userdocument
	// fmt.Println(v)
	// fmt.Println(c.Ctx.Input.Data())
	// json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	// fmt.Println(v)
	// if _, err := models.AddUserdocument(&v); err == nil {
	// c.Ctx.Output.SetStatus(201)
	// c.Data["json"] = v
	// } else {
	// c.Data["json"] = err.Error()
	// }
	// c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Userdocument by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Userdocument
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UserdocumentController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetUserdocumentById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Userdocument
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Userdocument
// @Failure 403
// @router / [get]
func (c *UserdocumentController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllUserdocument(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Userdocument
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Userdocument	true		"body for Userdocument content"
// @Success 200 {object} models.Userdocument
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UserdocumentController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Userdocument{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateUserdocumentById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Userdocument
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UserdocumentController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteUserdocument(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
