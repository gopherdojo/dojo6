package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// Result is Fortune result
type Result struct {
	Fortune string `json:"fotrune"`
	Work    string `json:"work"`
	Love    string `json:"love"`
	Health  string `json:"health"`
}

func main() {
	rand.Seed(time.Now().UnixNano())
	//TODO: ハンドラのテスト
	http.HandleFunc("/fortune/", drawFortune)
	http.ListenAndServe(":8080", nil)
}

func drawFortune(w http.ResponseWriter, r *http.Request) {
	result := draw()

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(result); err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, buf.String())
}

func draw() (result Result) {
	result = Result{
		Fortune: judgeFortune(),
		Work:    judge(),
		Love:    judge(),
		Health:  judge(),
	}
	return
}

func judgeFortune() (stage string) {
	// 今の日付を取得。日本で動かす場合はJSTが取得される
	yd := time.Now().YearDay()
	// 三ヶ日は全員幸せ
	if 1 <= yd && yd <= 3 {
		stage = "大吉"
	} else {
		randInt := rand.Intn(7)

		switch randInt {
		case 0:
			stage = "大吉"
		case 1:
			stage = "吉"
		case 2:
			stage = "中吉"
		case 3:
			stage = "小吉"
		case 4:
			stage = "末吉"
		case 5:
			stage = "凶"
		case 6:
			stage = "大凶"
		default:
			// TODO: ここにきたら本来はエラー
			stage = "大吉"
		}

	}
	return

}

func judge() (stage string) {
	randInt := rand.Intn(5)
	switch randInt {
	case 0:
		stage = "worst"
	case 1:
		stage = "bad"
	case 2:
		stage = "normal"
	case 3:
		stage = "good"
	case 4:
		stage = "best"
	default:
		// TODO: ここにきたら本来はエラー
		stage = "normal"
	}
	return
}
