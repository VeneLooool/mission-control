package control

import (
	"context"

	"github.com/VeneLooool/mission-control/internal/model"
)

type missionClient interface {
	GetMissionByID(ctx context.Context, missionID uint64) (model.Mission, error)
	UpdateMissionStatus(ctx context.Context, missionID uint64, status model.MissionStatus) (model.Mission, error)
}
