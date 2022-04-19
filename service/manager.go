package service

import (
	"fmt"
	"time"

	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/db"

	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/domain"
	"github.com/segmentio/ksuid"
)

type mgr struct {
	ds db.DataStore
}

func (m *mgr) CreateUser(usr *domain.User) *domain.Error {

	usr.CreatedAt = time.Now().UTC()
	usr.ID = ksuid.New().String()

	if len(usr.Name) < 3 {
		return &domain.Error{Code: 400, Message: "Name should be at least 3 characters long"}
	}

	m.ds.AddUser(usr)
	return nil

}

func (m *mgr) ViewUser(i string) (*domain.User, error) {

	return m.ds.ViewUser(i)
}

func (m *mgr) UpdateUser(usr *domain.User) error {

	fmt.Println(usr.ID)
	dbUser, ViewErr := m.ds.ViewUser(usr.ID)

	if ViewErr != nil {
		fmt.Println("1")
		return &domain.Error{Code: 404, Message: "User doesn't exist"}
	}

	dbUser.Address = usr.Address

	return m.ds.UpdateUser(dbUser)

}

func (m *mgr) DeleteUser(id string) error {

	_, ViewErr := m.ds.ViewUser(id)

	if ViewErr != nil {
		fmt.Println("1")
		return &domain.Error{Code: 404, Message: "User doesn't exist"}
	}

	return m.ds.DeleteUser(id)
}

func (m *mgr) ListUser(limit int64, fileteredMap map[string]interface{}) ([]*domain.User, error) {

	return m.ds.ListUsers(limit, fileteredMap)

}
