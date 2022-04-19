package main

import (
	"flag"
	"log"

	"github.com/go-openapi/loads"
	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/gen/restapi"
	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/gen/restapi/operations"

	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer"
	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/api/handlers"

	//_ "github.com/go-swagger/go-swagger/examples/GaganSimpleServer/db/inmemory"
	_ "github.com/go-swagger/go-swagger/examples/GaganSimpleServer/db/mongo"
)

func main() {

	var portFlag = flag.Int("port", 4000, "Port to run this service on")

	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API
	api := operations.NewGaganSimpleServerAPI(swaggerSpec)

	v := GaganSimpleServer.NewRunTime("Gagan")
	api.UsersFindUsersHandler = handlers.NewFindUser(v)
	api.UsersAddUserHandler = handlers.NewAddNewUser(v)
	api.UsersDeleteUserHandler = handlers.NewDeleteUser(v)
	api.UsersUpdateUserHandler = handlers.NewUpdateUser(v)

	server := restapi.NewServer(api)
	defer server.Shutdown()

	// parse flags
	flag.Parse()
	// set the port this service will be run on
	server.Port = *portFlag

	// TODO: Set Handle

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
