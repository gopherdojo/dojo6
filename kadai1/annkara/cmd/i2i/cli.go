package main

import "io"

const (
	exitCodeOK  = iota
	exitCodeErr = 10 + iota
)

type cli struct {
	outStream, errStream io.Writer
}

func (c *cli) run() int {

	return exitCodeOK
}
