package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gopherdojo/dojo6/kadai1/yashiken/convert"
	"github.com/gopherdojo/dojo6/kadai1/yashiken/info"
)

var extSrc = flag.String("s", "jpg", "Extension of source file.")
var extCnv = flag.String("d", "png", "Extension of converted file.")

func main() {
	flag.Parse()

	dir := flag.Arg(0)
	if dir == "" {
		fmt.Println("No argument is specified.")
		os.Exit(1)
	}

	// 指定した形式をもつファイルのパスを取得
	filepaths, err := info.GetFilePath(dir, *extSrc)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 変換方法を定義する構造体の生成
	c := convert.NewConverter(*extSrc, *extCnv)
	// 画像の変換を実行
	for _, filepath := range filepaths {

		err := c.Convert(filepath)
		if err != nil {
			fmt.Println(err)
		}
	}
}
