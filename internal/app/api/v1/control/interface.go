package control

import (
	"context"
)

type controlUC interface {
	SaveAnalyticResults(ctx context.Context, missionID uint64) error
}
