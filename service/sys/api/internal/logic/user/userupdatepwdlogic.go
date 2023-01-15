package logic

import (
	"context"
	"encoding/json"

	"rms/common/errorx"
	"rms/service/sys/api/internal/svc"
	"rms/service/sys/api/internal/types"
	"rms/service/sys/rpc/sysclient"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdatePwdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserUpdatePwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdatePwdLogic {
	return &UserUpdatePwdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserUpdatePwdLogic) UserUpdatePwd(req *types.UserUpdatePwdReq) error {
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}
	req.UpdateBy = userId

	var sysUser sysclient.UserUpdatePwdReq
	err = copier.Copy(&sysUser, req)
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	_, err = l.svcCtx.SysRpc.UserUpdatePwd(l.ctx, &sysUser)
	if err != nil {
		return err
	}

	return nil
}
