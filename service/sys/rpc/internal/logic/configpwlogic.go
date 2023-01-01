package logic

import (
	"context"

	"rms/service/sys/model"
	"rms/service/sys/rpc/internal/svc"
	"rms/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
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
			return nil, status.Error(100, "查询配置信息失败")
		}
		return nil, status.Error(500, err.Error())
	}

	configKey, configValue := "", ""
	for _, item := range res {
		configKey = item.ConfigKey.String
		configValue = item.ConfigValue.String
		break
	}

	return &sys.ConfigPwResp{
		ConfigKey:   configKey,
		ConfigValue: configValue,
	}, nil
}
