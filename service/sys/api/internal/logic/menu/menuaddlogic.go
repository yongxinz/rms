package menu

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

type MenuAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuAddLogic {
	return &MenuAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuAddLogic) MenuAdd(req *types.MenuAddReq) error {
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}
	req.CreateBy = userId
	req.UpdateBy = userId

	var sysMenu sysclient.MenuAddReq
	err = copier.Copy(&sysMenu, req)
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	_, err = l.svcCtx.SysRpc.MenuAdd(l.ctx, &sysMenu)
	if err != nil {
		return err
	}

	return nil
}
