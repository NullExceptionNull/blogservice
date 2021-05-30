package service

import (
	"blog-service/global"
	dao2 "blog-service/internal/order/dao"
	"context"
)

type Service struct {
	ctx context.Context
	dao *dao2.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao2.NewDao(global.DBEngine)
	return svc
}
