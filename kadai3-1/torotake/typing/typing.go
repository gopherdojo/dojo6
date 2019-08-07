/*
Package typing はタイピングゲームを実装したパッケージです。
Game型に
・問題の文字列のリスト
・入力(デフォルトでは標準入力を想定)
・出力(デフォルトでは標準出力を想定)
・制限時間
をセットしてRun()を呼ぶことで実行します。
問題はリストからランダムで出題され、入力から同じ文字列が入力されると正解でスコアが1加算されます。
*/
package typing

import (
	"bufio"
	crand "crypto/rand"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/rand"
	"time"
)

// Game 画像変換のオプション指定
type Game struct {
	Words   []string
	Input   io.Reader
	Output  io.Writer
	Score   int
	Timeout time.Duration
}

// Reason ゲームの終了結果を表す型
type Reason int

const (
	// Unknown 不明
	Unknown Reason = iota
	// Timeout タイムアウト
	Timeout
	// InputClosed 入力が閉じられた
	InputClosed
)

// Run optionsに指定された内容に従って画像を変換します
func (t *Game) Run() Reason {
	// 乱数シードの初期化
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())
	// タイムアウト
	timeoutCh := time.After(t.Timeout)

	ch := inputCh(t.Input)
	for {
		fmt.Fprintf(t.Output, "Let's input : %v\n", t.getProblem())
		select {
		case text, ok := <-ch:
			if !ok {
				// 入力が閉じられたので終了
				fmt.Fprintf(t.Output, "input closed\n")
				return InputClosed
			}
			if t.match(text) {
				t.Score++
				fmt.Fprintf(t.Output, "...OK!\n\n")
			} else {
				fmt.Fprintf(t.Output, "...NG!\n\n")
			}
		case <-timeoutCh:
			// タイムアウトで終了
			fmt.Fprintf(t.Output, "timeout\n")
			return Timeout
		}
	}
}

func (t *Game) getProblem() string {
	index := rand.Intn(len(t.Words))
	return t.Words[index]
}

func (t *Game) match(answer string) bool {
	for _, v := range t.Words {
		if answer == v {
			return true
		}
	}
	return false
}

func inputCh(r io.Reader) <-chan string {
	lines := make(chan string)
	go func() {
		defer close(lines)
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			lines <- scanner.Text()
		}
	}()
	return lines
}
