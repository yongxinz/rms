package logic

import (
	"context"

	"rms/common/errorx"
	"rms/service/sys/model"
	"rms/service/sys/rpc/internal/svc"
	"rms/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleRetrieveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleRetrieveLogic {
	return &RoleRetrieveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleRetrieveLogic) RoleRetrieve(in *sys.RoleRetrieveReq) (*sys.RoleRetrieveResp, error) {
	res, err := l.svcCtx.Role.FindOne(l.ctx, in.RoleId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.NewDefaultError(errorx.RoleIdErrorCode)
		}
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	rolemenu, err := l.svcCtx.RoleMenu.FindMenuIds(l.ctx, in.RoleId)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	var menuIds []int64
	for _, item := range rolemenu {
		menuIds = append(menuIds, item.MenuId)
	}

	return &sys.RoleRetrieveResp{
		RoleId:   res.RoleId,
		RoleKey:  res.RoleKey,
		RoleName: res.RoleName,
		Sort:     res.Sort,
		Status:   res.Status,
		Remark:   res.Remark,
		MenuIds:  menuIds,
	}, nil
}
