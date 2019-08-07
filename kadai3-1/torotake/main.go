package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gopherdojo/dojo6/kadai3-1/torotake/typing"
)

var (
	optWordsPath string
	optTimeout   int
	exitCode     int
)

func init() {
	// -f=[問題文ファイル] default : words.txt
	// -t=[制限時間(秒)] default : 30
	flag.StringVar(&optWordsPath, "f", "words.txt", "question words file path.")
	flag.IntVar(&optTimeout, "t", 30, "timeout (s)")
	flag.Parse()
}

func main() {
	exec()
	os.Exit(exitCode)
}

func exec() {
	words, err := loadWords(optWordsPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't load question words file.\n")
	}

	timeout, _ := time.ParseDuration(fmt.Sprintf("%ds", optTimeout))
	game := typing.Game{
		Words:   words,
		Input:   os.Stdin,
		Output:  os.Stdout,
		Timeout: timeout,
	}

	game.Run()

	fmt.Printf("Game over!\nYour score is %v\n", game.Score)
}

// 問題文ファイルの読み込み
// 1行が1つの問題。前後の空白はトリムし、空行は無視する。
func loadWords(path string) ([]string, error) {
	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	words := make([]string, 0)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if len(text) > 0 {
			words = append(words, text)
		}
	}

	return words, nil
}
