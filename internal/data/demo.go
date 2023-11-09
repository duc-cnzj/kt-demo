package data

import (
	"github.com/duc-cnzj/kt-demo/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type demoRepo struct {
	log  *log.Helper
	data *Data
}

func NewDemoRepo(logger log.Logger, data *Data) biz.DemoRepo {
	return &demoRepo{log: log.NewHelper(logger), data: data}
}

func (d *demoRepo) Get(input *biz.Input) (*biz.Resp, error) {
	d.log.Info("demoRepo called")
	return &biz.Resp{Name: input.Name}, nil
}
