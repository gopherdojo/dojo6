package main

import (
	"os"

	"github.com/gopherdojo/dojo6/kadai3/dobuzora/internal/game"
)

const timeLimit = 10

var (
	reader    = os.Stdin
	writer    = os.Stdout
	questions = []string{"public", "void", "func", "return", "continue"}
)

func main() {
	g := game.New(reader, writer, timeLimit, questions)
	g.Play()
}
