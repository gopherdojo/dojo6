package main

import (
	"math/rand"
	"time"
)

// For testing
var now = time.Now
var fortunes []string

func init() {
	rand.Seed(time.Now().UnixNano())
	fortunes = []string{"大吉", "吉", "中吉", "小吉", "末吉", "凶", "大凶"}
}

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
