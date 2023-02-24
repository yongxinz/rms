package dicttype

import (
	"context"

	"rms/service/sys/api/internal/svc"
	"rms/service/sys/api/internal/types"
	"rms/service/sys/rpc/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictTypeDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDictTypeDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictTypeDeleteLogic {
	return &DictTypeDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DictTypeDeleteLogic) DictTypeDelete(req *types.DictTypeDeleteReq) error {
	_, err := l.svcCtx.SysRpc.DictTypeDelete(l.ctx, &sysclient.DictTypeDeleteReq{
		DictIds: req.Ids,
	})
	return err
}
