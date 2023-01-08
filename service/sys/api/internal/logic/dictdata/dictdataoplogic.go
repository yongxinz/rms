package logic

import (
	"context"

	"rms/service/sys/api/internal/svc"
	"rms/service/sys/api/internal/types"
	"rms/service/sys/rpc/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictDataOpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDictDataOpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictDataOpLogic {
	return &DictDataOpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DictDataOpLogic) DictDataOp(req *types.DictDataOpReq) (resp []*types.DictDataOpResp, err error) {
	res, err := l.svcCtx.SysRpc.DictDataOp(l.ctx, &sysclient.DictDataOpReq{
		DictType: req.DictType,
	})
	if err != nil {
		return nil, err
	}

	var data []*types.DictDataOpResp
	for _, item := range res.Data {
		data = append(data, &types.DictDataOpResp{
			Label: item.Label,
			Value: item.Value,
		})
	}
	resp = data

	return
}
