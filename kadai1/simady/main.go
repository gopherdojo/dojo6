package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"

	"imgconv/converter"
)

var (
	in, out, src, dst string
)

func init() {
	flag.StringVar(&in, "in", "", "target directory.")
	flag.StringVar(&src, "src", "jpeg", "source extension. [jpeg,png]")
	flag.StringVar(&dst, "dst", "png", "destination extension. [jpeg,png]")
}

func main() {
	flag.Parse()
	c := converter.GetConverter(src, dst)
	err := filepath.Walk(in, func(path string, info os.FileInfo, err error) error {
		ext := filepath.Ext(path)
		if ext != "."+src {
			return nil
		}
		out = "dst/" + strings.ReplaceAll(filepath.Base(path), ext, "."+dst)
		return converter.Convert(path, c, out)
	})
	if err != nil {
		log.Fatalf("Error occurred in convert process. err = %v\n", err)
	}
}
