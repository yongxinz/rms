package role

import (
	"net/http"

	"rms/common/response"
	"rms/service/sys/api/internal/logic/role"
	"rms/service/sys/api/internal/svc"
	"rms/service/sys/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RoleDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleDeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := role.NewRoleDeleteLogic(r.Context(), svcCtx)
		err := l.RoleDelete(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			response.Response(w, nil, err)
		}
	}
}
