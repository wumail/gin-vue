package common

import (
	"fmt"
	"main/model"
	"net/url"

	"github.com/spf13/viper"

	"github.com/jinzhu/gorm"
)

//DB 全局DB变量
var DB *gorm.DB

//initDB 初始化数据库i
func initDB() *gorm.DB {
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	loc := viper.GetString("datasource.loc")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc))
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("faild to connect database err:" + err.Error())
	}

	db.AutoMigrate(&model.User{})

	return db
}

//GetDB 返回DB对象
func GetDB() *gorm.DB {
	DB = initDB()
	return DB
}
