package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gopherdojo/dojo6/kadai3-1/pei/pkg/typing"
	"github.com/gopherdojo/dojo6/kadai3-1/pei/pkg/wordsreader"
)

const (
	ExitCodeOK    = 0
	ExitCodeError = 1
)

func main() {
	os.Exit(execute())
}

func execute() int {
	fmt.Println("Start Typing Game!")

	wr := &wordsreader.WordsReader{FileName: "./textdata/words.txt"}
	words, err := wr.Read()
	if err != nil {
		fmt.Errorf("%v", err)
		return ExitCodeError
	}

	typingCh := typing.Question(words)
	timerCh := time.After(5 * time.Second)

	var (
		counter        int
		correctCounter int
	)

	for {
		counter++
		select {
		case isCorrect := <-typingCh:
			if isCorrect {
				correctCounter++
			}
		case <-timerCh:
			fmt.Printf("\nScore: %d/%d \n", correctCounter, counter)
			fmt.Println("End Typing Game!")
			return ExitCodeOK
		}
	}
}
