package data

import (
	"errors"
	"log"

	"github.com/DigitalInnovation/schnapi/entities"

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

func (m *Mongodal) GetAllUsers() ([]entities.User, error) {
	var user []entities.User

	ses := m.session.New()
	defer ses.Close()

	collection := ses.DB(m.databaseName).C(m.collectionName)
	err := collection.Find(nil).All(&user)
	if err != nil {
		log.Printf("Error finding user data in mongo: %v", err)
		return nil, err
	}

	return user, nil
}

func (m *Mongodal) CreateUser(user entities.User) error {
	ses := m.session.New()
	defer ses.Close()

	collection := ses.DB(m.databaseName).C(m.collectionName)
	err := collection.Insert(user)
	if err != nil {
		log.Printf("Error inserting user data in mongo: %v", err)
		return err
	}

	return nil
}

func (m *Mongodal) DeleteUser(id string) (error, int) {
	ses := m.session.New()
	defer ses.Close()

	if !bson.IsObjectIdHex(id) {
		return errors.New("Bad Request, Id must be a valid bson ObjectIdHex"), 400
	}

	userId := bson.ObjectIdHex(id)

	collection := ses.DB(m.databaseName).C(m.collectionName)
	err := collection.RemoveId(userId)
	if err != nil {
		log.Printf("Error deleting user with id: %v, %v", id, err)
		return err, 404
	}

	return nil, 204
}
