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
	words, err := getWords(s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("タイピングゲームを始めます。制限時間は%d分です。\n", t)
	for i := true; i && score < len(words); {
		qst := words[score]
		fmt.Println(qst)
		select {
		case ans := <-chrcv:
			if qst == ans {
				score++
			}
		case <-tm:
			fmt.Println("制限時間です。ゲームを終了します。")
			i = false
		}
	}
	fmt.Printf("あなたの正解数は%d問です。", score)
}

func getWords(s string) ([]string, error) {
	file, err := os.Open(s)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sl []string

	for scanner.Scan() {
		sl = append(sl, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}

	sl = shuffle(sl)

	return sl, nil
}

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

func shuffle(s []string) []string {
	rand.Seed(time.Now().UnixNano())
	for i := range s {
		j := rand.Intn(len(s))
		s[i], s[j] = s[j], s[i]
	}
	return s
}
