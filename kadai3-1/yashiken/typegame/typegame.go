package typegame

import (
	"bufio"
	"io"
	"math/rand"
	"os"
	"time"
)

// getWordsは指定したファイルの各行の文字列をスライスに
// 格納して返します。
func Words(s string) ([]string, error) {
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
func Input(r io.Reader) <-chan string {
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
func Shuffle(s []string) []string {
	rand.Seed(time.Now().UnixNano())
	for i := range s {
		j := rand.Intn(len(s))
		s[i], s[j] = s[j], s[i]
	}
	return s
}
