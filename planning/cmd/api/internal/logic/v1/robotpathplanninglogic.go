package v1

import (
	"apiDemo/planning/cmd/api/internal/svc"
	"apiDemo/planning/cmd/api/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type RobotPathPlanningLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRobotPathPlanningLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RobotPathPlanningLogic {
	return &RobotPathPlanningLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RobotPathPlanningLogic) RobotPathPlanning(req *types.RobotPlanningReq) (resp *types.RobotPlanningRes, err error) {
	// todo: add your logic here and delete this line

	return
}
