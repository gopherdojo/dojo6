package fortune

import (
	"fmt"
	"net/http"
	"math/rand"
	"time"
)

var fortuneList = []string{
	"大吉",
	"中吉",
	"吉",
	"小吉",
	"凶",
}

// Clock
type Clock struct {}

// Fortune
type Fortune struct {
	clock Clock
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GetCurrentTime return current time
func (c Clock) GetCurrentTime() time.Time {
	return time.Now()
}

// NewFortune return fortune instance
func NewFortune(c Clock) * Fortune {
	return &Fortune{clock: c}
}

func (f Fortune) drawingForNewYearDay() string {
	return "大吉"
}

func (f Fortune) defaultDrawing() string {
	return fortuneList[rand.Intn(len(fortuneList))]
}

func (f Fortune) isNewYearDay() bool {
	c := f.clock.GetCurrentTime()

	return c.Month() == 1 && 1 <= c.Day() && c.Day() <= 3
}

// Drawing return drawing result
func (f Fortune) Drawing() string {
	if f.isNewYearDay() {
		return f.drawingForNewYearDay()
	}
	return f.defaultDrawing()
}

// Handler
func (f Fortune) Handler(w http.ResponseWriter, r *http.Request) {
	result := f.Drawing
	fmt.Fprint(w, result)
}
