package control

import (
	"context"
	"log"
	"time"

	"github.com/VeneLooool/mission-control/internal/model"
)

type Cron struct {
	missionClient missionClient
	droneClient   droneClient
}

func New(missionClient missionClient, droneClient droneClient) *Cron {
	return &Cron{
		missionClient: missionClient,
		droneClient:   droneClient,
	}
}

func (c *Cron) Do(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := c.do(ctx); err != nil {
				log.Printf("failed to do: %s", err.Error())
			}
			time.Sleep(time.Second * 10)
		}
	}
}

func (c *Cron) do(ctx context.Context) error {
	log.Println("mission-control: start processing")

	now := time.Now().UTC()

	missions, err := c.missionClient.GetMissionsInStatus(ctx, []model.MissionStatus{
		model.MissionStatusCreated,
		model.MissionStatusScheduled,
	})
	if err != nil {
		return err
	}
	log.Printf("mission-control: mission count is - %d", len(missions))

	for _, mission := range missions {
		log.Printf("mission-control: processing mission id: %d, mission status: %s", mission.ID, mission.Status)

		if mission.Status == model.MissionStatusScheduled && now.Before(mission.StartedAt) {
			log.Printf("mission-control: skip mission id: %d; planned start at: %s", mission.ID, mission.StartedAt.String())
			continue
		}

		missionPlan, err := c.missionClient.GetMissionPlanByMissionID(ctx, mission.ID)
		if err != nil {
			log.Printf("mission-control: processing mission id: %d, failed to get mission plan: %s", mission.ID, err.Error())
			continue
		}

		if err = c.droneClient.StartDroneMission(ctx, mission.DroneID, missionPlan); err != nil {
			log.Printf("mission-control: processing mission id: %d, failed to startd drone: %s", mission.ID, err.Error())
			continue
		}
		log.Printf("mission-control: processing mission id: %d, drone has started: %d", mission.ID, mission.DroneID)
		
		if _, err = c.missionClient.UpdateMissionStatus(ctx, mission.ID, model.MissionStatusPending); err != nil {
			log.Printf("mission-control: processing mission id: %d, failed to update mission status: %s", mission.ID, err.Error())
		}

		log.Printf("mission-control: finish processing mission id: %d, mission status: %s", mission.ID, mission.Status)
	}

	log.Println("mission-control: finish processing")
	return nil
}
