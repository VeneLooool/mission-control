package missions_api

import (
	"github.com/VeneLooool/mission-control/internal/model"
	proto_model "github.com/VeneLooool/mission-control/internal/pb/missions-api/api/v1/model"
)

var (
	MissionTypeToProto = map[model.MissionType]proto_model.MissionType{
		model.MissionTypePatrol:   proto_model.MissionType_MISSION_TYPE_PATROL,
		model.MissionTypeResearch: proto_model.MissionType_MISSION_TYPE_RESEARCH,
	}
	MissionStatusToProto = map[model.MissionStatus]proto_model.MissionStatus{
		model.MissionStatusCreated:   proto_model.MissionStatus_MISSION_STATUS_CREATED,
		model.MissionStatusScheduled: proto_model.MissionStatus_MISSION_STATUS_SCHEDULED,
		model.MissionStatusPending:   proto_model.MissionStatus_MISSION_STATUS_PENDING,
		model.MissionStatusRunning:   proto_model.MissionStatus_MISSION_STATUS_RUNNING,
		model.MissionStatusAnalyse:   proto_model.MissionStatus_MISSION_STATUS_ANALYSE,
		model.MissionStatusCanceled:  proto_model.MissionStatus_MISSION_STATUS_CANCELED,
		model.MissionStatusWarning:   proto_model.MissionStatus_MISSION_STATUS_WARNING,
		model.MissionStatusFailed:    proto_model.MissionStatus_MISSION_STATUS_FAILED,
		model.MissionStatusSuccess:   proto_model.MissionStatus_MISSION_STATUS_SUCCESS,
	}

	MissionTypeToModel = map[proto_model.MissionType]model.MissionType{
		proto_model.MissionType_MISSION_TYPE_PATROL:   model.MissionTypePatrol,
		proto_model.MissionType_MISSION_TYPE_RESEARCH: model.MissionTypeResearch,
	}
	MissionStatusToModel = map[proto_model.MissionStatus]model.MissionStatus{
		proto_model.MissionStatus_MISSION_STATUS_CREATED:   model.MissionStatusCreated,
		proto_model.MissionStatus_MISSION_STATUS_SCHEDULED: model.MissionStatusScheduled,
		proto_model.MissionStatus_MISSION_STATUS_PENDING:   model.MissionStatusPending,
		proto_model.MissionStatus_MISSION_STATUS_RUNNING:   model.MissionStatusRunning,
		proto_model.MissionStatus_MISSION_STATUS_ANALYSE:   model.MissionStatusAnalyse,
		proto_model.MissionStatus_MISSION_STATUS_CANCELED:  model.MissionStatusCanceled,
		proto_model.MissionStatus_MISSION_STATUS_WARNING:   model.MissionStatusWarning,
		proto_model.MissionStatus_MISSION_STATUS_FAILED:    model.MissionStatusFailed,
		proto_model.MissionStatus_MISSION_STATUS_SUCCESS:   model.MissionStatusSuccess,
	}
)

func transformMissionStatusesToProto(statuses []model.MissionStatus) []proto_model.MissionStatus {
	result := make([]proto_model.MissionStatus, 0, len(statuses))
	for _, status := range statuses {
		result = append(result, MissionStatusToProto[status])
	}
	return result
}

func transformMissionsToModel(missions []*proto_model.Mission) []model.Mission {
	result := make([]model.Mission, 0, len(missions))
	for _, mission := range missions {
		result = append(result, transformMissionToModel(mission))
	}
	return result
}

func transformMissionToModel(mission *proto_model.Mission) model.Mission {
	if mission == nil {
		return model.Mission{}
	}

	return model.Mission{
		ID:        mission.GetId(),
		Name:      mission.GetName(),
		Type:      MissionTypeToModel[mission.GetType()],
		Status:    MissionStatusToModel[mission.GetStatus()],
		CreatedBy: mission.GetCreatedBy(),
		UpdatedBy: mission.GetUpdatedBy(),
		CreatedAt: mission.GetCreatedAt().AsTime(),
		StartedAt: mission.GetStartedAt().AsTime(),
		UpdatedAt: mission.GetUpdatedAt().AsTime(),
		FieldID:   mission.GetFieldId(),
		DroneID:   mission.GetDroneId(),
	}
}

func transformCoordinatesToModel(protoCoordinates []*proto_model.Coordinate) model.Coordinates {
	coordinates := make([]model.Coordinate, 0, len(protoCoordinates))
	for _, coordinate := range protoCoordinates {
		coordinates = append(coordinates, model.Coordinate{
			Latitude:  coordinate.Latitude,
			Longitude: coordinate.Longitude,
		})
	}
	return coordinates
}

func transformCoordinatesToProto(coordinates model.Coordinates) []*proto_model.Coordinate {
	protoCoordinates := make([]*proto_model.Coordinate, 0, len(coordinates))
	for _, coordinate := range coordinates {
		protoCoordinates = append(protoCoordinates, &proto_model.Coordinate{
			Latitude:  coordinate.Latitude,
			Longitude: coordinate.Longitude,
		})
	}
	return protoCoordinates
}

func transformMissionPlanToProto(plan model.MissionPlan) *proto_model.Plan {
	return &proto_model.Plan{
		Coordinates: transformCoordinatesToProto(plan.Coordinates),
	}
}

func transformMissionPlanToModel(protoPlan *proto_model.Plan) model.MissionPlan {
	return model.MissionPlan{
		Coordinates: transformCoordinatesToModel(protoPlan.GetCoordinates()),
	}
}
