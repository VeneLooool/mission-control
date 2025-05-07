package drone_events

import (
	"context"

	"github.com/VeneLooool/mission-control/internal/model"
)

type missionClient interface {
	GetMissionsInStatus(ctx context.Context, statuses []model.MissionStatus) ([]model.Mission, error)
	UpdateMissionStatus(ctx context.Context, missionID uint64, status model.MissionStatus) (model.Mission, error)
}

type analyticPublisher interface {
	Publish(ctx context.Context, task model.AnalyticTask) error
}
