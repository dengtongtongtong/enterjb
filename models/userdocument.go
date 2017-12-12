package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Userdocument struct {
	Id int64 `orm:"auto"`
	// 用户提交的文档ID
	Did string `orm:"size(128);column(documentid)"`
	// 用户ID
	Uid string `orm:"size(128);column(uid)"`
	// 汽车品牌型号
	VehicleBrandType string `orm:"size(128);column(vehicle_brand_type)"`
	// 车辆注册日期
	CarRegisterTime int `orm:"column(car_register_time)"`
	// 汽车(品牌)类型
	Cartype string `orm:"size(128);column(car_type)"`
	// 机动车号牌
	CarID string `orm:"size(128);column(carid)"`
	// 发动机号
	EngineID string `orm:"size(128);column(engineid)"`
	// 行驶证照片
	DrivingLisencePhotoPath string `orm:"size(128);column(driving_lisence_photo_path)"`
	// 车辆正面照片
	CarFrontPhotoPath string `orm:"size(128);column(car_front_photo_path)"`
	// 驾驶员姓名
	DriverName string `orm:"size(128);column(driver_name)"`
	// 驾照号
	DrivingLisence string `orm:"size(128);column(driving_lisence)"`
}

func init() {
	orm.RegisterModel(new(Userdocument))
}

// AddUserdocument insert a new Userdocument into database and returns
// last inserted Id on success.
func AddUserdocument(m *Userdocument) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUserdocumentById retrieves Userdocument by Id. Returns error if
// Id doesn't exist
func GetUserdocumentById(id int64) (v *Userdocument, err error) {
	o := orm.NewOrm()
	v = &Userdocument{Id: id}
	if err = o.QueryTable(new(Userdocument)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUserdocument retrieves all Userdocument matches certain condition. Returns empty list if
// no records exist
func GetAllUserdocument(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Userdocument))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Userdocument
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateUserdocument updates Userdocument by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserdocumentById(m *Userdocument) (err error) {
	o := orm.NewOrm()
	v := Userdocument{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUserdocument deletes Userdocument by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUserdocument(id int64) (err error) {
	o := orm.NewOrm()
	v := Userdocument{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Userdocument{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func UpdateOrInsertUserdocument(m *Userdocument) (err error) {
	return
}
