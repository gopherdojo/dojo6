package omikuji

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

// Omikuji Result
type Omikuji struct {
	Me    int    `json:"me"`
	Unsei string `json:"unsei"`
}

// Hiku a omikuji
func Hiku(t time.Time) (int, string) {

	if shogatsu(t) {
		return 0, "大吉"
	}

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
		return Hiku(t)
	}
	return me, unsei
}

func shogatsu(t time.Time) bool {

	if t.Month() == time.January {
		if (t.Day() == 1) || (t.Day() == 2) || (t.Day() == 3) {
			return true
		}
	}

	return false
}

// Handler provides Omikuji Handler
func Handler(w http.ResponseWriter, r *http.Request) {
	me, unsei := Hiku(time.Now())
	o := &Omikuji{Me: me, Unsei: unsei}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(o); err != nil {
		http.Error(w, "Omikuji Error", http.StatusInternalServerError)
	}
}
