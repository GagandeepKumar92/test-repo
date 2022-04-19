package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	gserver "github.com/go-swagger/go-swagger/examples/GaganSimpleServer"
	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/domain"
	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/gen/restapi/operations/users"
)

func NewAddNewUser(rt *gserver.Runtime) users.AddUserHandler {
	return &addUser{rt: rt}
}

type addUser struct {
	rt *gserver.Runtime
}

func (f *addUser) Handle(fup users.AddUserParams) middleware.Responder {

	usr := &domain.User{ID: fup.Body.ID, Name: *fup.Body.Name, Address: fup.Body.Address}

	if err := f.rt.GetManager().CreateUser(usr); err != nil {
		return users.NewAddUserBadRequest().WithPayload(asErrorResponse(err))
	}

	return users.NewAddUserCreated()

}
