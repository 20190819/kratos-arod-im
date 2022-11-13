package service

import (
	"context"

	pb "arod-im/api/logic/v1"
)

type LogicService struct {
	pb.UnimplementedLogicServer
}

func NewLogicService() *LogicService {
	return &LogicService{}
}

func (s *LogicService) CreateLogic(ctx context.Context, req *pb.CreateLogicRequest) (*pb.CreateLogicReply, error) {
	return &pb.CreateLogicReply{}, nil
}
func (s *LogicService) UpdateLogic(ctx context.Context, req *pb.UpdateLogicRequest) (*pb.UpdateLogicReply, error) {
	return &pb.UpdateLogicReply{}, nil
}
func (s *LogicService) DeleteLogic(ctx context.Context, req *pb.DeleteLogicRequest) (*pb.DeleteLogicReply, error) {
	return &pb.DeleteLogicReply{}, nil
}
func (s *LogicService) GetLogic(ctx context.Context, req *pb.GetLogicRequest) (*pb.GetLogicReply, error) {
	return &pb.GetLogicReply{}, nil
}
func (s *LogicService) ListLogic(ctx context.Context, req *pb.ListLogicRequest) (*pb.ListLogicReply, error) {
	return &pb.ListLogicReply{}, nil
}
