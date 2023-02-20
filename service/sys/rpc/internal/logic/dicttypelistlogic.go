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

type DictTypeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictTypeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictTypeListLogic {
	return &DictTypeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictTypeListLogic) DictTypeList(in *sys.DictTypeListReq) (*sys.DictTypeListResp, error) {
	res, err := l.svcCtx.DictType.FindAll(l.ctx, in.Page, in.Size)
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("get dicttype failed, params: %s, error: %s", reqStr, err.Error())
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	count, _ := l.svcCtx.DictType.Count(l.ctx)

	var data []*sys.DictTypeListData
	for _, item := range res {
		data = append(data, &sys.DictTypeListData{
			DictId:    item.DictId,
			DictName:  item.DictName.String,
			DictType:  item.DictType.String,
			Status:    item.Status.Int64,
			Remark:    item.Remark.String,
			CreatedAt: item.CreatedAt.Time.Format(globalkey.SysDateFormat),
		})
	}

	return &sys.DictTypeListResp{
		Count: count,
		Data:  data,
	}, nil
}
