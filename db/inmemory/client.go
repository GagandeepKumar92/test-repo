package inmemory

import (
	"fmt"
	"strings"

	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/db"
	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/domain"
)

func init() {
	db.RegisterDataStore("inmemory", NewClient)
	fmt.Println("Call is in Init")
}

func NewClient() (db.DataStore, error) {
	return &client{ds: make(map[string]*domain.User)}, nil
}

type client struct {
	ds map[string]*domain.User
}

func (c *client) ViewUser(id string) (*domain.User, error) {

	var userInfo *domain.User
	var ok bool
	if userInfo, ok = c.ds[id]; !ok {
		return nil, &domain.Error{Code: 404, Message: "User doesn't exist"}
	}

	return userInfo, nil
}

func (c *client) UpdateUser(user *domain.User) error {

	c.ds[user.ID] = user

	return nil
}

func (c *client) AddUser(user *domain.User) (string, error) {
	c.ds[user.ID] = user
	return user.ID, nil
}

func (c *client) ListUsers(limit int64, filteredMap map[string]interface{}) ([]*domain.User, error) {

	var userInfo = []*domain.User{}
	for _, value := range c.ds {
		userInfo = append(userInfo, value)
	}

	userInfoName := make([]*domain.User, 0)
	if filteredMap["name"] != nil {
		for i := 0; i < len(userInfo); i++ {
			if strings.ToUpper(fmt.Sprint(filteredMap["name"])) == strings.ToUpper(userInfo[i].Name) {
				userInfoName = append(userInfoName, userInfo[i])
			}
		}
		userInfo = userInfoName
	}

	userInfoLimit := make([]*domain.User, 0)
	if limit != 0 {
		if len(userInfo) > int(limit) {
			for i := 0; i < int(limit); i++ {
				userInfoLimit = append(userInfoLimit, userInfo[i])
			}
			userInfo = userInfoLimit
		}
	}

	fmt.Println("The value of userInfo is", userInfo)
	return userInfo, nil
}

func (c *client) DeleteUser(id string) error {
	if _, ok := c.ds[id]; ok {
		delete(c.ds, id)
		return nil
	}
	return nil
}
