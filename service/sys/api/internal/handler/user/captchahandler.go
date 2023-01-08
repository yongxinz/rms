package handler

import (
	"net/http"

	"rms/common/response"
	logic "rms/service/sys/api/internal/logic/user"
	"rms/service/sys/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.Captcha()
		if err != nil {
			httpx.Error(w, err)
		} else {
			response.Response(w, resp, err)
		}
	}
}
