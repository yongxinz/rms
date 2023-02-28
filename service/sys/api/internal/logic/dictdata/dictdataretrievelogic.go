package dictdata

import (
	"context"

	"rms/service/sys/api/internal/svc"
	"rms/service/sys/api/internal/types"
	"rms/service/sys/rpc/sysclient"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type DictDataRetrieveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDictDataRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictDataRetrieveLogic {
	return &DictDataRetrieveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DictDataRetrieveLogic) DictDataRetrieve(req *types.DictDataRetrieveReq) (resp *types.DictDataRetrieveResp, err error) {
	res, err := l.svcCtx.SysRpc.DictDataRetrieve(l.ctx, &sysclient.DictDataRetrieveReq{
		DictCode: req.DictCode,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.DictDataRetrieveResp{}
	err = copier.Copy(resp, res)

	return
}
