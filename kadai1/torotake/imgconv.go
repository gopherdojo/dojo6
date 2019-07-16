package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/gopherdojo/dojo6/kadai1/torotake/convert"
)

var (
	optInputFormat  string
	optOutputFormat string
	exitCode        int
)

func init() {
	// -i=[変換前形式] default : jpeg
	// -o=[変換後形式] default : png
	// 変換対象ディレクトリ
	flag.StringVar(&optInputFormat, "i", "jpeg", "input file format.")
	flag.StringVar(&optOutputFormat, "o", "png", "output file format.")
	flag.Parse()
}

func main() {
	exec()
	os.Exit(exitCode)
}

func exec() {
	args := flag.Args()
	inputFormat := getFormat(optInputFormat)
	outputFormat := getFormat(optOutputFormat)

	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "%s\n", "option error : no target directory")
		exitCode = 1
		return
	}

	if inputFormat == convert.UNKNOWN || outputFormat == convert.UNKNOWN || inputFormat == outputFormat {
		fmt.Fprintf(os.Stderr, "%s\n", "option error : invalid format")
		exitCode = 1
		return
	}

	files, err := convert.ListFiles(args[0], inputFormat)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		exitCode = 2
		return
	}

	convert.Convert(convert.Options{
		SrcFiles:     files,
		OutputFormat: outputFormat})
}

func getFormat(f string) convert.Format {
	f = strings.ToLower(f)
	if f == "jpg" || f == "jpeg" {
		return convert.JPEG
	} else if f == "png" {
		return convert.PNG
	} else if f == "gif" {
		return convert.GIF
	} else if f == "bmp" {
		return convert.BMP
	}

	return convert.UNKNOWN
}
