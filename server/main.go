package main

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
	"os"
	_ "server/controllers"
	_ "server/models"
	_ "server/routers"
)

func init() {
	// Log to file and console
	logs.SetLogger(logs.AdapterFile, `{"filename":"logs/server.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	logs.SetLogger(logs.AdapterConsole)

	// Connect to the database
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		os.Getenv("DB_POSTGRES_USERNAME"), os.Getenv("DB_POSTGRES_PASSWORD"), os.Getenv("DB_POSTGRES_HOST_AND_PORT"), os.Getenv("DB_POSTGRES_DBNAME")))

	// Initialize the ORM
	//orm.SetMaxIdleConns("default", 10)
	//orm.SetMaxOpenConns("default", 100)
	orm.RunSyncdb("default", false, true)

	// Enable function name and line number in logs
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
}

// https://botoom.hashnode.dev/rest-api-with-golang-and-postgresql-using-beego
func main() {
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	beego.Run()
}
