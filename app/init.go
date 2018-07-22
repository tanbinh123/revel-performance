package app

import (
	"github.com/revel/revel"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var (
	AppVersion string
	BuildTime string
	Db *gorm.DB
)

func init() {
	revel.Filters = []revel.Filter{
		revel.PanicFilter,
		revel.RouterFilter,
		revel.FilterConfiguringFilter,
		revel.ParamsFilter,
		revel.SessionFilter,
		revel.FlashFilter,
		revel.ValidationFilter,
		revel.I18nFilter,
		revel.InterceptorFilter,
		revel.CompressFilter,
		revel.ActionInvoker,
	}

	revel.OnAppStart(initDB)
}

func initDB() {
	db, err := gorm.Open("mysql", os.Getenv("DB_SOURCE"))

	if err != nil {
		revel.ERROR.Println("MySQL open error", err)
		panic(err)
	}

	db.DB()
	Db = db
}