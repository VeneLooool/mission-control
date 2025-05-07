package control

import (
	"context"

	desc "github.com/VeneLooool/mission-control/internal/pb/api/v1/control"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) SaveAnalyticResults(ctx context.Context, req *desc.SaveAnalyticResults_Request) (*emptypb.Empty, error) {
	err := i.controlUC.SaveAnalyticResults(ctx, req.GetMissionId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}
