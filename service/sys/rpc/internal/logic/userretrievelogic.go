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
		NickName: res.NickName.String,
		Phone:    res.Phone.String,
		Status:   res.Status.String,
		Email:    res.Email.String,
		Sex:      res.Sex.String,
		Remark:   res.Remark.String,
		RoleId:   res.RoleId.Int64,
		DeptId:   res.DeptId.Int64,
		PostId:   res.PostId.Int64,
	}, nil
}
