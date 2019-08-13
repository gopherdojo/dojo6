package main

import (
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

	const expected = `{"msg":"大吉"}`
	if s := string(b); s != expected {
		t.Fatalf("unexpected response: %s", s)
	}
}
