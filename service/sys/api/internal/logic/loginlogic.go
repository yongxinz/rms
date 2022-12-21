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

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("用户名或密码不能为空,请求信息失败,参数:%s", reqStr)
		return nil, errorx.NewDefaultError("用户名或密码不能为空")
	}

	res, err := l.svcCtx.SysRpc.Login(l.ctx, &sysclient.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{
		Code:             200,
		Success:          true,
		CurrentAuthority: res.CurrentAuthority,
		Expire:           res.Expire,
		Token:            res.Token,
	}, nil
}
