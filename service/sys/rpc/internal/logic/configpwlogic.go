package logic

import (
	"context"

	"rms/common/errorx"
	"rms/service/sys/model"
	"rms/service/sys/rpc/internal/svc"
	"rms/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigPwLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConfigPwLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigPwLogic {
	return &ConfigPwLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ConfigPwLogic) ConfigPw(in *sys.ConfigPwReq) (*sys.ConfigPwResp, error) {
	res, err := l.svcCtx.SysConfigModel.FindByKey(l.ctx, in.ConfigKey)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.NewDefaultError(errorx.ConfigErrorCode)
		}
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	configKey, configValue := "", ""
	for _, item := range res {
		configKey = item.ConfigKey
		configValue = item.ConfigValue
		break
	}

	return &sys.ConfigPwResp{
		ConfigKey:   configKey,
		ConfigValue: configValue,
	}, nil
}
