package main

import (
	"os"

	"typing-game/gamemaster"
)

const (
	timeLimit = 10
)

var (
	reader   = os.Stdin
	writer   = os.Stdout
	problems = []string{"apple", "bake", "cup", "dog", "egg", "fight", "green", "hoge", "idea", "japan"}
)

func main() {
	gamemaster.New(reader, writer, timeLimit, problems).Play()
}
