package country

import (
	"os"
	"reflect"
	"testing"

	"github.com/asdine/storm"
	"gopkg.in/mgo.v2/bson"
)

func TestMain(m *testing.M) {
	m.Run()
	os.Remove(dbPath)
}

func TestCRUD(t *testing.T) {
	t.Log("Create")
	c := &Country{
		ID:          bson.NewObjectId(),
		Name:        "Argentina",
		Population:  44490000,
		Capital:     "Buenos Aires",
		IsDemocracy: true,
		Flag: "",
		Map: ""
	}
	err := c.Save()
	if err != nil {
		t.Fatalf("Error saving a record: %s", err)
	}
	t.Log("Read")
	c2, err := One(c.ID)
	if err != nil {
		t.Fatalf("Error retrieving a record: %s", err)
	}
	if !reflect.DeepEqual(c2, c) {
		t.Error("Records do not match")
	}
	t.Log("Update")
	c.Population = 44500000
	err = c.Save()
	if err != nil {
		t.Fatalf("Error saving a record: %s", err)
	}
	c3, err := One(c.ID)
	if err != nil {
		t.Fatalf("Error retreieving a record: %s", err)
	}
	if !reflect.DeepEqual(c3, c) {
		t.Error("Records do not match")
	}
	t.Log("Delete")
	err = Delete(c.ID)
	if err != nil {
		t.Fatalf("error deleting a record: %s", err)
	}
	_, err = One(c.ID)
	if err == nil {
		t.Fatal("Record should no longer exist")
	}
	if err != storm.ErrNotFound {
		t.Fatalf("Error retrieving non existant record: %s", err)
	}

	t.Log("Read All")
	c2.ID = bson.NewObjectId()
	c3.ID = bson.NewObjectId()
	err = c.Save()
	if err != nil {
		t.Fatalf("Error saving a record: %s", err)
	}
	err = c2.Save()
	if err != nil {
		t.Fatalf("Error saving a record: %s", err)
	}
	err = c3.Save()
	if err != nil {
		t.Fatalf("Error saving a record: %s", err)
	}
	countries, err := All()
	if err != nil {
		t.Fatalf("Error reading all records: %s", err)
	}
	if len(countries) != 3 {
		t.Errorf("Wrong noumber of countries retrieved. Expected 3, Actual: %d", len(countries))
	}

}
