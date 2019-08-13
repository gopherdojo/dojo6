package omikuji_test

import (
	"testing"

	"github.com/gopherdojo/dojo6/kadai4/omikuji"
)

func Test_SpecificPeriodLottery(t *testing.T) {
	periods := map[int][]int{
		1: {1, 2, 3},
	}

	expect := "大吉"

	for m, days := range periods {
		for _, d := range days {
			for i := 0; i < 20; i++ {
				actual := omikuji.Do(m, d)
				if expect != actual {
					t.Errorf(`Omikuji error: expect="%s" actual="%s"`, expect, actual)
				}
			}
		}
	}
}
