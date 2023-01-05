package logic

import (
	"context"
	"encoding/json"

	"rms/common/errorx"
	"rms/service/sys/rpc/internal/svc"
	"rms/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserListLogic) UserList(in *sys.UserListReq) (*sys.UserListResp, error) {
	res, err := l.svcCtx.UserModel.FindAll1(l.ctx, in.Page, in.Size)
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("get userlist failed, params: %s, error: %s", reqStr, err.Error())
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	count, _ := l.svcCtx.UserModel.Count(l.ctx)

	var data []*sys.UserListData
	for _, user := range res {
		data = append(data, &sys.UserListData{
			UserId:    user.UserId,
			UserName:  user.Username,
			NickName:  user.NickName.String,
			Phone:     user.SysUser.Phone.String,
			Status:    user.SysUser.Status.String,
			Email:     user.SysUser.Email.String,
			Salt:      user.Salt.String,
			Avatar:    user.Avatar.String,
			Sex:       user.Sex.String,
			Remark:    user.Remark.String,
			CreateBy:  user.SysUser.CreateBy.Int64,
			CreatedAt: user.SysUser.CreatedAt.Time.Format("2006-01-02 15:04:05"),
			UpdateBy:  user.SysUser.UpdateBy.Int64,
			UpdatedAt: user.SysUser.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
			RoleId:    user.RoleId.Int64,
			DeptId:    user.SysUser.DeptId.Int64,
			PostId:    user.PostId.Int64,
			Dept: &sys.DeptListData{
				DeptId:    user.SysDept.DeptId,
				ParentId:  user.ParentId.Int64,
				DeptPath:  user.DeptPath.String,
				DeptName:  user.DeptName.String,
				Sort:      user.Sort.Int64,
				Leader:    user.Leader.String,
				Phone:     user.SysDept.Phone.String,
				Email:     user.SysDept.Email.String,
				Status:    user.SysDept.Status.Int64,
				CreateBy:  user.SysDept.CreateBy.Int64,
				CreatedAt: user.SysDept.CreatedAt.Time.Format("2006-01-02 15:04:05"),
				UpdateBy:  user.SysDept.UpdateBy.Int64,
				UpdatedAt: user.SysDept.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
			},
		})
	}

	return &sys.UserListResp{
		Count: count,
		List:  data,
	}, nil
}
