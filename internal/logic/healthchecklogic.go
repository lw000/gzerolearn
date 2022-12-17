package logic

import (
	"context"

	"demo/gzerolearn/internal/svc"
	"demo/gzerolearn/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HealthcheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHealthcheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HealthcheckLogic {
	return &HealthcheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HealthcheckLogic) Healthcheck(req *types.CheckingRequest) (resp *types.CheckingResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
