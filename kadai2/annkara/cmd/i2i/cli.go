package main

import (
	"flag"
	"fmt"
	"io"
	"log"

	"github.com/dojo6/kadai2/annkara/pkg/dir"
)

const (
	exitCodeOK  = 0
	exitCodeErr = 10
)

type cli struct {
	outStream, errStream io.Writer
}

func (c *cli) run(args []string) int {

	log.SetOutput(c.outStream)

	var (
		before string
		after  string
		debug  bool
	)

	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	flags.Usage = func() {
		fmt.Fprint(c.errStream, helpText)
	}

	flags.StringVar(&before, "b", "jpg", "変換前の画像形式を指定")
	flags.StringVar(&before, "before", "jpg", "変換前の画像形式を指定")
	flags.StringVar(&after, "a", "png", "変換後の画像形式を指定")
	flags.StringVar(&after, "after", "png", "変換後の画像形式を指定")
	flags.BoolVar(&debug, "debug", false, "")

	if err := flags.Parse(args[1:]); err != nil {
		if debug {
			log.Printf("failed: %+v\n", err)
		}
		return exitCodeErr
	}

	for _, v := range flags.Args() {
		err := dir.Walk(v, before, after)
		if err != nil {
			if debug {
				log.Printf("failed: %+v\n", err)
			}
			return exitCodeErr
		}
	}

	return exitCodeOK
}

var helpText = `Usage: i2i [options] directory

i2i は指定されたディレクトリ内の画像ファイルを変換するコマンドラインツールです。
オプションを指定しない場合には、JPEGファイルを対象にPNG形式へと変換します。
変換後のファイルは同一ディレクトリ内に出力され、変換前のファイルは削除されません。
JPEGまたはPNG形式をサポートします。

Options:
  -a, -after                     変換後の画像形式を指定
  -b, -before                    変換前の画像形式を指定
  -h, -help                      ヘルプを表示
`
