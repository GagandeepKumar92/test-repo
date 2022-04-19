package domain

import (
	"time"
)

type User struct {
	ID        string    `bson:"_id"`
	Name      string    `bson:"name"`
	Address   string    `bson:"address"`
	CreatedAt time.Time `bson:"created_at"`
}
