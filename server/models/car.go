package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Car struct {
	Id    int    `orm:"auto"`
	Make  string `orm:"size(100)"`
	Model string `orm:"size(100)"`
	Year  int
	Color string `orm:"size(50)"`
}

func init() {
	orm.RegisterModel(new(Car))
}
