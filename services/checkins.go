package services

import (
	"context"

	"github.com/shibafu528/utissue/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedCheckinsServer
	checkins map[uint64]*pb.Checkin
}

func NewCheckinsServer() pb.CheckinsServer {
	return &server{}
}

func (s *server) Create(ctx context.Context, request *pb.CheckinRequest) (*pb.CheckinResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented yet")
}
