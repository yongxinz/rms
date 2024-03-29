package menu

import (
	"net/http"

	"rms/common/response"
	logic "rms/service/sys/api/internal/logic/menu"
	"rms/service/sys/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func MenuRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewMenuRoleLogic(r.Context(), svcCtx)
		resp, err := l.MenuRole()
		if err != nil {
			httpx.Error(w, err)
		} else {
			response.Response(w, resp, err)
		}
	}
}
