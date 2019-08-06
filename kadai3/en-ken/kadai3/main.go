package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gopherdojo/dojo6/kadai3/en-ken/typing"
)

func main() {
	const timeoutSec = 10
	textDict := []string{"foo", "bar", "baz", "qux"}
	t := typing.NewTyping(textDict)

	chInput := inputRoutine(os.Stdin)
	chFinish := time.After(time.Duration(timeoutSec) * time.Second)

	execute(chInput, chFinish, os.Stdout, t)
}

// Typing is interface to typing.Typing
type Typing interface {
	GetNextText() string
	IsCorrect(input string) bool
}

func execute(chInput <-chan string, chFinish <-chan time.Time, stdout *os.File, t Typing) {

	score := 0
	for i := 1; ; i++ {
		fmt.Fprintf(stdout, "[%03d]: %v\n", i, t.GetNextText())
		fmt.Fprint(stdout, "type>>")
		select {
		case text := <-chInput:
			if t.IsCorrect(text) {
				score++
				fmt.Fprintln(stdout, "Correct!")
			} else {
				fmt.Fprintln(stdout, "Correct!")
			}
		case <-chFinish:
			fmt.Fprintln(stdout, "\nTime's up!!")
			fmt.Fprintf(stdout, "You Scored: %v\n", score)
			return
		}
	}
}

func inputRoutine(r io.Reader) <-chan string {
	ch := make(chan string)

	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
	}()

	return ch
}
