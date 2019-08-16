package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	handler(w, r)
	rw := w.Result()

	defer rw.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code: %d", rw.StatusCode)
	}

	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal("unexpected error")
	}

	res := &Response{}
	const expected = "大吉"
	if err := json.Unmarshal(b, res); err != nil {
		t.Fatalf("JSON unmarshall error: %v", err)
	}

	if res.Msg != string(expected) {
		t.Fatalf("unexpected response: %s", res.Msg)
	}
}
