package data

import (
	"log"

	"github.com/DigitalInnovation/vitamns/entities"

	"gopkg.in/mgo.v2"
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
		log.Printf("Error finding user data in mongo: %v", err)
		return nil, err
	}

	return meal, nil
}
