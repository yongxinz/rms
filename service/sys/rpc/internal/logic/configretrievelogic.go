package logic

import (
	"context"

	"rms/common/errorx"
	"rms/service/sys/model"
	"rms/service/sys/rpc/internal/svc"
	"rms/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigRetrieveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConfigRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigRetrieveLogic {
	return &ConfigRetrieveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ConfigRetrieveLogic) ConfigRetrieve(in *sys.ConfigRetrieveReq) (*sys.ConfigRetrieveResp, error) {
	res, err := l.svcCtx.SysConfigModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.NewDefaultError(errorx.PostIdErrorCode)
		}
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &sys.ConfigRetrieveResp{
		Id:          res.Id,
		ConfigName:  res.ConfigName.String,
		ConfigKey:   res.ConfigKey.String,
		ConfigValue: res.ConfigValue.String,
		ConfigType:  res.ConfigType.String,
		IsFrontend:  res.IsFrontend.String,
		Remark:      res.Remark.String,
	}, nil
}
