package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	divdl "github.com/gopherdojo/dojo6/kadai3-2/en-ken"

	"github.com/pkg/errors"
)

// Do is I/F to divdl.Do
type Do func(url string, fileName string, numOfDivision int) error

func outputError(err error) {
	os.Stderr.Write([]byte(fmt.Sprintf("%v", errors.WithStack(err))))
}

func main() {
	var (
		fileName      string
		numOfDivision int
	)

	flag.StringVar(&fileName, "o", "[remote-name]", "output file name")
	flag.IntVar(&numOfDivision, "n", 5, "number of downloading in parallel")
	flag.Parse()
	url := flag.Arg(0)
	// Set remote name as fileName
	if fileName == "[remote-name]" {
		fileName = path.Base(url)
	}

	err := Do(divdl.Do)(url, fileName, numOfDivision)
	if err != nil {
		outputError(err)
		return
	}

	os.Stdout.Write([]byte("Downloading done.\n"))
}
