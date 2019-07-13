package imgcnv

import (
	"io/ioutil"
	"path/filepath"
)

// AllFileInfo returns
// the paths of the files in the specified directory
// filtered by the specified extension.
func AllFileInfo(dirPath string, ext string) ([]string, error) {
	absPath, err := filepath.Abs(dirPath)
	if err != nil {
		return nil, err
	}

	return searchFiles(absPath, ext)
}

func searchFiles(dirPath string, ext string) ([]string, error) {

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	result := make([]string, 0) // allocate 0 items to merge 2 slices
	for _, file := range files {
		path := filepath.Join(dirPath, file.Name())
		if file.IsDir() {
			paths, _ := searchFiles(path, ext)
			result = append(result, paths...)
		} else {
			// ignore error
			isMatch, _ := filepath.Match("*"+ext, file.Name())
			if isMatch {
				result = append(result, path)
			}
		}
	}
	return result, nil
}
