package handler

import (
	"context"

	"github.com/d1agnoze/kitchen/services/common/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CommonHandler struct {
	pb.UnimplementedCommonServiceServer
}

func (x *CommonHandler) Ping(context.Context, *emptypb.Empty) (*pb.PingResponse, error) {
	return &pb.PingResponse{Message: "PONG"}, nil
}
func (x *CommonHandler) Puck(context.Context, *pb.MessageDto) (*pb.PuckResponse, error) {
  res := pb.PuckResponse{Result: "OK"}
	return &res, nil
}
