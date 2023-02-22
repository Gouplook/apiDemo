package v1

import (
	v1 "apiDemo/planning/cmd/api/internal/logic/v1"
	"apiDemo/planning/cmd/api/internal/svc"
	"apiDemo/planning/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func RobotPathPlanningHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RobotPlanningReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := v1.NewRobotPathPlanningLogic(r.Context(), svcCtx)
		resp, err := l.RobotPathPlanning(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
