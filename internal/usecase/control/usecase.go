package control

import (
	"context"
	"log"

	"github.com/VeneLooool/mission-control/internal/model"
)

type UseCases struct {
	missionClient missionClient
}

func New(missionClient missionClient) *UseCases {
	return &UseCases{
		missionClient: missionClient,
	}
}

func (u *UseCases) SaveAnalyticResults(ctx context.Context, missionID uint64) error {
	mission, err := u.missionClient.GetMissionByID(ctx, missionID)
	if err != nil {
		return err
	}

	if _, err = u.missionClient.UpdateMissionStatus(ctx, mission.ID, model.MissionStatusSuccess); err != nil {
		return err
	}
	log.Printf("Mission-control: successfully saved analytic results for mission %d", missionID)

	return nil
}
