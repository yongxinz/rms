package logic

import (
	"context"

	"rms/common/errorx"
	"rms/service/sys/model"
	"rms/service/sys/rpc/internal/svc"
	"rms/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptRetrieveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptRetrieveLogic {
	return &DeptRetrieveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeptRetrieveLogic) DeptRetrieve(in *sys.DeptRetrieveReq) (*sys.DeptRetrieveResp, error) {
	res, err := l.svcCtx.Dept.FindOne(l.ctx, in.DeptId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.NewDefaultError(errorx.PostIdErrorCode)
		}
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &sys.DeptRetrieveResp{
		DeptId:   res.DeptId,
		DeptName: res.DeptName.String,
		Phone:    res.Phone.String,
		Email:    res.Email.String,
		Leader:   res.Leader.String,
		Status:   res.Status.Int64,
		Sort:     res.Sort.Int64,
		ParentId: res.ParentId.Int64,
	}, nil
}
