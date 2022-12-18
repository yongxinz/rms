package logic

import (
	"context"
	"fmt"

	"rms/common/cryptx"
	"rms/service/sys/model"
	"rms/service/sys/rpc/internal/svc"
	"rms/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *sys.LoginRequest) (*sys.LoginResponse, error) {
	// 查询用户是否存在
	res, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	// 判断密码是否正确
	password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	fmt.Println(in.Password)
	fmt.Println(password)
	fmt.Println(res.Password)
	if password != res.Password {
		return nil, status.Error(100, "密码错误")
	}

	return &sys.LoginResponse{
		UserId:   res.UserId,
		Username: res.Username,
	}, nil
}
