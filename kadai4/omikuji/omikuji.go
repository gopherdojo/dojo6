package omikuji

import (
	"math/rand"
)

var omikuji = map[int]string{
	0: "凶",
	1: "末吉",
	2: "小吉",
	3: "中吉",
	4: "吉",
	5: "大吉",
}

func Do(month, day int) string {
	var i int
	// 1/1 ~ 1/3のみ大吉を出す
	if month == 1 && day >= 1 && day <= 3 {
		i = 5
	} else {
		i = rand.Intn(len(omikuji))
	}

	s, ok := omikuji[i]
	if !ok {
		panic("omikuji panic.")
	}

	return s
}
