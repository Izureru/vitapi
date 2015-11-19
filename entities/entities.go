package entities

import (
	"labix.org/v2/mgo/bson"
)

type Meal struct {
	Id        bson.ObjectId `bson:"_id,omitempty"`
	Name      string        `json:"name"`
	EPlan     string        `json:"ePlan"`
	Kcal      int           `json:"kcal"`
	Allergens string        `json:"allergens"`
	MTime     string        `json:"mTime"`
	Price     int           `json:"price"`
	Img       string        `json:"img"`
}

type Meals []Meal
