package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	err error
	DB  *gorm.DB
	Red *redis.Client
)

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
	// 添加日志
	newlogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢SQL
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	DB, err = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{Logger: newlogger})
	if err != nil {
		panic("failed to connect database")
	}
	//usr := models.UserBasic{}
	//DB.Find(&usr)
	//fmt.Println(usr)
}

func InitRedis() {
	Red = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})
	/*
		打印是否连接成功
		pong, err1 := Red.Ping().Result()
		if err != nil {
			fmt.Println("error :", err1)
		} else {
			fmt.Println("sucess,", pong)
		}
	*/
}

const (
	PublishKey = "websocket"
)

func Publish(ctx context.Context, channel string, msg string) error {
	err := Red.Publish(ctx, channel, msg).Err()
	fmt.Println("Publish,", msg)
	return err
}

func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := Red.PSubscribe(ctx, channel)
	msg, err := sub.ReceiveMessage(ctx)
	fmt.Println("Subscribe,", msg.Payload)
	return msg.String(), err
}
