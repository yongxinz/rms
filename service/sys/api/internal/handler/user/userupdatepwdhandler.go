package handler

import (
	"net/http"

	"rms/common/response"
	logic "rms/service/sys/api/internal/logic/user"
	"rms/service/sys/api/internal/svc"
	"rms/service/sys/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserUpdatePwdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserUpdatePwdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserUpdatePwdLogic(r.Context(), svcCtx)
		err := l.UserUpdatePwd(&req)
		response.Response(w, nil, err)
	}
}
