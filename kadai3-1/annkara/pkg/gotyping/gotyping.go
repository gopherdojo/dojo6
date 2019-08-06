package gotyping

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

type gotyping struct {
	outStream io.Writer
	errStream io.Writer
	limits    time.Duration
	result    int
}

// Run the gotyping
func Run(outStream, errStream io.Writer) error {

	var l int
	flag.IntVar(&l, "limits", 60, "制限時間")
	flag.Parse()

	g := gotyping{
		outStream: outStream,
		errStream: errStream,
		limits:    time.Duration(l) * time.Second,
	}

	err := g.start()

	if err != nil {
		return err
	}

	return nil
}

type data struct {
	in  string
	err error
}

func (g *gotyping) start() error {

	rand.Seed(time.Now().Unix())
	fmt.Fprintln(g.outStream, "=== gotyping start ===")

	in := make(chan data)
	input(in)

END:
	for {
		question := word()
		fmt.Fprintln(g.outStream, fmt.Sprintf("> %v", question))

		select {
		case <-time.After(g.limits):
			fmt.Fprintln(g.outStream, "=== gotyping finish ===")
			break END
		case answer := <-in:
			if answer.err != nil {
				return answer.err
			}
			if answer.in == question {
				g.result++
			}

		}
	}

	fmt.Fprintf(g.outStream, "Score: %v\n", g.result)
	return nil
}

func word() string {
	words := []string{"go", "Java", "C", "ruby", "perl", "assmbler"}
	return words[rand.Intn(len(words))]
}

func input(in chan data) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		in <- data{in: scanner.Text()}
	}
	if err := scanner.Err(); err != nil {
		in <- data{err: err}
	}
}
