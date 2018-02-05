package service

import (
	"encoding/json"
	"enterbj/models"
	"fmt"
	"gocommon/timeutils"
	"time"

	"github.com/astaxie/beego/orm"
)

type ApplyMentService struct {
}

func (s *ApplyMentService) obtainFinalSubmitImageIdWithSecret(data *models.ApplyMent) (imageId string) {
	appkey := "0791682354"
	imageId = appkey + s.obtainFinalSubmitImageId(data) + appkey
	return imageId
}

func (s *ApplyMentService) obtainFinalSubmitImageId(data *models.ApplyMent) (imageId string) {
	imageId = data.Inbjentrancecode + data.Inbjduration + data.Inbjtime + data.UidByBJJJ + data.Engineno + data.Cartypecode + data.Driverlicenseno + data.Carid + data.Timestamp
	return imageId
}

func (s *ApplyMentService) obtainDayStamp() (currentTime string) {
	t := time.Now()
	return fmt.Sprintf("%v-%v-%v", t.Year(), int(t.Month()), t.Day())
}

func (s *ApplyMentService) obtainTimeStamp() (currentTime string) {
	t := time.Now()
	return fmt.Sprintf("%v-%v-%v %v:%v:%v", t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func (s *ApplyMentService) generateApplyment(o *orm.Ormer, userInfo *models.UserInfo) (applyment models.ApplyMent, err error) {
	jsonUserInfo, _ := json.Marshal(userInfo)
	json.Unmarshal(jsonUserInfo, &applyment)
	applyment.Inbjduration = "02"
	applyment.Inbjtime = s.obtainDayStamp()
	applyment.Timestamp = s.obtainTimeStamp()
	imageId := s.obtainFinalSubmitImageIdWithSecret(&applyment)
	applyment.ImageId = imageId
	_, err = (*o).Insert(&applyment)
	return applyment, err
}

func (s *ApplyMentService) ObtainUnsignedApplyment() (applyment models.ApplyMent) {
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
		switch err {
		case orm.ErrNoRows:
			applyment = models.ApplyMent{}
			return applyment
		}
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
	applyment, err = s.generateApplyment(&o, &userInfo)
	o.Commit()
	return applyment
}

func (s *ApplyMentService) generateSignedApplyment(applyment *models.ApplyMent) {
	o := orm.NewOrm()
	v := models.ApplyMent{
		ImageId: applyment.ImageId,
	}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}
