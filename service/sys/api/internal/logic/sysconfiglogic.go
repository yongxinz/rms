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

	data := map[string]string{
		"sys_app_logo": res.SysAppLogo,
		"sys_app_name": res.SysAppName,
	}

	return &types.SysConfigResp{
		Data: data,
		Code: 200,
		Msg:  "查询成功",
	}, nil
}
