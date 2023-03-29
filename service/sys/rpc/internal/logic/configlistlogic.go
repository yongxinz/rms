package logic

import (
	"context"
	"encoding/json"

	"rms/common/errorx"
	"rms/service/sys/rpc/internal/svc"
	"rms/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConfigListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigListLogic {
	return &ConfigListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ConfigListLogic) ConfigList(in *sys.ConfigListReq) (*sys.ConfigListResp, error) {
	res, err := l.svcCtx.SysConfigModel.FindAll1(l.ctx, in.Page, in.Size)
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("get sysconfig failed, params: %s, error: %s", reqStr, err.Error())
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	count, _ := l.svcCtx.SysConfigModel.Count(l.ctx)

	var data []*sys.ConfigListData
	for _, item := range res {
		data = append(data, &sys.ConfigListData{
			Id:          item.Id,
			ConfigName:  item.ConfigName,
			ConfigKey:   item.ConfigKey,
			ConfigValue: item.ConfigValue,
			ConfigType:  item.ConfigType,
			IsFrontend:  item.IsFrontend,
			Remark:      item.Remark,
			CreatedAt:   item.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &sys.ConfigListResp{
		Count: count,
		Data:  data,
	}, nil
}
