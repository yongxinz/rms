package logic

import (
	"context"

	"rms/common/errorx"
	"rms/service/sys/rpc/internal/svc"
	"rms/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictDataDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictDataDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictDataDeleteLogic {
	return &DictDataDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictDataDeleteLogic) DictDataDelete(in *sys.DictDataDeleteReq) (*sys.DictDataDeleteResp, error) {
	err := l.svcCtx.DictData.DeleteMulti(l.ctx, in.DictIds)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &sys.DictDataDeleteResp{}, nil
}
