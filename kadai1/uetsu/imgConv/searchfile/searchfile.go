package searchfile

import (
	"io/ioutil"
	"path/filepath"
)

// RecursionFile 再帰的にファイル名の一覧を取得する
func RecursionFile(dirPath string) (filePathList []string) {

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			filePathList = append(filePathList, RecursionFile(filepath.Join(dirPath, file.Name()))...)
			continue
		}
		filePathList = append(filePathList, filepath.Join(dirPath, file.Name()))
	}

	return
}
