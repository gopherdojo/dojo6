package omikuji_test

import (
	"testing"
	"time"

	"github.com/gopherdojo/dojo6/kadai4/omikuji"
)

func Test_NormalTime(t *testing.T) {
	time := time.Date(2019, time.Month(8), 16, 0, 0, 0, 0, time.Local)
	omikuji := omikuji.Omikuji{time}
	expect := "大吉"
	actual := omikuji.Do()
	if expect != actual {
		t.Errorf(`Omikuji error: expect="%s" actual="%s"`, expect, actual)
	}
}

func Test_SpecificPeriod(t *testing.T) {
	periods := map[int][]int{
		1: {1, 2, 3},
	}

	expect := "大吉"

	for m, days := range periods {
		for _, d := range days {
			time := time.Date(2019, time.Month(m), d, 0, 0, 0, 0, time.Local)
			omikuji := omikuji.Omikuji{time}
			actual := omikuji.Do()
			if expect != actual {
				t.Errorf(`Omikuji error: expect="%s" actual="%s"`, expect, actual)
			}
		}
	}
}
