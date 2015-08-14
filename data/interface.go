package data

import (
	"github.com/DigitalInnovation/simplyactiveapi/entities"
)

type DAL interface {
	CreateDataConnection(url, database, collection string) error
	GetAllMeals() ([]entities.Meal, error)
	CreateMeal(meal entities.Meal) error
	DeleteMeal(id string) (error, int)
}
