package handler

import (
	"net/http"

	"rms/common/response"
	logic "rms/service/sys/api/internal/logic/user"
	"rms/service/sys/api/internal/svc"
	"rms/service/sys/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserAddLogic(r.Context(), svcCtx)
		err := l.UserAdd(&req)
		response.Response(w, nil, err)
	}
}
