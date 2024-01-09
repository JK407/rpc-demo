package logic

import (
	"context"
	"rpc-demo/internal/repository"

	"rpc-demo/internal/svc"
	"rpc-demo/rpc-demo"

	"github.com/zeromicro/go-zero/core/logx"
)

type HelloLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	helloRepo repository.HelloRepository
}

func NewHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HelloLogic {
	return &HelloLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		helloRepo: repository.NewHelloRepository(svcCtx.GormDB, svcCtx.RedisClient),
	}
}

func (l *HelloLogic) Hello(in *rpc_demo.HelloRequest) (*rpc_demo.HelloResponse, error) {
	// todo: add your logic here and delete this line
	resp, _ := l.helloRepo.Hello(l.ctx, in)
	return resp, nil
}
