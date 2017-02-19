package controllers

import (
	"github.com/eugene0707/gettek/models"
	"encoding/json"
	"strconv"
	"github.com/eugene0707/gettek/common"
	"fmt"
)

// Metric actions
type MetricsController struct {
	BaseApiController
}

// @Title Create
// @Description create metric
// @Param	body		body 	models.Metric	true		"The metric content"
// @Success 201 {object} models.Metric
// @Failure 403 body is empty
// @router / [post]
func (c *MetricsController) Post() {
	var newModel models.Metric
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &newModel)

	if err != nil {
		fmt.Println(err.Error())
	}

	err = common.CurrentDB().Create(&newModel).Error
	c.renderJSON(newModel, err, 201)
}

// @Title Get
// @Description find metric by id
// @Param	id		path 	string	true		"the metric ID you want to get"
// @Success 200 {object} models.Metric
// @Failure 404 :id not exists
// @router /:id [get]
func (c *MetricsController) Get() {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.renderJSON(nil, err, 500)
	}

	data, err := models.Find(&models.Metric{}, id)
	c.renderJSON(data, err, 200)

}

// @Title GetAll
// @Description get all metrics
// @Success 200 {object} models.Metric
// @router / [get]
func (c *MetricsController) GetAll() {
	data, err := models.All(&models.Metric{})
	c.renderJSON(data, err, 200)
}

// @Title Update
// @Description update the metric
// @Param	id		path 	string	true		"The metric ID you want to update"
// @Param	body		body 	models.Metric	true		"The body"
// @Success 200 {object} models.Metric
// @Failure 404 :id not exists
// @router /:id [put]
func (c *MetricsController) Put() {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.renderJSON(nil, err, 500)
	}

	_, err = models.Find(&models.Metric{}, id)
	if err != nil {
		c.renderJSON(nil, err, 500)
	}

	var editModel models.Metric
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
// @Description delete the metric
// @Param	id		path 	string	true		"The metric ID you want to delete"
// @Success 204 {string} delete success!
// @Failure 404 :id not exists
// @router /:id [delete]
func (c *MetricsController) Delete() {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		c.renderJSON(nil, err, 500)
	}

	data, err := models.Destroy(&models.Metric{}, id)
	c.renderJSON(data, err, 204)
}
