package menu

import (
	"net/http"

	"rms/common/response"
	"rms/service/sys/api/internal/logic/menu"
	"rms/service/sys/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func MenuListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := menu.NewMenuListLogic(r.Context(), svcCtx)
		resp, err := l.MenuList()
		if err != nil {
			httpx.Error(w, err)
		} else {
			response.Response(w, resp, err)
		}
	}
}
