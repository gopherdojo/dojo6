package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"

	"github.com/gopherdojo/dojo6/kadai3-2/pei/pkg/download"
)

const (
	exitCodeOk    = 0
	exitCodeError = 1

	splitNum = 4
)

type cliArgs struct {
	url, outputPath string
}

func (ca *cliArgs) validate() error {
	if ca.url == "" {
		return fmt.Errorf("No URL")
	}

	if ca.outputPath == "" {
		ca.outputPath = "./"
	}

	return nil
}

func main() {
	os.Exit(Run())
}

// Run runs download
func Run() int {
	ca := parseArgs()
	if err := ca.validate(); err != nil {
		fmt.Fprintln(os.Stderr, "Args error: ", err)
		return exitCodeError
	}

	downloader, err := download.NewDownloader(splitNum, ca.url, ca.outputPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Create downloader error: ", err)
		return exitCodeError
	}

	outputPath, err := downloader.Do()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Download error: ", err)
		return exitCodeError
	}

	var downloadType string
	if reflect.TypeOf(downloader) == reflect.TypeOf(&download.RangeDownloader{}) {
		downloadType = "Split Download"
	} else {
		downloadType = "Download"
	}
	fmt.Println("Download Type: ", downloadType)
	fmt.Println("Download completed. Output: ", outputPath)

	return exitCodeOk
}

func parseArgs() *cliArgs {
	var ca cliArgs
	flag.StringVar(&ca.outputPath, "o", "", "output path")
	flag.Parse()
	ca.url = flag.Arg(0)

	return &ca
}
