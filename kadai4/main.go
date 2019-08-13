package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gopherdojo/dojo6/kadai4/omikuji"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

type Response struct {
	Msg string `json:"msg"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	_, m, d := time.Now().Date()
	ret := omikuji.Do(int(m), d)
	res := Response{Msg: ret}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println("Error:", err)
	}
}
