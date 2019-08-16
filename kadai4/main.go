package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gopherdojo/dojo6/kadai4/omikuji"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

type Response struct {
	Msg string `json:"msg"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	omikuji := omikuji.Omikuji{time.Now()}
	ret, err := omikuji.Do()
	res := Response{}

	if err != nil {
		res.Msg = err.Error()
		log.Println("Error:", err)
	} else {
		res.Msg = ret
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println("Error:", err)
	}
}
