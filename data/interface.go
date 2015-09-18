package data

import (
	"github.com/DigitalInnovation/schnapi/entities"
)

type DAL interface {
	CreateDataConnection(url, database, collection string) error
	GetAllUsers() ([]entities.User, error)
	CreateUser(meal entities.User) error
	DeleteUser(id string) (error, int)
}
