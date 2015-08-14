package data

import (
	"errors"
	"log"

	"github.com/DigitalInnovation/simplyactiveapi/entities"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Mongodal struct {
	session        *mgo.Session
	databaseName   string
	collectionName string
}

func (m *Mongodal) CreateDataConnection(url, database, collection string) error {
	m.databaseName = database
	m.collectionName = collection

	s, err := mgo.Dial(url)
	if err != nil {
		log.Printf("There was an error connecting to Mongo: %v", err)
		return err
	}

	m.session = s
	return nil
}

func (m *Mongodal) close() {
	m.session.Close()
}

func (m *Mongodal) GetAllMeals() ([]entities.Meal, error) {
	var meal []entities.Meal

	ses := m.session.New()
	defer ses.Close()

	collection := ses.DB(m.databaseName).C(m.collectionName)
	err := collection.Find(nil).All(&meal)
	if err != nil {
		log.Printf("Error finding meal data in mongo: %v", err)
		return nil, err
	}

	return meal, nil
}

// func (m *Mongodal) GetLowStepsMeals() ([]entities.Meal, error) {

// 	var meal []entities.Meal

// 	ses := m.session.New()
// 	defer ses.Close()

// 	if !bson.IsObjectIdHex(id) {
// 		return errors.New("Bad Request, Id must be a valid bson ObjectIdHex"), 400
// 	}

// 	collection := ses.DB(m.databaseName).C(m.collectionName)
// 	err := collection.Find(bson.M{"expirydate": bson.M{"$gte": now}}).All(&message)
// 	if err != nil {
// 		log.Printf("Error finding message data in mongo: %v", err)
// 		return nil, err
// 	}

// 	return message, nil
// }

func (m *Mongodal) CreateMeal(meal entities.Meal) error {
	ses := m.session.New()
	defer ses.Close()

	collection := ses.DB(m.databaseName).C(m.collectionName)
	err := collection.Insert(meal)
	if err != nil {
		log.Printf("Error inserting meal data in mongo: %v", err)
		return err
	}

	return nil
}

func (m *Mongodal) DeleteMeal(id string) (error, int) {
	ses := m.session.New()
	defer ses.Close()

	if !bson.IsObjectIdHex(id) {
		return errors.New("Bad Request, Id must be a valid bson ObjectIdHex"), 400
	}

	mealId := bson.ObjectIdHex(id)

	collection := ses.DB(m.databaseName).C(m.collectionName)
	err := collection.RemoveId(mealId)
	if err != nil {
		log.Printf("Error deleting meal with id: %v, %v", id, err)
		return err, 404
	}

	return nil, 204
}
