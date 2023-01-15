package logic

import (
	"context"

	"rms/common/cryptx"
	"rms/common/errorx"
	"rms/service/sys/rpc/internal/svc"
	"rms/service/sys/rpc/sys"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdatePwdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserUpdatePwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdatePwdLogic {
	return &UserUpdatePwdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserUpdatePwdLogic) UserUpdatePwd(in *sys.UserUpdatePwdReq) (*sys.UserUpdatePwdResp, error) {
	sysUser, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.UserIdErrorCode)
	}

	err = copier.Copy(sysUser, in)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	sysUser.Password = cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	err = l.svcCtx.UserModel.Update(l.ctx, sysUser)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &sys.UserUpdatePwdResp{}, nil
}
