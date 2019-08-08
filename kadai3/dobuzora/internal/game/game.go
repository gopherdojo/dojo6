/*
Package game implements typing game for terminal.
*/
package game

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"math/rand"
	"time"
)

type Game struct {
	r                io.Reader
	w                io.Writer
	timeLimit        time.Duration
	questions        []string
	answerNum        int
	correctAnswerNum int
}

// New for game package
func New(reader io.Reader, writer io.Writer, timeLimit time.Duration, questions []string) *Game {
	return &Game{
		r:         reader,
		w:         writer,
		timeLimit: timeLimit,
		questions: questions,
	}
}

// Play is to start type game.
func (gm *Game) Play() {
	gm.displayRule()
	gm.shuffleQuestion()
	gm.playGame()
}

// displayRule is to display rules of type game.
func (gm *Game) displayRule() {
	fmt.Fprintf(gm.w, "画面に表示される英単語を入力しましょう!")
	fmt.Fprintf(gm.w, "制限時間は%dです\n", gm.timeLimit)
}

// playGame is type-game core logic.
func (gm *Game) playGame() {
	bc := context.Background()
	ctx, cancel := context.WithTimeout(bc, gm.timeLimit*time.Second)
	defer cancel()

	ch := gm.input()

gameLoop:
	for i := 0; i < len(gm.questions); i++ {
		question := gm.questions[i]
		fmt.Fprintf(gm.w, "%d問目: %s\n", i+1, question)
		fmt.Fprint(gm.w, ">")

		var in string
		var ok bool
		select {
		case in, ok = <-ch:
			if !ok {
				break gameLoop
			}
			gm.answerNum++
		case <-ctx.Done():
			break gameLoop
		}

		if in == question {
			gm.correctAnswerNum++
			fmt.Fprint(gm.w, "正解！")
		} else {
			fmt.Fprint(gm.w, "不正解...")
		}
		fmt.Fprintf(gm.w, " 現在の正答率:%d/%d\n", gm.correctAnswerNum, gm.answerNum)
	}
}

// displayResult is to display the result of type game.
func (gm *Game) displayResult() {
}

// shuffleQuestion is to shuffle given questions of typing game.
func (gm *Game) shuffleQuestion() {
	n := len(gm.questions)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		gm.questions[i], gm.questions[j] = gm.questions[j], gm.questions[i]
	}
}

func (gm *Game) input() <-chan string {
	ch := make(chan string)
	go func() {
		sc := bufio.NewScanner(gm.r)
		for sc.Scan() {
			ch <- sc.Text()
		}
		close(ch)
	}()
	return ch
}
