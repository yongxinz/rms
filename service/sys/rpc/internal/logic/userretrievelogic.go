package logic

import (
	"context"

	"rms/common/errorx"
	"rms/service/sys/model"
	"rms/service/sys/rpc/internal/svc"
	"rms/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRetrieveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRetrieveLogic {
	return &UserRetrieveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserRetrieveLogic) UserRetrieve(in *sys.UserInfoReq) (*sys.UserRetrieveResp, error) {
	res, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.NewDefaultError(errorx.UserIdErrorCode)
		}
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &sys.UserRetrieveResp{
		UserId:   res.UserId,
		Username: res.Username,
		Phone:    res.Phone,
		Status:   res.Status,
		Email:    res.Email,
		Sex:      res.Sex,
		Remark:   res.Remark,
		RoleId:   res.RoleId,
		DeptId:   res.DeptId,
		PostId:   res.PostId,
	}, nil
}
