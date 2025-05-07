package drone_events

import (
	"context"
	"log"

	"github.com/VeneLooool/mission-control/internal/model"
)

type Handler struct {
	missionClient     missionClient
	analyticPublisher analyticPublisher
}

func New(missionClient missionClient, analyticPublisher analyticPublisher) *Handler {
	return &Handler{
		missionClient:     missionClient,
		analyticPublisher: analyticPublisher,
	}
}

func (h *Handler) Handle(ctx context.Context, event model.Event) error {
	if event.EventType != model.EventTypeDroneChangeStatus {
		return nil
	}

	switch event.Drone.Status {
	case model.DroneStatusInMission:
		return h.HandleDroneInMission(ctx, event)
	case model.DroneStatusCharging:
		return h.HandleDroneCharging(ctx, event)
	}

	return nil
}

func (h *Handler) HandleDroneInMission(ctx context.Context, event model.Event) error {
	missions, err := h.missionClient.GetMissionsInStatus(ctx, []model.MissionStatus{model.MissionStatusPending})
	if err != nil {
		return err
	}
	mission, ok := model.Missions(missions).GetByDroneID(event.Drone.ID)
	if !ok {
		log.Printf("mission-control: not found mission for drone id: %d, drone event: %v", event.Drone.ID, event)
		return nil
	}

	if _, err = h.missionClient.UpdateMissionStatus(ctx, mission.ID, model.MissionStatusRunning); err != nil {
		return err
	}
	log.Printf("mission-control: updated mission status to running; id: %d", mission.ID)

	return nil
}

func (h *Handler) HandleDroneCharging(ctx context.Context, event model.Event) error {
	missions, err := h.missionClient.GetMissionsInStatus(ctx, []model.MissionStatus{model.MissionStatusRunning})
	if err != nil {
		return err
	}
	mission, ok := model.Missions(missions).GetByDroneID(event.Drone.ID)
	if !ok {
		log.Printf("mission-control: not found mission for drone id: %d, drone event: %v", event.Drone.ID, event)
		return nil
	}

	if err = h.analyticPublisher.Publish(ctx, model.GetAnalyticTaskByMission(mission)); err != nil {
		return err
	}
	log.Printf("mission-control: publish analytic task to analytic event; mission id: %d", mission.ID)

	if _, err = h.missionClient.UpdateMissionStatus(ctx, mission.ID, model.MissionStatusAnalyse); err != nil {
		return err
	}
	log.Printf("mission-control: updated mission status to analyse; mission id: %d", mission.ID)

	return nil
}
