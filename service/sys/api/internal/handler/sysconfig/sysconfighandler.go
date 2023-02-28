package sysconfig

import (
	"net/http"

	"rms/common/response"
	logic "rms/service/sys/api/internal/logic/sysconfig"
	"rms/service/sys/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SysConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewSysConfigLogic(r.Context(), svcCtx)
		resp, err := l.SysConfig()
		if err != nil {
			httpx.Error(w, err)
		} else {
			response.Response(w, resp, err)
		}
	}
}
