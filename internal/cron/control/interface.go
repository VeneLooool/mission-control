package control

import (
	"context"
	"github.com/VeneLooool/mission-control/internal/model"
)

type missionClient interface {
	GetMissionsInStatus(ctx context.Context, statuses []model.MissionStatus) ([]model.Mission, error)
	GetMissionPlanByMissionID(ctx context.Context, missionID uint64) (model.MissionPlan, error)
	UpdateMissionStatus(ctx context.Context, missionID uint64, status model.MissionStatus) (model.Mission, error)
}

type droneClient interface {
	StartDroneMission(ctx context.Context, droneID uint64, mission model.MissionPlan) error
}
