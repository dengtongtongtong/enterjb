package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ApplyMent struct {
	Id                int64  `orm:"auto" json:"-"`
	Uid               string `orm:"size(128);column(uid)" json:"uid"`
	UidByBJJJ         string `orm:"size(128);column(uid_by_bjjj)" json:"uid_by_bjjj"`
	Appsource         string `orm:"size(128);column(app_source)" json:"app_source"`
	Hiddentime        string `orm:"size(128);column(hidden_time)" json:"hidden_time"`
	Inbjentrancecode1 string `orm:"size(128);column(inbj_entrance_code1)" json:"inbj_entrance_code1"`
	Inbjentrancecode  string `orm:"size(128);column(inbj_entrance_code)" json:"inbj_entrance_code"`
	Inbjduration      string `orm:"size(128);column(inbj_duration)" json:"inbj_duration"`
	Inbjtime          string `orm:"size(128);column(inbj_time)" json:"inbj_time"`
	Deviceid          string `orm:"size(128);column(deviceid)" json:"deviceid"`
	Timestamp         string `orm:"size(128);column(apply_time_stamp)" json:"apply_time_stamp"`
	Licenseno         string `orm:"size(128);column(license_number)" json:"license_number"`
	Engineno          string `orm:"size(128);column(engine_number)" json:"engine_number"`
	Cartypecode       string `orm:"size(128);column(car_type_code)" json:"car_type_code"`
	Vehicletype       string `orm:"size(128);column(vehicle_type)" json:"vehicle_type"`
	Drivername        string `orm:"size(128);column(driver_name)" json:"driver_name"`
	Driverlicenseno   string `orm:"size(128);column(driver_license_number)" json:"driver_license_number"`
	Gpslon            string `orm:"size(128);column(gps_lon)" json:"gps_lon"`
	Gpslat            string `orm:"size(128);column(gps_lat)" json:"gps_lat"`
	Imei              string `orm:"size(128);column(imei)" json:"imei"`
	Carid             string `orm:"size(128);column(carid)" json:"carid"`
	Carmodel          string `orm:"size(128);column(car_model)" json:"car_model"`
	Carregtime        string `orm:"size(128);column(carreg_time)" json:"carreg_time"`
	EnvGrade          string `orm:"size(128);column(env_grade)" json:"env_grade"`
	ImageId           string `orm:"size(128);column(imageid)" json:"imageid"`
	Sign              string `orm:"size(128);column(sign)" json:"sign"`
	Platform          string `orm:"size(128);column(platform)" json:"platform"`

	Appkey       string `orm:"size(128);column(appkey)" json:"appkey"`
	Token        string `orm:"size(128);column(token)" json:"token"`
	Drivingphoto string `orm:"size(128);column(driver_photo)" json:"driver_photo"`
	Carphoto     string `orm:"size(128);column(car_photo)" json:"car_photo"`
	Personphoto  string `orm:"size(128);column(person_photo)" json:"person_photo"`
	Phoneno      string `orm:"size(128);column(phone_number)" json:"phone_number"`
	Imsi         string `orm:"size(128);column(imsi)" json:"imsi"`
	Code         string `orm:"size(128);column(code)" json:"code"`

	CreateTime int `orm:"column(create_time)"`
	UpdateTime int `orm:"column(update_time)"`
}

func init() {
	orm.RegisterModel(new(ApplyMent))
}

// AddApplyMent insert a new ApplyMent into database and returns
// last inserted Id on success.
func AddApplyMent(m *ApplyMent) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetApplyMentById retrieves ApplyMent by Id. Returns error if
// Id doesn't exist
func GetApplyMentById(id int64) (v *ApplyMent, err error) {
	o := orm.NewOrm()
	v = &ApplyMent{Id: id}
	if err = o.QueryTable(new(ApplyMent)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllApplyMent retrieves all ApplyMent matches certain condition. Returns empty list if
// no records exist
func GetAllApplyMent(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ApplyMent))
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

	var l []ApplyMent
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

// UpdateApplyMent updates ApplyMent by Id and returns error if
// the record to be updated doesn't exist
func UpdateApplyMentById(m *ApplyMent) (err error) {
	o := orm.NewOrm()
	v := ApplyMent{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteApplyMent deletes ApplyMent by Id and returns error if
// the record to be deleted doesn't exist
func DeleteApplyMent(id int64) (err error) {
	o := orm.NewOrm()
	v := ApplyMent{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ApplyMent{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
