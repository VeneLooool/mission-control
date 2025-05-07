package model

import (
	"time"
)

type Coordinates []Coordinate

type Coordinate struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type MissionPlan struct {
	Coordinates Coordinates `db:"coordinates"`
}

type MissionType string

func (mt MissionType) String() string {
	return string(mt)
}

const (
	MissionTypePatrol   MissionType = "patrol"
	MissionTypeResearch MissionType = "research"
)

type MissionStatus string

func (ms MissionStatus) String() string {
	return string(ms)
}

const (
	MissionStatusCreated   MissionStatus = "created"
	MissionStatusScheduled MissionStatus = "scheduled"
	MissionStatusPending   MissionStatus = "pending"
	MissionStatusRunning   MissionStatus = "Running"
	MissionStatusAnalyse   MissionStatus = "analyse"
	MissionStatusCanceled  MissionStatus = "canceled"
	MissionStatusWarning   MissionStatus = "warning"
	MissionStatusFailed    MissionStatus = "failed"
	MissionStatusSuccess   MissionStatus = "success"
)

type Missions []Mission

type Mission struct {
	ID        uint64        `db:"id"`
	Name      string        `db:"name"`
	Type      MissionType   `db:"type"`
	Status    MissionStatus `db:"status"`
	CreatedBy string        `db:"created_by"`
	UpdatedBy string        `db:"updated_by"`

	CreatedAt time.Time `db:"created_at"`
	StartedAt time.Time `db:"started_at"`
	UpdatedAt time.Time `db:"updated_at"`

	FieldID uint64 `db:"field_id"`
	DroneID uint64 `db:"drone_id"`
}

func (m *Mission) SetStatus(status MissionStatus) {
	m.Status = status
}
func (m *Mission) SetStartedAt(startedAt time.Time) {
	m.StartedAt = startedAt
}

func (ms Missions) GetByDroneID(droneID uint64) (Mission, bool) {
	for _, m := range ms {
		if m.DroneID == droneID {
			return m, true
		}
	}
	return Mission{}, false
}
