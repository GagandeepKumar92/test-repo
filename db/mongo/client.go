package mongo

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/db"
	"github.com/go-swagger/go-swagger/examples/GaganSimpleServer/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func init() {
	db.RegisterDataStore("mongo", NewClient)
}

func NewClient() (db.DataStore, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientCurrent, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		fmt.Println("The error is", err)
		return nil, err
	}

	err = clientCurrent.Ping(ctx, readpref.Primary())

	if err != nil {
		fmt.Println("The error is", err)
		return nil, err
	}

	return &client{dbc: clientCurrent.Database("users_db")}, nil
}

type client struct {
	dbc *mongo.Database
}

func (c *client) AddUser(user *domain.User) (string, error) {

	fmt.Println("In Add User 1")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	_, err := c.dbc.Collection("users").InsertOne(ctx, bson.D{
		{Key: "_id", Value: user.ID},
		{Key: "name", Value: user.Name},
		{Key: "address", Value: user.Address},
		{Key: "created_at", Value: user.CreatedAt},
	})

	fmt.Println("In Add User 2")
	if err != nil {
		fmt.Println("The error is", err)
		return "", err
	}
	fmt.Println("In Add User 3")

	return "", nil
}

func (c *client) ViewUser(id string) (*domain.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	fmt.Println(id, "My Id is there")

	var userInfo domain.User
	if err := c.dbc.Collection("users").FindOne(ctx, bson.M{"_id": id}).Decode(&userInfo); err != nil {
		fmt.Println(err, "error is there")
		return nil, &domain.Error{Code: 404, Message: "User doesn't exist"}
	}

	return &userInfo, nil
}

func (c *client) UpdateUser(user *domain.User) error {

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	fmt.Println("Update user is getting called 2")
	_, err := c.dbc.Collection("users").UpdateOne(
		ctx,
		bson.M{"_id": user.ID},
		bson.D{
			{"$set", bson.D{{"address", user.Address}}},
		},
	)

	fmt.Println("User Id = ", user.ID, " Address = ", user.Address)
	if err != nil {
		fmt.Println(user.ID, "Error")
		return err
	}

	return nil
}

func (c *client) ListUsers(limit int64, filteredMap map[string]interface{}) ([]*domain.User, error) {

	userInfo := make([]*domain.User, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var (
		cur *mongo.Cursor
		err error
	)

	options := options.Find().SetLimit(limit)
	cur, err = c.dbc.Collection("users").Find(ctx, applyFilter(filteredMap), options)
	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result domain.User
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		fmt.Println(&result)
		userInfo = append(userInfo, &result)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return userInfo, nil
}

func applyFilter(filterMap map[string]interface{}) map[string]interface{} {

	for k, v := range filterMap {
		switch mval := v.(type) {
		case string:
			// support searching by name using case-insensitive matching
			fmt.Println("v value is = ", v)
			if k == "name" {
				filterMap[k] = bson.M{"$regex": primitive.Regex{Pattern: "^" + regexp.QuoteMeta(mval) + "$", Options: "i"}}
			}

		}
	}
	return filterMap
}

func (c *client) DeleteUser(id string) error {

	fmt.Println(id, " Id ")
	_, err := c.dbc.Collection("users").DeleteOne(context.Background(), bson.M{"_id": id})
	fmt.Println(err, " Delete ")
	if err != nil {
		return err
	}

	return nil
}
