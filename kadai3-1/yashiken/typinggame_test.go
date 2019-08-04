package main

import (
	"bytes"
	"errors"
	"path/filepath"
	"reflect"
	"testing"
)

func TestGetWords(t *testing.T) {
	cases := []struct {
		name      string
		file      string
		output    []string
		expectErr error
	}{
		{
			name:      "case1",
			file:      "abc.txt",
			output:    []string{"a", "b", "c"},
			expectErr: nil,
		},
		{
			name:      "case2",
			file:      "empty-file.txt",
			output:    nil,
			expectErr: nil,
		},
		{
			name:      "case3",
			file:      "noexist.txt",
			output:    nil,
			expectErr: errors.New("something error happened"),
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			path := filepath.Join("testdata/" + c.file)
			sl, err := getWords(path)
			if !reflect.DeepEqual(sl, c.output) && err != c.expectErr {
				t.Errorf("Input file is %v. expected output is %v, but getWords returns %v", c.file, c.output, sl)
			}
		})
	}
}
func TestInput(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "case1",
			input:  "test",
			output: "test",
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			r := bytes.NewBuffer([]byte(c.input))
			if got := <-input(r); got != c.output {
				t.Errorf("Input is %v. Expected output is %v ,but input returns %v", c.input, c.output, got)
			}
		})
	}
}
func TestShuffle(t *testing.T) {
	cases := []struct {
		name   string
		input  []string
		output []string
	}{
		{
			name:   "case1",
			input:  []string{"a", "b", "c"},
			output: []string{"c", "a", "b"},
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			var (
				result  []string
				success bool
			)
			// ランダムにスライスの要素を入れ替えるため、十分な回数shuffleを実行
			for i := 1; i < 1000; i++ {
				if result = shuffle(c.input); reflect.DeepEqual(result, c.output) {
					success = true
					break
				}
			}
			if !success {
				t.Errorf("expect result %v cannot be gotten", c.output)
			}
		})
	}
}
