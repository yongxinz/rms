package logic

import (
	"context"

	"rms/common/errorx"
	"rms/service/sys/model"
	"rms/service/sys/rpc/internal/svc"
	"rms/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictDataRetrieveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictDataRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictDataRetrieveLogic {
	return &DictDataRetrieveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictDataRetrieveLogic) DictDataRetrieve(in *sys.DictDataRetrieveReq) (*sys.DictDataRetrieveResp, error) {
	res, err := l.svcCtx.DictData.FindOne(l.ctx, in.DictCode)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.NewDefaultError(errorx.PostIdErrorCode)
		}
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &sys.DictDataRetrieveResp{
		DictCode:  res.DictCode,
		DictLabel: res.DictLabel.String,
		DictValue: res.DictValue.String,
		DictType:  res.DictType.String,
		DictSort:  res.DictSort.Int64,
		Status:    res.Status.Int64,
		Remark:    res.Remark.String,
	}, nil
}
