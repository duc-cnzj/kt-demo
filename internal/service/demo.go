package service

import (
	"context"
	"github.com/duc-cnzj/kt-demo/internal/biz"

	pb "github.com/duc-cnzj/kt-demo/api/demo/v1"
)

type DemoService struct {
	pb.UnimplementedDemoServer
	db *biz.DemoBase
}

func NewDemoService(db *biz.DemoBase) *DemoService {
	return &DemoService{db: db}
}

func (s *DemoService) CreateDemo(ctx context.Context, req *pb.CreateDemoRequest) (*pb.CreateDemoReply, error) {
	if _, err := s.db.Get(&biz.Input{Name: req.Name}); err != nil {
		return nil, err
	}
	return &pb.CreateDemoReply{}, nil
}
func (s *DemoService) UpdateDemo(ctx context.Context, req *pb.UpdateDemoRequest) (*pb.UpdateDemoReply, error) {
	return &pb.UpdateDemoReply{}, nil
}
func (s *DemoService) DeleteDemo(ctx context.Context, req *pb.DeleteDemoRequest) (*pb.DeleteDemoReply, error) {
	return &pb.DeleteDemoReply{}, nil
}
func (s *DemoService) GetDemo(ctx context.Context, req *pb.GetDemoRequest) (*pb.GetDemoReply, error) {
	return &pb.GetDemoReply{}, nil
}
func (s *DemoService) ListDemo(ctx context.Context, req *pb.ListDemoRequest) (*pb.ListDemoReply, error) {
	return &pb.ListDemoReply{}, nil
}
