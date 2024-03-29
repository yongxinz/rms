package post

import (
	"context"
	"encoding/json"

	"rms/common/errorx"
	"rms/service/sys/api/internal/svc"
	"rms/service/sys/api/internal/types"
	"rms/service/sys/rpc/sysclient"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type PostUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostUpdateLogic {
	return &PostUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostUpdateLogic) PostUpdate(req *types.PostUpdateReq) error {
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}
	req.UpdateBy = userId

	var sysPost sysclient.PostUpdateReq
	err = copier.Copy(&sysPost, req)
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	_, err = l.svcCtx.SysRpc.PostUpdate(l.ctx, &sysPost)
	if err != nil {
		return err
	}

	return nil
}
