package logic

import (
	"context"

	"rms/common/errorx"
	"rms/service/sys/model"
	"rms/service/sys/rpc/internal/svc"
	"rms/service/sys/rpc/sys"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type DictTypeAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictTypeAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictTypeAddLogic {
	return &DictTypeAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictTypeAddLogic) DictTypeAdd(in *sys.DictTypeAddReq) (*sys.DictTypeAddResp, error) {
	var SysDictType = new(model.SysDictType)
	err := copier.Copy(SysDictType, in)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	_, err = l.svcCtx.DictType.Insert(l.ctx, SysDictType)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &sys.DictTypeAddResp{}, nil
}
