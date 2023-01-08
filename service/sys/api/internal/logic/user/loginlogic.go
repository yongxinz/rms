package logic

import (
	"context"
	"encoding/json"
	"strings"

	"rms/common/errorx"
	"rms/service/sys/api/internal/svc"
	"rms/service/sys/api/internal/types"
	"rms/service/sys/rpc/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	if len(strings.TrimSpace(req.Username)) == 0 {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("username cannot nil, params: %s", reqStr)
		return nil, errorx.NewDefaultError(errorx.AccountErrorCode)
	}

	if len(strings.TrimSpace(req.Password)) == 0 {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("password cannot be nil, params: %s", reqStr)
		return nil, errorx.NewDefaultError(errorx.PasswordErrorCode)
	}

	res, err := l.svcCtx.SysRpc.Login(l.ctx, &sysclient.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.LoginResp{
		CurrentAuthority: res.CurrentAuthority,
		Expire:           res.Expire,
		Token:            res.Token,
	}

	return
}
