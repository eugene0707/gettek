package controllers

import (
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	"net/http"
)

// Base Api operations
type BaseApiController struct {
	beego.Controller
}

func (c *BaseApiController) renderJSON(data interface{}, err error, status int) {

	c.Data["json"] = data
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			status = http.StatusNotFound
		case gorm.ErrInvalidSQL, gorm.ErrInvalidTransaction, gorm.ErrCantStartTransaction, gorm.ErrUnaddressable:
			status = http.StatusInternalServerError
		}
		c.Data["json"] = err.Error()
	}

	c.Ctx.Output.SetStatus(status)
	c.ServeJSON()

}
