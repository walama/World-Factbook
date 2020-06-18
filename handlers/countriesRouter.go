package handlers

import (
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

// CountriesRouter handles the /countries path
func CountriesRouter(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSuffix(r.URL.Path, "/")

	if path == "/countries" {

		switch r.Method {
		case http.MethodGet:

			countriesGetAll(w, r)
			return
		case http.MethodPost:
			countriesPostOne(w, r)
			return
		case http.MethodHead:
			countriesGetAll(w, r)
			return
		case http.MethodOptions:

			postOptionsResponse(w, []string{http.MethodGet, http.MethodPost, http.MethodHead, http.MethodOptions}, nil)
		default:
			postError(w, http.StatusMethodNotAllowed)
		}
	}

	path = strings.TrimPrefix(path, "/countries/")
	if !bson.IsObjectIdHex(path) {
		postError(w, http.StatusNotFound)
		return
	}

	id := bson.ObjectIdHex(path)
	switch r.Method {
	case http.MethodGet:
		countriesGetOne(w, r, id)
		return
	case http.MethodPut:
		countriesPutOne(w, r, id)
		return
	case http.MethodPatch:
		countriesPatchOne(w, r, id)
		return
	case http.MethodDelete:
		countriesDeleteOne(w, r, id)
		return
	case http.MethodHead:
		countriesGetOne(w, r, id)
		return
	case http.MethodOptions:

		postOptionsResponse(w, []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodPost, http.MethodHead, http.MethodOptions}, nil)

	default:
		postError(w, http.StatusMethodNotAllowed)
		return
	}
}
