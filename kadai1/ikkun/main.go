package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/gopherdojo/dojo6/kadai1/ikkun/imgconv"
)

const (
	ExitOK = iota
	ExitError
)

type CLI struct {
	outStream io.Writer
	errStream io.Writer
}

func main() {
	c := &CLI{os.Stdout, os.Stderr}
	os.Exit(c.run())
}

func (c *CLI) run() int {
	from := flag.String("f", "jpg", "from format(jpg|jpeg|png|gif) default jpg.")
	to := flag.String("t", "png", "to format(jpg|jpeg|png|gif) default png.")

	flag.Parse()

	if !validateFormat(*from) {
		fmt.Fprintln(c.errStream, fmt.Sprintf("unsupport format: %s", *from))
		return ExitError
	}

	if !validateFormat(*to) {
		fmt.Fprintln(c.errStream, fmt.Sprintf("unsupport format: %s", *to))
		return ExitError
	}

	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintln(c.errStream, "required directory assigned")
		return ExitError
	}

	dir := args[0]
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			paths = append(paths, path)
		}

		return nil
	})
	if err != nil {
		fmt.Fprintln(c.errStream, err)
		return ExitError
	}

	for _, path := range paths {
		img := imgconv.ImageFile{path, *from, *to}
		err = imgconv.Convert(img)
		if err != nil {
			// 対象ファイル外が含まれている場合は、標準出力するが処理自体は止めない
			fmt.Fprintf(c.outStream, "%s: %s\n", err, img.Source)
		}
	}

	return ExitOK
}

func validateFormat(f string) bool {
	switch f {
	case "jpg", "jpeg", "png", "gif":
		return true
	}

	return false
}
