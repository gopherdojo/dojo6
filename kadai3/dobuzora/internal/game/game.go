package game

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"time"
)

// ゲームの情報を管理
type Game struct {
	r                io.Reader
	w                io.Writer
	timeLimit        time.Duration
	questions        []string
	answerNum        int
	correctAnswerNum int
}

func New(reader io.Reader, writer io.Writer, timeLimit time.Duration, questions []string) *Game {
	return &Game{
		r:         reader,
		w:         writer,
		timeLimit: timeLimit,
		questions: questions,
	}
}

func (gm *Game) Play() {
	gm.displayRule()
	gm.playGame()
}

func (gm *Game) displayRule() {
	fmt.Fprintf(gm.w, "画面に表示される英単語を入力しましょう!")
	fmt.Fprintf(gm.w, "制限時間は%dです\n", gm.timeLimit)
}

func (gm *Game) playGame() {
	ctx, cancel := context.WithTimeout(context.Background(), gm.timeLimit*time.Second)
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

func (gm *Game) displayResult() {
}

func (gm *Game) input() <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(gm.r)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}
