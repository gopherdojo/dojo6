package omikuji_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/dojo6/kadai4/annkara/pkg/omikuji"
)

type result struct {
	r1 bool
	r2 bool
	r3 bool
	r4 bool
	r5 bool
	r6 bool
}

func (r *result) finished() bool {
	return r.r1 && r.r2 && r.r3 && r.r4 && r.r5 && r.r6
}

func TestOmikuji(t *testing.T) {
	n := time.Now().UnixNano()
	rand.Seed(n)

	r := result{}
	for {

		if r.finished() {
			return
		}

		me, unsei := omikuji.Draw()
		switch me {
		case 6:
			r.r6 = true
			if unsei != "大吉" {
				t.Errorf("me: %d, expected: %v, actual: %v", me, "大吉", unsei)
			}
		case 5, 4:
			if me == 5 {
				r.r5 = true
			} else {
				r.r4 = true
			}
			if unsei != "中吉" {
				t.Errorf("me: %d, expected: %v, actual: %v", me, "中吉", unsei)
			}
		case 3, 2:
			if me == 3 {
				r.r3 = true
			} else {
				r.r2 = true
			}
			if unsei != "小吉" {
				t.Errorf("me: %d, expected: %v, actual: %v", me, "小吉", unsei)
			}
		case 1:
			r.r1 = true
			if unsei != "凶" {
				t.Errorf("me: %d, expected: %v, actual: %v", me, "凶", unsei)
			}
		default:
			t.Fatalf("Invalid value %d", me)
		}
	}
}
