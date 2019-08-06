/*
infoパッケージは、指定されたディレクトリに格納されたイメージファイルの名前を取り出すためのパッケージです。
*/
package info

import (
	"io/ioutil"
	"path/filepath"
)

/*
GetFilePathは、指定されたディレクトリ配下に格納された、
特定の拡張子を持つ画像ファイルのファイルパス一覧を返します。
*/
func GetFilePath(dir, ext string) ([]string, error) {
	// ディレクトリ配下のファイル情報を取得
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	// ファイルパスを格納するためのスライスを生成
	paths := make([]string, 0)

	// ファイルパスを順次格納
	for _, file := range files {
		// ディレクトリがある場合、再帰的に処理
		if file.IsDir() {
			p, err := GetFilePath(filepath.Join(dir, file.Name()), ext)
			if err != nil {
				return nil, err
			}
			paths = append(paths, p...)
		}
		// スライスに追加
		path := filepath.Join(dir, file.Name())
		if filepath.Ext(path) == "."+ext {
			paths = append(paths, filepath.Join(dir, file.Name()))
		}
	}
	return paths, nil
}
