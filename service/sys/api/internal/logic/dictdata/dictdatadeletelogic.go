package dictdata

import (
	"context"

	"rms/service/sys/api/internal/svc"
	"rms/service/sys/api/internal/types"
	"rms/service/sys/rpc/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictDataDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDictDataDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictDataDeleteLogic {
	return &DictDataDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DictDataDeleteLogic) DictDataDelete(req *types.DictDataDeleteReq) error {
	_, err := l.svcCtx.SysRpc.DictDataDelete(l.ctx, &sysclient.DictDataDeleteReq{
		DictIds: req.Ids,
	})
	return err
}
