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

type DeptAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptAddLogic {
	return &DeptAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeptAddLogic) DeptAdd(in *sys.DeptAddReq) (*sys.DeptAddResp, error) {
	var SysDept = new(model.SysDept)
	err := copier.Copy(SysDept, in)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	res, err := l.svcCtx.Dept.Insert(l.ctx, SysDept)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	deptId, _ := res.LastInsertId()
	SysDept, _ = l.svcCtx.Dept.FindOne(l.ctx, deptId)
	parentId := SysDept.ParentId.Int64

	var deptPath string
	if parentId == 0 {
		deptPath = "/0/" + strconv.FormatInt(SysDept.DeptId, 10) + "/"
	} else {
		var parentDept = new(model.SysDept)
		parentDept, _ = l.svcCtx.Dept.FindOne(l.ctx, SysDept.ParentId.Int64)
		deptPath = parentDept.DeptPath + strconv.FormatInt(SysDept.DeptId, 10) + "/"
	}

	SysDept.DeptPath = deptPath
	err = l.svcCtx.Dept.Update(l.ctx, SysDept)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &sys.DeptAddResp{}, nil
}
