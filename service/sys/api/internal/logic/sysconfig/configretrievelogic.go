package sysconfig

import (
	"context"

	"rms/service/sys/api/internal/svc"
	"rms/service/sys/api/internal/types"
	"rms/service/sys/rpc/sysclient"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigRetrieveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfigRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigRetrieveLogic {
	return &ConfigRetrieveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfigRetrieveLogic) ConfigRetrieve(req *types.ConfigRetrieveReq) (resp *types.ConfigRetrieveResp, err error) {
	res, err := l.svcCtx.SysRpc.ConfigRetrieve(l.ctx, &sysclient.ConfigRetrieveReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.ConfigRetrieveResp{}
	err = copier.Copy(resp, res)

	return
}
