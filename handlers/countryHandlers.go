package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"worldfactbook/scratch_api/country"

	"github.com/asdine/storm"

	"gopkg.in/mgo.v2/bson"
)

func countriesGetAll(w http.ResponseWriter, r *http.Request) {

	countries, err := country.All()
	if err != nil {

		postError(w, http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodHead {
		postBodyResponse(w, http.StatusOK, jsonResponse{})
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"countries": countries})
}

func bodyToCountry(r *http.Request, c *country.Country) error {
	if r == nil {

		return errors.New("A request is required")
	}
	if r.Body == nil {

		return errors.New("Request body is empty")
	}
	if c == nil {

		return errors.New("a country is required")
	}
	bd, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(bd, c)
}

func countriesPostOne(w http.ResponseWriter, r *http.Request) {
	print("new country \n")
	c := new(country.Country)
	err := bodyToCountry(r, c)

	if err != nil {
		print("bad status")
		postError(w, http.StatusBadRequest)
		return
	}
	c.ID = bson.NewObjectId()
	err = c.Save()
	if err != nil {
		if err == country.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Location", "/users/"+c.ID.Hex())
	w.WriteHeader(http.StatusCreated)
}

func countriesGetOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {
	c, err := country.One(id)
	if err != nil {
		if err == storm.ErrNotFound {
			postError(w, http.StatusNotFound)
			return
		}
		postError(w, http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodHead {
		postBodyResponse(w, http.StatusOK, jsonResponse{})
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"country": c})
}

func countriesPutOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {
	c := new(country.Country)
	err := bodyToCountry(r, c)

	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	c.ID = id
	err = c.Save()
	if err != nil {
		if err == country.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"country": c})
}

func countriesPatchOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {
	c, err := country.One(id)
	if err != nil {
		if err == storm.ErrNotFound {
			postError(w, http.StatusNotFound)
			return
		}
		postError(w, http.StatusInternalServerError)
		return
	}
	err = bodyToCountry(r, c)

	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	c.ID = id
	err = c.Save()
	if err != nil {
		if err == country.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"country": c})
}

func countriesDeleteOne(w http.ResponseWriter, _ *http.Request, id bson.ObjectId) {
	err := country.Delete(id)
	if err != nil {
		if err == storm.ErrNotFound {
			postError(w, http.StatusNotFound)
			return
		}
		postError(w, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
