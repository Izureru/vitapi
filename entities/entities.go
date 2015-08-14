package entities

import (
	"labix.org/v2/mgo/bson"
)

type Meal struct {
	Id    bson.ObjectId `bson:"_id,omitempty"`
	Name  string        `json:"name"`
	Image string        `json:"image"`
}
