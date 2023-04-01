package service

import (
	"context"

	"blog-service/global"
	"blog-service/internal/dao"

	otgorm "github.com/eddycjy/opentracing-gorm"
)

/**
用于处理标签模块的业务逻辑
*/

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{
		ctx: ctx,
	}
	svc.dao = dao.New(otgorm.WithContext(svc.ctx, global.DBEngine))

	return svc
}
