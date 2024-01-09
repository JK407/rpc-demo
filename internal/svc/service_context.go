package svc

import (
	"github.com/JK407/middleware/common"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"rpc-demo/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	GormDB      *gorm.DB
	RedisClient *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	gormDb, err := common.InitGorm(c.MysqlClient.DSN)
	if err != nil {
		panic("连接mysql数据库失败, error=" + err.Error())
	}
	redisClient, err := common.InitRedis(c.RedisClient.Addr, c.RedisClient.Password, c.RedisClient.DB)
	if err != nil {
		panic("连接redis失败, error=" + err.Error())
	}
	return &ServiceContext{
		Config:      c,
		GormDB:      gormDb,
		RedisClient: redisClient,
	}
}
