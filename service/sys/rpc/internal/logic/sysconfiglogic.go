package logic

import (
	"context"

	"rms/service/sys/model"
	"rms/service/sys/rpc/internal/svc"
	"rms/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
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
			return nil, status.Error(100, "查询配置信息失败")
		}
		return nil, status.Error(500, err.Error())
	}

	sysAppLogo, sysAppName := "", ""
	for _, item := range res {
		if item.ConfigKey.String == "sys_app_logo" {
			sysAppLogo = item.ConfigValue.String
		} else if item.ConfigKey.String == "sys_app_name" {
			sysAppName = item.ConfigValue.String
		}
	}

	return &sys.SysConfigResp{
		SysAppLogo: sysAppLogo,
		SysAppName: sysAppName,
	}, nil
}
