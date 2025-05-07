package control

import desc "github.com/VeneLooool/mission-control/internal/pb/api/v1/control"

// Implementation is a Service implementation
type Implementation struct {
	desc.UnimplementedMissionControlServer

	controlUC controlUC
}

// NewService return new instance of Implementation.
func NewService(controlUC controlUC) *Implementation {
	return &Implementation{
		controlUC: controlUC,
	}
}
