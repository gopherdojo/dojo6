package omikuji

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

// Omikuji Result
type Omikuji struct {
	Me    int    `json:"me"`
	Unsei string `json:"unsei"`
}

// Draw a omikuji
func Draw() (int, string) {

	me := rand.Intn(7)
	var unsei string
	switch me {
	case 6:
		unsei = "大吉"
	case 5, 4:
		unsei = "中吉"
	case 3, 2:
		unsei = "小吉"
	case 1:
		unsei = "凶"
	default:
		return Draw()
	}
	return me, unsei
}

// Handler provides Omikuji Handler
func Handler(w http.ResponseWriter, r *http.Request) {
	me, unsei := Draw()
	o := &Omikuji{Me: me, Unsei: unsei}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(o); err != nil {
		log.Println("Error: ", err)
	}
}
