package logic

import (
	"context"

	"rms/common/cryptx"
	"rms/common/errorx"
	"rms/common/utils"
	"rms/service/sys/model"
	"rms/service/sys/rpc/internal/svc"
	"rms/service/sys/rpc/sys"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddLogic {
	return &UserAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserAddLogic) UserAdd(in *sys.UserAddReq) (*sys.UserAddResp, error) {
	_, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err == model.ErrNotFound {

		_, err := l.svcCtx.Dept.FindOne(l.ctx, in.DeptId)
		if err != nil {
			return nil, errorx.NewDefaultError(errorx.DeptIdErrorCode)
		}

		_, err = l.svcCtx.Post.FindOne(l.ctx, in.PostId)
		if err != nil {
			return nil, errorx.NewDefaultError(errorx.PostIdErrorCode)
		}

		var sysUser = new(model.SysUser)
		err = copier.Copy(sysUser, in)
		if err != nil {
			return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
		}

		sysUser.Avatar = utils.AvatarUrl()
		sysUser.Password = cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)

		_, err = l.svcCtx.UserModel.Insert(l.ctx, sysUser)
		if err != nil {
			return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
		}

		return &sys.UserAddResp{}, nil
	} else {
		return nil, errorx.NewDefaultError(errorx.AddUserErrorCode)
	}
}
