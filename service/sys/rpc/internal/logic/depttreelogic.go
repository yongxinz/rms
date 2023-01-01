package logic

import (
	"context"

	"rms/service/sys/model"
	"rms/service/sys/rpc/internal/svc"
	"rms/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptTreeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptTreeLogic {
	return &DeptTreeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeptTreeLogic) DeptTree(in *sys.DeptTreeReq) (*sys.DeptTreeResp, error) {
	depts, err := l.svcCtx.Dept.FindAll(l.ctx, 1, 1000)

	if err != nil {
		logx.WithContext(l.ctx).Errorf("查询DeptTree信息失败,异常:%s", err.Error())
		return nil, err
	}

	var data []*sys.DeptTreeData
	for _, item := range depts {
		if item.ParentId.Int64 != 0 {
			continue
		}
		m := sys.DeptTreeData{
			Id:       item.DeptId,
			Label:    item.DeptName.String,
			Children: []*sys.DeptTreeData{},
		}
		deptInfo := deptTreeCall(depts, &m)
		data = append(data, deptInfo)
	}

	return &sys.DeptTreeResp{
		Data: data,
	}, nil
}

func deptTreeCall(deptList []*model.SysDept, dept *sys.DeptTreeData) *sys.DeptTreeData {
	list := deptList

	min := make([]*sys.DeptTreeData, 0)
	for j := 0; j < len(list); j++ {
		if dept.Id != list[j].ParentId.Int64 {
			continue
		}

		mi := sys.DeptTreeData{
			Id:       list[j].DeptId,
			Label:    list[j].DeptName.String,
			Children: []*sys.DeptTreeData{},
		}
		ms := deptTreeCall(deptList, &mi)
		min = append(min, ms)
	}
	dept.Children = min

	return dept
}
