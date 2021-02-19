package main

import (
	"main/common"
	"main/routes"
	"os"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	InitConfig()
	db := common.GetDB()
	defer db.Close()
	r := gin.Default()
	r = routes.CollectRouter(r)
	// port := viper.GetString("server.port")
	// if port != "" {
	// 	panic(r.Run(":" + port))
	// }
	r.Run()
}

//InitConfig 读取配置文件
func InitConfig() {
	workDir, _ := os.Getwd()
	// log.Println(workDir)
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("faild to init config err:" + err.Error())
	}
}
