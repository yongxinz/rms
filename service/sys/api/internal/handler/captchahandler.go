package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"rms/service/sys/api/internal/logic"
	"rms/service/sys/api/internal/svc"
)

func CaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.Captcha()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
