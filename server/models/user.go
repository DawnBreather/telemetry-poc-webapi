package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	Id       int
	Username string `orm:"size(128);unique"`
	Password string `orm:"size(128)"`
	Email    string `orm:"size(128);unique"`
}

func init() {
	orm.RegisterModel(new(User))
}
