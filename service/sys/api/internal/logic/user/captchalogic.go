package logic

import (
	"context"

	"rms/common/captcha"
	"rms/service/sys/api/internal/svc"
	"rms/service/sys/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CaptchaLogic {
	return &CaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CaptchaLogic) Captcha() (resp *types.CaptchaResp, err error) {
	id, b64s, err := captcha.DriverDigitFunc()
	if err != nil {
		return nil, err
	}

	resp = &types.CaptchaResp{
		Id:   id,
		Data: b64s,
	}

	return
}
