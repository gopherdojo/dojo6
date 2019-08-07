package typing_test

import (
	"io/ioutil"
	"strings"
	"testing"
	"time"

	"github.com/gopherdojo/dojo6/kadai3-1/torotake/typing"
)

func TestGame_Run(t *testing.T) {
	answer := "hoge\nfoo\nhoge\nbar\nhoge\n"
	expectedScore := 3
	expectedDuration := time.Duration(5 * 1000 * 1000 * 1000) // 5sec
	durationMargin := time.Duration(500 * 1000 * 1000)        // 500msはOKとする

	// strings.NewReader(answer)だとすぐ全て読み取ってEOFで終わってしまうのでタイムアウトまで持たない
	// ダミーで無限に改行を読み取り続けるio.Readerを使う
	input := newInfiniteReader(answer)
	startTime := time.Now()

	g := typing.Game{
		Words:   []string{"hoge"},
		Input:   input,
		Output:  ioutil.Discard,
		Timeout: expectedDuration,
	}
	g.Run()

	duration := time.Since(startTime)
	if duration < (expectedDuration-durationMargin) || duration > (expectedDuration+durationMargin) {
		t.Errorf("Time error    expected:(%v +- %v) actual:%v", expectedDuration, durationMargin, duration)
	}

	if g.Score != expectedScore {
		t.Errorf("Score error   expected:%d actual:%d", expectedScore, g.Score)
	}
}

type InfiniteReader struct {
	Src *strings.Reader
}

func newInfiniteReader(src string) *InfiniteReader {
	return &InfiniteReader{strings.NewReader(src)}
}

func (r *InfiniteReader) Read(p []byte) (int, error) {
	n, err := r.Src.Read(p)
	if err != nil {
		// 元文字列を読み切った後はEOFが来るので、ダミーの改行を読み取ったことにする
		p[0] = '\n'
		n = 1
		err = nil
	}
	return n, err
}
