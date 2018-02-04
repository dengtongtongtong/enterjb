package controllers

import (
	"encoding/json"
	"enterbj/models"
	"strconv"

	"enterbj/service"

	"gocommon/webutils"

	"github.com/astaxie/beego"
)

//  ApplyMentController operations for ApplyMent
type ApplyMentController struct {
	beego.Controller
}

// URLMapping ...
func (c *ApplyMentController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetNextUnsigned)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create ApplyMent
// @Param	body		body 	models.ApplyMent	true		"body for ApplyMent content"
// @Success 201 {int} models.ApplyMent
// @Failure 403 body is empty
// @router / [post]
func (c *ApplyMentController) Post() {
	var v models.ApplyMent
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if _, err := models.AddApplyMent(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get ApplyMent by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ApplyMent
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ApplyMentController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetApplyMentById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get Next Unsigned
// @Description get ApplyMent
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.ApplyMent
// @Failure 403
// @router / [get]
func (c *ApplyMentController) GetNextUnsigned() {
	serviceApply := service.ApplyMentService{}
	applyment := serviceApply.ObtainUnsignedApplyment()
	c.Data["json"] = webutils.Success(applyment)
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the ApplyMent
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.ApplyMent	true		"body for ApplyMent content"
// @Success 200 {object} models.ApplyMent
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ApplyMentController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.ApplyMent{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateApplyMentById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the ApplyMent
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ApplyMentController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteApplyMent(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
