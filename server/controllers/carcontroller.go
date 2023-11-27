package controllers

import (
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	"server/models" // Replace with your actual app path
	"server/utils"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type CarController struct {
	beego.Controller
}

// @Title CreateCar
// @Description Add a new car to the system
// @router / [post]
// @Success 200 {object} map[string]string{"result":"success"}
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 500 Internal Server Error
func (c *CarController) CreateCar() {
	if !utils.IsAuthorized(c.Ctx.Request) {
		logs.Warning("Unauthorized attempt to create car")
		c.Ctx.Output.SetStatus(401)
		c.Data["json"] = map[string]string{"error": "Unauthorized"}
		c.ServeJSON()
		return
	}

	var car models.Car
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &car)
	if err != nil {
		logs.Error("Error parsing request data: ", err)
		c.Ctx.Output.SetStatus(400) // Bad Request
		c.Data["json"] = map[string]string{"error": "Bad request"}
		c.ServeJSON()
		return
	}

	o := orm.NewOrm()
	_, err = o.Insert(&car)
	if err != nil {
		logs.Error("Error creating car: ", err)
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Error creating car"}
	} else {
		logs.Info("Car created successfully: ", car.Id)
		c.Data["json"] = map[string]string{"result": "success"}
	}
	c.ServeJSON()
}

// @Title UpdateCar
// @Description Update an existing car entry
// @router /:carID [put]
// @Param carID path int true "Car ID"
// @Success 200 {object} map[string]string{"result":"success"}
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 404 Car not found
// @Failure 500 Internal Server Error
func (c *CarController) UpdateCar() {
	if !utils.IsAuthorized(c.Ctx.Request) {
		logs.Warning("Unauthorized attempt to update car")
		c.Ctx.Output.SetStatus(401)
		c.Data["json"] = map[string]string{"error": "Unauthorized"}
		c.ServeJSON()
		return
	}

	carID, err := strconv.Atoi(c.Ctx.Input.Param(":carID"))
	if err != nil {
		logs.Error("Error converting carID to integer: ", err)
		c.Ctx.Output.SetStatus(400) // Bad Request
		c.Data["json"] = map[string]string{"error": "Invalid car ID"}
		c.ServeJSON()
		return
	}

	var car models.Car
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &car)
	if err != nil {
		logs.Error("Error parsing request data: ", err)
		c.Ctx.Output.SetStatus(400) // Bad Request
		c.Data["json"] = map[string]string{"error": "Bad request"}
		c.ServeJSON()
		return
	}

	o := orm.NewOrm()
	car.Id = carID
	_, err = o.Update(&car)
	if err != nil {
		logs.Error("Error updating car with ID ", carID, ": ", err)
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Error updating car"}
	} else {
		logs.Info("Car updated successfully: ", carID)
		c.Data["json"] = map[string]string{"result": "success"}
	}
	c.ServeJSON()
}

// @Title DeleteCar
// @Description Remove a car from the system
// @router /:carID [delete]
// @Param carID path int true "Car ID"
// @Success 200 {object} map[string]string{"result":"success"}
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 404 Car not found
// @Failure 500 Internal Server Error
func (c *CarController) DeleteCar() {
	if !utils.IsAuthorized(c.Ctx.Request) {
		logs.Warning("Unauthorized attempt to delete car")
		c.Ctx.Output.SetStatus(401)
		c.Data["json"] = map[string]string{"error": "Unauthorized"}
		c.ServeJSON()
		return
	}

	carID, err := strconv.Atoi(c.Ctx.Input.Param(":carID"))
	if err != nil {
		logs.Error("Error converting carID to integer: ", err)
		c.Ctx.Output.SetStatus(400) // Bad Request
		c.Data["json"] = map[string]string{"error": "Invalid car ID"}
		c.ServeJSON()
		return
	}

	_, err = orm.NewOrm().Delete(&models.Car{Id: carID})
	if err != nil {
		logs.Error("Error deleting car with ID ", carID, ": ", err)
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Error deleting car"}
	} else {
		logs.Info("Car deleted successfully: Car ID ", carID)
		c.Data["json"] = map[string]string{"result": "success"}
	}
	c.ServeJSON()
}

// @Title GetCarDetails
// @Description Retrieve details of a specific car
// @router /:carID [get]
// @Param carID path int true "Car ID"
// @Success 200 {object} models.Car
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 404 Car not found
// @Failure 500 Internal Server Error
func (c *CarController) GetCarDetails() {
	if !utils.IsAuthorized(c.Ctx.Request) {
		logs.Warning("Unauthorized attempt to access car details")
		c.Ctx.Output.SetStatus(401)
		c.Data["json"] = map[string]string{"error": "Unauthorized"}
		c.ServeJSON()
		return
	}

	carID, err := strconv.Atoi(c.Ctx.Input.Param(":carID"))
	if err != nil {
		logs.Error("Error converting carID to integer: ", err)
		c.Ctx.Output.SetStatus(400) // Bad Request
		c.Data["json"] = map[string]string{"error": "Invalid car ID"}
		c.ServeJSON()
		return
	}

	car := models.Car{Id: carID}
	err = orm.NewOrm().Read(&car)
	if err != nil {
		if err == orm.ErrNoRows {
			logs.Info("Car not found with ID: ", carID)
			c.Ctx.Output.SetStatus(404)
			c.Data["json"] = map[string]string{"error": "Car not found"}
		} else {
			logs.Error("Error retrieving car from database: ", err)
			c.Ctx.Output.SetStatus(500)
			c.Data["json"] = map[string]string{"error": "Internal server error"}
		}
	} else {
		logs.Info("Car details retrieved for ID: ", carID)
		c.Data["json"] = car
	}
	c.ServeJSON()
}

// @Title ListCars
// @Description List all cars in the system
// @router /list [get]
// @Success 200 {array} models.Car
// @Failure 401 Unauthorized
// @Failure 500 Internal Server Error
func (c *CarController) ListCars() {
	if !utils.IsAuthorized(c.Ctx.Request) {
		logs.Warning("Unauthorized attempt to list cars")
		c.Ctx.Output.SetStatus(401)
		c.Data["json"] = map[string]string{"error": "Unauthorized"}
		c.ServeJSON()
		return
	}

	var cars []models.Car
	_, err := orm.NewOrm().QueryTable("car").All(&cars)
	if err != nil {
		logs.Error("Error retrieving cars: ", err)
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Error retrieving cars"}
	} else {
		logs.Info("Cars retrieved successfully")
		c.Data["json"] = cars
	}
	c.ServeJSON()
}
