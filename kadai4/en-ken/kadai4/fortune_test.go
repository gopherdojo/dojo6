package main_test

import (
	"fmt"
	"testing"
	"time"

	main "github.com/gopherdojo/dojo6/kadai4/en-ken/kadai4"
)

func TestGetLuck(t *testing.T) {

	cases := []struct {
		now      time.Time
		expected string
	}{
		{
			now:      time.Date(2018, time.December, 31, 0, 0, 0, 0, time.Local),
			expected: "凶",
		},
		{
			now:      time.Date(2018, time.December, 31, 23, 59, 59, 999, time.Local),
			expected: "凶",
		},
		{
			now:      time.Date(2019, time.January, 1, 0, 0, 0, 0, time.Local),
			expected: "大吉",
		},
		{
			now:      time.Date(2019, time.January, 2, 0, 0, 0, 0, time.Local),
			expected: "大吉",
		},
		{
			now:      time.Date(2019, time.January, 3, 0, 0, 0, 0, time.Local),
			expected: "大吉",
		},
		{
			now:      time.Date(2018, time.January, 3, 23, 59, 59, 999, time.Local),
			expected: "大吉",
		},
		{
			now:      time.Date(2019, time.January, 4, 0, 0, 0, 0, time.Local),
			expected: "凶",
		},
	}

	main.SetFortunes([]string{"凶", "凶", "凶"})

	for i, c := range cases {
		c := c
		t.Run(
			fmt.Sprintf("case[%v]", i),
			func(t *testing.T) {
				main.SetNow(func() time.Time {
					return c.now
				})

				if actual := main.GetFortune(); actual != c.expected {
					t.Errorf("actual:%v, expected:%v\n", actual, c.expected)
				}
			})
	}

}
