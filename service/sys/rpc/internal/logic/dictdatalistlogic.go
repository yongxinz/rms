package logic

import (
	"context"
	"encoding/json"

	"rms/common/errorx"
	"rms/common/globalkey"
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

	count, _ := l.svcCtx.DictType.Count(l.ctx)

	var data []*sys.DictDataListData
	for _, item := range res {
		data = append(data, &sys.DictDataListData{
			DictCode:  item.DictCode,
			DictLabel: item.DictLabel.String,
			DictValue: item.DictValue.String,
			DictType:  item.DictType.String,
			DictSort:  item.DictSort.Int64,
			Status:    item.Status.Int64,
			Remark:    item.Remark.String,
			CreatedAt: item.CreatedAt.Time.Format(globalkey.SysDateFormat),
		})
	}

	return &sys.DictDataListResp{
		Count: count,
		Data:  data,
	}, nil
}
