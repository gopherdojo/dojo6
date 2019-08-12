package main

import (
	"encoding/json"
	"io"
	"net/http"
)

// Encoder is I/F to json.Encoder
type Encoder interface {
	Encode(v interface{}) error
}

// NewEncoder is I/F to json.NewEncoder
type NewEncoder func(w io.Writer) Encoder

var newEncoder NewEncoder

func init() {
	// For testing
	newEncoder = func(w io.Writer) Encoder {
		return json.NewEncoder(w)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicaiton/json; charset=utf-8")

	enc := newEncoder(w)

	resp := &struct {
		Fortune string `json:"fortune"`
	}{
		Fortune: getFortune(),
	}

	if err := enc.Encode(resp); err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
	}
}
