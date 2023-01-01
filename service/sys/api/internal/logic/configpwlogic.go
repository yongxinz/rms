package logic

import (
	"context"

	"rms/service/sys/api/internal/svc"
	"rms/service/sys/api/internal/types"
	"rms/service/sys/rpc/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigPwLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigPwLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigPwLogic {
	return &ConfigPwLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfigPwLogic) ConfigPw(req *types.ConfigPwReq) (resp *types.ConfigPwResp, err error) {
	res, err := l.svcCtx.SysRpc.ConfigPw(l.ctx, &sysclient.ConfigPwReq{
		ConfigKey: req.ConfigKey,
	})
	if err != nil {
		return nil, err
	}

	data := map[string]string{
		"configKey":   res.ConfigKey,
		"configValue": res.ConfigValue,
	}

	resp = &types.ConfigPwResp{
		Data: data,
		Code: 200,
	}

	return
}
