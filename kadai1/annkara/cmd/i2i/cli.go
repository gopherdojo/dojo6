package main

import (
	"flag"
	"fmt"
	"io"
	"log"
)

const (
	exitCodeOK  = iota
	exitCodeErr = 10 + iota
)

type cli struct {
	outStream, errStream io.Writer
}

func (c *cli) run(args []string) int {

	log.SetOutput(c.outStream)

	var (
		after  string
		target string
	)

	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	flags.Usage = func() {
		fmt.Fprint(c.errStream, helpText)
	}
	flags.StringVar(&after, "a", "jpg", "変換対象の画像形式を指定")
	flags.StringVar(&after, "after", "jpg", "変換対象の画像形式を指定")
	flags.StringVar(&target, "t", "png", "変換後の画像形式を指定")
	flags.StringVar(&target, "target", "png", "変換後の画像形式を指定")

	if err := flags.Parse(args[1:]); err != nil {
		return exitCodeErr
	}

	return exitCodeOK
}

var helpText = `Usage: i2i [options] directory

i2i は指定されたディレクトリ内の画像ファイルを変換するコマンドラインツールです。
オプションを指定しない場合には、JPEGファイルを対象にPNG形式へと変換します。
変換後のファイルは同一ディレクトリ内に出力され、変換前のファイルは削除されません。
JPEGまたはPNG形式をサポートします。

Options:
  -a, --after                     変換後の画像形式を指定
  -h, --help                      ヘルプを表示
  -t, --target                    変換対象の画像形式を指定
`
