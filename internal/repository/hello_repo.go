package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	rpc_demo "rpc-demo/rpc-demo"
)

type HelloRepository interface {
	Hello(ctx context.Context, in *rpc_demo.HelloRequest) (*rpc_demo.HelloResponse, error)
}

type HelloRepositoryImpl struct {
	CourseDB    *gorm.DB
	CourseRedis *redis.Client
}

func (h HelloRepositoryImpl) Hello(ctx context.Context, in *rpc_demo.HelloRequest) (*rpc_demo.HelloResponse, error) {
	return &rpc_demo.HelloResponse{
		Msg: "hello " + in.Msg,
	}, nil
}

func NewHelloRepository(courseDB *gorm.DB, courseRedis *redis.Client) HelloRepository {
	return &HelloRepositoryImpl{
		CourseDB:    courseDB,
		CourseRedis: courseRedis,
	}
}
