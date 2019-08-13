package main

import (
	"encoding/json"
	"log"
	"net/http"
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

	res := Response{Msg: "json response"}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println("Error:", err)
	}
}

/*
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/isuzuki/omikuji/omikuji"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

type Lottery struct {
	Msg string `json:"msg"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	_, m, d := time.Now().Date()
	ret, ok := omikuji.Lottery(int(m), d)
	var res Lottery
	if ok {
		res = Lottery{Msg: ret}
	} else {
		res = Lottery{Msg: "エラーが発生しました。"}
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println("Error:", err)
	}
}
*/
