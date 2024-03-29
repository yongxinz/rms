package post

import (
	"context"

	"rms/service/sys/api/internal/svc"
	"rms/service/sys/api/internal/types"
	"rms/service/sys/rpc/sysclient"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type PostRetrieveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostRetrieveLogic {
	return &PostRetrieveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostRetrieveLogic) PostRetrieve(req *types.PostRetrieveReq) (resp *types.PostRetrieveResp, err error) {
	res, err := l.svcCtx.SysRpc.PostRetrieve(l.ctx, &sysclient.PostRetrieveReq{
		PostId: req.PostId,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.PostRetrieveResp{}
	err = copier.Copy(resp, res)

	return
}
