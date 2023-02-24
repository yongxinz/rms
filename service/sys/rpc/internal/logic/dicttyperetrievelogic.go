package logic

import (
	"context"

	"rms/common/errorx"
	"rms/service/sys/model"
	"rms/service/sys/rpc/internal/svc"
	"rms/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictTypeRetrieveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictTypeRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictTypeRetrieveLogic {
	return &DictTypeRetrieveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictTypeRetrieveLogic) DictTypeRetrieve(in *sys.DictTypeRetrieveReq) (*sys.DictTypeRetrieveResp, error) {
	res, err := l.svcCtx.DictType.FindOne(l.ctx, in.DictId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.NewDefaultError(errorx.PostIdErrorCode)
		}
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &sys.DictTypeRetrieveResp{
		DictId:   res.DictId,
		DictName: res.DictName.String,
		DictType: res.DictType.String,
		Status:   res.Status.Int64,
		Remark:   res.Remark.String,
	}, nil
}
