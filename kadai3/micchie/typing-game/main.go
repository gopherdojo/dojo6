package main

import (
	"context"
	"math/rand"
	"os"
	"time"

	"github.com/gopherdojo/dojo6/kadai3/micchie/typing-game/game"
	"github.com/jinzhu/configor"
)

// Config is a structure of config.yml.
var Config = struct {
	Limit time.Duration `default:"30"`
	Words []string
}{}

func main() {
	configor.Load(&Config, "config.yml")

	t := Config.Limit * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()

	list := Config.Words
	shuffle(list)

	g := game.NewGame(ctx, os.Stdout, os.Stdin, t, list)
	if err := g.Run(); err != nil {
		os.Exit(0)
	}
}

func shuffle(list []string) {
	rand.Seed(time.Now().UnixNano())
	for i := range list {
		j := rand.Intn(i + 1)
		list[i], list[j] = list[j], list[i]
	}
}
