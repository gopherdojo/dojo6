package main

import "os"

func main() {

	cli := &cli{
		outStream: os.Stdout,
		errStream: os.Stderr,
	}
	os.Exit(cli.run())
}
