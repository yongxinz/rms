package logic

import (
	"context"
	"encoding/json"

	"rms/common/errorx"
	"rms/service/sys/api/internal/svc"
	"rms/service/sys/api/internal/types"
	"rms/service/sys/rpc/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList(req *types.UserListReq) (resp *types.UserListResp, err error) {
	res, err := l.svcCtx.SysRpc.UserList(l.ctx, &sysclient.UserListReq{
		Page: req.PageIndex,
		Size: req.PageSize,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("get userlist error: %s, params: %s", err.Error(), string(data))
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	var list []*types.UserListData
	for _, user := range res.List {
		list = append(list, &types.UserListData{
			UserId:    user.UserId,
			UserName:  user.UserName,
			NickName:  user.NickName,
			Phone:     user.Phone,
			Status:    user.Status,
			Email:     user.Email,
			Salt:      user.Salt,
			Avatar:    user.Avatar,
			Sex:       user.Sex,
			Remark:    user.Remark,
			CreateBy:  user.CreateBy,
			CreatedAt: user.CreatedAt,
			UpdateBy:  user.UpdateBy,
			UpdatedAt: user.UpdatedAt,
			RoleId:    user.RoleId,
			DeptId:    user.DeptId,
			PostId:    user.PostId,
			Dept: types.DeptListData{
				DeptId:    user.Dept.DeptId,
				ParentId:  user.Dept.ParentId,
				DeptPath:  user.Dept.DeptPath,
				DeptName:  user.Dept.DeptName,
				Sort:      user.Dept.Sort,
				Leader:    user.Dept.Leader,
				Phone:     user.Dept.Phone,
				Email:     user.Dept.Email,
				Status:    user.Dept.Status,
				CreateBy:  user.Dept.CreateBy,
				CreatedAt: user.Dept.CreatedAt,
				UpdateBy:  user.Dept.UpdateBy,
				UpdatedAt: user.Dept.UpdatedAt,
			},
		})
	}

	resp = &types.UserListResp{
		List: list,
		Pagination: types.Pagination{
			PageIndex: req.PageIndex,
			PageSize:  req.PageSize,
			Count:     res.Count,
		},
	}

	return
}
