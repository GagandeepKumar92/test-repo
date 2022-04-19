package GaganSimpleServer

import (
	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/service"
)

func NewRunTime(an string) *Runtime {
	//return &Runtime{AppName: an, svc: service.NewManager("inmemory")}
	return &Runtime{AppName: an, svc: service.NewManager("mongo")}
}

type Runtime struct {
	AppName string
	svc     service.Manager
}

func (rt *Runtime) GetManager() service.Manager {
	return rt.svc
}

func (rt *Runtime) GetApplicationName() string {
	return rt.AppName
}
func (rt *Runtime) SetApplicationName(s string) {
	rt.AppName = s
}
