package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// For testing
var now = time.Now
var fortunes []string

func getFortune() string {
	i := rand.Int() % len(fortunes)
	f := fortunes[i]

	t := now()
	if t.Month() == 1 {
		if t.Day() >= 1 && t.Day() <= 3 {
			f = "大吉"
		}
	}

	return f
}
