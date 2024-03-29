package dept

import (
	"context"

	"rms/service/sys/api/internal/svc"
	"rms/service/sys/api/internal/types"
	"rms/service/sys/rpc/sysclient"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeptListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeptListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptListLogic {
	return &DeptListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeptListLogic) DeptList() (resp []*types.DeptListResp, err error) {
	res, err := l.svcCtx.SysRpc.DeptList(l.ctx, &sysclient.DeptListReq{})
	if err != nil {
		return nil, err
	}

	var data []*types.DeptListResp
	copier.Copy(&data, res.Data)
	resp = data

	return
}
