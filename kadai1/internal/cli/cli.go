package cli

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gopherdojo/dojo6/kadai1/internal/convert"
)

// ディレクトリをさがしてやる。
func searchDir(dirName string) {
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		path := filepath.Join(dirName, file.Name())
		if isDir(path) {
			searchDir(path)
		} else {
			err := convert.ConvertToPng(path)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

// isDir は ファイルがディレクトリかどうかを判断します
func isDir(dirName string) bool {
	dInfo, err := os.Stat(dirName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return dInfo.IsDir()
}

func Do() {
	// 第一引数を取得
	if len(os.Args) != 2 {
		fmt.Println("引数の数ちゃうで")
		os.Exit(1)
	}
	rootDir := os.Args[1]
	if !isDir(rootDir) {
		fmt.Println("ディレクトリやないとあかんで")
		os.Exit(1)
	}

	searchDir(rootDir)
}
