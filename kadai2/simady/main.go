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
	convertibleExts = []string{"jpg", "png", "gif"}
)

func main() {
	in := flag.String("in", "./", "input directory.")
	out := flag.String("out", "./", "output directory.")
	src := flag.String("src", "jpg", "source extension. [jpg, png, gif]")
	dst := flag.String("dst", "png", "destination extension. [jpg, png, gif]")
	flag.Parse()

	if !isSupported(*src) || !isSupported(*dst) {
		log.Fatalln("Unsupported extension.")
		flag.PrintDefaults()
	}
	if err := imgConvMain(*in, *out, *src, *dst); err != nil {
		log.Fatalf("Error occurred in convert process. err = %v\n", err)
	}
}

// imgConvMain メイン処理の実行.
func imgConvMain(in, out, src, dst string) error {
	conversionTargetFiles, err := getConversionTargetFiles(in, src)
	if err != nil {
		return err
	}
	c, err := converter.NewConverter(dst)
	if err != nil {
		return err
	}

	for _, file := range conversionTargetFiles {
		ext := filepath.Ext(file)
		out := out + "/" + strings.ReplaceAll(filepath.Base(file), ext, "."+dst)
		if c.Convert(file, out) != nil {
			return err
		}
	}
	return nil
}

// getConversionTargetFiles 変換対象のファイルを取得する.
func getConversionTargetFiles(dir, targetExt string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == "."+targetExt {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

// isSupported サポート対象の拡張子か判定する.
func isSupported(ext string) bool {
	for _, e := range convertibleExts {
		if ext == e {
			return true
		}
	}
	return false
}
