package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
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
		chrcv = input(os.Stdin)
	)
	// 単語リストの取得
	words, err := getWords(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// 単語リストの要素をランダムに入れ替え
	words = shuffle(words)

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

// getWordsは指定したファイルの各行の文字列をスライスに
// 格納して返します。
func getWords(s string) ([]string, error) {
	file, err := os.Open(s)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sl []string

	for scanner.Scan() {
		if scanner.Text() != "" {
			sl = append(sl, scanner.Text())
		}
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return sl, nil
}

// inputは新たなゴルーチンを生成し、そのゴルーチンが
// 標準入力の内容を受け取ったあとチャネルに格納して返します。
func input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
	}()
	return ch
}

// shuffleはスライスの要素をランダムに入れ替え、
// 要素を入れ替えたスライスを返します。
func shuffle(s []string) []string {
	rand.Seed(time.Now().UnixNano())
	for i := range s {
		j := rand.Intn(len(s))
		s[i], s[j] = s[j], s[i]
	}
	return s
}
