package logic

import (
	"context"

	"rms/common/errorx"
	"rms/service/sys/rpc/internal/svc"
	"rms/service/sys/rpc/sys"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type DictTypeUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictTypeUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictTypeUpdateLogic {
	return &DictTypeUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictTypeUpdateLogic) DictTypeUpdate(in *sys.DictTypeUpdateReq) (*sys.DictTypeUpdateResp, error) {
	sysDictType, err := l.svcCtx.DictType.FindOne(l.ctx, in.DictId)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.UserIdErrorCode)
	}

	err = copier.Copy(sysDictType, in)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}
	err = l.svcCtx.DictType.Update(l.ctx, sysDictType)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &sys.DictTypeUpdateResp{}, nil
}
