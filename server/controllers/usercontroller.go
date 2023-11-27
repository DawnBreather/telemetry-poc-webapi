package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"server/models" // replace with your actual app path
	"strconv"
	"time"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	// Import additional packages as needed
)

type UserController struct {
	beego.Controller
}

// LoginCredentials struct for login
type LoginCredentials struct {
	Username string
	Password string
}

// @Title CreateUser
// @Description Create a new user
// @router / [post]
// @Success 200 {object} map[string]string{"result":"success"}
// @Failure 400 Bad request
// @Failure 500 Internal Server Error
func (c *UserController) CreateUser() {
	fmt.Printf("%+v\n", c.Ctx.Input)
	var user models.User
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if err != nil {
		logs.Error("Error unmarshaling request data %s: %v", string(c.Ctx.Input.RequestBody), err)
		c.Ctx.Output.SetStatus(400) // Bad Request
		c.Data["json"] = map[string]string{"error": "Bad request"}
		c.ServeJSON()
		return
	}

	// Hashing the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logs.Error("Error hashing password: ", err)
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Error hashing password"}
		c.ServeJSON()
		return
	}
	user.Password = string(hashedPassword)

	o := orm.NewOrm()
	_, err = o.Insert(&user)
	if err != nil {
		logs.Error("Error creating user: ", err)
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Error creating user"}
	} else {
		logs.Info("User created successfully: ", user.Id)
		c.Data["json"] = map[string]string{"result": "success"}
	}
	c.ServeJSON()
}

// @Title Login
// @Description Authenticate a user and return a token
// @router /login [post]
// @Success 200 {object} map[string]string{"token":"YOUR_TOKEN"}
// @Failure 400 Bad request
// @Failure 401 Unauthorized
// @Failure 500 Internal Server Error
func (c *UserController) Login() {
	var credentials LoginCredentials
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &credentials)
	if err != nil {
		logs.Error("Error parsing login credentials: ", err)
		c.Ctx.Output.SetStatus(400) // Bad Request
		c.Data["json"] = map[string]string{"error": "Invalid request data"}
		c.ServeJSON()
		return
	}

	o := orm.NewOrm()
	user := models.User{Username: credentials.Username}
	err = o.Read(&user, "Username")
	if err == orm.ErrNoRows {
		logs.Warning("Login attempt for non-existing user: ", credentials.Username)
		c.Ctx.Output.SetStatus(401) // Unauthorized
		c.Data["json"] = map[string]string{"error": "Invalid credentials"}
		c.ServeJSON()
		return
	}

	// Check if user exists and password is correct
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		logs.Warning("Invalid login attempt for user: ", credentials.Username)
		c.Ctx.Output.SetStatus(401) // Unauthorized
		c.Data["json"] = map[string]string{"error": "Invalid credentials"}
		c.ServeJSON()
		return
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte("YourSigningSecret")) // Use a secret from your environment
	if err != nil {
		logs.Error("Error generating JWT token: ", err)
		c.Ctx.Output.SetStatus(500) // Internal Server Error
		c.Data["json"] = map[string]string{"error": "Error generating token"}
		c.ServeJSON()
		return
	}

	logs.Info("User logged in: ", credentials.Username)
	c.Data["json"] = map[string]string{"token": tokenString}
	c.ServeJSON()
}

// @Title UpdateUser
// @Description Update an existing user's information
// @router /:userID [put]
// @Param userID path int true "User ID"
// @Success 200 {object} map[string]string{"result":"success"}
// @Failure 400 Bad request
// @Failure 404 User not found
// @Failure 500 Internal Server Error
func (c *UserController) UpdateUser() {
	userID, err := strconv.Atoi(c.Ctx.Input.Param(":userID"))
	if err != nil {
		logs.Error("Error converting user ID to integer: ", err)
		c.Ctx.Output.SetStatus(400) // Bad Request
		c.Data["json"] = map[string]string{"error": "Invalid user ID"}
		c.ServeJSON()
		return
	}

	var user models.User
	json.Unmarshal(c.Ctx.Input.RequestBody, &user)

	o := orm.NewOrm()
	user.Id = userID

	// Check if the user exists before updating
	existingUser := models.User{Id: userID}
	err = o.Read(&existingUser)
	if err == orm.ErrNoRows {
		logs.Warning("Attempt to update non-existing user: ", userID)
		c.Ctx.Output.SetStatus(404) // Not Found
		c.Data["json"] = map[string]string{"error": "User not found"}
		c.ServeJSON()
		return
	}

	// Update the user information
	numUpdated, err := o.Update(&user)
	if err != nil {
		logs.Error("Error updating user: ", err)
		c.Ctx.Output.SetStatus(500) // Internal Server Error
		c.Data["json"] = map[string]string{"error": "Error updating user"}
	} else {
		if numUpdated == 0 {
			logs.Warning("No changes detected during user update: ", userID)
			c.Data["json"] = map[string]string{"result": "No changes detected"}
		} else {
			logs.Info("User updated successfully: ", userID)
			c.Data["json"] = map[string]string{"result": "success"}
		}
	}
	c.ServeJSON()
}

// @Title DeleteUser
// @Description Remove a user from the system
// @router /:userID [delete]
// @Param userID path int true "User ID"
// @Success 200 {object} map[string]string{"result":"success"}
// @Failure 400 Bad request
// @Failure 404 User not found
// @Failure 500 Internal Server Error
func (c *UserController) DeleteUser() {
	userID, err := strconv.Atoi(c.Ctx.Input.Param(":userID"))
	if err != nil {
		logs.Error("Error converting user ID to integer: ", err)
		c.Ctx.Output.SetStatus(400) // Bad Request
		c.Data["json"] = map[string]string{"error": "Invalid user ID"}
		c.ServeJSON()
		return
	}

	o := orm.NewOrm()
	numDeleted, err := o.Delete(&models.User{Id: userID})
	if err != nil {
		logs.Error("Error deleting user: ", err)
		c.Ctx.Output.SetStatus(500) // Internal Server Error
		c.Data["json"] = map[string]string{"error": "Error deleting user"}
	} else {
		if numDeleted == 0 {
			logs.Warning("Attempted to delete non-existing user: ", userID)
			c.Ctx.Output.SetStatus(404) // Not Found
			c.Data["json"] = map[string]string{"error": "User not found"}
		} else {
			logs.Info("User deleted successfully: ", userID)
			c.Data["json"] = map[string]string{"result": "success"}
		}
	}
	c.ServeJSON()
}

// @Title GetUserDetails
// @Description Retrieve detailed information for a specific user
// @router /:userID [get]
// @Param userID path int true "User ID"
// @Success 200 {object} models.User
// @Failure 400 Bad request
// @Failure 404 User not found
// @Failure 500 Internal Server Error
func (c *UserController) GetUserDetails() {
	userID, err := strconv.Atoi(c.Ctx.Input.Param(":userID"))
	if err != nil {
		logs.Error("Error converting user ID to integer: ", err)
		c.Ctx.Output.SetStatus(400) // Bad Request
		c.Data["json"] = map[string]string{"error": "Invalid user ID"}
		c.ServeJSON()
		return
	}

	o := orm.NewOrm()
	user := models.User{Id: userID}
	err = o.Read(&user)
	if err == orm.ErrNoRows {
		logs.Warning("Attempt to retrieve details of non-existing user: ", userID)
		c.Ctx.Output.SetStatus(404) // Not Found
		c.Data["json"] = map[string]string{"error": "User not found"}
		c.ServeJSON()
		return
	} else if err != nil {
		logs.Error("Error retrieving user details: ", err)
		c.Ctx.Output.SetStatus(500) // Internal Server Error
		c.Data["json"] = map[string]string{"error": "Error retrieving user details"}
		c.ServeJSON()
		return
	}

	c.Data["json"] = user
	c.ServeJSON()
}

// @Title ListUsers
// @Description List all users in the system
// @router /list [get]
// @Success 200 {array} []models.User
// @Failure 500 Internal Server Error
func (c *UserController) ListUsers() {
	o := orm.NewOrm()
	users := []models.User{}
	_, err := o.QueryTable("user").All(&users)
	if err != nil {
		logs.Error("Error retrieving users: ", err)
		c.Ctx.Output.SetStatus(500) // Internal Server Error
		c.Data["json"] = map[string]string{"error": "Error retrieving users"}
		c.ServeJSON()
		return
	}

	c.Data["json"] = users
	c.ServeJSON()
}
