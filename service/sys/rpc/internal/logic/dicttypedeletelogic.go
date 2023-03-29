package logic

import (
	"context"

	"rms/common/errorx"
	"rms/service/sys/rpc/internal/svc"
	"rms/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictTypeDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictTypeDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictTypeDeleteLogic {
	return &DictTypeDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictTypeDeleteLogic) DictTypeDelete(in *sys.DictTypeDeleteReq) (*sys.DictTypeDeleteResp, error) {
	err := l.svcCtx.DictType.DeleteMulti(l.ctx, in.Ids)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &sys.DictTypeDeleteResp{}, nil
}
