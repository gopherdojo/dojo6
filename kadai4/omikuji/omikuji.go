package omikuji

import (
	"fmt"
	"math/rand"
	"time"
)

var omikuji = map[int]string{
	0: "凶",
	1: "末吉",
	2: "小吉",
	3: "中吉",
	4: "吉",
	5: "大吉",
}

type Omikuji struct {
	Time time.Time
}

func (o *Omikuji) SetSeed(seed int64) {
	rand.Seed(seed)
}

type OmikujiError struct {
	Msg string
}

func (err *OmikujiError) Error() string {
	return fmt.Sprintf(err.Msg)
}

func (o *Omikuji) Do() (string, error) {
	var i int

	_, m, d := o.Time.Date()
	// 1/1 ~ 1/3のみ大吉を出す
	if int(m) == 1 && d >= 1 && d <= 3 {
		i = 5
	} else {
		i = rand.Intn(len(omikuji))
	}

	s, ok := omikuji[i]
	if !ok {
		return "", &OmikujiError{"おみくじが引けませんでした。"}
	}

	return s, nil
}
