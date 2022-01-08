package record

import (
	bsonpr "go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Entity represents the record domain model in the database.
type Entity struct{
	Id bsonpr.ObjectID `bson:"_id"`
	CreatedAt time.Time `bson:"createdAt"`
	Counts []int `json:"counts"`
	Key string `bson:"key"`
	Value string `bson:"value"`

	TotalCount int `bson:"totalCount"`
}
