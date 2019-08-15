package fortune_test

import (
	"testing"
	"time"

	"github.com/gopherdojo/dojo6/kadai4/pei/pkg/fortune"
)

type MockClock struct {
	currentTime time.Time
}

func (mc MockClock) GetCurrentTime() time.Time {
	return mc.currentTime
}

func TestFortune_Drawing(t *testing.T) {
	cases := []struct {
		clock time.Time
	}{
		{time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)},
		{time.Date(2019, 1, 3, 23, 59, 59, 0, time.UTC)},
	}

	for _, c := range cases {
		mc := &MockClock{c.clock}
		f := fortune.NewFortune(mc)
		if actual := f.Drawing(); actual != "大吉" {
			t.Errorf("unexpected result: %s on %v", actual, mc.GetCurrentTime())
		}
	}
}
