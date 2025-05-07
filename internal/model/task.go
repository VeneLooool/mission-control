package model

import "strconv"

type AnalyticTaskType string

func (att AnalyticTaskType) String() string {
	return string(att)
}

const (
	AnalyticTaskTypeAnalysePatrol   AnalyticTaskType = "task_analyse_patrol"
	AnalyticTaskTypeAnalyseResearch AnalyticTaskType = "task_analyse_research"
)

type AnalyticTask struct {
	Type      AnalyticTaskType `json:"task_type"`
	MissionID uint64           `json:"mission_id"`
}

func (at AnalyticTask) GetEventKey() string {
	return strconv.FormatUint(at.MissionID, 10)
}

func GetAnalyticTaskByMission(mission Mission) AnalyticTask {
	var taskType AnalyticTaskType

	switch mission.Type {
	case MissionTypeResearch:
		taskType = AnalyticTaskTypeAnalyseResearch
	default:
		taskType = AnalyticTaskTypeAnalysePatrol
	}

	return AnalyticTask{
		Type:      taskType,
		MissionID: mission.ID,
	}
}
