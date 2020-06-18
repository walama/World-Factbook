package country

import (
	"errors"

	"gopkg.in/mgo.v2/bson"

	"github.com/asdine/storm"
)

// Country is a structure that defines a country in the database
type Country struct {
	ID          bson.ObjectId `json:"id"`
	Name        string        `json:"name"`
	Population  int           `json:"population"`
	Capital     string        `json:"capital"`
	IsDemocracy bool          `json:"isDemocracy"`
	Flag        string        `json:"flag"`
	Map         string        `json:"map"`
}

const (
	dbPath = "countries.db"
)

// errors
var (
	ErrRecordInvalid = errors.New("record is invalid")
)

// All retrieves all countries from the database
func All() ([]Country, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	countries := []Country{}
	err = db.All(&countries)
	if err != nil {
		return nil, err
	}
	return countries, nil
}

// One retrieves one country from the db by name
func One(id bson.ObjectId) (*Country, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	c := new(Country)
	err = db.One("ID", id, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

//Delete removes a country from the db
func Delete(id bson.ObjectId) error {
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	c := new(Country)
	err = db.One("ID", id, c)
	if err != nil {
		return err
	}
	return db.DeleteStruct(c)
}

// Save updates or creates a given country in the db
func (c *Country) Save() error {
	if err := c.validate(); err != nil {
		return err
	}
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Save(c)
}

func (c *Country) validate() error {
	if c.ID == "" {
		return ErrRecordInvalid
	}
	return nil
}
