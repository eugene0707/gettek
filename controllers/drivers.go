package controllers

import (
	"github.com/eugene0707/gettek/models"
	"encoding/json"
	"strconv"
	"github.com/eugene0707/gettek/common"
	"fmt"
)

// Driver actions
type DriversController struct {
	BaseApiController
}

// @Title Create
// @Description create driver
// @Param	body		body 	models.Driver	true		"The driver content"
// @Success 201 {object} models.Driver
// @Failure 403 body is empty
// @router / [post]
func (c *DriversController) Post() {
	var newModel models.Driver
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &newModel)

	if err != nil {
		fmt.Println(err.Error())
	}

	err = common.CurrentDB().Create(&newModel).Error
	c.renderJSON(newModel, err, 201)
}

// @Title Get
// @Description find driver by id
// @Param	id		path 	string	true		"the driver ID you want to get"
// @Success 200 {object} models.Driver
// @Failure 404 :id not exists
// @router /:id [get]
func (c *DriversController) Get() {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.renderJSON(nil, err, 500)
	}

	data, err := models.Find(&models.Driver{}, id)
	c.renderJSON(data, err, 200)

}

// @Title GetAll
// @Description get all drivers
// @Success 200 {object} models.Driver
// @router / [get]
func (c *DriversController) GetAll() {
	data, err := models.All(&models.Driver{})
	c.renderJSON(data, err, 200)
}

// @Title Update
// @Description update the driver
// @Param	id		path 	string	true		"The driver ID you want to update"
// @Param	body		body 	models.Driver	true		"The body"
// @Success 200 {object} models.Driver
// @Failure 404 :id not exists
// @router /:id [put]
func (c *DriversController) Put() {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.renderJSON(nil, err, 500)
	}

	_, err = models.Find(&models.Driver{}, id)
	if err != nil {
		c.renderJSON(nil, err, 500)
	}

	var editModel models.Driver
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &editModel)

	if err != nil {
		c.renderJSON(nil, err, 500)
	} else {
		editModel.ID = id
	}

	err = common.CurrentDB().Save(&editModel).Error
	c.renderJSON(editModel, err, 200)
}

// @Title Delete
// @Description delete the driver
// @Param	id		path 	string	true		"The driver ID you want to delete"
// @Success 204 {string} delete success!
// @Failure 404 :id not exists
// @router /:id [delete]
func (c *DriversController) Delete() {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.renderJSON(nil, err, 500)
	}

	data, err := models.Destroy(&models.Driver{}, id)
	c.renderJSON(data, err, 204)
}
