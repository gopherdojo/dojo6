package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gopherdojo/dojo6/kadai3-1/yashiken/typegame"
)

var t int
var s string

// フラグのパース
func init() {
	flag.IntVar(&t, "t", 1, "Time limit")
	flag.StringVar(&s, "s", "wordlist.txt", "wordlist")
	flag.Parse()
}

func main() {
	var (
		tm    = time.After(time.Duration(t) * time.Minute)
		score = 0
		chrcv = typegame.Input(os.Stdin)
	)
	// 単語リストの取得
	words, err := typegame.Words(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// 単語リストの要素をランダムに入れ替え
	words = typegame.Shuffle(words)

	fmt.Printf("タイピングゲームを始めます。制限時間は%d分です。\n", t)
	for i := true; i && score < len(words); {
		qst := words[score]
		fmt.Println(qst)
		select {
		// チャネルを通して標準入力から文字列を受け取った場合、
		// 正誤判定を行い、正解ならscoreをインクリメント
		case ans := <-chrcv:
			if qst == ans {
				score++
			}
		// チャネルを通して制限時間の終了を通知されたとき、
		// ループを脱出
		case <-tm:
			fmt.Println("制限時間です。ゲームを終了します。")
			i = false
		}
	}
	fmt.Printf("あなたの正解数は%d問です。", score)
}
