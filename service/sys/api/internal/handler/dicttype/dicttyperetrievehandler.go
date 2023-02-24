package dicttype

import (
	"net/http"

	"rms/common/response"
	"rms/service/sys/api/internal/logic/dicttype"
	"rms/service/sys/api/internal/svc"
	"rms/service/sys/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DictTypeRetrieveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DictTypeRetrieveReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := dicttype.NewDictTypeRetrieveLogic(r.Context(), svcCtx)
		resp, err := l.DictTypeRetrieve(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			response.Response(w, resp, err)
		}
	}
}
