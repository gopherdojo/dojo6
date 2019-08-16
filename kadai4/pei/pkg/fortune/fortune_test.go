package fortune_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gopherdojo/dojo6/kadai4/pei/pkg/fortune"
)

func TestFortune_Handler(t *testing.T) {
	f := fortune.NewFortune(fortune.DefaultClock{})
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	f.Handler(w, r)

	rw := w.Result()
	defer rw.Body.Close()
	if rw.StatusCode != http.StatusOK {
		t.Error("unexpected status code")
	}

	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Error("unexpected error")
	}

	dr := &fortune.DrawingResult{}
	if err := json.Unmarshal(b, &dr); err != nil {
		t.Error("failed json unmarshal")
	}

	fortuneList := f.GetFortuneList()
	contain := false
	for _, v := range fortuneList {
		if v == dr.Result {
			contain = true
			break
		}
	}

	if !contain {
		t.Errorf("unexpected response: %s", string(b))
	}
}

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
