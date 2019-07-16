package Converter

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// 変換元ディレクトリと変換後ディレクトリの情報を保持する構造体
type FileInfo struct {
	// 変換元ディレクトリの情報
	Base DirType
	// 変換後ディレクトリの情報
	Dist DirType
}

// ディレクトリ別の情報を保持する構造体
type DirType struct {
	// ディレクトリ名
	DirName   string
	// ディレクトリに格納されている画像ファイルの拡張子
	Extension string
	// ディレクトリに格納されている画像ファイルのパス
	FilePaths []string
}

// ディレクトリ情報のインスタンスを作成
func createFileInfo() *FileInfo {
	a := FileInfo{
		DirType{},
		DirType{},
	}
	a.setArgs()

	// 変換元ディレクトリ内の全てのファイルのパスをセット
	ss, err := walkFilePath(a.Base.DirName)
	if err != nil {
		fmt.Println(err)
	}
	a.Base.FilePaths = ss

	return &a
}

// CLIで渡された引数を構造体に格納する
func (a *FileInfo) setArgs() {
	//　フラグをセット
	var (
		baseExtension = flag.String("base", "jpg", "base extension")
		distExtension = flag.String("dist", "png", "dist extension")
	)

	// フラグをパース
	flag.Parse()

	// 変換元ファイルの情報をセットする
	a.Base.DirName = flag.Arg(0)
	a.Base.Extension = *baseExtension

	// 変換後ファイルの情報をセットする
	a.Dist.DirName = flag.Arg(1)
	a.Dist.Extension = *distExtension

}

// ディレクトリ内にある全てのファイルのパスを取得する
func walkFilePath(dirname string) ([]string, error) {
	var s []string

	err := filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".jpg" || filepath.Ext(path) == ".png" {
			s = append(s, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	fmt.Println(s)
	return s, nil
}
