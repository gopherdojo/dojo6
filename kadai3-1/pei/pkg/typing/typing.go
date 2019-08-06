package typing

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Question receives words and returns isCorrectCh
func Question(words []string) <-chan bool {
	isCorrectCh := make(chan bool)

	go func() {
		stdin := bufio.NewScanner(os.Stdin)
		rand.Seed(time.Now().UnixNano())

		for {
			word := words[rand.Intn(len(words))]
			fmt.Println(word)

			stdin.Scan()
			answer := stdin.Text()
			if word == answer {
				isCorrectCh <- true
			} else {
				isCorrectCh <- false
			}
		}
		close(isCorrectCh)
	}()

	return isCorrectCh
}
