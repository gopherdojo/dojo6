package main

import (
	"net/http"
	"math/rand"
	"time"

	"github.com/dojo6/kadai4/annkara/pkg/omikuji"
)

func main() {
	n := time.Now().UnixNano()
	rand.Seed(n)

	http.HandleFunc("/", omikuji.Handler)
	http.ListenAndServe(":8080", nil)
}
