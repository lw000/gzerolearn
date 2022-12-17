package handler

import (
	"net/http"

	"demo/gzerolearn/internal/logic"
	"demo/gzerolearn/internal/svc"
	"demo/gzerolearn/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func HealthcheckHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CheckingRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewHealthcheckLogic(r.Context(), svcCtx)
		resp, err := l.Healthcheck(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
