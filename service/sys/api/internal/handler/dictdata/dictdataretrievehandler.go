package dictdata

import (
	"net/http"

	"rms/common/response"
	"rms/service/sys/api/internal/logic/dictdata"
	"rms/service/sys/api/internal/svc"
	"rms/service/sys/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DictDataRetrieveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DictDataRetrieveReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := dictdata.NewDictDataRetrieveLogic(r.Context(), svcCtx)
		resp, err := l.DictDataRetrieve(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			response.Response(w, resp, err)
		}
	}
}
