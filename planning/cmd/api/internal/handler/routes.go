// Code generated by goctl. DO NOT EDIT.
package handler

import (
	v1 "apiDemo/planning/cmd/api/internal/handler/v1"
	"apiDemo/planning/cmd/api/internal/svc"
	"net/http"
	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/robotPathPlanning",
				Handler: v1.RobotPathPlanningHandler(serverCtx),
			},
		},
		rest.WithPrefix("/planning/v1/path"),
	)
}