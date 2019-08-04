package game

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"time"
)

// Game はtタイピングゲームのコンテキストや入出力などの情報を格納します.
type Game struct {
	Context   context.Context
	Output    io.Writer
	Input     io.Reader
	TimeLimit time.Duration
	Words     []string
	Score     Score
}

// Score はゲーム結果を格納します.
type Score struct {
	Count         int
	CorrectNumber int
}

// NewGame は Gme の構造体を新しく作ります.
func NewGame(ctx context.Context, w io.Writer, r io.Reader, t time.Duration, words []string) *Game {
	return &Game{
		Context:   ctx,
		Output:    w,
		Input:     r,
		TimeLimit: t,
		Words:     words,
	}
}

// Run は TimeLimit に経過時間が達するまで Words をランダムに表示します.
// Word と同じ文字列を入力すると, correct!! 異なった文字列を入力すると, incorrect!! と表示されます.
func (g *Game) Run() error {
	fmt.Fprintln(g.Output, "===== Typing Game Start =====")

	scan := make(chan string)
	go func() {
		scanner := bufio.NewScanner(g.Input)
		defer close(scan)
		for scanner.Scan() {
			scan <- scanner.Text()
		}
	}()

	for {
		word := g.Words[g.Score.Count]
		fmt.Fprintln(g.Output, fmt.Sprintf("> %v", word))
		g.Score.Count++

		switch {
		case g.Score.Count == len(g.Words):
			fmt.Fprintln(g.Output, fmt.Sprintf(
				"\n===== Typing Game Finished =====\n%vwords completed. %v times corrected.\n",
				g.Score.Count,
				g.Score.CorrectNumber,
			))
			return nil

		}

		select {
		case <-g.Context.Done():
			fmt.Fprintln(g.Output, fmt.Sprintf(
				"\n===== Typing Game Finished =====\n%v have passed. %v times corrected.\n",
				g.TimeLimit,
				g.Score.CorrectNumber,
			))
			return nil
		case input := <-scan:
			switch {
			case g.Judgment(word, input):
				fmt.Fprintln(g.Output, "correct!!")
			default:
				fmt.Fprintln(g.Output, "incorrect!!")
			}
		}
	}
}

// Judgment は Word と同じ文字列が入力されているかを判断し, 正解数をインクリメントします.
func (g *Game) Judgment(word, input string) bool {
	if word == input {
		g.Score.CorrectNumber++
		return true
	}
	return false
}
