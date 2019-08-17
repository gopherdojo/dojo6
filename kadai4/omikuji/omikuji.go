package omikuji

import (
	"math/rand"
	"time"
)

var omikuji = []string{
	"凶",
	"末吉",
	"小吉",
	"中吉",
	"吉",
	"大吉",
}

type Omikuji struct {
	Time time.Time
}

func (o *Omikuji) SetSeed(seed int64) {
	rand.Seed(seed)
}

func (o *Omikuji) Do() string {
	var i int

	_, m, d := o.Time.Date()
	// 1/1 ~ 1/3のみ大吉を出す
	if int(m) == 1 && d >= 1 && d <= 3 {
		return "大吉"
	}

	i = rand.Intn(len(omikuji))
	return omikuji[i]
}
