package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"rms/service/sys/api/internal/logic"
	"rms/service/sys/api/internal/svc"
)

func DeepTreeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewDeepTreeLogic(r.Context(), svcCtx)
		resp, err := l.DeepTree()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
