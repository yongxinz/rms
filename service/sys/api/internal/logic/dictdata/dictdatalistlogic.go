package dictdata

import (
	"context"
	"encoding/json"

	"rms/common/errorx"
	"rms/service/sys/api/internal/svc"
	"rms/service/sys/api/internal/types"
	"rms/service/sys/rpc/sysclient"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type DictDataListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDictDataListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictDataListLogic {
	return &DictDataListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DictDataListLogic) DictDataList(req *types.DictDataListReq) (resp *types.DictDataListResp, err error) {
	res, err := l.svcCtx.SysRpc.DictDataList(l.ctx, &sysclient.DictDataListReq{
		Page:     req.PageIndex,
		Size:     req.PageSize,
		DictType: req.DictType,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("get dictdata error: %s, params: %s", err.Error(), string(data))
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	var dictdata types.DictDataListData
	var list []types.DictDataListData
	for _, v := range res.Data {
		copier.Copy(&dictdata, &v)
		list = append(list, dictdata)
	}

	resp = &types.DictDataListResp{
		List: list,
		Pagination: types.Pagination{
			PageIndex: req.PageIndex,
			PageSize:  req.PageSize,
			Count:     res.Count,
		},
	}

	return
}
