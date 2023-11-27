package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MiscController struct {
	beego.Controller
}

// @Title HealthCheck
// @Description Check if the API is running and reachable
// @router /health [get]
// @Success 200 {object} map[string]string{"status":"healthy"}
// @Failure 500 Internal Server Error
func (c *MiscController) HealthCheck() {
	c.Data["json"] = map[string]string{"status": "healthy"}
	c.ServeJSON()
}

// @Title APIVersion
// @Description Provide the current version of the API
// @router /version [get]
// @Success 200 {object} map[string]string{"version":"1.0.0"}
func (c *MiscController) APIVersion() {
	// Replace "1.0.0" with your actual API version
	c.Data["json"] = map[string]string{"version": "1.0.0"}
	c.ServeJSON()
}
