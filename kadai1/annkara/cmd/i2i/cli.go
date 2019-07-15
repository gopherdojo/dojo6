package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/dojo6/kadai1/annkara/pkg/image"
)

const (
	exitCodeOK  = iota
	exitCodeErr = 10 + iota
)

type cli struct {
	outStream, errStream io.Writer
}

func (c *cli) walk(root, target, format string) error {

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		n := info.Name()
		if strings.HasSuffix(n, target) {
			origin, err := os.Open(path)
			if err != nil {
				return err
			}
			defer origin.Close()

			// 拡張子を含まない出力用ファイル名
			n := filepath.Base(n[:len(n)-len(filepath.Ext(n))])
			dir := filepath.Dir(path)
			out, err := os.Create(filepath.Join(dir, n+"."+format))
			if err != nil {
				return err
			}

			err = image.Convert(origin, out, format)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (c *cli) run(args []string) int {

	log.SetOutput(c.outStream)

	var (
		after  string
		target string
		debug  bool
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
	flags.BoolVar(&debug, "debug", false, "")

	if err := flags.Parse(args[1:]); err != nil {
		if debug {
			log.Printf("failed: %+v\n", err)
		}
		return exitCodeErr
	}

	for _, v := range flags.Args() {
		err := c.walk(v, target, after)
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
  -h, -help                      ヘルプを表示
  -t, -target                    変換対象の画像形式を指定
`
