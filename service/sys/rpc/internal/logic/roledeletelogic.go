package logic

import (
	"context"

	"rms/common/errorx"
	"rms/service/sys/rpc/internal/svc"
	"rms/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleDeleteLogic {
	return &RoleDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleDeleteLogic) RoleDelete(in *sys.RoleDeleteReq) (*sys.RoleDeleteResp, error) {
	for _, roleId := range in.RoleIds {
		l.svcCtx.RoleMenu.DeleteByRoleId(l.ctx, roleId)
	}
	err := l.svcCtx.Role.DeleteMulti(l.ctx, in.RoleIds)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &sys.RoleDeleteResp{}, nil
}
