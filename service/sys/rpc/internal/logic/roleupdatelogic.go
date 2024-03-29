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

type RoleUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleUpdateLogic {
	return &RoleUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleUpdateLogic) RoleUpdate(in *sys.RoleUpdateReq) (*sys.RoleUpdateResp, error) {
	sysRole, err := l.svcCtx.Role.FindOne(l.ctx, in.RoleId)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.RoleIdErrorCode)
	}

	err = copier.Copy(sysRole, in)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	for _, menuId := range in.MenuIds {
		l.svcCtx.RoleMenu.DeleteByRoleId(l.ctx, in.RoleId)
		l.svcCtx.RoleMenu.Insert(l.ctx, &model.SysRoleMenu{RoleId: in.RoleId, MenuId: menuId})
	}

	return &sys.RoleUpdateResp{}, nil
}
