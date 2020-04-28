package db

import (
	"fmt"
	"simple-crud/conf"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.C.Database.User,
		conf.C.Database.Password,
		conf.C.Database.Host,
		conf.C.Database.Port,
		conf.C.Database.DB,
	)
	var err error
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	DB.DB().SetConnMaxLifetime(30 * time.Second)
	DB.DB().SetMaxIdleConns(200)
	DB.DB().SetMaxOpenConns(100)
	DB.BlockGlobalUpdate(true)
	DB.LogMode(true)
}
