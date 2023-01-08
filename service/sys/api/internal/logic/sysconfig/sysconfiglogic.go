package logic

import (
	"context"

	"rms/service/sys/api/internal/svc"
	"rms/service/sys/api/internal/types"
	"rms/service/sys/rpc/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysConfigLogic {
	return &SysConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysConfigLogic) SysConfig() (resp *types.SysConfigResp, err error) {
	res, err := l.svcCtx.SysRpc.SysConfig(l.ctx, &sysclient.SysConfigReq{})
	if err != nil {
		return nil, err
	}

	return &types.SysConfigResp{
		SysAppName: res.SysAppName,
		SysAppLogo: res.SysAppLogo,
	}, nil
}
