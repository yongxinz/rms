package logic

import (
	"context"
	"encoding/json"

	"rms/service/sys/api/internal/svc"
	"rms/service/sys/api/internal/types"
	"rms/service/sys/rpc/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	userId, _ := l.ctx.Value("userId").(json.Number).Int64()
	res, err := l.svcCtx.SysRpc.UserInfo(l.ctx, &sysclient.UserInfoRequest{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}

	return &types.UserInfoResponse{
		Code: res.Code,
		Data: types.UserInfo{
			Avatar:       res.Avatar,
			Introduction: res.Introduction,
			Name:         res.Name,
			UserName:     res.UserName,
			UserId:       res.UserId,
			DeptId:       res.DeptId,
			Buttons:      res.Buttons,
			Roles:        res.Roles,
			Permissions:  res.Permissions,
		},
	}, nil
}
