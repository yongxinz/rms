package handler

import (
	"net/http"

	"rms/common/response"
	logic "rms/service/sys/api/internal/logic/user"
	"rms/service/sys/api/internal/svc"
	"rms/service/sys/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserDeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserDeleteLogic(r.Context(), svcCtx)
		err := l.UserDelete(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			response.Response(w, nil, err)
		}
	}
}
