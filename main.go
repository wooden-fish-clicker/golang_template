package main

import (
	"github.com/wooden-fish-clicker/golang_template/configs"
	"github.com/wooden-fish-clicker/golang_template/pkg/db"
	"github.com/wooden-fish-clicker/golang_template/pkg/logger"
	"github.com/wooden-fish-clicker/golang_template/pkg/redis"
)

func init() {
	configs.Setup()
	logger.Setup()
	db.ConnectMongoDB()
	redis.ConnectRedis()
}
func main() {

}
