package db

import (
	"fmt"

	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/domain"
)

type DataStore interface {
	UpdateUser(*domain.User) error
	AddUser(*domain.User) (string, error)
	ListUsers(int64, map[string]interface{}) ([]*domain.User, error)
	DeleteUser(string) error
	ViewUser(string) (*domain.User, error)
}
type DatastoreFactory func() (DataStore, error)

var factories map[string]DatastoreFactory

func RegisterDataStore(key string, value DatastoreFactory) {
	if factories == nil {
		factories = make(map[string]DatastoreFactory)
	}
	factories[key] = value
}

func NewDataStore(dbType string) (DataStore, error) {
	fmt.Println("The length is", len(factories))
	return factories[dbType]()
}
