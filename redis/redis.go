package redis

import (
	"github.com/go-redis/redis"

	"github.com/ihahoo/go-api-lib/config"
	"github.com/ihahoo/go-api-lib/log"
)

// ConnectDB 连接redis
func ConnectDB(opt *redis.Options) (*redis.Client, error) {
	client := redis.NewClient(opt)
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

// Connect 用配置文件连接数据库
func Connect(db int, prePath string) *redis.Client {
	host := config.GetString(prePath + "host")
	port := config.GetString(prePath + "port")
	password := config.GetString(prePath + "password")

	client, err := Conn(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       db,
	})
	if err != nil {
		logger := log.GetLog()
		logger.WithFields(logger.Fields{"func": "redis.Client"}).Fatal(err)
	}
	return client
}

// Conn 用配置文件的默认参数连接数据库
func Conn() *redis.Client {
	return Connect("redis.")
}
