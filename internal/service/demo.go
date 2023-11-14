package service

import (
	"context"
	"github.com/duc-cnzj/kt-demo/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"time"

	pb "github.com/duc-cnzj/kt-demo/api/demo/v1"
)

type DemoService struct {
	pb.UnimplementedDemoServer
	db  *biz.DemoBase
	log *log.Helper
}

func NewDemoService(db *biz.DemoBase, logger log.Logger) *DemoService {
	return &DemoService{db: db, log: log.NewHelper(log.With(logger, "module", "svc/demo"))}
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
func (s *DemoService) ListDemo(req *pb.ListDemoRequest, ds pb.Demo_ListDemoServer) error {
	s.log.Info("ListDemo")
	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			return nil
		default:
		}
		ds.Send(&pb.ListDemoReply{Data: req.Data})
	}
}
