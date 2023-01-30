package dept

import (
	"net/http"

	"rms/common/response"
	"rms/service/sys/api/internal/logic/dept"
	"rms/service/sys/api/internal/svc"
	"rms/service/sys/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeptAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeptAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := dept.NewDeptAddLogic(r.Context(), svcCtx)
		err := l.DeptAdd(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			response.Response(w, nil, err)
		}
	}
}
