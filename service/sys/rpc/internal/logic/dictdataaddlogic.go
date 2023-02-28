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

type DictDataAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictDataAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictDataAddLogic {
	return &DictDataAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictDataAddLogic) DictDataAdd(in *sys.DictDataAddReq) (*sys.DictDataAddResp, error) {
	var SysDictData = new(model.SysDictData)
	err := copier.Copy(SysDictData, in)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	_, err = l.svcCtx.DictData.Insert(l.ctx, SysDictData)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &sys.DictDataAddResp{}, nil
}
