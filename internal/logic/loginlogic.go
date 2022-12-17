package logic

import (
	"context"

	"demo/gzerolearn/internal/svc"
	"demo/gzerolearn/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	if len(req.Name) == 0 {
		return
	}

	if len(req.Password) == 0 {
		return
	}

	// todo: add your logic here and delete this line
	resp = &types.LoginResponse{Token: "12313123123123123123123"}
	return
}
