package dept

import (
	"net/http"

	"rms/common/response"
	"rms/service/sys/api/internal/logic/dept"
	"rms/service/sys/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeptListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := dept.NewDeptListLogic(r.Context(), svcCtx)
		resp, err := l.DeptList()
		if err != nil {
			httpx.Error(w, err)
		} else {
			response.Response(w, resp, err)
		}
	}
}
