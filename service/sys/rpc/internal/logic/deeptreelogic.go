package logic

import (
	"context"

	"rms/service/sys/model"
	"rms/service/sys/rpc/internal/svc"
	"rms/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeepTreeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeepTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeepTreeLogic {
	return &DeepTreeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeepTreeLogic) DeepTree(in *sys.DeepTreeReq) (*sys.DeepTreeResp, error) {
	depts, err := l.svcCtx.Dept.FindAll(l.ctx, 1, 1000)

	if err != nil {
		logx.WithContext(l.ctx).Errorf("查询DeepTree信息失败,异常:%s", err.Error())
		return nil, err
	}

	var data []*sys.DeepTreeData
	for _, item := range depts {
		if item.ParentId.Int64 != 0 {
			continue
		}
		m := sys.DeepTreeData{
			Id:       item.DeptId,
			Label:    item.DeptName.String,
			Children: []*sys.DeepTreeData{},
		}
		deptInfo := deptTreeCall(depts, &m)
		data = append(data, deptInfo)
	}

	return &sys.DeepTreeResp{
		Data: data,
	}, nil
}

func deptTreeCall(deptList []*model.SysDept, dept *sys.DeepTreeData) *sys.DeepTreeData {
	list := deptList

	min := make([]*sys.DeepTreeData, 0)
	for j := 0; j < len(list); j++ {
		if dept.Id != list[j].ParentId.Int64 {
			continue
		}

		mi := sys.DeepTreeData{
			Id:       list[j].DeptId,
			Label:    list[j].DeptName.String,
			Children: []*sys.DeepTreeData{},
		}
		ms := deptTreeCall(deptList, &mi)
		min = append(min, ms)
	}
	dept.Children = min

	return dept
}
