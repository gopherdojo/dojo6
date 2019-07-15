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
	exitCodeOK  = 0
	exitCodeErr = 10
)

type cli struct {
	outStream, errStream io.Writer
}

func (c *cli) walk(root, before, after string) error {

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		n := info.Name()
		if strings.HasSuffix(n, before) {
			origin, err := os.Open(path)
			if err != nil {
				return err
			}
			defer origin.Close()

			// 拡張子を含まない出力用ファイル名
			n := filepath.Base(n[:len(n)-len(filepath.Ext(n))])
			dir := filepath.Dir(path)
			out, err := os.Create(filepath.Join(dir, n+"."+after))
			if err != nil {
				return err
			}

			err = image.Convert(origin, out, after)
			if err != nil {
				// 変換処理に失敗した場合、不要なファイルが作成されてしまうため、削除する
				// ファイルを閉じた後でないと、Windowsの場合削除できないのでここでCloseする
				out.Close()
				e := os.Remove(filepath.Join(dir, n+"."+after))
				if e != nil {
					return e
				}
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
		err := c.walk(v, before, after)
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
