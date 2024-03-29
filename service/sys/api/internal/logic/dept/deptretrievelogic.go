package dept

import (
	"context"

	"rms/service/sys/api/internal/svc"
	"rms/service/sys/api/internal/types"
	"rms/service/sys/rpc/sysclient"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeptRetrieveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeptRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptRetrieveLogic {
	return &DeptRetrieveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeptRetrieveLogic) DeptRetrieve(req *types.DeptRetrieveReq) (resp *types.DeptRetrieveResp, err error) {
	res, err := l.svcCtx.SysRpc.DeptRetrieve(l.ctx, &sysclient.DeptRetrieveReq{
		DeptId: req.DeptId,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.DeptRetrieveResp{}
	err = copier.Copy(resp, res)

	return
}
