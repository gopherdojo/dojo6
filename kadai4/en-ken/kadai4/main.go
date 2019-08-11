package main

import (
	"encoding/json"
	"net/http"
)

func init() {
	fortunes = []string{"大吉", "吉", "中吉", "小吉", "末吉", "凶", "大凶"}
}

type response struct {
	Fortune string `json:"fortune"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicaiton/json; charset=utf-8")

	enc := json.NewEncoder(w)

	resp := &response{
		Fortune: getFortune(),
	}

	if err := enc.Encode(resp); err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
