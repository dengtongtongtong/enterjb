package service

import (
	"encoding/json"
	"enterbj/models"
	"fmt"
	"gocommon/timeutils"

	"github.com/astaxie/beego/orm"
)

type ApplyMentService struct {
}

func (s *ApplyMentService) getFinalSubmitImageId(data *models.ApplyMent) (imageId string) {
	imageId = data.Inbjentrancecode + data.Inbjduration + data.Inbjtime + data.UidByBJJJ + data.Engineno + data.Cartypecode + data.Driverlicenseno + data.Carid + data.Timestamp
	return imageId
}

func (s *ApplyMentService) generateApplyment(o *orm.Ormer, userInfo *models.UserInfo) (applyment models.ApplyMent) {
	jsonUserInfo, _ := json.Marshal(userInfo)
	json.Unmarshal(jsonUserInfo, &applyment)
	applyment.Inbjduration = "02"
	// applyment.
	return applyment
}

func (s *ApplyMentService) ObtainUnsignedApplyment() (applyMent string) {
	timeCurrent := timeutils.CurrentTimestamp()
	o := orm.NewOrm()
	o.Begin()
	var userInfo models.UserInfo
	err := o.Raw(`select 
			user_info.uid,
			user_info.uid_by_bjjj,
			user_info.phone,
			user_info.license_number,
			user_info.engine_number,
			user_info.car_type_code,
			user_info.vehicle_type,
			user_info.driver_name,
			user_info.driver_license_number,
			user_info.carid,
			user_info.car_model,
			user_info.carreg_time 
		from 
			user_info,
			user_status,
			apply_status 
		where 
			user_status.expire_time > ? and 
			user_status.status = 0 and 
			apply_status.status = 0 and 
			apply_status.next_apply_time < ? and 
			user_info.uid = user_status.uid and 
			user_status.uid = apply_status.uid 
		order by 
			last_apply_time, user_status.id 
		asc limit 1 
	`, timeCurrent, timeCurrent).QueryRow(&userInfo)
	fmt.Println("userinfo", userInfo.UidByBJJJ)
	if err != nil {
		o.Rollback()
	}
	_, err = o.Raw(`update
			apply_status
		set
			last_apply_time = ?,
			status= 1
	`, timeCurrent).Exec()
	if err != nil {
		o.Rollback()
	}
	applyment := s.generateApplyment(&o, &userInfo)
	fmt.Println("applyment Engineno", applyment.Engineno)
	o.Commit()
	return "hello dengtongtong"
}
