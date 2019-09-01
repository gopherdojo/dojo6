package main_test

import (
	"encoding/json"
	"fmt"
	"io"
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
		t.Errorf("Response error:%v", err)
	}

	var d data
	if err = json.NewDecoder(res.Body).Decode(&d); err != nil {
		t.Errorf("Data format error:%v", err)
	}
}

type EncoderMock struct {
}

func NewEncoderMock(w io.Writer) main.Encoder {
	return &EncoderMock{}
}

func (enc *EncoderMock) Encode(v interface{}) error {
	return fmt.Errorf("Internal server error")
}

func TestServerInternalServerError(t *testing.T) {
	main.SetNewEncoder(NewEncoderMock)
	defer main.SetNewEncoder(main.NewEncoder(func(w io.Writer) main.Encoder {
		return json.NewEncoder(w)
	}))

	ts := httptest.NewServer(http.HandlerFunc(main.Handler))
	defer ts.Close()

	res, _ := http.Get(ts.URL)
	if res.StatusCode != 500 {
		t.Errorf("Response code error:%v", res.StatusCode)
	}

}

func TestHandler(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	main.SetFortunes([]string{"大凶"})

	main.Handler(w, r)

	if w.Code != 200 {
		t.Errorf("Invalid code: %v", w.Code)
	}
	if w.Header().Get("Content-Type") != "applicaiton/json; charset=utf-8" {
		t.Errorf("Insufficient headers: %v", w.Header())
	}
	body := w.Body
	dec := json.NewDecoder(body)
	var d data
	if err := dec.Decode(&d); err != nil || d.Fortune != "大凶" {
		t.Errorf("Failed to decode: %v", err)
	}
}
