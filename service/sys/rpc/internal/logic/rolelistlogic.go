package logic

import (
	"context"
	"encoding/json"

	"rms/common/errorx"
	"rms/common/globalkey"
	"rms/service/sys/rpc/internal/svc"
	"rms/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleListLogic) RoleList(in *sys.RoleListReq) (*sys.RoleListResp, error) {
	res, err := l.svcCtx.Role.FindAll(l.ctx, in.Page, in.Size)
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("get rolelist failed, params: %s, error: %s", reqStr, err.Error())
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	count, _ := l.svcCtx.Role.Count(l.ctx)

	var data []*sys.RoleListData
	for _, item := range res {
		data = append(data, &sys.RoleListData{
			RoleId:    item.RoleId,
			RoleName:  item.RoleName.String,
			RoleKey:   item.RoleKey.String,
			RoleSort:  item.RoleSort.Int64,
			Flag:      item.Flag.String,
			Remark:    item.Remark.String,
			Admin:     item.Admin.Int64,
			DataScope: item.DataScope.String,
			Status:    item.Status.String,
			CreateBy:  item.CreateBy.Int64,
			UpdateBy:  item.UpdateBy.Int64,
			CreatedAt: item.CreatedAt.Time.Format(globalkey.SysDateFormat),
			UpdatedAt: item.UpdatedAt.Time.Format(globalkey.SysDateFormat),
		})
	}

	return &sys.RoleListResp{
		Count: count,
		Data:  data,
	}, nil
}
