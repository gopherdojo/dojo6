package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"./convert"
)

func main() {
	// 第一引数を取得
	if len(os.Args) != 2 {
		fmt.Println("引数の数ちゃうで")
		os.Exit(1)
	}
	rootDir := os.Args[1]
	dInfo, err := os.Stat(rootDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if !dInfo.IsDir() {
		fmt.Println("ディレクトリやないとあかんで")
		os.Exit(1)
	}

	files, err := ioutil.ReadDir(rootDir)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		path := filepath.Join(dInfo.Name(), file.Name())
		err := convert.ConvertToPng(path)
		if err != nil {
			fmt.Println(err)
		}
	}
}
