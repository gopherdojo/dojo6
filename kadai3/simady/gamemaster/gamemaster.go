package gamemaster

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"time"
)

type GameMaster interface {
	Play()
}

type gameMaster struct {
	// 入力元
	r io.Reader
	// 出力先
	w io.Writer
	// 制限時間(分)
	timeLimit time.Duration
	// 問題
	problems []string
	// 回答数
	answerNum int
	// 正答数
	correctAnswerNum int
}

// New GameMasterを生成する.
func New(reader io.Reader, writer io.Writer, timeLimit time.Duration, problems []string) GameMaster {
	return &gameMaster{
		r:         reader,
		w:         writer,
		timeLimit: timeLimit,
		problems:  problems,
	}
}

func (gm *gameMaster) Play() {
	gm.displayRule()
	gm.game()
	gm.displayResult()
}

// displayRule ルールを表示する.
func (gm *gameMaster) displayRule() {
	fmt.Fprintln(gm.w, "【タイピングゲーム】画面に表示される英単語をできるだけ多く入力しましょう！")
	fmt.Fprintf(gm.w, "制限時間は%d秒です。\n", gm.timeLimit)
}

// game ゲームを行う.
func (gm *gameMaster) game() {
	ctx, cancel := context.WithTimeout(context.Background(), gm.timeLimit*time.Second)
	defer cancel()

	ch := gm.input()

gameLoop:
	for i := 0; i < len(gm.problems); i++ {
		problem := gm.problems[i]
		fmt.Fprintf(gm.w, "%d問目: %s\n", i+1, problem)
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

		if in == problem {
			gm.correctAnswerNum++
			fmt.Fprint(gm.w, "正解！")
		} else {
			fmt.Fprint(gm.w, "不正解...")
		}
		fmt.Fprintf(gm.w, " 現在の正答率:%d/%d\n", gm.correctAnswerNum, gm.answerNum)
	}
}

// displayResult 結果を表示する.
func (gm *gameMaster) displayResult() {
	fmt.Fprintln(gm.w)
	if gm.answerNum == len(gm.problems) {
		fmt.Fprintln(gm.w, "全問回答しました！")
	} else {
		fmt.Fprintln(gm.w, "タイムアップ！")
	}
	fmt.Fprintln(gm.w, "***")
	fmt.Fprintf(gm.w, "%d問中%d問正解\n", gm.answerNum, gm.correctAnswerNum)
	fmt.Fprintln(gm.w, "***")
	fmt.Fprintln(gm.w, "お疲れ様でした！")
}

func (gm *gameMaster) input() <-chan string {
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
