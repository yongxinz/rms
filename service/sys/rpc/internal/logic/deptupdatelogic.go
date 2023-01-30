package logic

import (
	"context"
	"strconv"

	"rms/common/errorx"
	"rms/service/sys/model"
	"rms/service/sys/rpc/internal/svc"
	"rms/service/sys/rpc/sys"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeptUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptUpdateLogic {
	return &DeptUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeptUpdateLogic) DeptUpdate(in *sys.DeptUpdateReq) (*sys.DeptUpdateResp, error) {
	sysDept, err := l.svcCtx.Dept.FindOne(l.ctx, in.DeptId)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.UserIdErrorCode)
	}

	err = copier.Copy(sysDept, in)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	var deptPath string
	parentId := sysDept.ParentId.Int64
	if parentId == 0 {
		deptPath = "/0/" + strconv.FormatInt(sysDept.DeptId, 10) + "/"
	} else {
		var parentDept = new(model.SysDept)
		parentDept, _ = l.svcCtx.Dept.FindOne(l.ctx, sysDept.ParentId.Int64)
		deptPath = parentDept.DeptPath + strconv.FormatInt(sysDept.DeptId, 10) + "/"
	}
	sysDept.DeptPath = deptPath

	err = l.svcCtx.Dept.Update(l.ctx, sysDept)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &sys.DeptUpdateResp{}, nil
}
