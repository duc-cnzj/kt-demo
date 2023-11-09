package biz

import (
	v1 "github.com/duc-cnzj/kt-demo/api/demo/v1"
	"github.com/go-kratos/kratos/v2/log"
)

type Input struct {
	Name string
}

type Resp struct {
	Name string
}

type DemoRepo interface {
	Get(input *Input) (*Resp, error)
}

type DemoBase struct {
	repo DemoRepo
	log  *log.Helper
}

func NewDemoBase(repo DemoRepo, logger log.Logger) *DemoBase {
	return &DemoBase{repo: repo, log: log.NewHelper(logger)}
}

func (b *DemoBase) Get(input *Input) (*Resp, error) {
	if input.Name == "abc" {
		return nil, v1.ErrorDemoNotFound("err %s", "abc")
	}
	return b.repo.Get(input)
}
