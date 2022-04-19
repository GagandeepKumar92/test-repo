package service

import (
	"fmt"

	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/db"
	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/domain"
)

type Manager interface {
	CreateUser(usr *domain.User) *domain.Error
	UpdateUser(usr *domain.User) error
	ListUser(limit int64, fileteredMap map[string]interface{}) ([]*domain.User, error)
	DeleteUser(id string) error
	ViewUser(id string) (*domain.User, error)
}

func NewManager(dbType string) Manager {

	ds, err := db.NewDataStore(dbType)
	if err != nil {
		fmt.Println("The err is", err)
		return nil
	}

	return &mgr{ds: ds}

}
