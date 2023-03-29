package logic

import (
	"context"
	"encoding/json"

	"rms/common/errorx"
	"rms/service/sys/rpc/internal/svc"
	"rms/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictDataListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictDataListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictDataListLogic {
	return &DictDataListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictDataListLogic) DictDataList(in *sys.DictDataListReq) (*sys.DictDataListResp, error) {
	res, err := l.svcCtx.DictData.FindAll(l.ctx, in.DictType, in.Page, in.Size)
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("get dicttype failed, params: %s, error: %s", reqStr, err.Error())
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	count, _ := l.svcCtx.DictData.Count(l.ctx, in.DictType)

	var data []*sys.DictDataListData
	for _, item := range res {
		data = append(data, &sys.DictDataListData{
			DictId:    item.DictId,
			DictLabel: item.DictLabel,
			DictValue: item.DictValue,
			DictType:  item.DictType,
			DictSort:  item.DictSort,
			Status:    item.Status,
			Remark:    item.Remark,
			CreatedAt: item.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &sys.DictDataListResp{
		Count: count,
		Data:  data,
	}, nil
}
