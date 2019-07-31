/*
cli はコマンドライン引数を処理します。
*/
package cli

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	cnv "github.com/gopherdojo/dojo6/kadai2/dobuzora/internal/convert"
)

// isDir は ディレクトリかどうかを判断します
func isDir(dirName string) bool {
	dInfo, err := os.Stat(dirName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return dInfo.IsDir()
}

// ファイル名から拡張子を取ったファイル名を返します
func getFileNameWithoutExt(path string) string {
	// Fixed with a nice method given by mattn-san
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

// imageConvert はディレクトリ名を受け取り、そのディレクトリ内のファイルを変換します
func imageConvert(dirName string) []cnv.ConvertFile {
	var fileList []cnv.ConvertFile

	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	//
	for _, file := range files {
		path := filepath.Join(dirName, file.Name())
		if isDir(path) {
			flist := imageConvert(path)
			fileList = append(fileList, flist...)
		} else {
			ext := filepath.Ext(path)
			if ext == ".jpeg" || ext == ".jpg" {
				newPath := filepath.Dir(path) + "/" + getFileNameWithoutExt(path) + ".png"
				fileList = append(fileList, cnv.ConvertFile{Old: path, New: newPath})
			}
		}
	}
	return fileList
}

// Do は ファイルを走査し、convertパッケージを用いてJPEをPNG形式に変換します
func Do() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Wrong number of arguments")
		os.Exit(1)
	}
	rootDir := os.Args[1]
	if !isDir(rootDir) {
		fmt.Fprintln(os.Stderr, "Not a directory")
		os.Exit(1)
	}

	hoge := imageConvert(rootDir)
	for _, v := range hoge {
		v.ConvertToPng()
		os.Remove(v.Old)
	}
}
