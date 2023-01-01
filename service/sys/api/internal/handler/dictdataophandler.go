package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"rms/service/sys/api/internal/logic"
	"rms/service/sys/api/internal/svc"
	"rms/service/sys/api/internal/types"
)

func DictDataOpHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DictDataOpReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDictDataOpLogic(r.Context(), svcCtx)
		resp, err := l.DictDataOp(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
