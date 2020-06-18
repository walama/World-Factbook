package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
	"worldfactbook/scratch_api/country"

	"gopkg.in/mgo.v2/bson"
)

func TestBodyToUser(t *testing.T) {
	valid := &country.Country{
		ID:          bson.NewObjectId(),
		Name:        "Argentina",
		Population:  44490000,
		Capital:     "Buenos Aires",
		IsDemocracy: true,
		Img:         "https://cdn.britannica.com/83/183583-050-B79EFF03/World-Data-Locator-Map-Argentina.jpg",
	}
	js, err := json.Marshal(valid)
	if err != nil {
		t.Errorf("Error Marshalling a valid country: %s", err)
	}
	ts := []struct {
		txt string
		r   *http.Request
		c   *country.Country
		err bool
		exp *country.Country
	}{
		{
			txt: "nil request",
			err: true,
		},
		{
			txt: "empty request body",
			r:   &http.Request{},
			err: true,
		},
		{
			txt: "empty country",
			r: &http.Request{
				Body: ioutil.NopCloser(bytes.NewBufferString("{}")),
			},
			err: true,
		},
		{
			txt: "malformed data",
			r: &http.Request{
				Body: ioutil.NopCloser(bytes.NewBufferString(`{"id":12}`)),
			},
			c:   &country.Country{},
			err: true,
		},
		{
			txt: "valid request",
			r: &http.Request{
				Body: ioutil.NopCloser(bytes.NewBuffer(js)),
			},
			c:   &country.Country{},
			exp: valid,
		},
	}

	for _, tc := range ts {
		t.Log(tc.txt)
		err := bodyToCountry(tc.r, tc.c)
		if tc.err {
			if err == nil {
				t.Error("Expected error, got none.")
			}
			continue
		}
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
			continue
		}
		if !reflect.DeepEqual(tc.c, tc.exp) {
			t.Error("Unmarshalled data is different:")
			t.Error(tc.c)
			t.Error(tc.exp)
		}
	}
}
