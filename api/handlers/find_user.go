package handlers

import (
	"log"

	"github.com/go-openapi/runtime/middleware"
	gserver "github.com/go-swagger/go-swagger/examples/GaganSimpleServer"
	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/gen/models"
	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/gen/restapi/operations/users"
)

func NewFindUser(rt *gserver.Runtime) users.FindUsersHandler {
	return &findUser{rt: rt}
}

type findUser struct {
	rt *gserver.Runtime
}

func (f *findUser) Handle(fup users.FindUsersParams) middleware.Responder {

	us, err := f.rt.GetManager().ListUser(*fup.Limit, filteredMap(fup))

	if err != nil {
		log.Fatal(err)
	}

	usResponse := []*models.User{}
	for _, usr := range us {
		usResponse = append(usResponse, asUserResponse(usr))
	}

	return users.NewFindUsersOK().WithPayload(usResponse)
}

func filteredMap(fup users.FindUsersParams) map[string]interface{} {
	filterMap := make(map[string]interface{})

	if fup.Name != nil {
		filterMap["name"] = *fup.Name
	}

	return filterMap
}
