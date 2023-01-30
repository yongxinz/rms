package dept

import (
	"net/http"

	"rms/common/response"
	"rms/service/sys/api/internal/logic/dept"
	"rms/service/sys/api/internal/svc"
	"rms/service/sys/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeptRetrieveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeptRetrieveReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := dept.NewDeptRetrieveLogic(r.Context(), svcCtx)
		resp, err := l.DeptRetrieve(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			response.Response(w, resp, err)
		}
	}
}
