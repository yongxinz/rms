package logic

import (
	"context"

	"rms/common/errorx"
	"rms/service/sys/model"
	"rms/service/sys/rpc/internal/svc"
	"rms/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysConfigLogic {
	return &SysConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysConfigLogic) SysConfig(in *sys.SysConfigReq) (*sys.SysConfigResp, error) {
	res, err := l.svcCtx.SysConfigModel.FindAll(l.ctx)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.NewDefaultError(errorx.ConfigErrorCode)
		}
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	sysAppLogo, sysAppName := "", ""
	for _, item := range res {
		if item.ConfigKey == "sys_app_logo" {
			sysAppLogo = item.ConfigValue
		} else if item.ConfigKey == "sys_app_name" {
			sysAppName = item.ConfigValue
		}
	}

	return &sys.SysConfigResp{
		SysAppLogo: sysAppLogo,
		SysAppName: sysAppName,
	}, nil
}
