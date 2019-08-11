package main_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	main "github.com/gopherdojo/dojo6/kadai4/en-ken/kadai4"
)

type data struct {
	Fortune string `json:"fortune"`
}

func TestServer(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(main.Handler))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Errorf("%v", err)
	}

	var d data
	if err = json.NewDecoder(res.Body).Decode(&d); err != nil {
		t.Errorf("%v", err)
	}
}
