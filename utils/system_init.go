package utils

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("config app:", viper.Get("app"))
	fmt.Println("config app:", viper.Get("mysql"))
}

func InitMysql() {
	DB, err = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//usr := models.UserBasic{}
	//DB.Find(&usr)
	//fmt.Println(usr)
}
