package logic

import (
	"context"

	"rms/common/errorx"
	"rms/service/sys/model"
	"rms/service/sys/rpc/internal/svc"
	"rms/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuRetrieveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuRetrieveLogic {
	return &MenuRetrieveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuRetrieveLogic) MenuRetrieve(in *sys.MenuRetrieveReq) (*sys.MenuRetrieveResp, error) {
	res, err := l.svcCtx.Menu.FindOne(l.ctx, in.MenuId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.NewDefaultError(errorx.PostIdErrorCode)
		}
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &sys.MenuRetrieveResp{
		MenuId:     res.MenuId,
		MenuName:   res.MenuName.String,
		Title:      res.Title.String,
		MenuType:   res.MenuName.String,
		Visible:    res.Visible.String,
		NoCache:    res.NoCache.Int64,
		Component:  res.Component.String,
		Path:       res.Path.String,
		Paths:      res.Paths,
		Breadcrumb: res.Breadcrumb.String,
		Permission: res.Permission.String,
		Icon:       res.Icon.String,
		IsFrame:    res.IsFrame,
		Action:     res.Action.String,
		Sort:       res.Sort.Int64,
		ParentId:   res.ParentId.Int64,
	}, nil
}
