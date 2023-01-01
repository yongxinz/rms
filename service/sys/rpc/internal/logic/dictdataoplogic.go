package logic

import (
	"context"
	"encoding/json"

	"rms/service/sys/rpc/internal/svc"
	"rms/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictDataOpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictDataOpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictDataOpLogic {
	return &DictDataOpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictDataOpLogic) DictDataOp(in *sys.DictDataOpReq) (*sys.DictDataOpResp, error) {
	res, err := l.svcCtx.DictData.FindByDictType(l.ctx, in.DictType)
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询DictDataOp信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var data []*sys.DictDataOp
	for _, item := range res {
		data = append(data, &sys.DictDataOp{
			Label: item.DictLabel.String,
			Value: item.DictValue.String,
		})
	}

	return &sys.DictDataOpResp{
		Data: data,
	}, nil
}
