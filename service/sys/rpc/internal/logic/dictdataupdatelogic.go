package logic

import (
	"context"

	"rms/common/errorx"
	"rms/service/sys/rpc/internal/svc"
	"rms/service/sys/rpc/sys"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type DictDataUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictDataUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictDataUpdateLogic {
	return &DictDataUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictDataUpdateLogic) DictDataUpdate(in *sys.DictDataUpdateReq) (*sys.DictDataUpdateResp, error) {
	SysDictData, err := l.svcCtx.DictData.FindOne(l.ctx, in.DictId)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.UserIdErrorCode)
	}

	err = copier.Copy(SysDictData, in)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}
	err = l.svcCtx.DictData.Update(l.ctx, SysDictData)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &sys.DictDataUpdateResp{}, nil
}
