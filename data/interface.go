package data

import (
	"github.com/DigitalInnovation/vitamns/entities"
)

type DAL interface {
	CreateDataConnection(url, database, collection string) error
	GetAllMeals() ([]entities.Meal, error)
}
