package main

import (
	"time"
)

var GetFortune = getFortune
var Handler = handler

func SetNow(n func() time.Time) {
	now = n
}

func SetFortunes(f []string) {
	fortunes = f
}

func SetNewEncoder(n NewEncoder) {
	newEncoder = n
}
