package dicttype

import (
	"context"

	"rms/service/sys/api/internal/svc"
	"rms/service/sys/api/internal/types"
	"rms/service/sys/rpc/sysclient"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type DictTypeRetrieveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDictTypeRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictTypeRetrieveLogic {
	return &DictTypeRetrieveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DictTypeRetrieveLogic) DictTypeRetrieve(req *types.DictTypeRetrieveReq) (resp *types.DictTypeRetrieveResp, err error) {
	res, err := l.svcCtx.SysRpc.DictTypeRetrieve(l.ctx, &sysclient.DictTypeRetrieveReq{
		DictId: req.DictId,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.DictTypeRetrieveResp{}
	err = copier.Copy(resp, res)

	return
}
