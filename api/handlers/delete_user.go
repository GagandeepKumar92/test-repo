package handlers

import (
	"fmt"

	"github.com/go-openapi/runtime/middleware"
	gserver "github.com/go-swagger/go-swagger/examples/GaganSimpleServer"
	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/domain"
	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/gen/restapi/operations/users"
)

func NewDeleteUser(rt *gserver.Runtime) users.DeleteUserHandler {
	return &deleteUser{rt: rt}
}

type deleteUser struct {
	rt *gserver.Runtime
}

func (d *deleteUser) Handle(del users.DeleteUserParams) middleware.Responder {

	err := d.rt.GetManager().DeleteUser(del.ID)
	fmt.Println("The Delete status from handler", err)
	if err != nil {
		derr, ok := err.(domain.Err)
		fmt.Println("The typecasted Delete status from handler", derr.StatusCode())
		if ok {
			switch derr.StatusCode() {
			case 404:
				return users.NewDeleteUserNotFound().WithPayload(asErrorResponse(err.(*domain.Error)))
			}
		} else {
			return users.NewDeleteUserDefault(500).WithPayload(asErrorResponse(&domain.Error{Message: "Internal Server Error"}))
		}
	}

	return users.NewDeleteUserNoContent()
}
