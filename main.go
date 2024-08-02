package main

import (
	"template/configs"
	"template/pkg/db"
	"template/pkg/logger"
	"template/pkg/redis"
)

func init() {
	configs.Setup()
	logger.Setup()
	db.ConnectMongoDB()
	redis.ConnectRedis()
}
func main() {

}
