package entities

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type User struct {
	Id       bson.ObjectId `bson:"_id,omitempty"`
	Name     string        `json:"name"`
	Wardrobe []Clothes     `json:"wardrobe"`
}

type Users []User

type Clothes struct {
	Name     string    `json:"name"`
	Type     string    `json:"type"`
	Image    string    `json:"image"`
	Colour   string    `json:"colour"`
	Lastworn time.Time `json:"lastworn"`
	Wearing  bool      `json:"wearing"`
}
