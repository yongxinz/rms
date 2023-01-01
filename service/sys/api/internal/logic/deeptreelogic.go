package logic

import (
	"context"

	"rms/service/sys/api/internal/svc"
	"rms/service/sys/api/internal/types"
	"rms/service/sys/rpc/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeepTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeepTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeepTreeLogic {
	return &DeepTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeepTreeLogic) DeepTree() (resp *types.DeepTreeResp, err error) {
	res, err := l.svcCtx.SysRpc.DeepTree(l.ctx, &sysclient.DeepTreeReq{})
	if err != nil {
		return nil, err
	}

	var data []*types.DeptTreeData
	for _, item := range res.Data {
		children := make([]types.DeptTreeData, 0)
		for _, dept := range item.Children {
			children = append(children, types.DeptTreeData{
				Id:    dept.Id,
				Label: dept.Label,
			})
		}
		data = append(data, &types.DeptTreeData{
			Id:       item.Id,
			Label:    item.Label,
			Children: children,
		})
	}
	resp = &types.DeepTreeResp{
		Code: 200,
		Data: data,
	}

	return
}
