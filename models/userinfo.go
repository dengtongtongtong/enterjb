package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type UserInfo struct {
	Id              int64  `orm:"auto" json:"-"`
	Uid             string `orm:"size(128);column(uid)" json:"uid"`
	Phone           string `orm:"size(128);column(phone)" json:"phone"`
	UidByBJJJ       string `orm:"size(128);column(uid_by_bjjj)" json:"uid_by_bjjj"`
	Licenseno       string `orm:"size(128);column(license_number)" json:"license_number"`
	Engineno        string `orm:"size(128);column(engine_number)" json:"engine_number"`
	Cartypecode     string `orm:"size(128);column(car_type_code)" json:"car_type_code"`
	Vehicletype     string `orm:"size(128);column(vehicle_type)" json:"vehicle_type"`
	Drivername      string `orm:"size(128);column(driver_name)" json:"driver_name"`
	Driverlicenseno string `orm:"size(128);column(driver_license_number)" json:"driver_license_number"`
	// Gpslon          string `orm:"size(128);column(uid_by_bjjj)"`
	// Gpslat          string `orm:""`
	// Imei            string
	Carid      string `orm:"size(128);column(carid)" json:"carid"`
	Carmodel   string `orm:"size(128);column(car_model)" json:"car_model"`
	Carregtime string `orm:"size(128);column(carreg_time)" json:"carreg_time"`
	// EnvGrade        string
	// ImageId         string
	// Sign            string
	// Platform        string
}

func (u *UserInfo) TableIndex() [][]string {
	return [][]string{
		[]string{"UidByBJJJ"},
		[]string{"Uid"},
	}
}

func init() {
	orm.RegisterModel(new(UserInfo))
}

// AddUserInfo insert a new UserInfo into database and returns
// last inserted Id on success.
func AddUserInfo(m *UserInfo) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUserInfoById retrieves UserInfo by Id. Returns error if
// Id doesn't exist
func GetUserInfoByUId(uid string) (v *UserInfo, err error) {
	o := orm.NewOrm()
	v = &UserInfo{Uid: uid}
	if err = o.QueryTable(new(UserInfo)).Filter("Uid", uid).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUserInfo retrieves all UserInfo matches certain condition. Returns empty list if
// no records exist
func GetAllUserInfo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UserInfo))
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

	var l []UserInfo
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

// UpdateUserInfo updates UserInfo by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserInfoById(m *UserInfo) (err error) {
	o := orm.NewOrm()
	v := UserInfo{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUserInfo deletes UserInfo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUserInfo(id int64) (err error) {
	o := orm.NewOrm()
	v := UserInfo{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&UserInfo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
