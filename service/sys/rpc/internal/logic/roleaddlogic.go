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

type RoleAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleAddLogic {
	return &RoleAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleAddLogic) RoleAdd(in *sys.RoleAddReq) (*sys.RoleAddResp, error) {
	var SysRole = new(model.SysRole)
	err := copier.Copy(SysRole, in)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	res, err := l.svcCtx.Role.Insert(l.ctx, SysRole)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	roleId, _ := res.LastInsertId()
	for _, menuId := range in.MenuIds {
		l.svcCtx.RoleMenu.DeleteByRoleId(l.ctx, roleId)
		l.svcCtx.RoleMenu.Insert(l.ctx, &model.SysRoleMenu{RoleId: roleId, MenuId: menuId})
	}

	return &sys.RoleAddResp{}, nil
}
