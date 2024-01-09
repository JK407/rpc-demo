package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	MysqlClient struct {
		DSN string
	}
	RedisClient struct {
		Addr     string
		Password string
		DB       int
	}
}
