package controller

import (
	"pangolin/core"
	"pangolin/web/service"
)

var srv *Services

type Services struct {
	Base *service.Base
}

func NewServices() {
	srv = &Services{
		Base: service.NewBase(core.CoreEngine),
	}
}
