package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"server/controllers" // Replace with your actual app path
)

func init() {
	nsUser := beego.NewNamespace("/user",
		beego.NSRouter("/login", &controllers.UserController{}, "post:Login"),
		beego.NSRouter("/create", &controllers.UserController{}, "post:CreateUser"),
		beego.NSRouter("/update/:userID", &controllers.UserController{}, "put:UpdateUser"),
		beego.NSRouter("/delete/:userID", &controllers.UserController{}, "delete:DeleteUser"),
		beego.NSRouter("/:userID", &controllers.UserController{}, "get:GetUserDetails"),
		beego.NSRouter("/", &controllers.UserController{}, "get:ListUsers"),
	)

	nsCar := beego.NewNamespace("/car",
		beego.NSRouter("/create", &controllers.CarController{}, "post:CreateCar"),
		beego.NSRouter("/update/:carID", &controllers.CarController{}, "put:UpdateCar"),
		beego.NSRouter("/delete/:carID", &controllers.CarController{}, "delete:DeleteCar"),
		beego.NSRouter("/:carID", &controllers.CarController{}, "get:GetCarDetails"),
		beego.NSRouter("/", &controllers.CarController{}, "get:ListCars"),
	)

	nsMisc := beego.NewNamespace("/misc",
		beego.NSRouter("/health", &controllers.MiscController{}, "get:HealthCheck"),
		beego.NSRouter("/version", &controllers.MiscController{}, "get:APIVersion"),
	)

	// Register namespaces
	beego.AddNamespace(nsUser, nsCar, nsMisc)
}
