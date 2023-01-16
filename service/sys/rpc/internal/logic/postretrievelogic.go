package logic

import (
	"context"

	"rms/common/errorx"
	"rms/service/sys/model"
	"rms/service/sys/rpc/internal/svc"
	"rms/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostRetrieveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostRetrieveLogic {
	return &PostRetrieveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PostRetrieveLogic) PostRetrieve(in *sys.PostRetrieveReq) (*sys.PostRetrieveResp, error) {
	res, err := l.svcCtx.Post.FindOne(l.ctx, in.PostId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.NewDefaultError(errorx.PostIdErrorCode)
		}
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &sys.PostRetrieveResp{
		PostId:   res.PostId,
		PostName: res.PostName.String,
		PostCode: res.PostCode.String,
		Sort:     res.Sort.Int64,
		Status:   res.Status.Int64,
		Remark:   res.Remark.String,
	}, nil
}
