package missions_api

import (
	"context"
	"fmt"

	"github.com/VeneLooool/mission-control/internal/config"
	"github.com/VeneLooool/mission-control/internal/model"
	"github.com/VeneLooool/mission-control/internal/pb/missions-api/api/v1/missions"
	"github.com/VeneLooool/mission-control/internal/pb/missions-api/api/v1/planner"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	missionApi missions.MissionsClient
	plannerApi planner.PlannerClient
}

func New(ctx context.Context, cfg *config.MissionApiClientConfig) (*Client, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", cfg.Host, cfg.GrpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, errors.Wrap(err, "grpc.NewClient()")
	}

	return &Client{
		missionApi: missions.NewMissionsClient(conn),
		plannerApi: planner.NewPlannerClient(conn),
	}, nil
}

func (c *Client) GetMissionByID(ctx context.Context, missionID uint64) (model.Mission, error) {
	resp, err := c.missionApi.GetMissionByID(ctx, &missions.GetMissionByID_Request{Id: missionID})
	if err != nil {
		return model.Mission{}, err
	}
	return transformMissionToModel(resp.GetMission()), nil
}

func (c *Client) GetMissionsInStatus(ctx context.Context, statuses []model.MissionStatus) ([]model.Mission, error) {
	resp, err := c.missionApi.GetMissionsInStatuses(ctx, &missions.GetMissionsInStatuses_Request{
		Statuses: transformMissionStatusesToProto(statuses),
	})
	if err != nil {
		return nil, err
	}
	return transformMissionsToModel(resp.GetMissions()), nil
}

func (c *Client) GetMissionPlanByMissionID(ctx context.Context, missionID uint64) (model.MissionPlan, error) {
	resp, err := c.plannerApi.GetMissionPlanByMissionID(ctx, &planner.GetMissionPlanByMissionID_Request{MissionId: missionID})
	if err != nil {
		return model.MissionPlan{}, err
	}
	return transformMissionPlanToModel(resp.GetPlan()), nil
}

func (c *Client) UpdateMissionStatus(ctx context.Context, missionID uint64, status model.MissionStatus) (model.Mission, error) {
	resp, err := c.missionApi.UpdateMissionStatus(ctx, &missions.UpdateMissionStatus_Request{
		Id:     missionID,
		Status: MissionStatusToProto[status],
	})
	if err != nil {
		return model.Mission{}, err
	}
	return transformMissionToModel(resp.GetMission()), nil
}
